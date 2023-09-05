package bookmark

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nao-United92/book-matcher/model"
)

/*
* ブックマークデータの構造体
 */
type Bookmark struct {
	BookmarkId int
	UserId     int
	ItemId     string
	BookInfo   []string
	CreateDate time.Time
}

/*
* ブックマークフォーム表示
* 【リクエスト】
*   c: リクエストとレスポンスのコンテキストを表すオブジェクト
*   db: データベース接続オブジェクト
 */
func ShowBookmarkForm(c *gin.Context, db *gorm.DB) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)

	// セッションからユーザーIDを取得、数値としてセット
	user_id := fmt.Sprintf("%v", session.Get("user_id"))

	// セッションから取得したユーザーIDをキーにブックマークデータを検索
	conditions := "user_id = ?"
	bookmarks, err := model.GetBookmarkDatas(db, conditions, user_id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "ブックマークデータの検索中にエラーが発生しました。"})
		return
	}

	// Book_infoのアンシリアライズに伴い、検索データを別の構造体配列へセット
	var recreateBookmarks []Bookmark
	for _, bookmark := range bookmarks {
		recreateBookmark := Bookmark{
			BookmarkId: bookmark.Bookmark_id,
			UserId:     bookmark.User_id,
			ItemId:     bookmark.Item_id,
			BookInfo:   unserialize(bookmark.Book_info),
			CreateDate: bookmark.Create_date,
		}
		recreateBookmarks = append(recreateBookmarks, recreateBookmark)
	}

	// ブックマークフォームを表示
	c.HTML(http.StatusOK, "bookmark.html", gin.H{
		"Bookmarks": recreateBookmarks,
	})
}

/*
* ブックマークへの書籍データ登録
* 【リクエスト】
*   c: リクエストとレスポンスのコンテキストを表すオブジェクト
*   db: データベース接続オブジェクト
 */
func AddBookmark(c *gin.Context, db *gorm.DB) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)

	// セッションからユーザーIDを取得、数値としてセット
	user_id := fmt.Sprintf("%v", session.Get("user_id"))

	// 書籍データより商品IDをセット
	item_id := c.PostForm("id")

	// 書籍データをシリアライズ(例：「フィールド名: データ」)
	serializedData := fmt.Sprintf("Title: %s, Subtitle: %s, Authors: %v, Description: %s, SmallThumbnail: %s, Publisher: %s, PublishedDate: %s, Saleability: %s, WebReaderLink: %s, AccessViewStatus: %s", c.PostForm("title"), c.PostForm("subtitle"), c.PostForm("authors"), c.PostForm("description"), c.PostForm("smallthumbnail"), c.PostForm("publisher"), c.PostForm("publisheddate"), c.PostForm("saleability"), c.PostForm("webreaderlink"), c.PostForm("accessviewstatus"))

	// ブックマークデータの新規登録
	bookmark, err := model.CreateBookmarkData(db, user_id, item_id, string(serializedData))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "ブックマークデータの登録中にエラーが発生しました。"})
		return
	}

	// ブックマークデータの新規登録が成功した場合
	if bookmark != nil {
		// ブックマークの登録結果をJSONで返す
		c.JSON(http.StatusOK, "追加しました。")
	}
}

/*
* ブックマークの削除
* 【リクエスト】
*   c: リクエストとレスポンスのコンテキストを表すオブジェクト
*   db: データベース接続オブジェクト
 */
func DeleteBookmark(c *gin.Context, db *gorm.DB) {
	// ブックマークIDを文字列 ⇨ 数値へ型変換 ⇨ ブックマークの削除
	if converted_bookmark_id, err := strconv.Atoi(c.PostForm("bookmark_id")); err == nil {
		// ブックマークの削除(※ delete_flag 更新による論理削除)
		err := model.DeleteBookmarkData(db, "bookmark_id = ?", converted_bookmark_id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "ブックマークの削除中にエラーが発生しました。"})
			return
		}

		// ブックマークの削除結果をJSONで返す
		c.JSON(http.StatusOK, "削除しました。")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "ブックマークIDの数値への型変換中にエラーが発生しました。"})
		return
	}
}

// シリアライズされたデータをアンシリアライズする関数
func unserialize(data string) []string {
	var bookInfo []string

	// データをカンマとコロンで分割
	fields := strings.Split(data, ", ")

	// 「Title」+「Vol.」で１要素とする
	var title string
	for _, field := range fields {
		if strings.HasPrefix(field, "Title: ") {
			title = strings.TrimPrefix(field, "Title: ")
		} else if strings.HasPrefix(field, "Vol. ") {
			title += ", Vol. " + strings.TrimPrefix(field, "Vol. ")
		}
	}
	bookInfo = append(bookInfo, title)

	// その他データをフィールドごとに取得
	for _, field := range fields {
		if strings.HasPrefix(field, "Subtitle: ") {
			bookInfo = append(bookInfo, strings.TrimPrefix(field, "Subtitle: "))
		} else if strings.HasPrefix(field, "Authors: ") {
			bookInfo = append(bookInfo, strings.TrimSuffix(strings.TrimPrefix(field, "Authors: ["), "]"))
		} else if strings.HasPrefix(field, "Description: ") {
			bookInfo = append(bookInfo, strings.TrimPrefix(field, "Description: "))
		} else if strings.HasPrefix(field, "SmallThumbnail: ") {
			bookInfo = append(bookInfo, strings.TrimPrefix(field, "SmallThumbnail: "))
		} else if strings.HasPrefix(field, "Publisher: ") {
			bookInfo = append(bookInfo, strings.TrimPrefix(field, "Publisher: "))
		} else if strings.HasPrefix(field, "PublishedDate: ") {
			bookInfo = append(bookInfo, strings.TrimPrefix(field, "PublishedDate: "))
		} else if strings.HasPrefix(field, "Saleability: ") {
			bookInfo = append(bookInfo, strings.TrimPrefix(field, "Saleability: "))
		} else if strings.HasPrefix(field, "WebReaderLink: ") {
			bookInfo = append(bookInfo, strings.TrimPrefix(field, "WebReaderLink: "))
		} else if strings.HasPrefix(field, "AccessViewStatus: ") {
			bookInfo = append(bookInfo, strings.TrimPrefix(field, "AccessViewStatus: "))
		}
	}
	return bookInfo
}

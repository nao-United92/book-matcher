package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nao-United92/book-matcher/model"
)

// 書籍データの構造体
type BookInfo struct {
	Id               string   `json:"id"`
	Title            string   `json:"title"`
	Subtitle         string   `json:"subtitle"`
	Authors          []string `json:"authors"`
	Description      string   `json:"description"`
	SmallThumbnail   string   `json:"smallThumbnail"`
	Publisher        string   `json:"publisher"`
	PublishedDate    string   `json:"publishedDate"`
	Saleability      string   `json:"saleability"`
	WebReaderLink    string   `json:"webReaderLink"`
	AccessViewStatus string   `json:"accessViewStatus"`
}

/*
*	書籍検索
* 【リクエスト】
*		c: リクエストとレスポンスのコンテキストを表すオブジェクト
 */
func SearchBooks(c *gin.Context) {
	// クエリパラメータ取得
	q := c.Query("q")                                      // 検索文字列
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) // デフォルトページ番号
	if err != nil {
		page = 1
	}
	orderBy := c.Query("orderBy") // 順序種別

	// 取得データをキーに書籍検索、検索結果の画面表示
	RenderPage(c, q, page, orderBy)
}

/*
*	書籍詳細データの表示
* 【リクエスト】
*		c: リクエストとレスポンスのコンテキストを表すオブジェクト
*		db: データベース接続オブジェクト
 */
func ShowBooksDetailInfo(c *gin.Context, db *gorm.DB) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)
	user_id := session.Get("user_id")

	// クエリパラメータ取得(※ 配列データは、置換および分割によりループ表示できるよう整形)
	id := c.DefaultQuery("id", "")
	authors_origin := c.Query("authors")
	authors_origin = strings.ReplaceAll(authors_origin, "[", "")
	authors_origin = strings.ReplaceAll(authors_origin, "]", "")
	authors := strings.Split(authors_origin, ",")
	accessViewStatus := c.DefaultQuery("accessViewStatus", "")
	var webReaderLink string
	if accessViewStatus == "SAMPLE" {
		webReaderLink = c.DefaultQuery("webReaderLink", "")
	} else {
		webReaderLink = ""
	}
	fromPage := c.DefaultQuery("fromPage", "")

	// ブックマークへ追加済みかどうかのチェック(検索データあり：追加済、検索データなし：未追加)
	conditions := []interface{}{"user_id = ? AND item_id = ?", user_id, id}
	bookmark, err := model.GetBookmarkData(db, conditions, "")
	if err != nil {
		fmt.Errorf("ブックマークデータの検索中にエラーが発生しました")
	}
	isBookmark := false
	if bookmark != nil {
		isBookmark = true
	}

	// 取得データの画面表示
	c.HTML(http.StatusOK, "search_books_detail.html", gin.H{
		"UserId":           user_id,
		"Id":               id,
		"SmallThumbnail":   c.DefaultQuery("smallThumbnail", ""),
		"Title":            c.DefaultQuery("title", ""),
		"Subtitle":         c.DefaultQuery("subtitle", ""),
		"Authors":          authors,
		"Description":      c.DefaultQuery("description", ""),
		"Publisher":        c.DefaultQuery("publisher", ""),
		"PublishedDate":    c.DefaultQuery("publishedDate", ""),
		"Saleability":      c.DefaultQuery("saleability", ""),
		"WebReaderLink":    webReaderLink,
		"AccessViewStatus": accessViewStatus,
		"IsBookmark":       isBookmark,
		"FromPage":         fromPage,
	})
}

/*
*	書籍検索、検索結果の画面表示
* 【リクエスト】
*		c: リクエストとレスポンスのコンテキストを表すオブジェクト
*		q: フォームで入力された検索文字列
*		page: ページ番号
*		orderBy: 順序種別
 */
func RenderPage(c *gin.Context, q string, page int, orderBy string) {
	// 書籍検索
	books, totalPages, err := SearchBooksInfo(c, q, page, orderBy)
	if err != nil {
		c.HTML(http.StatusBadRequest, "search_books.html", gin.H{
			"Error": "書籍データの検索中にエラーが発生しました。",
		})
		return
	}

	// ページネーション用のページ番号スライスの作成
	var pages []int
	for i := 1; i <= totalPages; i++ {
		if page > 10 {
			pages = append(pages, i+(((page/10)%10)*10))
		} else {
			pages = append(pages, i)
		}
	}

	// 書籍検索結果、ページ番号の画面表示
	c.HTML(http.StatusOK, "search_books.html", gin.H{
		"Books":       books,
		"TotalPages":  totalPages,
		"Pages":       pages,
		"CurrentPage": page,
		"Previous":    page - 1,
		"Next":        page + 1,
	})
}

/*
*	書籍詳細データの検索
* 【リクエスト】
*		c: リクエストとレスポンスのコンテキストを表すオブジェクト
*		q: フォームで入力された検索文字列
*		page: ページ番号
*		orderBy: 順序種別
*
* 【レスポンス】
*		[]BookInfo: 書籍検索データ
*		int: ページ総数
*		error: エラー内容
 */
func SearchBooksInfo(c *gin.Context, q string, page int, orderBy string) ([]BookInfo, int, error) {
	// Google Books APIのエンドポイントを指定
	apiEndpoint := "https://www.googleapis.com/books/v1/volumes"

	// APIリクエスト
	url := fmt.Sprintf("%s?q=%s&maxResults=10&startIndex=%d&orderBy=%s", apiEndpoint, url.QueryEscape(q), (page-1)*10, orderBy)
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("Failed to send API request: %v", err)
	}
	defer resp.Body.Close()

	// APIレスポンスの読み込み
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("Failed to read API response: %v", err)
	}

	// APIレスポンスデータの解析、取得
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, 0, fmt.Errorf("Failed to parse API response: %v", err)
	}

	// 画面表示する書籍データの抽出、リスト化
	var books []BookInfo
	if items, ok := data["items"].([]interface{}); ok {
		for _, item := range items {
			var book BookInfo
			if id, ok := item.(map[string]interface{})["id"].(string); ok {
				book.Id = id
			}

			if volumeInfo, ok := item.(map[string]interface{})["volumeInfo"].(map[string]interface{}); ok {
				book.Title = getString(volumeInfo, "title")
				book.Subtitle = getString(volumeInfo, "subtitle")
				book.Authors = getStringSlice(volumeInfo, "authors")
				book.Description = getString(volumeInfo, "description")
				book.Publisher = getString(volumeInfo, "publisher")
				book.PublishedDate = getString(volumeInfo, "publishedDate")

				// サムネイル画像URLを取得
				if imageLinks, ok := volumeInfo["imageLinks"].(map[string]interface{}); ok {
					book.SmallThumbnail = getString(imageLinks, "smallThumbnail")
				}
			}

			if saleInfo, ok := item.(map[string]interface{})["saleInfo"].(map[string]interface{}); ok {
				book.Saleability = getString(saleInfo, "saleability")
			}

			if accessInfo, ok := item.(map[string]interface{})["accessInfo"].(map[string]interface{}); ok {
				book.AccessViewStatus = getString(accessInfo, "accessViewStatus")
				if book.AccessViewStatus == "SAMPLE" {
					book.WebReaderLink = getString(accessInfo, "webReaderLink")
				}
			}

			books = append(books, book)
		}
	}

	totalItems, _ := data["totalItems"].(float64)
	totalPages := int(totalItems)/10 + 1

	return books, totalPages, nil
}

/*
*	マップから指定したキーの文字列値を取得
* 【リクエスト】
*		m: マッピングデータ
*		key: マップデータから取得するキー
*
* 【レスポンス】
*		string: 取得文字列値
 */
func getString(m map[string]interface{}, key string) string {
	if value, ok := m[key].(string); ok {
		return value
	}
	return ""
}

/*
*	マップから指定したキーの文字列のスライスを取得
* 【リクエスト】
*		m: マッピングデータ
*		key: マップデータから取得するキー
*
* 【レスポンス】
*		[]string: 取得文字列のスライス
 */
func getStringSlice(m map[string]interface{}, key string) []string {
	if values, ok := m[key].([]interface{}); ok {
		var result []string
		for _, value := range values {
			if s, ok := value.(string); ok {
				result = append(result, s)
			}
		}
		return result
	}
	return nil
}

package profile

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nao-United92/book-matcher/model"
)

/*
* プロフィールフォーム表示
* 【リクエスト】
*		c: リクエストとレスポンスのコンテキストを表すオブジェクト
*		db: データベース接続オブジェクト
 */
func ShowProfileForm(c *gin.Context, db *gorm.DB) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)

	// セッションから取得したメールアドレスをキーにユーザーデータを検索
	conditions := []interface{}{"address = ? AND delete_flag = 0", session.Get("address")}
	user, err := model.GetUserData(db, conditions, "")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "ユーザーデータの検索中にエラーが発生しました。"})
		return
	}

	// プロフィール画像をセット
	var imagePath string
	if user.Image != "" {
		imagePath = "/static/images/" + user.Image
	} else {
		imagePath = "/static/images/kkrn_icon_user_11.png"
	}

	// プロフィールフォームを表示
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"UserId":       user.User_id,
		"Username":     user.User_name,
		"Address":      user.Address,
		"ImagePath":    imagePath,
		"Introduction": user.Introduction,
	})
}

/*
* プロフィールデータの更新
* 【リクエスト】
*		c: リクエストとレスポンスのコンテキストを表すオブジェクト
*		db: データベース接続オブジェクト
 */
func UpdateProfileInfo(c *gin.Context, db *gorm.DB) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)

	// フォームデータを取得
	user_name := c.PostForm("user_name")
	address := c.PostForm("address")
	introduction := c.PostForm("introduction")

	// 自身を除くユーザーのメールアドレスを検索
	conditions := "address != ? AND delete_flag = 0"
	address_list, err := model.GetUserSpecificData(db, "address", conditions, address)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "ユーザーデータの検索中にエラーが発生しました。"})
		return
	}

	// メールアドレスの存在チェック
	if isAddressExists(address, address_list) {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "既に同じメールアドレスが登録されています。"})
		return
	}

	// メールアドレスのフォーマットチェック
	if !isValidAddress(address) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "メールアドレスのフォーマットが不正です。"})
		return
	}

	// ユーザーIDの型変換(文字列 ⇨ 数値)
	if converted_user_id, err := strconv.Atoi(c.PostForm("user_id")); err == nil {
		// プロフィールデータの更新
		updates := map[string]interface{}{
			"User_name":    user_name,
			"Address":      address,
			"Introduction": introduction,
			"Update_date":  time.Now(),
		}
		updatedUser, err := model.UpdateUserData(db, converted_user_id, updates)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "プロフィールデータの更新中にエラーが発生しました。"})
			return
		}

		// セッションへ新しいユーザーデータを保存
		session.Set("user_id", converted_user_id)
		session.Set("address", address)
		session.Save()

		// プロフィールデータの更新結果をJSONで返す
		c.JSON(http.StatusOK, updatedUser)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "ユーザーIDの数値への型変換中にエラーが発生しました。"})
		return
	}
}

/*
*	アカウントの削除
* 【リクエスト】
*		c: リクエストとレスポンスのコンテキストを表すオブジェクト
*		db: データベース接続オブジェクト
 */
func DeleteAccount(c *gin.Context, db *gorm.DB) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)

	// ユーザーIDを文字列 ⇨ 数値へ型変換 ⇨ アカウントの削除
	if converted_user_id, err := strconv.Atoi(c.PostForm("user_id")); err == nil {
		// アカウントの削除(※ delete_flag 更新による論理削除)
		updates := map[string]interface{}{
			"Delete_date": time.Now(),
			"Delete_flag": 1,
		}
		updatedUser, err := model.UpdateUserData(db, converted_user_id, updates)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "アカウントの削除中にエラーが発生しました。"})
			return
		}

		// セッションデータを削除
		session.Clear()

		// セッションの破棄を確定
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// アカウントの削除結果をJSONで返す
		c.JSON(http.StatusOK, updatedUser)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "ユーザーIDの数値への型変換中にエラーが発生しました。"})
		return
	}
}

/*
*	メールアドレスの存在チェック
* 【リクエスト】
*		input_address: フォームで入力されたメールアドレス
*		address_list: DBから取得したチェック対象のメールアドレス
*
* 【レスポンス】
*		bool: true / false
 */
func isAddressExists(input_address string, address_list []string) bool {
	for _, address := range address_list {
		if address == input_address {
			return true
		}
	}
	return false
}

/*
*	メールアドレスのフォーマットチェック
* 【リクエスト】
*		address: チェック対象のメールアドレス
*
* 【レスポンス】
*		bool: true / false
 */
func isValidAddress(address string) bool {
	// RFC 5322に基づく正規表現パターンを定義
	checkPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 正規表現パターンに一致するかどうかを確認
	match, err := regexp.MatchString(checkPattern, address)
	if err != nil {
		return false
	}
	return match
}

package login

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nao-United92/book-matcher/model"
	"golang.org/x/crypto/bcrypt"
)

/*
* ログインフォーム表示
* 【リクエスト】
*   c: リクエストとレスポンスのコンテキストを表すオブジェクト
 */
func ShowLoginForm(c *gin.Context) {
	// ログインフォームを表示
	c.HTML(http.StatusOK, "login.html", nil)
}

/*
* ログイン
* 【リクエスト】
*   c: リクエストとレスポンスのコンテキストを表すオブジェクト
*   db: データベース接続オブジェクト
 */
func Login(c *gin.Context, db *gorm.DB) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)

	// フォームデータを取得
	address := c.PostForm("address")
	password := c.PostForm("password")

	// ログインフォームの入力内容と一致するユーザーデータの存在チェック
	conditions := []interface{}{"address = ? AND delete_flag = 0", address}
	user, err := model.GetUserData(db, conditions, "")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "メールアドレスに該当するユーザーは存在しません。"})
		return
	}

	// ハッシュ化されたパスワードと元のパスワードを比較
	err = CompareHashAndPassword(user.Password, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "パスワードが不一致です。"})
		return
	}

	// セッションへユーザーデータを保存
	session.Set("user_id", user.User_id)
	session.Set("address", user.Address)
	session.Save()

	// メニューページへリダイレクト
	c.Redirect(http.StatusSeeOther, "/")
}

/*
* 暗号(Hash)化および入力されたパスワードの比較
* 【リクエスト】
*   hashedPassword: ハッシュ化されたパスワード
*   password: ハッシュ化前のパスワード
*
* 【レスポンス】
*   error: エラー内容
 */
func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

/*
* ログアウト
* 【リクエスト】
*   c: リクエストとレスポンスのコンテキストを表すオブジェクト
 */
func Logout(c *gin.Context) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)

	// セッションデータを削除
	session.Clear()

	// セッションの破棄を確定
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

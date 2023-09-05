package signup

import (
	"net/http"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nao-United92/book-matcher/model"
	"golang.org/x/crypto/bcrypt"
)

/*
* 新規登録フォーム表示
* 【リクエスト】
*   c: リクエストとレスポンスのコンテキストを表すオブジェクト
 */
func ShowSignupForm(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

/*
* ユーザーデータの新規登録
* 【リクエスト】
*   c: リクエストとレスポンスのコンテキストを表すオブジェクト
*   db: データベース接続オブジェクト
 */
func Signup(c *gin.Context, db *gorm.DB) {
	// セッションのデフォルトストアからセッションオブジェクトを取得
	session := sessions.Default(c)

	// フォームデータを取得
	user_name := c.PostForm("username")
	address := c.PostForm("address")
	password := c.PostForm("password")

	// チェック対象のメールアドレスを取得
	var address_list []string
	if err := db.Model(&model.User{}).Pluck("address", &address_list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "ユーザーデータの検索中にエラーが発生しました。"})
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

	// パスワードが半角英数・記号：8～16文字を満たすか否かのチェック
	valid := ValidatePassword(password)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "パスワードのフォーマットが不正です。"})
		return
	}

	// パスワードをハッシュ化
	hashedPassword, err := PasswordHash(password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "パスワードの暗号化中にエラーが発生しました。"})
		return
	}

	// ユーザーデータの新規登録
	user, err := model.CreateUserData(db, user_name, address, hashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "ユーザーデータの登録中にエラーが発生しました。"})
		return
	}

	// ユーザーデータの新規登録が成功した場合
	if user != nil {
		// セッションへ登録したユーザーデータを保存
		session.Set("user_id", user.User_id)
		session.Set("address", address)
		session.Save()

		// メニューページへリダイレクト
		c.Redirect(http.StatusSeeOther, "/")
	}
}

/*
* メールアドレスの存在チェック
* 【リクエスト】
*   input_address: フォームで入力されたメールアドレス
*   address_list: DBから取得したチェック対象のメールアドレス
*
* 【レスポンス】
*   bool: true / false
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
* メールアドレスのフォーマットチェック
* 【リクエスト】
*   address: チェック対象のメールアドレス
*
* 【レスポンス】
*   bool: true / false
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

/*
* パスワードが半角英数・記号：8～16文字を満たすか否かのチェック
* 【リクエスト】
*   password: メールアドレス
*
* 【レスポンス】
*   bool: チェック結果
 */
func ValidatePassword(password string) bool {
	// 正規表現でパスワードが半角英数・記号：8～16文字か否かのチェック
	return regexp.MustCompile(`^[!-~]{8,16}$`).MatchString(password)
}

/*
* パスワードの暗号(Hash)化
* 【リクエスト】
*   password: ハッシュ化対象のメールアドレス
*
* 【レスポンス】
*   string: ハッシュ化されたパスワード
*   error: エラー内容
 */
func PasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

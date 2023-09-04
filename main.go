package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/nao-United92/book-matcher/controller/bookmark"
	"github.com/nao-United92/book-matcher/controller/login"
	"github.com/nao-United92/book-matcher/controller/password"
	"github.com/nao-United92/book-matcher/controller/profile"
	"github.com/nao-United92/book-matcher/controller/search"
	"github.com/nao-United92/book-matcher/controller/signup"
)

// クッキーストアの作成、セキュリティを強化したクッキーの管理が可能
// セキュアクッキー：情報を暗号化・保護し、改ざんを防ぐために署名を付けるなどのセキュリティ対策が施されたクッキー
var store = cookie.NewStore([]byte("secret"))

func main() {
	// デフォルトのGinエンジンを生成
	r := gin.Default()

	// 事前にテンプレートをレンダリング
	r.LoadHTMLGlob("view/*")

	// 静的ファイルのディレクトリを指定
	r.Static("/static", "./static")

	// セッションの設定
	r.Use(sessions.Sessions("mysession", store))

	// 設定ファイルから環境変数の情報を取得
	envLoad()

	// 初期画面表示のハンドラー
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu.html", gin.H{})
	})

	// メニュー画面 ログインボタンクリック時のハンドラー
	r.GET("/login", login.ShowLoginForm)

	// ログインフォーム ログインボタンクリック時のハンドラー
	r.POST("/login", func(c *gin.Context) {
		// DB接続
		db := dbConnect()

		// ログイン処理
		login.Login(c, db)

		// DB接続切断
		defer db.Close()
	})

	// メニュー画面 ログアウトボタンクリック時のハンドラー
	r.PUT("/login", func(c *gin.Context) {
		// ログアウト処理
		login.Logout(c)
	})

	// ログインフォーム パスワード再設定リンククリック時のハンドラー
	r.GET("/password", password.ShowPasswordResetForm)

	// メニュー画面 新規登録ボタンクリック時のハンドラー
	r.GET("/signup", signup.ShowSignupForm)

	// 新規登録フォーム 登録ボタンクリック時のハンドラー
	r.POST("/signup", func(c *gin.Context) {
		// DB接続
		db := dbConnect()

		// 新規登録処理
		signup.Signup(c, db)

		// DB接続切断
		defer db.Close()
	})

	// 書籍検索ボタン、ページネーションボタンクリック時のハンドラー
	r.GET("/search", search.SearchBooks)

	// 書籍画像、書籍タイトルのリンククリック時のハンドラー
	r.GET("/search_detail", func(c *gin.Context) {
		// DB接続
		db := dbConnect()

		// ブックマークフォーム表示処理
		search.ShowBooksDetailInfo(c, db)

		// DB接続切断
		defer db.Close()
	})

	// マイページまたはプロフィールメニュークリック時のハンドラー
	r.GET("/profile", func(c *gin.Context) {
		// DB接続
		db := dbConnect()

		// プロフィールフォーム表示処理
		profile.ShowProfileForm(c, db)

		// DB接続切断
		defer db.Close()
	})

	// プロフィール保存ボタン / アカウント削除ボタンクリック時のハンドラー
	r.PUT("/profile", func(c *gin.Context) {
		// DB接続
		db := dbConnect()

		// 処理種別によって呼び出すメソッド制御
		switch c.PostForm("process_type") {
		case "UpdateProfileInfo":
			// プロフィールデータの更新処理
			profile.UpdateProfileInfo(c, db)
		case "DeleteAccount":
			// アカウントの削除処理
			profile.DeleteAccount(c, db)
		}

		// DB接続切断
		defer db.Close()
	})

	// ブックマークメニュークリック時のハンドラー
	r.GET("/bookmark", func(c *gin.Context) {
		// DB接続
		db := dbConnect()

		// ブックマークフォーム表示処理
		bookmark.ShowBookmarkForm(c, db)

		// DB接続切断
		defer db.Close()
	})

	// 書籍詳細フォーム ブックマーク追加ボタンクリック時のハンドラー
	r.POST("/bookmark", func(c *gin.Context) {
		// DB接続
		db := dbConnect()

		// ブックマークへの書籍データ登録処理
		bookmark.AddBookmark(c, db)

		// DB接続切断
		defer db.Close()
	})

	// ブックマークフォーム ブックマーク削除ボタンクリック時のハンドラー
	r.PUT("/bookmark", func(c *gin.Context) {
		// DB接続
		db := dbConnect()

		// ブックマークの削除処理
		bookmark.DeleteBookmark(c, db)

		// DB接続切断
		defer db.Close()
	})

	// サーバーを起動
	r.Run(":8080")
}

// .env ファイル読み込み
func envLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env target")
	}
}

// DB 接続
func dbConnect() *gorm.DB {
	// 環境変数から接続文字列を作成
	connect := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))

	// 環境変数からDBMSの種類を取得
	db, err := gorm.Open(os.Getenv("DBMS"), connect)
	if err != nil {
		panic(err.Error())
	}
	return db
}

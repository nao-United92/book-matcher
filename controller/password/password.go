package password

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
* パスワード再設定フォーム表示
* 【リクエスト】
*		c: リクエストとレスポンスのコンテキストを表すオブジェクト
 */
func ShowPasswordResetForm(c *gin.Context) {
	// パスワード再設定フォーム表示
	c.HTML(http.StatusOK, "password_reset.html", nil)
}

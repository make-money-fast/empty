package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/make-money-fast/empty/templates"
)

var (
	redirectString = ``
)

func Redirect(ctx *gin.Context) {
	url := ctx.Query("url")
	if url == "" {
		ctx.Redirect(302, "/")
		return
	}
	templates.Html(ctx, "redirect", gin.H{
		"url": url,
	})
}

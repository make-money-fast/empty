package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/make-money-fast/empty/templates"
)

func Index(ctx *gin.Context) {
	templates.Html(ctx, "index", gin.H{
		"Foo": "bar",
	})
}

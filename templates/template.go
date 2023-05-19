package templates

import (
	"github.com/gin-gonic/gin"
	"github.com/make-money-fast/empty/config"
)

func Html(ctx *gin.Context, name string, data gin.H) {
	if data == nil {
		data = gin.H{}
	}
	conf := config.Load()
	data["seo"] = conf.Seo

	ctx.HTML(200, name, data)
}

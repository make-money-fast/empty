package ext

import "github.com/gin-gonic/gin"

const (
	defaultLang = "en"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang, _ := ctx.Cookie("lang")

		if lang == "" {
			q := ctx.Request.URL.Query().Get("lang")
			if q != "" {
				lang = ctx.Query("lang")
			}
		}

		if lang == "" {
			lang = defaultLang
		}

		ctx.Set("lang", lang)
		ctx.Next()
	}
}

func GetLanguage(ctx *gin.Context) string {
	l, ok := ctx.Get("lang")
	if ok {
		return l.(string)
	}
	return defaultLang
}

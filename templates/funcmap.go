package templates

import (
	"github.com/gin-gonic/gin"
	"github.com/make-money-fast/empty/config"
	"github.com/make-money-fast/empty/pkg/ext"
	"html/template"
)

var FuncMap = template.FuncMap{
	"asset": asset,
	"css":   css,
	"js":    js,
	"html":  html,
	"t":     t,
}

func asset(data string) string {
	c := config.Load()
	return c.Seo.CDN + "/static/" + data
}

func css(data string) string {
	c := config.Load()
	return c.Seo.CDN + "/static/css/" + data
}

func js(data string) string {
	c := config.Load()
	return c.Seo.CDN + "/static/js" + data
}

func html(s string) template.HTML {
	return template.HTML(s)
}

func t(s string) string {
	return s
}

type FuncMapService struct {
	ctx  *gin.Context
	lang *Language
}

func NewFuncMapService(ctx *gin.Context) *FuncMapService {
	return &FuncMapService{
		ctx:  ctx,
		lang: NewLanguage(),
	}
}

func (s *FuncMapService) FuncMap() template.FuncMap {
	return s.merge(FuncMap, s.funcMap())
}

func (s *FuncMapService) funcMap() template.FuncMap {
	return template.FuncMap{
		"t": s.t,
	}
}

func (s *FuncMapService) merge(fms ...template.FuncMap) template.FuncMap {
	var base = make(template.FuncMap)
	for _, fm := range fms {
		for k, v := range fm {
			base[k] = v
		}
	}
	return base
}

func (s *FuncMapService) t(key string) string {
	var lm langMap
	// 判断环境
	if gin.Mode() == gin.ReleaseMode {
		lm = s.lang.get()
	} else {
		lm, _ = s.lang.DynamicLoad()
	}

	lang := ext.GetLanguage(s.ctx)

	return lm.t(lang, key)
}

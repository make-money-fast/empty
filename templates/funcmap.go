package templates

import (
	"github.com/make-money-fast/empty/config"
	"html/template"
)

var FuncMap = template.FuncMap{
	"asset": asset,
	"css":   css,
	"js":    js,
	"html":  html,
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

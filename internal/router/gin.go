package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/make-money-fast/empty/config"
	ctrl2 "github.com/make-money-fast/empty/internal/ctrl"
	"github.com/make-money-fast/empty/pkg/ext"
	"github.com/make-money-fast/empty/templates"
	"github.com/thinkerou/favicon"
	"log"
)

func Start(conf *config.Config) error {
	g := gin.Default()

	g.SetFuncMap(templates.FuncMap)
	g.LoadHTMLGlob("templates/*.gohtml")
	g.Static("static", "./static")
	g.Use(favicon.New("./static/favicon.ico"))

	ext.InitLogger()
	g.Use(ext.Logger())

	for route, data := range conf.Seo.Links {
		g.GET(route, func(ctx *gin.Context) {
			ctx.Header("Content-Type", "text/plain")
			ctx.String(200, data)
		})
	}
	// 重定向页面.
	g.GET("/redirect", ctrl2.Redirect)
	g.GET("/", ctrl2.Index)

	addr := fmt.Sprintf("%s:%d", conf.Http.Addr, conf.Http.Port)
	log.Println("http://" + addr)
	return g.Run(addr)
}

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
	"net/http"
	"strings"
)

func Start(conf *config.Config) error {
	g := gin.Default()

	templates.NewLanguage().Load(config.Load().Http.LangFiles...)
	g.HTMLRender = templates.NewHtmlRender(config.Load().Http.TemplatePath)
	g.Static("static", "./static")
	g.Use(favicon.New("./static/favicon.ico"))
	g.Use(ext.I18nMiddleware())

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

	g.GET("/login.html", ctrl2.Login)
	g.GET("/register.html", ctrl2.Register)

	addr := fmt.Sprintf("%s:%d", conf.Http.Addr, conf.Http.Port)
	log.Println("http://" + addr)

	s := &myServer{
		g: g,
	}
	return http.ListenAndServe(addr, s)
}

type myServer struct {
	g *gin.Engine
}

func (s *myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 处理path的语言.
	path := r.URL.Path
	langs := templates.NewLanguage()
	for _, l := range langs.Lang() {
		if strings.HasSuffix(path, "/"+l) {
			path = strings.TrimSuffix(path, "/"+l)
			if path == "" {
				path = "/"
			}
			r.URL.Path = path
			q := r.URL.Query()
			q.Set("lang", l)
			r.URL.RawQuery = q.Encode()
			break
		}
	}
	s.g.ServeHTTP(w, r)
}

package templates

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/make-money-fast/empty/config"
	"html/template"
	"sync"
)

func Html(ctx *gin.Context, name string, data gin.H) {
	if data == nil {
		data = gin.H{}
	}
	conf := config.Load()
	data["seo"] = conf.Seo
	data["_ctx"] = ctx

	ctx.HTML(200, name, data)
}

type Render struct {
	path     string
	o        sync.Once
	template *template.Template
}

func NewHtmlRender(path string) *Render {
	return &Render{
		path: path,
	}
}

func (r *Render) Instance(s string, a any) render.Render {
	var (
		tpl *template.Template
		err error
	)
	fm := NewFuncMapService(a.(gin.H)["_ctx"].(*gin.Context))
	if gin.Mode() == gin.ReleaseMode {
		r.o.Do(func() {
			tpl, err = template.New("_root_").Funcs(fm.FuncMap()).ParseGlob(r.path)
			if err != nil {
				panic(err)
			}
			r.template = tpl
		})
	} else {
		tpl, err = template.New("_root_").Funcs(fm.FuncMap()).ParseGlob(r.path)
		if err != nil {
			panic(err)
		}
	}

	return &render.HTML{
		Template: tpl,
		Name:     s,
		Data:     a,
	}
}

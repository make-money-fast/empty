package templates

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

type langMap map[string]map[string]string

func (l langMap) t(lang string, key string) string {
	ll, ok := l[lang]
	if ok {
		return ll[key]
	}
	return ""
}

var (
	langOnce sync.Once
	_lang    *Language
)

type Language struct {
	lang langMap

	files []string

	langs []string
}

func NewLanguage() *Language {
	langOnce.Do(func() {
		_lang = &Language{}
	})
	return _lang
}

func (l *Language) Load(files ...string) {
	if len(l.files) == 0 {
		l.files = files
	}
	res, langs := l.DynamicLoad()
	if len(l.langs) == 0 {
		l.langs = langs
	}
	l.lang = res
}

func (l *Language) DynamicLoad() (langMap, []string) {
	var (
		res   = make(langMap)
		langs []string
	)
	for _, file := range l.files {
		base := filepath.Base(file)
		langName := strings.Split(base, ".")[0]
		var data = make(map[string]string)
		bs, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(bs, &data)
		if err != nil {
			panic(err)
		}
		lName := strings.ToLower(langName)
		langs = append(langs, lName)
		res[lName] = data
	}
	return res, langs
}

func (l *Language) get() langMap {
	return l.lang
}

func (l *Language) Lang() []string {
	return l.langs
}

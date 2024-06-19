package config

type Config struct {
	Http *Http `json:"http"`
	DB   *Db   `json:"db"`
	Seo  *Seo  `json:"seo"`
}

type Http struct {
	Addr         string   `json:"addr"`
	Port         int      `json:"port"`
	TemplatePath string   `json:"templatePath"`
	LangFiles    []string `json:"langFiles"`
}

type FriendLink struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Direct bool   `json:"direct"` // 是否 直接打开， 否则通过中间页.
}

type Db struct {
	Dsn    string `json:"dsn"`    // 连接
	Driver string `json:"driver"` // 默认 sqlite3
}

type Seo struct {
	FullHost     string            `json:"full_host"`     // 全路径域名  带 http/https
	Host         string            `json:"host"`          // 域名: 不带 http/https
	CDN          string            `json:"cdn"`           // 静态文件域名
	HeaderScript string            `json:"header_script"` // 头部脚本
	FooterScript string            `json:"footer_script"` // 底部脚本
	Title        string            `json:"title"`         // 标题
	KeyWords     string            `json:"key_words"`     // 关键词
	Description  string            `json:"description"`   // 描述
	Copyright    string            `json:"copyright"`     // 版权描述
	Links        map[string]string `json:"links"`         // 自定义路由
	FriendLink   []FriendLink      `json:"friend_link"`   // 友情链接
}

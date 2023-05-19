package config

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"sync"
)

var (
	o     sync.Once
	_conf *Config
)

func Load() *Config {
	o.Do(func() {
		data, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Fatal("未找到配置", err)
		}
		var conf Config
		err = yaml.Unmarshal(data, &conf)
		if err != nil {
			log.Fatal("配置解析失败", err)
		}
		fmt.Println(string(data))
		_conf = &conf
	})
	return _conf
}

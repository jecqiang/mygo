package mygo

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

//程序配置
type Config struct {
	HttpPort int   `yaml:"httpPort"`
	Mysql    Mysql `yaml:"mysql"`
}

//mysql数据库配置
type Mysql struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Database  string `yaml:"database"`
	Character string `yaml:"character"`
}

//单例模式
var (
	G_config *Config
)

//初始化配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		config  Config
	)
	content, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return
	}
	G_config = &config
	return
}

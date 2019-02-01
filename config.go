package mygo

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

//程序配置
type Config struct {
	Serve Serve `yaml:"serve"`
	Mysql Mysql `yaml:"mysql"`
	Log   Log   `yaml:"log"`
	Smtp  Smtp  `yaml:"smtp"`
}

type Smtp struct {
	Enable int    `yaml:"enable"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
}

//服务配置
type Serve struct {
	Port          int `yaml:"port"`
	Read_timeout  int `yaml:"read_timeout"`
	Write_timeout int `yaml:"write_timeout"`
	Idle_timeout  int `yaml:"idle_timeout"`
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

//日志配置
type Log struct {
	Path  string `yaml:"path"`
	Level int    `yaml:"level"`
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

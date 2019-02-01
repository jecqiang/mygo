package main

import (
	"fmt"
	"flag"
	"runtime"
	"os"
	"github.com/jecqiang/mygo"
	"log"
)

var (
	help       bool   //启动帮助
	configFile string //配置文件路径
	version    bool   //是否查看版本
	logFile    string
)

//初始化
func init() {
	flag.BoolVar(&help, "h", false, "This help")
	flag.StringVar(&configFile, "c", "./config/mygo.yaml", "Set configuration file")
	flag.BoolVar(&version, "v", false, "Version number")

	flag.Parse()

	//初始化线程数量
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func outputInfo(tag string, value interface{}) {
	fmt.Printf("%-18s    %v\n", tag+":", value)
}

func run()  {
	var (
		err error
		f *os.File
	)
	outputInfo("Service", "mygo")
	outputInfo("Version", mygo.VERSION)
	//初始化配置
	err = mygo.InitConfig(configFile)
	if err != nil {
		panic(err)
	}
	outputInfo("Config Loaded", configFile)

	//设置logger
	if mygo.G_config.Log.Path != "" && mygo.G_config.Log.Path != "stdout" {
		f, err = os.OpenFile(mygo.G_config.Log.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}
	outputInfo("Log", mygo.G_config.Log.Path)

	mg := mygo.NewMygo()
	mg.RegisterMail()
	mg.RegisterDb()
	defer mygo.Db.Close()

	fmt.Println("Start Running...")

}

func main() {
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if version {
		fmt.Printf("mygo/%s\n", mygo.VERSION)
		os.Exit(0)
	}
	run()
}
package main

import (
	"fmt"
	"flag"
	"runtime"
	"os"
	"github.com/jecqiang/mygo"
	"github.com/dreamans/syncd"
)

var (
	help       bool   //启动帮助
	configFile string //配置文件路径
	version    bool   //是否查看版本
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

func welcome() {
	fmt.Println(" _ __ ___   _   _   __ _   ___")
	fmt.Println("| '_ ` _ \\ | | | | / _` | / _ \\")
	fmt.Println("| | | | | || |_| || (_| || (_) |")
	fmt.Println("|_| |_| |_| \\__, | \\__, | \\___/")
	fmt.Println("			 __/ |  __/ |")
	fmt.Println("			|___/  |___/")
	fmt.Println("")
	outputInfo("Service", "mygo")
	outputInfo("Version", syncd.VERSION)
}

func outputInfo(tag string, value interface{}) {
	fmt.Printf("%-18s    %v\n", tag + ":", value)
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
	var(
		err error
	)
	welcome()

	//初始化配置
	err = mygo.InitConfig(configFile)
	if err != nil {
		goto ERR
	}

	fmt.Println("Start Running...")

	return
ERR:
	fmt.Println(err)
}

// +build dev

package main

import (
	"flag"
	"os"
	"fmt"
	"os/exec"
	"log"
)

var (
	goos   string
	goarch string
	goarm  string
	race   bool
)

//跨平台交叉编译
//CGO_ENABLED = 0 表示设置CGO工具不可用
//GOOS 程序构建环境的目标操作系统
//GOARCH 表示程序构建环境的目标计算架构
//GOARM 表示使用的浮点运算协处理器版本号，只对arm平台有用，可选值有5，6，7
//在windows 32位系统下运行 : CGO_ENABLED=0 GOOS=windows GOARCH=386 go build test.go
//在windows 64位系统下运行 : CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go
//在OS X 32位系统下运行 : CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build test.go
//在OS X 64位系统下运行 : CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build test.go
//在Linux 32位系统下运行 : CGO_ENABLED=0 GOOS=linux GOARCH=386 go build test.go
//在Linux 64位系统下运行 : CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build test.go
func init() {
	flag.StringVar(&goos, "goos", "", "GOOS for which to build")
	flag.StringVar(&goarch, "goarch", "", "GOARCH for which to build")
	flag.StringVar(&goarm, "goarm", "", "GOARM for which to build")
	//是否同时检测数据竞争状态
	flag.BoolVar(&race, "race", false, "Enable detect data race status")
}

func main() {
	var (
		gopath string
		args   []string
		env    []string
		cmd    *exec.Cmd
		err    error
	)
	flag.Parse()
	gopath = os.Getenv("GOPATH")
	args = []string{
		"build",
		"-asmflags", fmt.Sprintf("-trimpath=%s", gopath),
		"-gcflags", fmt.Sprintf("-trimpath=%s", gopath),
	}
	if race {
		args = append(args, "-race")
	}

	env = os.Environ()
	env = append(env, "GOOS="+goos, "GOARCH="+goarch, "GOARM="+goarm)
	if !race {
		env = append(env, "CGO_ENABLED=0")
	}

	cmd = exec.Command("go", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = env
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

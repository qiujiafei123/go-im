package main

import (
	"flag"
	"fmt"
	"go-im/web"
	"os"
	"os/signal"
	"syscall"
)

var (
	 quit = make(chan os.Signal, 1)
	 service string
)

func main() {
	// 创建一个接收 Linux 中断信号的管道
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)
	// 接收外部参数
	flag.StringVar(&service, "service", "", "使用 service 启动各个模块")
	flag.Parse()

	fmt.Printf("Start run %s service \n", service)

	switch service {
	case "web":
		fmt.Println("web start")
		web.New().Run()
	default:
		fmt.Println("没有对应的服务")
		quit <- syscall.SIGQUIT
	}

	<- quit
	fmt.Println("Service exit")
}
package main

import (
	"flag"
	"fmt"
	"github.com/benmanns/goworker"
	"os"
)

var logfile string //日志文件

func main(){
	fmt.Println("ansyWork start success")
	var pidFile string
	flag.StringVar( &pidFile, "pid", "/var/log/ansyTask.pid", "pid file" ) //接收-pid参数 为pid保存的文件路径
	flag.StringVar( &logfile, "log", "/var/log/ansyTask/ansyTask.log", "log file" ) //接收-log参数 为log日志保存的文件路径
	if !flag.Parsed() {
		flag.Parse()
	}
	//获取当前进程的pid,将pid保存到pidFile
	pid := os.Getpid()
	err := WritePid( pidFile, pid )
	if err != nil {
		fmt.Println("Error:", err )
		os.Exit(1)
	}
	//运行gowoker
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
	goworker.Close()
}
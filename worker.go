package main

import (
	"flag"
	"fmt"
	"github.com/benmanns/goworker"
	"os"
)

var logfile string //日志文件

var afterChan chan int //定时任务缓冲队列，防止同一时间并发太高

var maxAfterChanNum int //定时任务缓冲队列最大长度

func main(){
	var pidFile string
	flag.StringVar( &pidFile, "pid", "/var/log/ansyTask.pid", "pid file" ) //接收-pid参数 为pid保存的文件路径
	flag.StringVar( &logfile, "log", "/var/log/ansyTask/ansyTask.log", "log file" ) //接收-log参数 为log日志保存的文件路径
	flag.IntVar( &maxAfterChanNum, "maxAfterChanNum", 15, "max after time task num in chan" ) //定时任务的缓存队列长度
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
	fmt.Println("ansyWork start success")

	afterChan = make(chan int, maxAfterChanNum)
	fmt.Println(fmt.Sprintf("afterChan 长度：%d", cap(afterChan)))
	//运行gowoker
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
	goworker.Close()
}
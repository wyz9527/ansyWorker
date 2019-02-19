package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	myInfo *log.Logger
	myWarning *log.Logger
	myError * log.Logger
)

func setMylog(){
	//设置日志的格式
	log.SetPrefix("【ansyTask】") //设置日志的前缀
	log.SetFlags(log.Ldate|log.Lshortfile) //显示日期，文件名和行号
	t := time.Now()
	y,m,d := t.Date()
	ymd := fmt.Sprintf("%d%d%d", y,m,d)
	errFile,err:=os.OpenFile( logfile+"."+ymd, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}

	myInfo = log.New(io.MultiWriter(os.Stdout,errFile),"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	myWarning = log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	myError = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)

}

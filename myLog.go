package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)


func init()  {
	//设置日志的格式
	log.SetPrefix("【ansyTask】") //设置日志的前缀
	log.SetFlags(log.Ldate|log.Lshortfile) //显示日期，文件名和行号
}

//记录信息日志
func infoLog( msg ...interface{} ){
	t := time.Now()
	y,m,d := t.Date()
	ymd := fmt.Sprintf("%d%d%d", y,m,d)
	errFile,err:=os.OpenFile( logfile+"."+ymd, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	defer errFile.Close()
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}

	myInfo := log.New(io.MultiWriter(os.Stdout,errFile),"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	myInfo.Println( msg... )
}

//记录错误日志
func errorLog( msg ...interface{} ){
	t := time.Now()
	y,m,d := t.Date()
	ymd := fmt.Sprintf("%d%d%d", y,m,d)
	errFile,err:=os.OpenFile( logfile+"."+ymd, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	defer errFile.Close()
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}

	myError := log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)
	myError.Println( msg... )

}

//记录警告日志
func warnLog(msg ...interface{})  {
	t := time.Now()
	y,m,d := t.Date()
	ymd := fmt.Sprintf("%d%d%d", y,m,d)
	errFile,err:=os.OpenFile( logfile+"."+ymd, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	defer errFile.Close()
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}

	myWarn := log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)
	myWarn.Println( msg... )
}
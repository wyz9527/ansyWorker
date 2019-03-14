package main

import (
	"encoding/json"
	"fmt"
	"github.com/benmanns/goworker"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

//在初始化里面进行业务注册
func init() {
	goworker.Register("AnsyTask", ansyTaskWorker)
}


func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			os.Stdout.Write(d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
	// never reached
	panic(true)
	return nil, nil
}


//写文件
func WritePid(name string, pid int) error {
	return  ioutil.WriteFile(name, []byte(fmt.Sprintln(pid)),0666)
}

//异步任务 执行的方法
func ansyTaskWorker(queue string, args ...interface{}) error {
	var err error
	var cmd  *exec.Cmd
	//打印日志
	go infoLog(fmt.Sprintf( "From Redis Key : %s; Args: %v\n", queue, args[0] ))
	//myInfo.Printf("From Redis Key : %s; Args: %v\n", queue, args[0])
	//解析数据
	params := make(map[string]string)
	data, _ := json.Marshal(args[0])
	json.Unmarshal(data, &params)

	taskType := params["type"] //脚本类型 php|shell
	dir := params["dir"] //执行文件所在目录
	mainFile := params["mainFile"] //执行的文件

	var cmdArgs string  //参数
	cmdArgs = ""
	if _, ok := params["cmdArgs"]; ok{
		cmdArgs = params["cmdArgs"] //参数
	}
	switch taskType {
		case "php":
			phpbin := params["phpbin"] //php命令文件
			action := params["action"] //控制器-方法-动作
			cmd = exec.Command( phpbin, mainFile, action, cmdArgs )
		case "shell":
			//执行shell脚本
			cmd = exec.Command( "/bin/bash", mainFile, cmdArgs )
	}
	cmd.Dir = dir
	var stdout,stderr []byte
	var errStdout, errStderr error

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	cmd.Start()

	go func() {
		stdout,errStdout = copyAndCapture(os.Stdout, stdoutIn)
	}()

	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		go errorLog( err )
		//myError.Println( err )
	}

	if errStderr != nil || errStdout != nil {
		go errorLog(errStdout, errStderr)
		//myError.Println(errStdout, errStderr)
	}

	outStr, errStr := string(stdout), string(stderr)
	//myInfo.Printf("\nout:\n%s\n", outStr)
	infoLog( fmt.Sprintf( "\nout:\n%s\n", outStr ) )
	if errStr != "" {
		errorLog( fmt.Sprintf( "\nerr:\n%s\n", errStr ))
		//myError.Printf("\nerr:\n%s\n", errStr)
	}
	return nil
}
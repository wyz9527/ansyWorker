
# ansyWorker

goworker+python+redis 实现一个异步任务处理

## 配置
### config.json配置文件

```{```
  ```"uri" : "redis://localhost:6379/",  //redis信息```
  ```"queues": "900sui:ansyTaskQueue", //redis队列名称```
  ```"connections" : 10, //redis最大连接数```
  ```"concurrency" : 5, //任务并行数```
  ```"namespace"   : "kcResque:", //redis队列命名空间```
  ```"interval"    : 3, //时间间隔 单位 秒```
  ```"pid"         : "/var/log/ansyTask/dange/pid/ansyTask.pid", //pid存放文件```
  ```"log"         : "/var/log/ansyTask/dange/logs/dange.log" //日志文件```
```}```

## 编译go程序
`go build`

## awctls.py设置权限
`chmod +x awctls.py`

##启动
`./awctls.py start`

##停止
`./awctls.py stop`

##重启
`./awctls.py restart`


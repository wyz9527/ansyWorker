
# ansyWorker

goworker+python+redis 实现一个异步任务处理

## 配置
### config.json配置文件

```javascript
{
  "uri" : "redis://localhost:6379/",  //redis信息
  "queues": "900sui:ansyTaskQueue", //redis队列名称
  "connections" : 10, //redis最大连接数
  "concurrency" : 5, //任务并行数
  "namespace"   : "kcResque:", //redis队列命名空间
  "interval"    : 3, //时间间隔 单位 秒
  "pid"         : "/var/log/ansyTask/dange/pid/ansyTask.pid", //pid存放文件
  "log"         : "/var/log/ansyTask/dange/logs/dange.log" //日志文件
}
```

## 编译go程序
```shell
go build
```

## awctls.py设置权限
```shell
chmod +x awctls.py
```

## 启动
```shell
./awctls.py start
```

## 停止
```shell
./awctls.py stop
```

## 重启
```shell
./awctls.py restart
```

## PHP测试
使用的tp框架
```php
<?php
    function testAnsyTask(){
        $mRedis = CacheRedisQueue::getInstance();
        //shell脚本
        $data = [
            'class' => 'AnsyTask',
            'args'  => [
                [
                    'type'  => 'shell',
                    'dir'   => '/data/www/shell/',
                    'mainFile' => 'test.sh'
                ]
            ]
        ];
        $mRedis->rPush( 'kcResque:queue:900sui:ansyTaskQueue', json_encode( $data ) );
        
        //php 脚本
        $data = [
            'class' => 'AnsyTask',
            'args'  => [
                [
                    'type'      => 'php',
                    'dir'       => '/lamp/kctz/jkd11.9/',
                    'mainFile'  => 'index.php',
                    'phpbin'    => '/usr/local/php2/bin/php',
                    'action'    => '/Batch-Once-Test',
                    'cmdArgs'   => 'orderId=1&goodsId=2'
                ]
            ]
        ];

        $mRedis->rPush( 'kcResque:queue:900sui:ansyTaskQueue', json_encode( $data ) );
    }
```

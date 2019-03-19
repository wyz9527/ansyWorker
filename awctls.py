#!/usr/bin/python
# -*- coding: UTF-8 -*-

import json,os,sys,signal,time,getopt

try:
	opts, args = getopt.gnu_getopt(sys.argv[1:],"hc:", ["help","config="])
except:
	print 'Error: awctls.py -c <config> '
	sys.exit(0)

argLen = len(sys.argv)
if argLen < 2:
	print "arg start|restart|stop need one"
	sys.exit(0) #退出程序

dir 		=  os.getcwd()
configFile 	= dir + "/config.json"
for opt, arg in opts:
	if opt in ("-h", "--help"):
		print 'awctls.py start|stop|restart -c <config>'
		sys.exit(0)
	elif opt in ("-c", "--config"):
		configFile = arg

#检查进程是否存在
def check_pid(pid):
	""" Check For the existence of a unix pid. """
	try:
		os.kill(pid, 0)
	except OSError:
		return False
	else:
		return True

#启动
def ansyStart():
	print "ansyTask start..."
	dir =  os.getcwd()
	awConf = readConfig()
	if os.path.exists(awConf['pid']):
		with open(awConf['pid']) as f:
			pid = f.read()
		if pid != "":
			if check_pid( int(pid) ):
				print pid + " server has runing; continued"
				return True
	cmd = u'%s/ansyTask -queues=%s -uri=%s -connections=%s -concurrency=%s -namespace=%s -interval=%s -use-number=true -exit-on-complete=false -pid=%s -log=%s & '%(dir,awConf['queues'],awConf['uri'],awConf['connections'],awConf['concurrency'],awConf['namespace'],awConf['interval'],awConf['pid'], awConf['log'])
	print "Run :" + cmd
	print  os.system(cmd)

#停止
def ansyStop():
	print "ansyTask stoping..."
	awConf = readConfig()
	if os.path.exists(awConf['pid']) == False :
		print "not find pid file:"+awConf['pid']
		return False

	with open(awConf['pid']) as f:
		pid = f.read()
	if check_pid( int(pid) ):
		try:
			print os.kill(int(pid),signal.SIGQUIT)
			print  "stop success"
		except OSError,e:
			print  e.message
	else:
		print pid + "server is not run"

#读取配置文件
def readConfig():
	print configFile
	with open(configFile) as f:
		awConf = json.loads( f.read() )
	return awConf


#接收命令行参数，做出相应的动作
if sys.argv[1] == 'start':
	ansyStart()
elif sys.argv[1] == 'stop':
	ansyStop()
elif sys.argv[1] == 'restart':
	ansyStop()
	time.sleep(2)
	ansyStart()



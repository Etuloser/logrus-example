package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// 转换成json格式方便logstash或者Splunk解析
	// log.SetFormatter(&log.JSONFormatter{})

	// 格式化默认配置
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	// 开启调用发起函数
	// log.SetReportCaller(true)

	// TODO:
	log.SetOutput(os.Stderr)

	// 简单使用
	log.Info("some message")

	// 日志包含提示字段
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	// 字段重用
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
	
	// 致命异常, 会返回code 1后退出程序
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

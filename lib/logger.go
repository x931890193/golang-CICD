package lib

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

//var logger *zap.SugaredLogger
var Logger *logrus.Entry

func InitLogger() {
	log := logrus.New()
	fileName := "log/xj.log"
	if os.Getenv("PROGRAM_ENV") == "pro" {
		name, _ := os.Hostname()
		fileName = fmt.Sprintf("log/%s_xj.log", name)
		logger := &lumberjack.Logger{
			// 日志输出文件路径
			Filename: fileName,
			// 日志文件最大 size, 单位是 MB
			MaxSize: 300, // megabytes
			// 最大过期日志保留的个数
			MaxBackups: 30,
			// 保留过期文件的最大时间间隔,单位是天
			// 是否需要压缩滚动日志, 使用的 gzip 压缩
			Compress: true, // disabled by default
		}

		log.SetOutput(logger)
		log.SetFormatter(&ecslogrus.Formatter{})
	} else {
		log.SetOutput(os.Stdout)
		log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
	}
	Logger = log.WithField("psm", "xj")
}

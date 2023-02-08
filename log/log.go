package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wyt/GinStudy/conf"
)

var Logger logrus.Logger

func init() {
	Logger = *logrus.StandardLogger()

	// 设置全局Logger的日志格式
	Logger.SetFormatter(&logrus.JSONFormatter{
		// 设置日期格式
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	// 设置全局Logger的日志级别
	var lv logrus.Level
	lv.UnmarshalText([]byte(conf.LogLevel))
	Logger.SetLevel(lv)

	//设置日志输出流(控制台输出)
	Logger.SetOutput(os.Stdout)

	file, err := os.OpenFile(conf.LogPath, os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		Logger.WithField("err", err).Error("fail to create or open log file")
		panic(err)
	}

	//控制台和文件复合输出流
	fileAndStdoutStream := io.MultiWriter(os.Stdout, file)

	//设置全局Logger日志输出流(同时输出控制台和文件)
	Logger.SetOutput(fileAndStdoutStream)
}

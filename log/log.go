package log

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wyt/GinStudy/conf"

	"github.com/lestrrat-go/file-rotatelogs"
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

	// file, err := os.OpenFile(conf.LogPath, os.O_APPEND|os.O_CREATE, 0666)

	// if err != nil {
	// 	Logger.WithField("err", err).Error("fail to create or open log file")
	// 	panic(err)
	// }

	//设置输出文件(切割文件)
	logf, err := rotatelogs.New(
		conf.LogPath+"_%Y%m%d.log",
		rotatelogs.WithLinkName(conf.LogPath),
		//最多保留30天的日志
		rotatelogs.WithMaxAge(time.Hour*24*30),
		//24小时分割一次文件
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		Logger.WithField("err", err).Error("failed to create rotatelogs: %s", err)
		panic(err)
	}

	//控制台和文件复合输出流
	// fileAndStdoutStream := io.MultiWriter(os.Stdout, file)
	fileAndStdoutStream := io.MultiWriter(os.Stdout, logf)

	//设置全局Logger日志输出流(同时输出控制台和文件)
	Logger.SetOutput(fileAndStdoutStream)
}

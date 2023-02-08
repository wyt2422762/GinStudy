package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wyt/GinStudy/log"
)

// 记录请求日志中间件
func WriteLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求时间
		reqTime := time.Now()

		c.Next()

		// 请求耗时(毫秒)
		spendTime := time.Since(reqTime).Milliseconds()

		// 状态码
		statusCode := c.Writer.Status()

		//客户端ip
		clientIP := c.ClientIP()

		// 请求url
		url := c.Request.URL

		// 请求方法
		method := c.Request.Method

		errs := c.Errors

		log.Logger.WithFields(logrus.Fields{
			"reqTime":    reqTime,
			"url":        url,
			"method":     method,
			"clientIP":   clientIP,
			"spendTime":  spendTime,
			"statusCode": statusCode,
			"errs":       errs.String(),
		}).Info("请求访问日志")
	}
}

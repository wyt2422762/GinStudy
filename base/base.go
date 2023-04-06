package base

import (
	"net/http"

	"github.com/gin-gonic/gin"

	cusErr "github.com/wyt/GinStudy/error"
)

// 返回数据格式
type Resp struct {
	Code    int         `json:"code" example:"200"`
	Msg     string      `json:"msg" example:"成功"`
	Success bool        `json:"success" example:"false"`
	Data    interface{} `json:"data" `
}

// 基础handler
type BaseHandler struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type handlerFunc func(c *gin.Context) error

// 统一错误处理
func Wrapper(handler handlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			if e, ok := err.(cusErr.CusError); ok {
				c.Status(e.Code)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			c.Error(err)
			return
		}
	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	// c.Status(http.StatusNotFound)
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusOK,
		"msg":     "404",
		"success": false,
	})
}

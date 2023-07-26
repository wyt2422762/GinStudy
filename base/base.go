package base

import (
	"net/http"

	"github.com/gin-gonic/gin"

	cusErr "github.com/wyt/GinStudy/error"
)

// 返回数据格式
type Resp struct {
	Code    int    `json:"code" example:"200"`
	Msg     string `json:"msg" example:"成功"`
	Success bool   `json:"success" example:"false"`
	Data    any    `json:"data" `
}

// 构造错误返回数据方法
func ErrorResp(code int, msg string) Resp {
	return Resp{
		Code:    code,
		Msg:     msg,
		Success: false,
	}
}

// 构造成功返回数据方法
func SuccessResp(data any) Resp {
	return Resp{
		Code:    200,
		Msg:     "成功",
		Success: true,
		Data:    data,
	}
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
				c.AbortWithStatusJSON(http.StatusOK, Resp{
					Code:    e.Code,
					Msg:     e.Msg,
					Success: false,
				})
			} else {
				c.AbortWithStatusJSON(http.StatusOK, Resp{
					Code:    http.StatusInternalServerError,
					Msg:     "系统内部错误",
					Success: false,
				})
			}
		}
	}
}

// 404, 405处理
func HandleNotFound(code int, msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, ErrorResp(code, msg))
	}
}

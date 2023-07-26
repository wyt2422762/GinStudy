package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wyt/GinStudy/base"
	cusErr "github.com/wyt/GinStudy/error"
)

// panic异常捕获
func ErrorHanlder() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		if err != nil {
			fmt.Printf("自定义错误处理, %t\n", err)
			switch err.(type) {
			case cusErr.CusError:
				e := err.(cusErr.CusError)
				c.AbortWithStatusJSON(http.StatusOK, base.ErrorResp(e.Code, e.Msg))
			default:
				c.AbortWithStatusJSON(http.StatusOK, base.ErrorResp(http.StatusInternalServerError, "系统内部错误"))
			}
		}
	}
}

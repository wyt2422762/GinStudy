package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/base"
)

type jwtTestHandler struct {
	base.BaseHandler
}

var JwtTest = &jwtTestHandler{}

// jwtTest方法
func (jwt *jwtTestHandler) JwtTest(c *gin.Context) error {
	fmt.Println("jwtTest")
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"msg":     "jwtTest",
		"success": false,
	})
	return nil
}

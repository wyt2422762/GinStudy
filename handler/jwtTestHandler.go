package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//jwtTest方法
func JwtTest(c *gin.Context) {
	fmt.Println("jwtTest")
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"msg":     "jwtTest",
		"success": false,
	})
}

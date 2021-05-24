package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func SimpleHandler(c *gin.Context) {
	fmt.Printf("收到请求，请求地址%s\n", c.FullPath())
	c.String(http.StatusOK, "从服务端返回的数据")
}

// url: /user/:name
func SimpleUrlParamHandler(c *gin.Context) {
	//获取参数
	name := c.Param("name")
	c.String(http.StatusOK, "参数 = " + name)
}

// url: /user?name=xxx&age=xx
func SimplParamHandler(c *gin.Context) {
	//获取参数
	name := c.Query("name")
	age := c.Query("age")
	c.String(http.StatusOK, "name = " + name + " age = " + age)
}
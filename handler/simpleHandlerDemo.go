package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/base"
)

type simpleHandler struct {
	base.BaseHandler
}

var Simple = &simpleHandler{}

func (s *simpleHandler) SimpleHandler(c *gin.Context) error {
	fmt.Printf("收到请求，请求地址%s\n", c.FullPath())
	c.String(http.StatusOK, "从服务端返回的数据")
	return nil
}

// url: /user/:name
func (s *simpleHandler) SimpleUrlParamHandler(c *gin.Context) error {
	//获取参数
	name := c.Param("name")
	c.String(http.StatusOK, "参数 = "+name)
	return nil
}

// url: /user?name=xxx&age=xx
func (s *simpleHandler) SimplParamHandler(c *gin.Context) error {
	//获取参数
	name := c.Query("name")
	age := c.Query("age")
	c.String(http.StatusOK, "name = "+name+" age = "+age)
	return nil
}

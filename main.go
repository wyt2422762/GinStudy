package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/conf"
	cusErr "github.com/wyt/GinStudy/error"
	"github.com/wyt/GinStudy/handler"
	"github.com/wyt/GinStudy/middlewares"
)

func main() {
	// simpleDemo()
	// simpleRouterDemo()
	// simpleRouterGroupDemo()
	simpleTemplateDemo()
}

//简单示例
func simpleDemo() {
	fmt.Println("Gin简单示例")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Printf("收到请求，请求地址%s\n", c.FullPath())
	})
	r.Run(":" + strconv.Itoa(conf.HttpPort))
}

//简单路由示例
func simpleRouterDemo() {
	fmt.Println("Gin Router 简单示例")
	r := gin.Default()
	//最简单的请求路由
	r.GET("/", handler.SimpleHandler)
	//带路径参数的路由
	r.GET("/user/:name", handler.SimpleUrlParamHandler)
	//?后面传参的路由
	r.GET("/user", handler.SimplParamHandler)
	r.Run(":" + strconv.Itoa(conf.HttpPort))
}

//简单路由分组示例
func simpleRouterGroupDemo() {
	fmt.Println("Gin Router 简单示例")
	r := gin.Default()

	//index路由组
	index := r.Group("/")
	{
		//最简单的请求路由
		index.GET("", handler.SimpleHandler)
	}

	//user路由组
	user := r.Group("/user")
	{
		//带路径参数的路由
		user.GET("/:name", handler.SimpleUrlParamHandler)
		//?后面传参的路由
		user.GET("", handler.SimplParamHandler)
	}

	r.Run(":" + strconv.Itoa(conf.HttpPort))
}

type HandlerFunc func(c *gin.Context) error

//统一错误处理
func wrapper(handler HandlerFunc) func(c *gin.Context) {
    return func(c *gin.Context) {
        err := handler(c)
		if err != nil {
			if e, ok := err.(cusErr.CusError); ok {
				c.Status(e.Code)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			fmt.Println("出错啦: ", err.Error())
			return
		}
    }
}

//404处理
func HandleNotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
    fmt.Println("404")
}

//简单template示例
func simpleTemplateDemo() {
	fmt.Println("Gin Router 简单示例")
	r := gin.Default()
	//设置中间件(跨域问题)
	r.Use(middlewares.Cors())
	//404处理
	r.NoMethod(HandleNotFound)
    r.NoRoute(HandleNotFound)
	//index路由组
	index := r.Group("/")
	{
		//最简单的请求路由
		index.GET("", handler.SimpleHandler)
	}
	//user路由组
	user := r.Group("/user")
	{
		//带路径参数的路由
		user.GET("/:name", handler.SimpleUrlParamHandler)
		//?后面传参的路由
		user.GET("", handler.SimplParamHandler)
	}
	//模板资源位置
	r.LoadHTMLGlob("templates/**/*")
	//静态资源位置
	r.Static("/static", "./static")
	//page/auth路由组
	page := r.Group("/page/auth")
	{
		page.GET("/toLogin", wrapper(handler.GotoLoginPage))
		page.POST("/login", wrapper(handler.Login))
	}
	// jwtTest路由组
	jwtTest := r.Group("/jwtTest", middlewares.JwtAuth())
	{
		jwtTest.GET("", handler.JwtTest)
	}
	r.Run(":" + strconv.Itoa(conf.HttpPort))
}


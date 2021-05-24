package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/conf"
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

//简单template示例
func simpleTemplateDemo() {
	fmt.Println("Gin Router 简单示例")
	r := gin.Default()
	//设置中间件(跨域问题)
	r.Use(middlewares.Cors())
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
		page.GET("/toLogin", handler.GotoLoginPage)
		page.POST("/login", handler.Login)
	}
	// jwtTest路由组
	jwtTest := r.Group("/jwtTest", middlewares.JwtAuth())
	{
		jwtTest.GET("", handler.JwtTest)
	}
	r.Run(":" + strconv.Itoa(conf.HttpPort))
}

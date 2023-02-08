package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/base"
	"github.com/wyt/GinStudy/middlewares"
)

// test-router
func testRouter(r *gin.Engine) {
	//最简单的请求路由
	// r.GET("/test", test)

	testGroup := r.Group("/test", middlewares.WriteLog())
	{
		testGroup.GET("", test)
	}
}

// 测试 gin-swagger
// @Summary 测试 gin-swagger
// @Schemes
// @Description 测试 gin-swagger
// @Tags 测试 gin-swagger
// @Accept json
// @Produce json
// @param id query string false "id"
// @Success 200 {object} base.Resp
// @Router /test [get]
func test(c *gin.Context) {
	fmt.Printf("收到请求，请求地址%s\n", c.FullPath())
	id := c.Query("id")
	c.JSON(http.StatusOK, base.Resp{
		Code:    http.StatusOK,
		Msg:     "从服务端返回的数据：" + id,
		Success: true,
	})
}

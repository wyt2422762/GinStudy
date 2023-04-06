package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/base"
	"github.com/wyt/GinStudy/handler"
	"github.com/wyt/GinStudy/middlewares"
)

// test-router
func testRouter(r *gin.Engine) {
	//最简单的请求路由
	// r.GET("/test", test)

	testGroup := r.Group("/test", middlewares.WriteLog())
	{
		testGroup.GET("/1", base.Wrapper(handler.Test.Test))
		testGroup.GET("/2/:id", base.Wrapper(handler.Test.Test2))
	}

}

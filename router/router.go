package router

import "github.com/gin-gonic/gin"

func BuildRouter(r *gin.Engine) {
	//swagger-router
	swagRouter(r)
	//test-router
	testRouter(r)
}
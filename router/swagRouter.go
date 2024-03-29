package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	docs "github.com/wyt/GinStudy/docs"
)

// swagger-router
func swagRouter(r *gin.Engine) {
	//基础url
	docs.SwaggerInfo.BasePath = ""
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/base"
	cusErr "github.com/wyt/GinStudy/error"
)

type testHandler struct {
	base.BaseHandler
}

var Test = &testHandler{}

// 测试 gin-swagger
// @Summary 测试 gin-swagger
// @Schemes
// @Description 测试 gin-swagger
// @Tags 测试 gin-swagger
// @Accept json
// @Produce json
// @param id query string false "id"
// @Security ApiKeyAuth
// @Success 200 {object} base.Resp
// @Router /test/1 [get]
func (ts *testHandler) Test(c *gin.Context) error {
	fmt.Printf("收到请求，请求地址%s\n", c.FullPath())
	id := c.Query("id")
	c.JSON(http.StatusOK, base.Resp{
		Code:    http.StatusOK,
		Msg:     "从服务端返回的数据：" + id,
		Success: true,
	})
	return nil
}

// 测试 gin-swagger
// @Summary 测试 gin-swagger
// @Schemes
// @Description 测试 gin-swagger
// @Tags 测试 gin-swagger
// @Accept json
// @Produce json
// @param id path string true "id"
// @Success 200 {object} base.Resp
// @Router /test/2/{id} [get]
func (ts *testHandler) Test2(c *gin.Context) error {
	fmt.Printf("收到请求，请求地址%s\n", c.FullPath())
	// panic(cusErr.NewCusError(http.StatusBadRequest, "参数错误"))
	
	id := c.Param("id")

	if id == "999" {
		return cusErr.NewCusError(400, "黑名单")
	}

	c.JSON(http.StatusOK, base.Resp{
		Code:    http.StatusOK,
		Msg:     "从服务端返回的数据：" + id,
		Success: true,
		Data: 123,
	})

	return nil
}

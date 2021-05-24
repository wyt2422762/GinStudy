package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/model"
)

func GotoLoginPage(c *gin.Context) {
	title := c.Param("title")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": title,
	})
}

//登录方法
func Login(c *gin.Context) {
	//普通方式获取
	/* username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("username: %s, password: %s\n", username, password) */

	//参数绑定实体类
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	fmt.Printf("username: %s, password: %s\n", user.Username, user.Password)
	c.HTML(http.StatusOK, "index.html", nil)
}

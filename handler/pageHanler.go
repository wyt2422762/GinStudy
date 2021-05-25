package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/jwt"
	"github.com/wyt/GinStudy/model"
	cusErr "github.com/wyt/GinStudy/error"
)

func GotoLoginPage(c *gin.Context) error {
	title := c.Param("title")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": title,
	})
	return nil
}

//登录方法
func Login(c *gin.Context) error {
	//普通方式获取
	/* username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("username: %s, password: %s\n", username, password) */

	//参数绑定实体类
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("error: ", err.Error())
		return cusErr.NewCusError(http.StatusBadRequest, "参数错误")
	}

	//生成token
	loginUser := model.LoginUser{
		CacheKey: model.LOGIN_TOKEN_KEY + user.Username + ":" + time.Now().String(),
		User: user,
	}
	token, err := jwt.CreateToken(loginUser)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return cusErr.NewCusError(http.StatusBadRequest, "系统内部错误-生成token有误")
	}
	//存redis
	//redis.xxxxxx

	fmt.Printf("username: %s, password: %s\n", user.Username, user.Password)
	c.HTML(http.StatusOK, "index.html", token)
	return nil
}

package middlewares

import (
	"fmt"
	"net/http"
	_ "net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/wyt/GinStudy/jwt"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if strings.Trim(token, " ") == "" {
			fmt.Println("token为空")
			c.JSON(http.StatusUnauthorized, gin.H{
                "code": http.StatusUnauthorized,
				"success": false,
                "msg":  "token为空",
            })
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil{
			fmt.Println("token认证失败")
			c.JSON(http.StatusUnauthorized, gin.H{
                "code": http.StatusUnauthorized,
				"success": false,
                "msg":  "token认证失败",
            })
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
        c.Set("claims", claims)
	}
}

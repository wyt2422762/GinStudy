package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/wyt/GinStudy/model"
)

//token过期时间
const tokenExpireDuration = time.Hour * 2
//密钥
var sercet = []byte("wyt")

//自定义载荷
type CusClaims struct {
	CacheKey string `json:"cacheKey"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"pwssword"`
	jwt.StandardClaims
}

//生成token
func CreateToken(user model.LoginUser) (string, error){
	cusClaims := CusClaims{
		user.CacheKey,
		user.ID,
		user.Username,
		user.Password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(),
			Issuer:    "wyt", //签发者
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cusClaims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(sercet)
}

//解析token
func ParseToken(tokenString string) (*CusClaims, error){
	token, err := jwt.ParseWithClaims(tokenString, &CusClaims{}, func(token *jwt.Token)(i interface{}, err error){
		return sercet, nil
	})
	if err != nil{
		fmt.Println("Token验证失败")
		return nil, err
	}
	if claims, ok := token.Claims.(*CusClaims); ok && token.Valid{
		return claims, nil
	}
	return nil, errors.New("invalid token")
}



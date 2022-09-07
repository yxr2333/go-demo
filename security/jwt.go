package security

import (
	"errors"
	"net/http"
	"sheep/demo/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var signSecret = []byte("hellosheep")

// var signSecret = "hellosheep"

// 设置过期时间
const TokenExpireDuration = time.Hour * 24

type UserClaims struct {
	UID      uint   `json:"uid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(id uint, username string) (string, error) {
	claims := UserClaims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "sheep", // 签发人
		},
	}
	/*
		不同的签名方法，需要不同的密钥进行签名
		有的需要传入string，有的需要[]byte
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signSecret)
}

func ParseToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return signSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无法解析的token")
}

func JWTAuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("请求头中未找到有效token"))
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("token格式有误"))
			c.Abort()
			return
		}
		c.Set("uid", claims.UID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

package service

import (
	"github.com/gin-gonic/gin"
)

func TestService(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":     200,
		"msg":      "Hello,World",
		"uid":      c.GetString("uid"),
		"username": c.GetString("username"),
	})
}

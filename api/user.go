package api

import (
	"sheep/demo/security"
	"sheep/demo/service"

	"github.com/gin-gonic/gin"
)

func LoadUserApi(r *gin.Engine) {
	user := r.Group("/user")
	{
		normal := user.Group("/n")
		{
			normal.POST("/register", service.DoRegister)
			normal.POST("/login", service.DoLogin)
			normal.PUT("/up", service.UpdateUserBaseInfo)
		}
		privilege := user.Group("/p", security.JWTAuthMiddleWare())
		{
			privilege.GET("/all", service.FindAllUsers)
			privilege.DELETE("/del/:id", service.DeleteOneUser)
		}
	}
}

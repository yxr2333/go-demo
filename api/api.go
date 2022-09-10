package api

import (
	"sheep/demo/security"
	"sheep/demo/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("", security.JWTAuthMiddleWare(), service.TestService)
	LoadUserApi(r)
	LoadRoleApi(r)
	LoadCommonApi(r)
	LoadCardApi(r)
	LoadFileApi(r)
	LoadUserClassApi(r)
	return r
}

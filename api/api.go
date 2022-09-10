package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	LoadUserApi(r)
	LoadRoleApi(r)
	LoadCommonApi(r)
	LoadCardApi(r)
	LoadFileApi(r)
	LoadUserClassApi(r)
	return r
}

package api

import (
	"sheep/demo/service"

	"github.com/gin-gonic/gin"
)

func LoadRoleApi(r *gin.Engine) {
	role := r.Group("/role")
	{
		role.POST("", service.SaveRole)
	}
}

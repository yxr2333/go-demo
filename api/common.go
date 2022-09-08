package api

import (
	"sheep/demo/service"

	"github.com/gin-gonic/gin"
)

func LoadCommonApi(r *gin.Engine) {
	common := r.Group("/common")
	{
		common.POST("/gen/code", service.GenerateCaptchaCode)
		common.POST("/ver/code", service.VerifyCaptchaCode)
	}
}

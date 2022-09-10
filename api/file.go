package api

import (
	"sheep/demo/service"

	"github.com/gin-gonic/gin"
)

func LoadFileApi(r *gin.Engine) {
	file := r.Group("/file")
	{
		file.POST("/upload/one", service.UploadFile)
	}
}

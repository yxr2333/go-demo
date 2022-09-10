package api

import (
	"sheep/demo/service"

	"github.com/gin-gonic/gin"
)

func LoadUserClassApi(r *gin.Engine) {
	userClass := r.Group("/user_class")
	{
		userClass.POST("", service.SaveOneClass)
		userClass.GET("/all", service.FindAllClasses)
		userClass.GET("/all/students", service.FindAllStudentsByClassID)
	}
}

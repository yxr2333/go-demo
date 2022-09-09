package api

import (
	"sheep/demo/service"

	"github.com/gin-gonic/gin"
)

func LoadCardApi(r *gin.Engine) {
	card := r.Group("/card")
	{
		card.POST("", service.CreateOneIdCard)
	}
}

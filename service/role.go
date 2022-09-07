package service

import (
	"fmt"
	"net/http"
	"sheep/demo/common"
	"sheep/demo/model"
	"sheep/demo/sql"

	"github.com/gin-gonic/gin"
)

func SaveRole(c *gin.Context) {
	name := c.Query("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("请输入角色名称"))
		return
	}
	role := model.UserRole{
		Rolename: name,
	}
	result := sql.DB.Create(&role)
	if result.Error != nil {
		fmt.Printf("create role error: %s", result.Error.Error())
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("新增角色失败"))
		return
	}
	c.JSON(200, common.SuccessReturnWithMsg("新增角色成功"))
}

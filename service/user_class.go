package service

import (
	"fmt"
	"sheep/demo/common"
	"sheep/demo/model"
	"sheep/demo/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveOneClass(c *gin.Context) {
	name := c.Query("name")
	if len(name) <= 0 {
		c.JSON(400, common.ErrorReturnWithMsg("参数错误"))
		return
	}
	result := sql.DB.Create(&model.Class{ClassName: name})
	if result.Error != nil {
		fmt.Printf("插入异常: %v\n", result.Error)
		c.JSON(400, common.ErrorReturnWithMsg("创建失败"))
		return
	}
	c.JSON(200, common.SuccessReturnWithMsg("创建成功"))
}

func FindAllClasses(c *gin.Context) {
	var classes []model.Class
	result := sql.DB.Find(&classes)
	if result.Error != nil {
		fmt.Printf("查询异常: %v\n", result.Error)
		c.JSON(400, common.ErrorReturnWithMsg("查询失败"))
		return
	}
	c.JSON(200, common.SuccessReturnWithMsgAndData("查询成功", classes))
}

func FindAllStudentsByClassID(c *gin.Context) {
	id := c.Query("cid")
	if len(id) <= 0 {
		c.JSON(400, common.ErrorReturnWithMsg("参数错误"))
		return
	}
	cid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, common.ErrorReturnWithMsg("参数错误"))
		return
	}
	class := model.Class{
		Model: gorm.Model{
			ID: uint(cid),
		},
	}
	sql.DB.Model(model.Class{}).Preload("Users").Find(&class)
	c.JSON(200, common.SuccessReturnWithMsgAndData("查询成功", class.Users))
}

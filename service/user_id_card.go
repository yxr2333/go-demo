package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sheep/demo/common"
	"sheep/demo/model"
	"sheep/demo/sql"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOneIdCard(c *gin.Context) {
	raw, _ := c.GetRawData()
	var param common.CreateUserIDCardParam
	json.Unmarshal(raw, &param)
	if len(param.BankName) == 0 || param.UID <= 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("数据格式有误"))
		return
	}
	rand.Seed(time.Now().UnixNano())
	card := model.UserIDCard{
		UserID:     param.UID,
		BankName:   param.BankName,
		CardNumber: strconv.Itoa(rand.Intn(99999)),
	}
	result := sql.DB.Create(&card)
	if result.Error != nil {
		fmt.Printf("插入异常: %v", result.Error)
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("暂时无法发卡"))
		return
	}
	c.JSON(200, common.SuccessReturnWithMsg("发卡成功"))
}

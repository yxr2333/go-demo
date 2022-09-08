package service

import (
	"fmt"
	"net/http"
	"sheep/demo/common"
	"sheep/demo/store"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateCaptchaCode(c *gin.Context) {
	code := common.RandomStr(5)
	reqId := common.RandomStr(8)
	err := store.RedisClient.Set("CODE_"+reqId, code, time.Minute*2).Err()
	if err != nil {
		fmt.Println("set code value failed", err)
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("暂时无法生成验证码"))
		return
	}
	c.JSON(200, gin.H{
		"reqId": reqId,
		"code":  code,
	})
}

func VerifyCaptchaCode(c *gin.Context) {
	reqId := c.Query("reqId")
	code := c.Query("code")
	if len(reqId) == 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("未找到请求编号"))
		return
	}
	if len(code) == 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("未找到验证码"))
		return
	}
	codeKey := "CODE_" + reqId
	result, err := store.RedisClient.Get(codeKey).Result()
	if err != nil {
		fmt.Printf("find value in redis error: %v\n", err)
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("未获取验证码或验证码已过期"))
		return
	}
	if code == result {
		store.RedisClient.Del(codeKey)
		c.JSON(200, common.SuccessReturnWithMsg("验证成功"))
	} else {
		c.JSON(http.StatusBadRequest, common.SuccessReturnWithMsg("验证码错误"))

	}

}

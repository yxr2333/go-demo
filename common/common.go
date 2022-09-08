package common

import (
	"math/rand"
	"net/http"
	convert "strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessReturn(msg string) ApiResult {
	return ApiResult{
		Code: http.StatusOK,
	}
}

func SuccessReturnWithMsg(msg string) ApiResult {
	return ApiResult{
		Code: http.StatusOK,
		Msg:  msg,
	}
}

func SuccessReturnWithMsgAndData(msg string, data interface{}) ApiResult {
	return ApiResult{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	}
}

func ErrorReturnWithMsg(msg string) ApiResult {
	return ApiResult{
		Code: http.StatusBadRequest,
		Msg:  msg,
	}
}

func ErrorReturnWithMsgAndData(msg string, data interface{}) ApiResult {
	return ApiResult{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: data,
	}
}

// 随机生成字符串

func RandomStr(n int) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r1.Intn(len(letters))]
	}
	return string(b)
}

// 生成分页条件
func CreatePaginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageNum, err := convert.Atoi(c.DefaultQuery("pageNum", "1"))
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorReturnWithMsg("请输入正确的页码"))
			c.Abort()
		}
		pageSize, err := convert.Atoi(c.DefaultQuery("pageSize", "1"))
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorReturnWithMsg("请输入正确的页大小"))
			c.Abort()
		}
		if pageNum == 0 {
			pageNum = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func IsNum(s string) bool {
	_, err := convert.ParseFloat(s, 64)
	return err == nil
}

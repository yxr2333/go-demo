package service

import (
	"context"
	"fmt"
	"net/http"
	"sheep/demo/common"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("暂时无法上传文件"))
		return
	}
	now := time.Now()
	name := now.Format("2006-01-02T15:04:05") + file.Filename
	fileHandle, err := file.Open()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("暂时无法上传文件"))
		return
	}
	defer fileHandle.Close()
	_, err = common.OSSClient.Object.Put(context.Background(), name, fileHandle, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("暂时无法上传文件"))
		return
	}
	c.JSON(http.StatusOK, common.SuccessReturnWithMsgAndData("上传成功", gin.H{
		"url": "https://insurence-1304011999.cos.ap-shanghai.myqcloud.com/" + name,
	}))
}

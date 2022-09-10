package main

import (
	"fmt"
	"sheep/demo/api"
	"sheep/demo/common"
	"sheep/demo/config"
	"sheep/demo/model"
	"sheep/demo/sql"
	"sheep/demo/store"
)

func main() {
	sql.InitSqlConnection()
	err := store.InitRedisClient()
	if err != nil {
		fmt.Printf("redis连接失败,错误%v\n", err)
		return
	}
	err = config.InitAllConfigs()
	if err != nil {
		fmt.Printf("读取配置文件出错,错误%v\n", err)
		return
	}
	err = common.InitCosClient()
	if err != nil {
		fmt.Printf("cos连接失败,错误%v\n", err)
		return
	}
	// InitTables()
	r := api.InitRoutes()
	r.Run(":8081")
}

func InitTables() {
	sql.DB.AutoMigrate(&model.UserRole{}, &model.User{}, &model.Class{}, &model.UserIDCard{})
}

package main

import (
	"fmt"
	"sheep/demo/api"
	"sheep/demo/model"
	"sheep/demo/sql"
	"sheep/demo/store"
)

func main() {
	sql.InitSqlConnection()
	err := store.InitRedisClient()
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}
	InitTables()
	r := api.InitRoutes()
	r.Run(":8081")
}

func InitTables() {
	sql.DB.AutoMigrate(&model.UserRole{}, &model.User{})
}

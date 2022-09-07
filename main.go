package main

import (
	"sheep/demo/api"
	"sheep/demo/model"
	"sheep/demo/sql"
)

func main() {
	sql.InitSqlConnection()
	InitTables()
	r := api.InitRoutes()
	r.Run(":8081")
}


func InitTables() {
	sql.DB.AutoMigrate(&model.UserRole{}, &model.User{})
}

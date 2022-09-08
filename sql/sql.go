package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitSqlConnection() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 200,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

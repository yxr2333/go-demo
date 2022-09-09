package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string       `gorm:"uniqueIndex" json:"username"`
	Password    string       `json:"password"`
	Salt        string       `json:"salt"`
	UserRoleID  uint         `json:"userRoleId" gorm:"foreignKey: ID"`
	UserIDCards []UserIDCard `json:"userIdCards"`
}

type UserRole struct {
	gorm.Model
	Rolename string `gorm:"uniqueIndex"`
	Users    []User
}

type UserIDCard struct {
	gorm.Model
	CardNumber string `json:"number"`
	BankName   string `json:"bankName"`
	UserID     uint   `json:"userId" gorm:"foreignKey: ID"`
}

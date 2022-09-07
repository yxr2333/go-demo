package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"uniqueIndex" json:"username"`
	Password   string `json:"password"`
	Salt       string `json:"salt"`
	UserRoleID uint   `json:"userRoleId"`
}

type UserRole struct {
	gorm.Model
	Rolename string `gorm:"uniqueIndex"`
	Users    []User
}

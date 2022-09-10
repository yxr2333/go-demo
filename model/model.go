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
	UserClasses []Class      `json:"userClasses" gorm:"many2many:user_classes;"`
}

type Class struct {
	gorm.Model
	ClassName string `gorm:"uniqueIndex" json:"className"`
	Users     []User `json:"users" gorm:"many2many:user_classes;"`
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

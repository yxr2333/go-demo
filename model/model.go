package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"uniqueIndex" json:"username"`
	Password   string `json:"password"`
	Salt       string `json:"salt"`
	UserRoleID uint   `json:"userRoleId" gorm:"foreignKey: ID"`
}

// func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
// 	result := tx.First(&u)
// 	fmt.Println("result rows:", result.RowsAffected)
// 	if result.RowsAffected <= 0 {
// 		return errors.New("不能删除不存在的数据")
// 	}
// 	return
// }

type UserRole struct {
	gorm.Model
	Rolename string `gorm:"uniqueIndex"`
	Users    []User
}

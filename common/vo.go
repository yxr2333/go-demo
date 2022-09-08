package common

type UserLoginParam struct {
	Username string
	Password string
}

type UpdateUserBaseInfoParam struct {
	ID         uint
	Username   string `json:"username"`
	UserRoleID uint   `json:"userRoleId" gorm:"foreignKey: ID"`
}

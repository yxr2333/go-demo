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

type CreateUserIDCardParam struct {
	BankName string `json:"bankName"`
	UID      uint   `json:"uid"`
}

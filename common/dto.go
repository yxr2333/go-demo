package common

type UserBaseInfo struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	UserRoleID uint   `json:"userRoleId"`
}

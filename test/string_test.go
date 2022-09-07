package test

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"sheep/demo/common"
	"sheep/demo/model"
	"sheep/demo/security"
	"testing"
)

func TestString(t *testing.T) {
	var user model.User
	user.ID = 10
	println(user.Username)
	println(len(user.Username))
	user.Username = "yxr"
	println(len(user.Username))

}

func TestUserLogin(t *testing.T) {
	password := "123456"
	str := common.RandomStr(10)
	salt := base64.StdEncoding.EncodeToString([]byte(str))
	println("str:", str)
	encoder1 := md5.New()
	mdResult1 := encoder1.Sum([]byte(password + salt))
	dbPwd1 := hex.EncodeToString(mdResult1)

	encoder2 := md5.New()
	mdResult2 := encoder2.Sum([]byte(password + salt))
	dbPwd2 := hex.EncodeToString(mdResult2)

	ss, _ := base64.StdEncoding.DecodeString(salt)
	encoder3 := md5.New()
	mdResult3 := encoder3.Sum([]byte(password + string(ss)))
	dbPwd3 := hex.EncodeToString(mdResult3)

	println("ss:", string(ss))
	if dbPwd1 == dbPwd2 && dbPwd1 == dbPwd3 {
		println("yes")
	} else {
		println("no")
	}
}

func TestCreateToken(t *testing.T) {
	token, err := security.CreateToken(1, "yxr")
	if err != nil {
		println(err.Error())
	} else {
		println(token)
	}
}

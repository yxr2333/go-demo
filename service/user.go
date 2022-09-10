package service

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sheep/demo/common"
	"sheep/demo/model"
	"sheep/demo/security"
	"sheep/demo/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DoRegister(c *gin.Context) {
	b, _ := c.GetRawData()
	var m model.User
	_ = json.Unmarshal(b, &m)
	if len(m.Username) == 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("请输入用户名"))
		return
	}
	if len(m.Password) == 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("请输入密码"))
		return
	}
	// 生成随机盐值
	randomStr := common.RandomStr(10)
	salt := base64.StdEncoding.EncodeToString([]byte(randomStr))
	m.Salt = salt
	// 对密码+盐值组成的新字符串进行md5加密
	encoder := md5.New()
	mdResult := encoder.Sum([]byte(m.Password + salt))
	// 将加密后的密码设置回去
	m.Password = hex.EncodeToString(mdResult)
	/*
		调用数据库连接，保存数据
		注册的时候设置默认的角色为普通用户
	*/
	m.UserRoleID = 2
	result := sql.DB.Create(&m)
	if result.Error != nil {
		fmt.Printf("error: %v", result.Error.Error())
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("注册失败"))
		return
	}
	c.JSON(http.StatusOK, common.SuccessReturnWithMsg("注册成功"))
}

func DoLogin(c *gin.Context) {
	data, _ := c.GetRawData()
	var param common.UserLoginParam
	_ = json.Unmarshal(data, &param)
	// fmt.Printf("param: %v\n", param)
	if len(param.Username) == 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("请输入用户名"))
		return
	}
	if len(param.Password) == 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("请输入密码"))
		return
	}
	var dbUser model.User
	sql.DB.Where("username = ?", param.Username).First(&dbUser)
	// 获得加密的盐值
	_, err := base64.StdEncoding.DecodeString(dbUser.Salt)
	if err != nil {
		fmt.Printf("decode error: %v\n", err.Error())
		return
	}
	// salt := string(s)
	// fmt.Println(salt)
	// 使用查询到的盐值对传入的密码进行加密
	encoder := md5.New()
	mdResult := encoder.Sum([]byte(param.Password + dbUser.Salt))
	pwd := hex.EncodeToString(mdResult)
	// 判断加密后的密码是否匹配
	// fmt.Printf("pwd: %s, dbPwd: %s\n", pwd, dbUser.Password)
	if pwd == dbUser.Password {
		token, err := security.CreateToken(dbUser.Model.ID, dbUser.Username)
		if err != nil {
			fmt.Printf("create token error: %s\n", err.Error())
			c.JSON(http.StatusOK, common.ErrorReturnWithMsg("分发token失败"))
			return
		}
		c.JSON(http.StatusOK, common.SuccessReturnWithMsgAndData("登录成功", gin.H{
			"token": token,
		}))
	} else {
		c.JSON(http.StatusOK, common.ErrorReturnWithMsg("用户名或密码输入错误"))
	}

}

func FindAllUsers(c *gin.Context) {
	var users []model.User
	result := sql.DB.Scopes(common.CreatePaginate(c)).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("查询失败"))
		return
	}
	var datas []common.UserBaseInfo
	for _, value := range users {
		temp := common.UserBaseInfo{
			ID:         value.Model.ID,
			Username:   value.Username,
			UserRoleID: value.UserRoleID,
		}
		datas = append(datas, temp)
	}
	c.JSON(200, common.SuccessReturnWithMsgAndData("查询成功", datas))
}

func DeleteOneUser(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if len(id) == 0 || !common.IsNum(id) || err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("输入正确的编号"))
		return
	}
	result := sql.DB.Delete(&model.User{}, uid)
	if result.Error != nil {
		fmt.Printf("delete error: %v", result.Error)
		return
	}
	if result.RowsAffected <= 0 {
		c.JSON(http.StatusBadRequest, common.SuccessReturnWithMsg("数据不存在"))
		return
	}
	c.JSON(200, common.SuccessReturnWithMsg("删除成功"))
}

func UpdateUserBaseInfo(c *gin.Context) {
	data, _ := c.GetRawData()
	var info common.UserBaseInfo
	err := json.Unmarshal(data, &info)
	if err != nil {
		fmt.Printf("更新时获取数据异常：%v\n", err)
		return
	}
	user := model.User{
		Model: gorm.Model{
			ID: info.ID,
		},
		Username:   info.Username,
		UserRoleID: info.UserRoleID,
	}
	result := sql.DB.Model(&user).Updates(user)
	fmt.Println("受影响的行数为：", result.RowsAffected)
	c.JSON(200, common.SuccessReturnWithMsg("更新成功"))
}

func InsertMany(c *gin.Context) {
	data, _ := c.GetRawData()
	var users []model.User
	err := json.Unmarshal(data, &users)
	if err != nil {
		fmt.Printf("反序列化异常：%v\n", err)
		return
	}
	sql.DB.Create(&users)
	c.JSON(200, common.SuccessReturnWithMsg("添加成功"))
}

func FindOne(c *gin.Context) {
	uid := c.Param("id")
	println(uid)
	id, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("请输入正确的编号"))
		return
	}
	var user model.User
	rows := sql.DB.Preload("UserIDCards").Preload("UserClasses").Find(&user, id).RowsAffected
	if rows <= 0 {
		c.JSON(http.StatusBadRequest, common.ErrorReturnWithMsg("暂未找到对应的信息"))
		return
	}
	c.JSON(200, common.SuccessReturnWithMsgAndData("查询成功", user))
}

package model

import (
	"flamingo/database"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

// User 用户类
type User struct {
	gorm.Model
	UniqueId uuid.UUID
	Name     string `json:"name"`
	Gender   uint   `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// LoginReq 登录请求参数类
type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// Register 插入用户，先检查是否存在用户，如果没有则存入
func Register(mobile string, password string) error {
	if CheckUser(mobile) {
		return fmt.Errorf("该手机号已经注册！")
	}

	newUser := User{
		UniqueId: uuid.New(),
		Name:     mobile,
		Gender:   0,
		Mobile:   mobile,
		Password: password,
	}

	newUser.SetPassword(password)

	return database.DB.Create(&newUser).Error
}

// CheckUser 检查用户是否存在
func CheckUser(mobile string) bool {
	count := 0
	database.DB.Model(&User{}).Where("mobile = ?", mobile).Count(&count)
	result := false
	if count > 0 {
		result = true
	}
	return result
}

// LoginCheck 登录验证
func LoginCheck(loginRequest LoginRequest) (bool, User, error) {
	resultBool := false
	user := User{}
	//	这里进行密码校验
	if err := database.DB.Where("mobile = ?", loginRequest.Mobile).First(&user).Error; err != nil {
		return resultBool, user, err
	} else {
		//	查到了一条记录
		if user.CheckPassword(loginRequest.Password) {
			resultBool = true
		}else{
			return resultBool, user, fmt.Errorf("用户名或密码错误！")
		}

		return resultBool, user, nil
	}

}

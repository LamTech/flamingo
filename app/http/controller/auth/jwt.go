package api

import (
	myjwt "flamingo/app/http/middleware/jwt"
	"flamingo/database/model"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// LoginReq 登录请求参数类
type LoginReq struct {
	Phone    string `json:"mobile"`
	Password string `json:"password"`
}

func GetDataByTime(c *gin.Context) {
	isPass := c.GetBool("isPass")
	if !isPass {
		return
	}
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

type LoginResult struct {
	Token string `json:"token"`
	model.User
}

// 登录
func Login(c *gin.Context) {
	var loginReq LoginReq
	if c.BindJSON(&loginReq) == nil {
		isPass, user, err := LoginCheck(loginReq)
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败" + err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
		return
	}

}


// CheckUser 检查用户是否存在
func CheckUser(phone string) bool {
	result := false
	return result
}



// LoginCheck 登录验证
func LoginCheck(loginReq LoginReq) (bool, model.User, error) {

	resultUser := model.User{}
	resultBool := false

	err := fmt.Errorf("")

	return resultBool, resultUser, err
}


// Register 注册用户，先检查是否存在用户，如果没有则注册
func Register(phone string, password string) error {
	if CheckUser(phone) {
		return fmt.Errorf("用户已存在！")
	}

	return fmt.Errorf("")
}

// 生成令牌
func generateToken(c *gin.Context, user model.User) {
	j := &myjwt.JWT{
		[]byte("flamingo"),
	}

	claims := myjwt.CustomClaims{
		user.Model.ID,
		user.UserName,
		user.Mobile,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "flamingo",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

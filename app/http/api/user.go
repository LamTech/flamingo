package api

import (
	mjwt "flamingo/app/http/middleware/jwt"
	"flamingo/database/model"
	"log"
	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 注册信息
type RegistInfo struct {
	Mobile string `json:"mobile"`
	PassWord   string `json:"password" binding:"required,min=6,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=40"`
}

// Register 注册用户
func RegisterUser(c *gin.Context) {
	var registerInfo RegistInfo
	if bindErr := c.BindJSON(&registerInfo); bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    bindErr.Error(),
		})

		return
	}

	err := model.Register(registerInfo.Mobile, registerInfo.PassWord)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "注册成功！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "注册失败：" + err.Error(),
		})
	}

}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	model.User
}

// Login 登录
func Login(c *gin.Context) {
	var loginReq model.LoginRequest
	if c.BindJSON(&loginReq) == nil {
		isPass, user, err := model.LoginCheck(loginReq)
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败," + err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
	}
}

// 生成令牌
func generateToken(c *gin.Context, user model.User) {
	j := &mjwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := mjwt.CustomClaims{
		user.Name,
		user.Name,
		user.Mobile,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
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

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*mjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}
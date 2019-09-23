package api

import (
	mjwt "flamingo/app/http/middleware/jwt"
	"flamingo/database/model"
	"flamingo/util/response"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"time"
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
	if bindErr := c.ShouldBindJSON(&registerInfo); bindErr != nil {
		response.JsonError(c,response.RequireParam,bindErr.Error())
		return
	}

	err := model.Register(registerInfo.Mobile, registerInfo.PassWord)
	if err == nil {
		response.JsonSuccess(c,nil)
	} else {
		response.JsonError(c,response.MobileExist,err.Error())
	}

}

// LoginResult 登录结果结构
type LoginResult struct {
	LoginUser
	AccessToken string `json:"access_token"`
}

type LoginUser struct {
	UniqueId uuid.UUID
	Name     string `json:"name"`
	Gender   uint   `json:"gender"`
	Mobile   string `json:"mobile"`
}

// Login 登录
func Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if parseErr := c.ShouldBindJSON(&loginRequest); parseErr != nil {
		response.JsonError(c,response.ParseJsonError,parseErr.Error())
		return
	}

	isPass, user, err := model.LoginCheck(loginRequest)
	if isPass {
		generateToken(c, user)
	} else {
		response.JsonError(c,response.ParseJsonError,err.Error())
		return
	}
}

// 生成令牌
func generateToken(c *gin.Context, user model.User) {
	j := &mjwt.JWT{
		[]byte(os.Getenv("Flamingo")),
	}
	claims := mjwt.CustomClaims{
		user.UniqueId,
		user.Name,
		user.Mobile,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), 			// 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), 			// 过期时间一小时
			Issuer:    os.Getenv("Flamingo"),					//签名的发行者
		},
	}

	accessToken, err := j.CreateToken(claims)

	if err != nil {
		response.JsonError(c,response.CreateTokenError,err.Error())
		return
	}

	log.Println(accessToken)

	data := LoginResult{
		LoginUser:  BuildUser(user),
		AccessToken: accessToken,
	}

	response.JsonSuccess(c,data)
	return
}

func BuildUser(user model.User) LoginUser {
	return LoginUser{
		UniqueId : user.UniqueId,
		Name : user.Name,
		Gender : user.Gender,
		Mobile : user.Mobile,
	}
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
package router

import (
	"flamingo/app/http/api"
	"flamingo/app/http/middleware"
	"flamingo/app/http/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	r.POST("/api/login/account", api.Login)
	r.POST("/api/register", api.RegisterUser)

	// 路由
	v1 := r.Group("/api/v1")
	v1.Use(jwt.JWTAuth())
	{
		v1.GET("/userinfo", api.GetUserInfo)
	}
	return r
}

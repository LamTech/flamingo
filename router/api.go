package router

import (
	fctrl "flamingo/app/http/controller"
	cjwt "flamingo/app/http/controller/auth"
	"flamingo/app/http/middleware"
	mjwt "flamingo/app/http/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	r.POST("/login",cjwt.Login)
	r.POST("/register",cjwt.Register)

	// 路由
	v1 := r.Group("/api/v1")
	v1.Use(mjwt.JWTAuth())
	{
		v1.GET("ping", fctrl.Ping)
	}
	return r
}

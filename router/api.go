package router

import (
	fctrl "flamingo/app/http/controller"
	"flamingo/app/http/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", fctrl.Ping)
	}
	return r
}

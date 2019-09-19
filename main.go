package main

import (
	"flamingo/config"
	"flamingo/router"
)

func main(){
	config.Load()

	// 装载路由
	r := router.NewRouter()
	r.Run(":3000")
}
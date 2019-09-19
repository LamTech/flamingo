package config

import (
	"flamingo/database"
	//"flamingo/cache"
	//"flamingo/database"
	"flamingo/util"
	"github.com/joho/godotenv"
	"os"
)

func Load(){
	//	从本地读取环境变量
	godotenv.Load()
	//	设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))
	//	连接数据库
	database.Database()
	//	连接 redis
	//cache.Redis()

}

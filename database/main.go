package database

import (
	"flamingo/database/migrations"
	"flamingo/util"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database() {

	if os.Getenv("DB_USERNAME") == "" {
		util.Log().Panic("数据库用户名缺失！ \n")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		util.Log().Panic("数据库密码缺失！ \n")
	}

	if os.Getenv("DB_DATABASE") == "" {
		util.Log().Panic("数据库名称缺失！ \n")
	}

	databaseConfig := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(databaseConfig)
	db, err := gorm.Open("mysql", databaseConfig)
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功 \n", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	if os.Getenv("DB_MIGRATE") == "true" {
		migration()
	}

}

func migration() {
	// 迁移模式
	DB.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&migrations.User{})
}

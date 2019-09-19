package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName string
	Gender   string `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string
	Status   string
	Avatar   string `gorm:"size:1000"`
}

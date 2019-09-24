package migrations

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UniqueId uuid.UUID `gorm:"primary_key"`
	Name     string `json:"name"`
	Gender   uint `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

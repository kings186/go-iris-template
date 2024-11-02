package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique"`
	Password    string
	PhoneNumber string
	Gender      string // 假设性别用字符串表示，如 "male" 或 "female"
	// 创建时间和修改时间，GORM 会自动处理
	CreatedAt time.Time
	UpdatedAt time.Time
}

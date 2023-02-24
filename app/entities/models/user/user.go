package user

import (
	"gorm.io/gorm"
	"skeleton/bootstrap/database"
	"time"
)

type User struct {
	// 管理员ID
	Id int64 `gorm:"primaryKey"`
	// 用户名
	Username string
	// 密码
	Password string
	// 上次登陆IP
	LastLoginIp string
	// 上次登陆时间
	LastLoginAt time.Time `gorm:"autoCreateTime" `
	// 创建时间
	CreatedAt time.Time `gorm:"autoCreateTime" `
}

func (user User) Create() (User, *gorm.DB) {
	result := database.DB.Model(User{}).Create(&user)

	if result.Error != nil {
		return User{}, result
	}
	return user, result
}

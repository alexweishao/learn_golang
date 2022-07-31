package models

import "github.com/jinzhu/gorm"

// UserLogin 用户登录信息表
type UserLogin struct {
	gorm.Model
	UserName string `form:"username"`
	Password string `form:"password"`
}

// BlogMessage 信息表
type BlogMessage struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `gorm:"type:text"`
	Tag     string `json:"tag"`
}

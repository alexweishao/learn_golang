package models

import "github.com/jinzhu/gorm"

// UserLogin 用户登录信息表
type UserLogin struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}

// BlogMessage 信息表
type BlogMessage struct {
	gorm.Model
	Tag     string `form:"tag" gorm:"tag" json:"tag"`
	Title   string `form:"title" gorm:"title" json:"title"`
	Content string `form:"content" gorm:"content" json:"content""`
}

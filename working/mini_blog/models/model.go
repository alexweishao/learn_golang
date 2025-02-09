package models

// UserLogin 用户登录信息表
type UserLogin struct {
	//gorm.Model
	ID       int    `form:"id" gorm:"primary_key" json:"id"`
	UserName string `form:"username" gorm:"column:username" json:"username" binding:"required"`
	Password string `form:"password" gorm:"password" json:"password" binding:"required"`
}

// BlogMessage 信息表
type BlogMessage struct {
	//gorm.Model
	ID      int    `form:"id"  json:"id"`
	Tag     string `form:"tag" gorm:"tag" json:"tag" binding:"required"`
	Title   string `form:"title" gorm:"title" json:"title" binding:"required"`
	Content string `form:"content" gorm:"content" json:"content" binding:"required"`
	UID     int    `form:"uid"  json:"uid" binding:"required"`
}

// UserMessage 用户信息表
type UserMessage struct {
	ID          int    `form:"id" gorm:"primary_key" json:"id"`
	Name        string `form:"name" gorm:"name" json:"name"`
	Age         int    `form:"age" gorm:"age" json:"age"`
	PhoneNumber string `form:"phoneNumber" gorm:"phone_number" json:"phone_number"`
	Address     string `form:"address" gorm:"address" json:"address"`
	Email       string `form:"email" gorm:"email" json:"email"`
}

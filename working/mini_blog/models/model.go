package models

// UserLogin 用户登录信息表
type UserLogin struct {
	//gorm.Model
	ID       int    `form:"id" gorm:"primary_key" json:"id"`
	UserName string `form:"username" gorm:"username" json:"username"`
	Password string `form:"password" gorm:"password" json:"password"`
}

// BlogMessage 信息表
type BlogMessage struct {
	//gorm.Model
	ID      int    `form:"id"  json:"id"`
	Tag     string `form:"tag" gorm:"tag" json:"tag"`
	Title   string `form:"title" gorm:"title" json:"title"`
	Content string `form:"content" gorm:"content" json:"content""`
}

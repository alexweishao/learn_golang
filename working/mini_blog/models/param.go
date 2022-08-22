package models

type RegisterUser struct {
	UserName string `form:"username" binding:"required,gte=6,lte=20"`
	Password string `form:"password" binding:"required,gte=6,lte=18"`
}

type LoginUser struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AllArticleList struct {
	Page     int `form:"page" binding:"required"`
	LimitNum int `form:"limit_num" binding:"required"`
}

type UserArticleList struct {
	UID      int `form:"uid"  json:"uid" binding:"required"`
	Page     int `form:"page" binding:"required"`
	LimitNum int `form:"limit_num" binding:"required"`
}

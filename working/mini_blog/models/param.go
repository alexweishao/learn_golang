package models

type RegisterUser struct {
	UserName string `form:"username" binding:"required,gte=6,lte=20"`
	Password string `form:"password" binding:"required,gte=6,lte=18"`
}

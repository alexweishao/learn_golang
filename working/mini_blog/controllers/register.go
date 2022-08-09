package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

//到注册页面
func GoRegister(ctx *gin.Context) {

	//返回html
	ctx.HTML(http.StatusOK, "register.html", nil)
}

//处理注册
func Register(ctx *gin.Context) {
	db := dao.GetDB()

	//Bind()函数绑定数据到结构体
	var reqData models.UserLogin
	ctx.Bind(&reqData)
	//获取表单信息
	username := reqData.UserName
	password := reqData.Password

	//数据验证
	if len(username) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名需要11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	//判断用户名是否存在
	var user models.UserLogin
	db.Where("user_name = ?", username).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名已存在",
		})
		return
	}
	newUser := models.UserLogin{
		UserName: username,
		Password: password,
	}
	db.Create(&newUser) //将用户数据写入数据库表中
	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
	//ctx.Redirect(http.StatusMovedPermanently, "login")
}

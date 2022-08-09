package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func GoLogin(ctx *gin.Context) {
	//返回html
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func Login(ctx *gin.Context) {
	db := dao.GetDB()

	//获取表单信息

	var reqData models.UserLogin
	err := ctx.Bind(&reqData)
	if err != nil {
		return
	}
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

	//判断用户名+密码是否存在数据库中
	var user models.UserLogin
	db.Where("user_name=? and password=?", username, password).Find(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名或密码不正确",
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
	})
	//ctx.Redirect(http.StatusMovedPermanently, "/addarticles")
}

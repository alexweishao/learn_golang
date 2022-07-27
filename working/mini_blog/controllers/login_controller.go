package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini_blog/database"
	"mini_blog/models"
	"net/http"
)

func LoginGet(ctx *gin.Context) {
	//返回html
	ctx.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(ctx *gin.Context) {

	db := database.GetDB()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUserLogin models.UserLogin
	ctx.Bind(&requestUserLogin)
	//获取表单信息
	//username := requestUserLogin.UserName
	//password := requestUserLogin.Password
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println("username:", username)
	fmt.Println("password:", password)

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
		fmt.Println(user.ID)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	//判断密码是否正确
	/*if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}*/

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
	})
}

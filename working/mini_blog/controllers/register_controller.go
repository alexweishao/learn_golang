package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini_blog/database"
	"mini_blog/models"
	"net/http"
)

func RegisterGet(ctx *gin.Context) {

	//返回html
	ctx.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

//处理注册

func RegisterPost(ctx *gin.Context) {

	db := database.GetDB()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUserLogin models.UserLogin
	err := ctx.Bind(&requestUserLogin)
	if err != nil {
		return
	}
	//username := requestUserLogin.UserName
	//password := requestUserLogin.Password
	//获取表单信息
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username, password)

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
			"message": "用户已存在",
		})
		return
	}

	//密码加密
	/*createPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}*/

	//创建用户
	newUserLogin := models.UserLogin{
		UserName: username,
		Password: password,
		//Password: string(createPassword),
	}
	db.Create(&newUserLogin)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})

}

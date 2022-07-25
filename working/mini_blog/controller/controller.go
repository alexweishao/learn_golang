package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToLogin(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", nil) //登录页面模板渲染

}

// DoLogin  登录
func DoLogin(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username, password)

	ctx.Redirect(http.StatusFound, "/index.html")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录成功"})

}

func Redirect(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// ToRegister 注册
func ToRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil) //注册页面模板渲染
}

// DoRegister 注册
func DoRegister(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//数据验证
	if len(username) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名必须为11位",
		})
		return
	}
	if len(password) != 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码必须为6位",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "注册成功"})
}

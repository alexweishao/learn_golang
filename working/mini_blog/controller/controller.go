package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToLogin(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", gin.H{"title": "登录"}) //登录页面模板渲染

}

// DoLogin  登录
func DoLogin(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username, password)
	ctx.String(http.StatusOK, "添加成功")

}

// Register 注册
func Register(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "login.html")        //注册完成 点击注册按钮 跳转到登录页面
	ctx.HTML(http.StatusOK, "register.html", gin.H{"title": "注册"}) //注册页面模板渲染
}

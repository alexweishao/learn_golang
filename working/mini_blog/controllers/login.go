package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"mini_blog/util"
	"net/http"
)

func Login(ctx *gin.Context) {
	db := dao.GetDB()

	//获取表单信息

	var reqData models.LoginUser
	err := ctx.ShouldBind(&reqData)
	if err != nil {
		returnErr(ctx, err, "参数错误")
		return
	}

	username := reqData.UserName
	pwd := util.GetMd5String(reqData.Password)
	//id := reqData.ID

	//判断用户名+密码是否存在数据库中
	var user models.UserLogin
	db.Table("user_login").Where("username=? and password=?", username, pwd).Find(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名或密码不正确",
		})
		return
	}
	id := user.ID
	fmt.Println(user)
	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"id":      id,
	})
}

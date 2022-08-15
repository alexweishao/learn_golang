package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
	"strings"
)

func UserMessage(ctx *gin.Context) {
	var userMessage models.UserMessage
	ctx.Bind(&userMessage)

	name := userMessage.Name
	age := userMessage.Age
	phoneNumber := userMessage.PhoneNumber
	address := userMessage.Address
	email := userMessage.Email

	//手机号和邮箱格式验证
	if len(phoneNumber) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号码需要11位",
		})
		return
	}

	if !strings.Contains(email, "@") {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "邮箱号码需包含@字符",
		})
		return
	}
	dao.DB.Where("name = ?", name).First(&userMessage)
	if userMessage.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号已存在",
		})
		return
	}

	newUser := models.UserMessage{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		Address:     address,
		Email:       email,
	}
	dao.DB.Create(&newUser) //将用户数据写入数据库表中
	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "个人信息登记成功",
	})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini_blog/dao"
	"mini_blog/models"
	"mini_blog/util"
	"net/http"
)

//处理注册
func Register(ctx *gin.Context) {
	db := dao.GetDB()

	//Bind()函数绑定数据到结构体
	var reqData models.RegisterUser
	err := ctx.ShouldBind(&reqData)
	if err != nil {
		returnErr(ctx, err, "入参错误")
		return
	}
	//获取表单信息
	username := reqData.UserName
	password := reqData.Password

	//数据验证
	if len(username) < 6 || len(username) > 20 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名需要小于6 大于20 位",
		})
		return
	}
	if len(password) < 6 || len(username) > 18 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码必须大于等6位小于等18位",
		})
		return
	}

	//判断用户名是否存在
	var user models.UserLogin
	err = db.Table("user_login").Where("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	if err != nil {
		returnErr(ctx, err, "db.Where(\"user_name = ?\", username).First(&user).Error ")
		return
	}
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名已存在",
		})
		return
	}
	pwd := util.GetMd5String(password)
	newUser := models.UserLogin{
		UserName: username,
		Password: pwd,
	}
	err = db.Table("user_login").Create(&newUser).Error //将用户数据写入数据库表中
	if err != nil {
		returnErr(ctx, err, "db.Create(&newUser).Error ")
		return
	}
	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

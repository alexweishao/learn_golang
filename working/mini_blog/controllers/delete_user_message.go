package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func DeleteUserMessage(ctx *gin.Context) {
	dao.GetDB()
	id := ctx.Param("id")
	var LoginMessage models.UserLogin
	num := dao.DB.Debug().Where("id=?", id).Delete(&LoginMessage).RowsAffected
	if num == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "该用户不存在！",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户信息已成功删除",
	})
}

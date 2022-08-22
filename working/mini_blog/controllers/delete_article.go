package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func DeleteArticle(ctx *gin.Context) {
	db := dao.GetDB()

	id := ctx.Param("id")
	var reqDelete models.BlogMessage //reqDelete是结构体
	//var articles []models.BlogMessage

	num := db.Table("blog_message").Where("id = ?", id).Delete(&reqDelete).RowsAffected
	if num == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "delete fail!",
			"data": "id不存！",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "delete success.",
		//"data": articles,
	})
}

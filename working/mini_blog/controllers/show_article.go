package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func ShowArticle(ctx *gin.Context) {

	db := dao.GetDB()

	var articles []models.BlogMessage

	//Debug()返回SQL语句

	db.Find(&articles)
	if err := db.Find(&articles).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "fail",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": articles,
		})
	}
}

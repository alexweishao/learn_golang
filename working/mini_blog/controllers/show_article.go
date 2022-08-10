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
	if err := db.Find(&articles); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"msg":     "success",
			"message": "blog发布成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "success",
			"data": articles,
		})
	}

	ctx.JSON(200, gin.H{
		"articles": articles,
		"message":  "查询成功",
	})
}

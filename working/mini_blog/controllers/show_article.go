package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
)

func ShowArticle(ctx *gin.Context) {

	db := dao.GetDB()
	var allArticle models.BlogMessage
	db.Find(&allArticle)
	fmt.Println(allArticle)

	//返回结果

}

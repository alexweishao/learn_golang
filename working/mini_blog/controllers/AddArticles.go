package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func GoAddArticle(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "addarticles.html", "nil")
}

func AddArticle(ctx *gin.Context) {

	db := dao.GetDB()
	var reqArticleMessage models.BlogMessage
	err := ctx.BindJSON(&reqArticleMessage)
	if err != nil {
		return
	}
	title := reqArticleMessage.Title
	tag := reqArticleMessage.Tag
	content := reqArticleMessage.Content

	newMessage := models.BlogMessage{
		Title:   title,
		Tag:     tag,
		Content: content,
	}
	db.Create(&newMessage) //将用户数据写入数据库表中

	ctx.Redirect(http.StatusMovedPermanently, "index")
}

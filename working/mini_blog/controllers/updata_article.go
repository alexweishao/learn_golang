package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func UpdateArticle(ctx *gin.Context) {
	db := dao.GetDB()
	var reqArticleMessage models.BlogMessage
	err := ctx.Bind(&reqArticleMessage)
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

	db.First(&newMessage)
	newMessage.Title = title
	newMessage.Tag = tag
	newMessage.Content = content
	db.Model(&newMessage).Update(title, tag, content)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

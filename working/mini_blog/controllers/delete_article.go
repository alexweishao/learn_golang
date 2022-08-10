package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func DeleteArticle(ctx *gin.Context) {
	db := dao.GetDB()
	db.Where("id = ?", 1).Delete(&models.BlogMessage{})

	db.Delete(&models.BlogMessage{}, "tag = ?", 001)
	/*var reqArticleMessage models.BlogMessage
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
	db.Delete(&newMessage)*/

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}

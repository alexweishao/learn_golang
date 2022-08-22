package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func PublishArticle(ctx *gin.Context) {

	db := dao.GetDB()
	var reqArticleMessage models.BlogMessage
	err := ctx.Bind(&reqArticleMessage)
	if err != nil {
		return
	}
	title := reqArticleMessage.Title
	tag := reqArticleMessage.Tag
	content := reqArticleMessage.Content
	uid := reqArticleMessage.UID

	newMessage := models.BlogMessage{
		Title:   title,
		Tag:     tag,
		Content: content,
		UID:     uid,
	}

	if err = db.Table("blog_message").Create(&newMessage).Error; err != nil { //newMessage 字段插入db表中
		returnErr(ctx, err, "db.Create(&newUser).Error ")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}

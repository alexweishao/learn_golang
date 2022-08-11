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

	newMessage := models.BlogMessage{
		Title:   title,
		Tag:     tag,
		Content: content,
	}
	//db.Create(&newMessage) 将blog内容写入数据库blog_messages表中

	if err = db.Debug().Create(&newMessage).Error; err != nil { //newMessage 字段插入db表中
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "fail",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

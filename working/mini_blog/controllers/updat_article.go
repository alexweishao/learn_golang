package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func UpdateArticle(ctx *gin.Context) {

	db := dao.GetDB()
	tag, ok := ctx.Params.Get("tag")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "tag不存在",
		})
		return
	}
	var article models.BlogMessage

	if err := db.Where("tag=?", tag).First(&article).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := ctx.BindJSON(&article)
	if err != nil {
		return
	}

	title := article.Title
	tag = article.Tag
	content := article.Content

	newMessage := models.BlogMessage{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	if err := db.Save(&newMessage).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1004,
			"msg":     "success",
			"message": "blog发布失败",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "success",
			"data": article,
		})
	}
}

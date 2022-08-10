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

	//判断文章tag是否重复
	/*var articleTag models.BlogMessage
	db.Where("tag=?", tag).Find(&articleTag)
	if articleTag.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "文章标签重复",
		})
		return
	}*/

	newMessage := models.BlogMessage{
		Title:   title,
		Tag:     tag,
		Content: content,
	}
	//db.Create(&newMessage) //将blog内容写入数据库blog_messages表中

	if err = db.Create(&newMessage).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"msg":     "success",
			"message": "blog发布成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "success",
			"data": newMessage,
		})
	}
	//返回结果
	/*ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "blog发布成功",
	})*/
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

type message models.BlogMessage

func GetAMessage(id string) (blogmessage *message, err error) {
	blogmessage = new(message)
	dao.DB.Where("id=?", id).First(blogmessage)
	return
}

func UpdateAMessage(blogmessage *message) (err error) {
	err = dao.DB.Save(blogmessage).Error
	return
}

func UpdateArticle(ctx *gin.Context) {

	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	blogmessage, err := GetAMessage(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Bind(&blogmessage)
	if err = UpdateAMessage(blogmessage); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, blogmessage)
	}

	/*db := dao.GetDB()

	var article models.BlogMessage

	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "id invalid!",
		})
		return
	}

	ctx.Bind(&article)

	messageId := article.ID
	newMessage := models.BlogMessage{
		ID: messageId,
	}

	if err := db.Where("id=?", id).First(&article).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Save(&newMessage).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1004,
			"msg":     "success",
			"message": "update fail！",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "update success.",
			"data": newMessage,
		})
	}*/
}

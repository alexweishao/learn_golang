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
	db.Create(&newMessage) //将blog内容写入数据库blog_messages表中

	/*//判断用户名+密码是否存在数据库中
	var user models.UserLogin
	db.Where("user_name=? and password=?", username, password).Find(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名或密码不正确",
		})
		return
	}*/

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "blog数据写入成功",
	})
}

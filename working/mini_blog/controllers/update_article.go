package controllers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
	"strconv"
)

func UpdateArticle(ctx *gin.Context) {
	db := dao.GetDB()

	id := ctx.Param("id") //获取前端指定的id
	var article models.BlogMessage
	db.Table("blog_message").Where("id=?", id).First(&article)
	//SELECT * FROM `blog_messages`  WHERE (id='2') ORDER BY `blog_messages`.`id` ASC LIMIT 1

	err := ctx.Bind(&article)
	if err != nil {
		return
	}
	id = strconv.Itoa(article.ID)
	title := article.Title
	tag := article.Tag
	content := article.Content
	uid := article.UID

	newMessage := models.BlogMessage{
		ID:      article.ID,
		Title:   title,
		Tag:     tag,
		Content: content,
		UID:     uid,
	}

	if err := db.Table("blog_message").Save(&newMessage).Error; err != nil {
		returnErr(ctx, err, "db.Table('blog_message').Save(&newMessage).Error; err != nil")
		/*ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"msg":     "success",
			"message": "update fail！",
		})*/
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "update success.",
		"data": article,
	})

}

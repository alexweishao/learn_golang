package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
	"strconv"
)

func UpdateArticle(ctx *gin.Context) {

	db := dao.GetDB()

	id := ctx.Param("id") //获取前端指定的id
	fmt.Println(id)
	var article models.BlogMessage
	db.Debug().Where("id=?", id).First(&article)
	//SELECT * FROM `blog_messages`  WHERE (id='2') ORDER BY `blog_messages`.`id` ASC LIMIT 1

	/*ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}*/
	ctx.ShouldBind(&article)

	err := ctx.Bind(&article)
	if err != nil {
		return
	}
	id = strconv.Itoa(article.ID)
	title := article.Title
	tag := article.Tag
	content := article.Content

	newMessage := models.BlogMessage{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	if err := db.Save(&newMessage).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"msg":     "success",
			"message": "update fail！",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "update success.",
			"data": article,
		})
	}
}

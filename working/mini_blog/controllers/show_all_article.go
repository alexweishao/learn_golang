package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mini_blog/dao"
	"mini_blog/models"
	"net/http"
)

func ShowAllArticle(ctx *gin.Context) {

	db := dao.GetDB()

	var reqData models.AllArticleList
	err := ctx.ShouldBind(&reqData)
	if err != nil {
		returnErr(ctx, err, "入参错误")
		return
	}

	page := reqData.Page
	if page <= 0 {
		page = 1
	}
	limitNum := reqData.LimitNum
	if limitNum <= 0 || limitNum > 10 {
		limitNum = 10
	}

	offset := (page - 1) * limitNum
	var data []models.BlogMessage
	err = db.Table("blog_message").
		Order("id DESC").Offset(offset).Limit(limitNum).Find(&data).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	if err != nil {
		returnErr(ctx, err, "查询数据异常")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "success",
		"list":      data,
		"page":      page,
		"limit_num": limitNum,
	})

}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {

	//返回html
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"code":  200,
		"title": "初始页面",
	})
}

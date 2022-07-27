package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {

	//返回html
	ctx.HTML(http.StatusOK, "register.html", gin.H{
		"code":  200,
		"title": "注册页",
	})
}

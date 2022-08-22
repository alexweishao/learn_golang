package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func returnErr(ctx *gin.Context, err error, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    123,
		"message": message,
		"data":    err,
	})
	return
}

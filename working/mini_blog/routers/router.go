package routers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/controllers"
)

func ParentRouter(router *gin.Engine) {
	router.GET("/", controllers.Index) //默认接口为注册页面

	//router.GET("/login", controllers.GoLogin) //登录GET接口
	router.POST("/login", controllers.Login) //登录POST接口

	//router.GET("/register", controllers.GoRegister) //注册GET接口
	router.POST("/register", controllers.Register) //注册POST接口

}

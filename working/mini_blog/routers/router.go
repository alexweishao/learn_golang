package routers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/controllers"
)

func ParentRouter(router *gin.Engine) {

	router.LoadHTMLGlob("views/*") //前端页面模板解析

	router.Static("/static", "./static")    //加载静态文件
	router.GET("/", controllers.GoRegister) //默认接口为注册页面

	router.GET("/login", controllers.GoLogin) //登录GET接口
	router.POST("/login", controllers.Login)  //登录POST接口

	router.GET("/register", controllers.GoRegister) //注册GET接口
	router.POST("/register", controllers.Register)  //注册POST接口

	router.GET("/index", controllers.Index) //首页

}

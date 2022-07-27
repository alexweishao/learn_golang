package routers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/controllers"
)

func ParentRouter(router *gin.Engine) {
	router.GET("/", controllers.Index) //默认接口为注册页面

	router.GET("/login", controllers.LoginGet)   //登录GET接口
	router.POST("/login", controllers.LoginPost) //登录POST接口

	router.GET("/register", controllers.RegisterGet)   //注册GET接口
	router.POST("/register", controllers.RegisterPost) //注册POST接口

	//router.GET("/user_index_page", controllers.UserIndexPage) //注册接口
	//router.GET("/publish_message", controllers.PublishMessage)            //发布消息
	//router.GET("/all_message_list", controllers.AllMessageList)           //全部消息列表
	//router.GET("/personal_message_list", controllers.PersonalMessageList) //个人消息列表
	//router.GET("/message_detail", controllers.MessageDetail)              //消息详情
}

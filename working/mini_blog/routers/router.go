package routers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/controller"
)

func ParentRouter(router *gin.Engine) {
	router.GET("/", controller.Register) //默认接口为注册页面

	router.GET("/to_login", controller.ToLogin)  //登录接口
	router.POST("/do_login", controller.DoLogin) //登录接口

	router.GET("/register", controller.Register) //注册接口

	//router.GET("/user_index_page", controller.UserIndexPage) //注册接口
	//router.GET("/publish_message", controller.PublishMessage)            //发布消息
	//router.GET("/all_message_list", controller.AllMessageList)           //全部消息列表
	//router.GET("/personal_message_list", controller.PersonalMessageList) //个人消息列表
	//router.GET("/message_detail", controller.MessageDetail)              //消息详情
}

package routers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/controller"
)

func ParentRouter(router *gin.Engine) {
	router.GET("/", controller.ToRegister) //默认接口为注册页面

	router.GET("/to_login", controller.ToLogin)  //登录接口
	router.POST("/do_login", controller.DoLogin) //登录接口
	router.GET("/redirect", controller.Redirect)

	router.GET("/to_register", controller.ToRegister)  //登录接口
	router.POST("/do_register", controller.DoRegister) //注册接口

	//router.GET("/user_index_page", controller.UserIndexPage) //注册接口
	//router.GET("/publish_message", controller.PublishMessage)            //发布消息
	//router.GET("/all_message_list", controller.AllMessageList)           //全部消息列表
	//router.GET("/personal_message_list", controller.PersonalMessageList) //个人消息列表
	//router.GET("/message_detail", controller.MessageDetail)              //消息详情
}

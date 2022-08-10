package routers

import (
	"github.com/gin-gonic/gin"
	"mini_blog/controllers"
)

func ParentRouter(router *gin.Engine) {

	/*//router.LoadHTMLGlob("views/*") //前端页面模板解析
	//router.Static("/static", "./static")    //加载静态文件
	//router.GET("/", controllers.Register) //默认接口为注册页面
	//router.GET("/login", controllers.GoLogin) //登录GET接口
	//router.GET("/register", controllers.GoRegister) //注册GET接口
	//router.GET("/addArticle", controllers.GoAddArticle)
	//router.POST("/toaddArticle", controllers.AddArticle)
	*/
	router.POST("/register", controllers.Register)               //注册POST接口
	router.POST("/login", controllers.Login)                     //登录POST接口
	router.POST("/publish_articles", controllers.PublishArticle) //发布文章接口
	router.POST("/delete_articles", controllers.DeleteArticle)   //删除文章接口
	router.POST("/update_articles", controllers.UpdateArticle)   //更新文章接口
	router.POST("/show_articles", controllers.ShowArticle)       //更新文章接口

	router.POST("/get_all_articles", controllers.Register)   //获取全部文章接口
	router.POST("/get_person_message", controllers.Register) //获取个人消息列表
	router.POST("/message_detail", controllers.Register)     //消息详情
}

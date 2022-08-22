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
	//router.POST("/toadArticle", controllers.AddArticle)
	*/
	router.POST("/register", controllers.Register)                           //注册POST接口
	router.POST("/login", controllers.Login)                                 //登录POST接口
	router.POST("/user_message", controllers.UserMessage)                    //登记个人信息接口
	router.DELETE("/delete_user_message/:id", controllers.DeleteUserMessage) //删除用户注册信息接口
	router.POST("/publish_articles", controllers.PublishArticle)             //发布文章接口
	router.GET("/all/article/list", controllers.ShowAllArticle)              //查看所有文章列表接口
	router.GET("/user/article/list/", controllers.ShowUserArticle)           //查看用户文章列表接口
	router.DELETE("/delete_articles/:id", controllers.DeleteArticle)         //删除文章接口
	router.PUT("/update_articles/:id", controllers.UpdateArticle)            //更新文章接口

}

package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mini_blog/database"
	"mini_blog/routers"
)

func main() {
	//获取初始化的数据库
	db := database.InitMySQL()
	//延迟关闭数据库连接，防止内存泄漏
	defer db.Close()

	//默认路由引擎
	router := gin.Default()

	//前端页面模板解析
	router.LoadHTMLGlob("views/*")
	//加载静态文件
	router.Static("/static", "./static")

	//路由分离
	routers.ParentRouter(router)

	//启动http服务 默认端口号为127.0.0.1:8080
	router.Run(":8080")
}

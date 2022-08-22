package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mini_blog/dao"
	"mini_blog/routers"
)

func main() {
	db := dao.InitDB() //获取初始化数据库
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db) //关闭数据库连接
	router := gin.Default() //默认路由引擎

	routers.ParentRouter(router) //路由分离

	err := router.Run(":8086") //启动http服务 默认端口号为127.0.0.1:8080
	if err != nil {
		panic(err)
	}
}

package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"mini_blog/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB { //定义数据库参数
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "weibo"
	username := "root"
	password := "kk123456yy"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset) //字段拼接
	db, err := gorm.Open(driverName, args) //连接数据库
	if err != nil {
		log.Fatal("连接数据库失败，err:" + err.Error())
	}
	db.AutoMigrate(&models.UserLogin{}, &models.BlogMessage{}) //自动创建表 UserLogin BlogMessage

	DB = db

	return db
}
func GetDB() *gorm.DB {
	return DB
}

package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"mini_blog/models"
)

var DB *gorm.DB

func InitMySQL() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "weibo"
	username := "root"
	password := "kk123456yy"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		log.Fatal("连接数据库失败，err:" + err.Error())
	}
	/*方式二 gorm 连接数据库
	//用户名：密码@tcp（ip:port）/数据库名？charset=utf8&parseTime=true&loc=Local
	db, err := gorm.Open("mysql", "root:kk123456yy@tcp(localhost:3306)/weibo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接，防止内存泄漏*/

	db.AutoMigrate(&models.UserLogin{}) //自动创建表

	/*db.Create(&models.UserLogin{UserName: "zhangSan", Password: "1000"}) //插入数据

	var user models.UserLogin
	db.First(&user, 1) // 查询id为1的product
	//db.First(&user, "username = ?", "zhangSan") // 查询username为zhangSan的user
	fmt.Println(user)
	db.Model(&user).Update("password", "2000")
	fmt.Println(user)*/
	/*db.Delete(&user)
	fmt.Println(user)*/
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}

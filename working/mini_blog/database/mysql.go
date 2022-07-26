package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// UserLoginMessage 用户登录信息表
type UserLoginMessage struct {
	Id         uint
	UserName   string
	Password   string
	Status     int
	Createtime time.Time
}

func InitMySQL() (err error) {
	//用户名：密码@tcp（ip:port）/数据库名？charset=utf8&parseTime=true&loc=Local
	db, err := gorm.Open("mysql", "root:kk123456yy@tcp(localhost:3306)/weibo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接，防止内存泄漏

	db.AutoMigrate(&UserLoginMessage{}) //自动创建表

	db.Create(&UserLoginMessage{Id: 1, UserName: "zhangSan", //插入数据
		Password: "1000", Status: 0, Createtime: time.Now()})

	var user UserLoginMessage
	db.First(&user, 1) // 查询id为1的product
	//db.First(&user, "username = ?", "zhangSan") // 查询username为zhangSan的user
	fmt.Println(user)
	db.Model(&user).Update("password", "2000")
	fmt.Println(user)
	/*db.Delete(&user)
	fmt.Println(user)*/

	return
}

package models

import (
	"github.com/jinzhu/gorm"
)

// User 用户基础信息表
type User struct {
	Id    uint   `gorm:"primary_key" json:"id" binding:"required" db:"id"`
	Name  string `json:"name" db:"name"`
	Age   int    `json:"age" db:"age"`
	Adder string `json:"adder" db:"adder"`
	Pic   string `json:"pic" db:"pic"`
}

// Message 用户发送消息表
type Message struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

// UserLoginMessage 用户登录信息表
type UserLoginMessage struct {
	Id       uint   `gorm:"primary_key" json:"id" binding:"required" db:"id"`
	UserName string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password"`
}

func InitMySQL() (err error) {
	//用户名：密码@tcp（ip:port）/数据库？charset=utf8&parseTime=true&loc=Local
	db, err := gorm.Open("mysql", "root:kk123456yy@tcp(localhost:3306)/weibo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接，防止内存泄漏

	/*//创建三个表  方式一
	if !db.HasTable(&User{}) {
	  	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 user '用户基础信息表'").
	  		CreateTable(&User{}).Error; err != nil {
	  		panic(err)
	  	}
	  }
	  if !db.HasTable(&UserLoginMessage{}) {
	  	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 userLoginMessage '用户登录信息表'").
	  		CreateTable(&UserLoginMessage{}).Error; err != nil {
	  		panic(err)
	  	}
	  }
	  if !db.HasTable(&models.Message{}) {
	  	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 message '用户发送消息表'").
	  		CreateTable(&models.Message{}).Error; err != nil {
	  		panic(err)
	  	}
	  }

	//方式二
	//db.CreateTable(&User{}) //表名称为users
	//db.CreateTable(&UserLoginMessage{})
	//db.CreateTable(&Message{})
	//指定表名称
	//db.Table("user").CreateTable(&User{}) //创建表名为user

	//删除表
	//db.DropTable("user") //或者下面方式
	//db.DropTable(&User{})

	//检查表是否存在
	b := db.HasTable("users")
	  fmt.Println(b)
	  b2 := db.HasTable(&User{})
	  fmt.Println(b2)


	//方式三  自动迁移表  仅仅会创建表 并添加缺少的列和索引，不会改变现有列的类型或删除未使用的列以保护数据
	db.AutoMigrate(&User{})
	  db.AutoMigrate(&UserLoginMessage{})
	  db.AutoMigrate(&Message{})*/

	db.AutoMigrate(&User{}, &UserLoginMessage{}, &Message{}) //自动创建表

	//创建表时添加表后缀
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &UserLoginMessage{}, &Message{})

	return
}

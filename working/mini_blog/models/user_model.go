package models

import "github.com/jinzhu/gorm"

// UserLogin 用户登录信息表
type UserLogin struct {
	gorm.Model
	UserName string `gorm:"varchar(20);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

/*type UserLoginMessage struct {
	gorm.Model
	Id       int
	Username string
	Password string
	//Status   int // 0 正常状态， 1删除
}*/

//--------------数据库操作-----------------

//插入
/*func InsertUser(user UserLoginMessage) (int64, error) {
	return UserLoginMessage{0, user.Username, user.Password, user.Status, user.Createtime}, nil
	/*database.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
	user.Username, user.Password, user.Status, user.Createtime)
}*/

/*//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	//row := database.QueryRowDB(sql)
	id := 1
	//row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}*/

package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
)

func SignUp() {
	// 1.判断用户是否存在
	mysql.QueryUserByUsername()
	// 2. 生成UID
	snowflake.GenID()
	//保存到数据库
	mysql.SignUp()

	//
}

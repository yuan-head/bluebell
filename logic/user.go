package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2. 生成UID
	userID := snowflake.GenID()

	// 3. 构造User对象
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password, // 密码需要加密处理
	}
	//保存到数据库
	return mysql.InsertUser(user)
}

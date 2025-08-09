package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "wainlin_study"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// 把每一步数据库操作都封装成一个函数
// 待logic层根据业务需求调用
func CheckUserExist(username string) (err error) {
	sqlStr := `SELECT COUNT(user_id) FROM user WHERE username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 用于将用户信息插入到数据库中
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)

	// 执行sql语句入库
	sqlStr := `INSERT INTO user (user_id, username, password) VALUES (?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	if err != nil {
		return err
	}
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(oPassword + secret))
	// 计算哈希值并转换为十六进制字符串
	return hex.EncodeToString(h.Sum(nil))
}

func Login(user *models.User) (err error) {
	// 1. 判断用户是否存在
	oPassword := user.Password
	sqlStr := `SELECT user_id, username, password FROM user WHERE username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// 2. 验证密码
	// 如果用户存在，user.Password已经被赋值为数据库中的密码
	// 注意：此时user.Password是加密后的密码
	// 因此需要对输入的密码进行加密后再进行比较
	password := encryptPassword(oPassword)
	if user.Password != password {
		return ErrorInvalidPassword
	}
	// 3. 登录成功，返回nil
	return

}

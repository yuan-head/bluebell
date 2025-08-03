package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "wainlin_study"

// 把每一步数据库操作都封装成一个函数
// 待logic层根据业务需求调用
func CheckUserExist(username string) (err error) {
	sqlStr := `SELECT COUNT(user_id) FROM user WHERE username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
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

package controller

import (
	"bluebell/middlerwares"
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 从上下文中获取当前用户的ID
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlerwares.CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

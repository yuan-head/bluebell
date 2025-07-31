package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// 处理注册请求函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数绑定失败，返回错误信息
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	fmt.Println("SignUpHandler p:", p)
	fmt.Println(p)

	// 2. 业务处理
	logic.SignUp()
	fmt.Println("test")
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

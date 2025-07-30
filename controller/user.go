package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理注册请求函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验

	// 2. 业务处理
	logic.SignUp()
	// 3. 返回响应
	c.JSON(http.StatusOK, "ok")
}

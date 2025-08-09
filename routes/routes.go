package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlerwares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(zap.L()), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", middlerwares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录的用户，判断请求头中是否有 有效的 JWT Token
		// 把认证操作封装到 JWTAuthMiddleware 中间件中
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "not found",
		})
		zap.L().Error("404 not found", zap.String("path", c.Request.URL.Path))
	})
	return r
}

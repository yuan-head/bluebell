package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(zap.L()), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "not found",
		})
		zap.L().Error("404 not found", zap.String("path", c.Request.URL.Path))
	})
	return r
}

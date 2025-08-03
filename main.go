package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"bluebell/settings"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("load config failed, err:#{err}")
		return
	}
	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Println("init logger failed, err:#{err}")
		return
	}
	defer zap.L().Sync() // 确保日志在程序退出前被刷新
	zap.L().Debug("logger init success...")
	// 3. 初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql failed, err:#{err}")
		return
	}
	defer mysql.Close() // 确保在程序退出时关闭MySQL连接
	// 4. 初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("init redis failed, err:#{err}")
		return
	}
	defer redis.Close() // 确保在程序退出时关闭Redis连接
	// 初始化翻译器

	err := controller.InitTrans("zh")
	if err != nil {
		fmt.Println("init trans failed, err:#{err}")
		return
	} // 初始化翻译器，设置语言为中文
	// 初始化雪花算法
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil { // 初始化雪花算法
		fmt.Printf("init snowflake failed, err: %v\n", err)
		return
	}
	// 5. 注册路由
	r := routes.SetupRouter()
	// 6. 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")), // 从配置文件中读取端口号
		Handler: r,
	}
	fmt.Println("🚀 Listening on port:", viper.GetInt("port"))
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("server start failed", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIG
	// kill -9 发送 syscall.SIGKILL 信号，这个信号是无法被捕获的
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞，直到接收到信号
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保在函数退出时取消上下文
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server close failed", zap.Error(err))
	} else {
		zap.L().Info("Server exited gracefully")
	}
}

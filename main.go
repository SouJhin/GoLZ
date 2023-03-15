package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"server/dao/mysql"
	"server/dao/redis"
	"server/routes"
	"server/settings"
)

func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return
	}
	//初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//mysql
	if err := mysql.Init(); err != nil {
		fmt.Printf("mysql init failed =====> 🚀🚀🚀 %v\n", err)
	}
	defer mysql.Close()
	//redis
	if err := redis.Init(); err != nil {
		fmt.Printf("mysql init failed =====> 🚀🚀🚀 %v\n", err)
	}
	defer redis.Close()
	//注册路由
	r := routes.SetUp()
	//启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Singal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}

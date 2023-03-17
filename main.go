package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"

	"syscall"
	"time"

	"server/controllers"
	"server/dao/mysql"
	"server/dao/redis"
	"server/logger"
	"server/pkg/snowflake"
	"server/routes"
	"server/settings"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"os"
)

func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("err =====> 🚀27🚀🚀 %v\n", err)
		return
	}
	//初始化日志
	if err := logger.Init(settings.Conf.LogConfig, "dev"); err != nil {
		fmt.Printf("err =====> 🚀🚀32🚀 %v\n", err)
		return
	}
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {

		}
	}(zap.L())
	zap.L().Debug("logger init success...")
	//mysql
	if err := mysql.InitDB(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("mysql init failed =====> 🚀🚀🚀 %v\n", err)
	}
	defer mysql.Close()
	//redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("mysql init failed =====> 🚀🚀🚀 %v\n", err)
	}
	defer redis.Close()
	snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId)
	//初始化校验器
	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Fatal("错误翻译初始化错误...")
		return
	}
	//注册路由
	r := routes.SetUp()
	//启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("server listen err:%s", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在此阻塞
	<-quit

	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)

	defer channel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}

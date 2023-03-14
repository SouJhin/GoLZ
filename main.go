package main

import (
	"fmt"

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
	}
	//mysql
	if err := mysql.Init(); err != nil {
		fmt.Printf("mysql init failed =====> 🚀🚀🚀 %v\n", err)
	}
	//redis
	//注册路由
	//启动服务
}

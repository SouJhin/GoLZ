package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("The settings files has been modified")
	})
	return
}

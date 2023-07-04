package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")  // 还可以在工作目录中查找配置
	err = viper.ReadInConfig() // read config
	if err != nil {
		fmt.Printf("settings viper.ReadInConfig() failed,err:%v\n", err)
	}
	viper.WatchConfig()
	//callback func
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("yaml changed")
	})

	return
}

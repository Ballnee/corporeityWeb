package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/corporeityWeb/") // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME/corporeityWeb") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")                   // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()                 // read config
	if err != nil {
		fmt.Printf("settings viper.ReadInConfig() failed,err:%v\n", err)
	} else {
		// 配置文件被找到，但产生了另外的错误
		fmt.Printf("settings viper.ReadInConfig() but err,err:%v\n", err)
	}
	viper.WatchConfig()
	//callback func
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("yaml changed")
	})

	return
}

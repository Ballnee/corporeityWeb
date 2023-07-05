package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// BMI
var BMImapBoy map[int][3]float32
var BMImapGirl map[int][3]float32

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
	BMImapBoy = make(map[int][3]float32)
	BMImapGirl = make(map[int][3]float32)
	BMImapBoy[1] = [3]float32{20.4, 18.2, 13.4}
	BMImapBoy[2] = [3]float32{20.5, 18.5, 13.6}
	BMImapBoy[3] = [3]float32{22.2, 19.5, 13.8}
	BMImapBoy[4] = [3]float32{22.7, 20.2, 14.1}
	BMImapBoy[5] = [3]float32{24.2, 21.5, 14.3}
	BMImapBoy[6] = [3]float32{24.6, 21.9, 14.6}

	BMImapGirl[1] = [3]float32{19.3, 17.4, 13.2}
	BMImapGirl[2] = [3]float32{20.3, 17.9, 13.4}
	BMImapGirl[3] = [3]float32{21.2, 18.7, 13.5}
	BMImapGirl[4] = [3]float32{22.1, 19.5, 13.6}
	BMImapGirl[5] = [3]float32{23.0, 20.6, 13.7}
	BMImapGirl[6] = [3]float32{23.7, 20.9, 14.1}

	return
}

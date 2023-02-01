package main

import (
	"corporeit/backend/dao/mysql"
	"corporeit/backend/logger"
	"corporeit/backend/pkg/snowflake"
	"corporeit/backend/routes"
	"corporeit/backend/settings"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {

	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed. err: %v\n", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(viper.GetString("app.mode")); err != nil {
		fmt.Printf("Init logger failed")
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")

	//3.初始化mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed. err: %v\n", err)
		return
	}
	defer mysql.Close()

	if err := snowflake.Init(uint16(viper.GetInt("app.machine_id"))); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	r := routes.Setup()

	err := r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}

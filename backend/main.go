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
	//初始化雪花id
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

//中间件测试
//func test() {
//	r := gin.Default()
//	r.GET("/signup", func(context *gin.Context) {
//		context.JSON(200, gin.H{
//			"res": "hello",
//		})
//	})
//	r.Use(func1)
//
//	v1 := r.Group("/v1", func2)
//	{
//		v1.GET("/test", func3)
//
//	}
//
//	v2 := r.Group("/v2")
//	v2.Use(func2)
//	{
//		v2.GET("/test1", func3)
//		v2.GET("/test2", func4)
//
//	}
//	fmt.Println("v1 ==", v1.Handlers)
//	fmt.Println("r == ", r.Handlers)
//	fmt.Println("v2==", v2.Handlers)
//	r.Run(":8888")
//
//}
//
//var cnt = 0
//
//func func1(c *gin.Context) {
//
//	fmt.Println("func1 before", cnt)
//	//c.Next()
//	//cnt++
//	fmt.Println("func1 after", cnt)
//}
//
//func func2(c *gin.Context) {
//	fmt.Println("func2")
//}
//
//func func3(c *gin.Context) {
//	fmt.Println(func3)
//	fmt.Println("func3")
//}
//func func4(c *gin.Context) {
//	fmt.Println(func4)
//	fmt.Println("func4")
//}

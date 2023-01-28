package main

import (
	"corporeit/backend/settings"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed. err: %v\n", err)
		return
	}

	fmt.Println("hello")
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

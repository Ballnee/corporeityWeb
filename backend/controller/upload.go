package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	name := c.PostForm("name")
	price := c.PostForm("price")
	fmt.Println(name, price)
	err := c.SaveUploadedFile(file, "./backend/"+file.Filename)
	if err != nil {
		fmt.Println(err)
	}
	Response(c, "success", "")
	return
}

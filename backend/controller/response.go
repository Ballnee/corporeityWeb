package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*

	error msg      string
	response data  {}
*/
//统一的返回给前端的数据结构
type ResponseData struct {
	Msg  interface{}
	Data interface{}
}

func Response(c *gin.Context, msg interface{}, data interface{}) {
	rd := ResponseData{
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

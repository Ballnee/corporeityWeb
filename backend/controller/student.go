package controller

import (
	"corporeit/backend/logic"
	"corporeit/backend/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func StudentAddHandler(c *gin.Context) {
	p := model.ParamAddStu{}
	if err := c.ShouldBindJSON(&p); err != nil {
		fmt.Println(p)
		fmt.Println(err)
		zap.L().Error("add student params invalid")
		Response(c, "add student params invalid", "")
		return
	}
	//get userid
	userId, err := GetCurrentUserId(c)
	if err != nil {
		Response(c, "ErrorUserNotLoginIn", "")
		return
	}
	if err := logic.AddStudent(&p, userId); err != nil {
		Response(c, "add student err", "")
		return
	}
	Response(c, "add student success", "")
	return
}

func GetCurrentUserId(c *gin.Context) (userId uint64, err error) {
	uid, ok := c.Get("userId")
	if !ok {
		err = errors.New("ErrorUserNotLoginIn")
		return
	}
	userId, ok = uid.(uint64)
	if !ok {
		err = errors.New("ErrorUserNotLoginIn")
		return
	}
	return

}

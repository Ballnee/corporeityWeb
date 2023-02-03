package controller

import (
	"corporeit/backend/logic"
	"corporeit/backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SignupHandler(c *gin.Context) {
	var p model.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		fmt.Println(p)
		zap.L().Error("signup with invalid params")
		//response
		c.JSON(http.StatusOK, gin.H{
			"msg": "register failed due to params",
		})
		return
	}

	err := logic.SignUp(&p)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user already exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "signup success",
	})

}

func LoginHandler(c *gin.Context) {
	var p model.ParamLoginIn
	fmt.Println(p.Username)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login with invalid params")
		c.JSON(http.StatusOK, gin.H{
			"msg": "login failed due to params",
		})
		return
	}
	token, err := logic.Login(&p)

	if err != nil {
		zap.L().Error("login err")
		c.JSON(http.StatusOK, gin.H{
			"msg": "username or password has err",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "login success",
		"data": token,
	})
	return

}

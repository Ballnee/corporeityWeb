package controller

import (
	"corporeit/backend/logic"
	"corporeit/backend/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SignupHandler(c *gin.Context) {
	var p model.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
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

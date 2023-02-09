package controller

import (
	"corporeit/backend/logic"
	"corporeit/backend/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ClassGetByIdHandler(c *gin.Context) {
	p := model.ParamGetClassById{}
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("class get by classId param err")
		Response(c, "class get by classId param err", "")
		return
	}
	err, class := logic.ClassGetById(p.ClassId)
	if err != nil {
		Response(c, "get class by id err", "")
		return
	}
	Response(c, "success", class)
	return
}

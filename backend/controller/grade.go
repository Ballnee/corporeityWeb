package controller

import (
	"corporeit/backend/logic"
	"corporeit/backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GradeAddHandler(c *gin.Context) {
	grade := &model.ParamAddGrade{}
	if err := c.ShouldBindJSON(grade); err != nil {
		fmt.Println("upload grades", err)
		zap.L().Error("grade add param err")
		Response(c, "grade add param err", "")
		return
	}

	err := logic.GradeAdd(grade)
	if err != nil {
		fmt.Println(err)
		Response(c, "grade add database err", "")
		zap.L().Error("grade add database err")
		return
	}
	Response(c, "grade add success", "")
	return
}

func GradeGetHandler(c *gin.Context) {
	err, grades := logic.QueryAllGrade()
	if err != nil {
		zap.L().Error("GradeGetHandler err")
		Response(c, "query all grades failed", "")
		return
	}
	Response(c, "query all grades success", grades)
	return

}

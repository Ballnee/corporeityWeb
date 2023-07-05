package routes

import (
	"corporeit/backend/controller"
	"corporeit/backend/logger"
	"corporeit/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.Cors())
	v1.POST("/signup", controller.SignupHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.POST("/student", controller.StudentAddHandler)
		v1.GET("/student", controller.StudentGetAllHandler)
	}

	{
		v1.POST("/grade", controller.GradeAddHandler)
		v1.GET("/grade", controller.GradeGetHandler)
	}

	{
		v1.GET("/class", controller.ClassGetByIdHandler)
		v1.POST("/class", controller.ClassAnalysisHandler)
	}
	{

	}
	return r
}

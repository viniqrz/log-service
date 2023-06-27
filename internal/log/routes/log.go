package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniqrz/gin-crud/internal/use_cases/create_log"
)

func GetLogRoutes(rg *gin.RouterGroup) {
	createLogUseCase := *create_log.NewCreateLogUseCase()
	createLogUseCaseController := create_log.NewCreateLogUseCaseController(createLogUseCase)

	logRouter := rg.Group("/log")

	logRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	logRouter.POST("/", func(c *gin.Context) {
		createLogUseCaseController.Execute(c)
	})
}
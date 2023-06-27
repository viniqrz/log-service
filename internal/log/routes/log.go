package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniqrz/gin-crud/internal/log/dto"
)

func GetLogRoutes(rg *gin.RouterGroup) {
	logRouter := rg.Group("/log")

	logRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	logRouter.POST("/", func(c *gin.Context) {
		var req dto.CreateLogInput
        c.BindJSON(&req)

		fmt.Println(req)
		println("here")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
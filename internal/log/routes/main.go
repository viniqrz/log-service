package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Run() {
	GetRoutes()
	r.Run()
}

func GetRoutes() {
	v1 := r.Group("/")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "Hello World!",
		})
	})

	GetLogRoutes(v1)
}
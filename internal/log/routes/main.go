package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniqrz/gin-crud/internal/log/middleware"
)

var r = gin.Default()

func Run(db *sql.DB) {
	r.Use(middleware.HandleExceptions)

	v1 := r.Group("/")

	GetLogRoutes(v1, db)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Hello World!",
		})
	})

	r.Run()
}

func HandleRoute(fn func(c *gin.Context)) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		fn(ctx)
	}
}

package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Run(db *sql.DB) {
	v1 := r.Group("/")

	GetLogRoutes(v1, db)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "Hello World!",
		})
	})

	r.Run()
}

func HandleRoute(fn func(c *gin.Context)) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		fn(ctx)
	};
}
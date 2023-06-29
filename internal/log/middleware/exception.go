package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/viniqrz/gin-crud/internal/log/exception"
)

func HandleExceptions(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		println(exception.GetStatusCode(err))
		c.JSON(exception.GetStatusCode(err), gin.H{"message": err.Error()})
		return
	}
}
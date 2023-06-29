package delete_log

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteLogUseCaseController struct {
	useCase *DeleteLogUseCase
}

func NewDeleteLogUseCaseController(useCase *DeleteLogUseCase) *DeleteLogUseCaseController {
	return &DeleteLogUseCaseController{
		useCase: useCase,
	}
}

func (c *DeleteLogUseCaseController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.useCase.Execute(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Log deleted successfully"})
}

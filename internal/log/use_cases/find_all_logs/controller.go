package find_all_logs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindAllLogsUseCaseController struct {
	useCase *FindAllLogsUseCase
}

func NewFindAllLogsUseCaseController(useCase *FindAllLogsUseCase) *FindAllLogsUseCaseController {
	return &FindAllLogsUseCaseController{
		useCase: useCase,
	}
}

func (c *FindAllLogsUseCaseController) Execute(ctx *gin.Context) {
	res, err := c.useCase.Execute()
	if (err != nil) {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, res)
}
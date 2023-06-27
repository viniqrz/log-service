package create_log

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniqrz/gin-crud/internal/log/dto"
)

type CreateLogUseCaseController struct {
	useCase CreateLogUseCase
}

func NewCreateLogUseCaseController(useCase CreateLogUseCase) *CreateLogUseCaseController {
	return &CreateLogUseCaseController{
		useCase: useCase,
	}
}

func (c *CreateLogUseCaseController) Execute(context *gin.Context) {
	var req *dto.CreateLogInput
	err := context.BindJSON(&req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON malformed",
		})
		return
	}
	context.JSON(http.StatusCreated, c.useCase.Execute(req))
}

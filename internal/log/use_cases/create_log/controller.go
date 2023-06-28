package create_log

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/viniqrz/gin-crud/internal/log/dto"
)

type CreateLogUseCaseController struct {
	useCase *CreateLogUseCase
}

func NewCreateLogUseCaseController(useCase *CreateLogUseCase) *CreateLogUseCaseController {
	return &CreateLogUseCaseController{
		useCase: useCase,
	}
}

func (c *CreateLogUseCaseController) Execute(ctx *gin.Context) {
	var dto *dto.CreateLogInput
	ctx.ShouldBindBodyWith(&dto, binding.JSON)

	level := dto.Level

	if (level != "ERROR" && level != "WARNING" && level != "INFO" && level != "DEBUG") {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid level"})
		ctx.Abort()
		return
	}

	res, err := c.useCase.Execute(dto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	fmt.Println(res)

	ctx.JSON(http.StatusCreated, res)
}

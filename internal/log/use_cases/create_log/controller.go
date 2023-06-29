package create_log

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/viniqrz/gin-crud/internal/log/dto"
	"github.com/viniqrz/gin-crud/internal/log/exception"
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
		ctx.Error(errors.New(exception.UnprocessableEntityErrorMessage))
		ctx.Abort()
		return
	}

	res, err := c.useCase.Execute(dto)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	fmt.Println(res)

	ctx.JSON(http.StatusCreated, res)
}

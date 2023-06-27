package create_log

import (
	"github.com/viniqrz/gin-crud/internal/log/dto"
	"github.com/viniqrz/gin-crud/internal/log/entity"
	"github.com/viniqrz/gin-crud/internal/log/transformer"
)

type CreateLogUseCase struct {}

func NewCreateLogUseCase() *CreateLogUseCase {
	return &CreateLogUseCase{}
}

func (c *CreateLogUseCase) Execute(input *dto.CreateLogInput) *entity.Log {
	return transformer.TransformLogCreateInputToDomain(input)
}
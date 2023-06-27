package transformer

import (
	"github.com/viniqrz/gin-crud/internal/log/dto"
	"github.com/viniqrz/gin-crud/internal/log/entity"
)


func TransformLogCreateInputToDomain(input dto.CreateLogInput) entity.Log {
	return *entity.NewLog(input.Message)
}

func TransformLogDomainToCreateOutput(log entity.Log) *dto.CreateLogOutput {
	return &dto.CreateLogOutput{
		ID:        log.ID,
		Message:   log.Message,
		CreatedAt: log.CreatedAt,
	}
}

package create_log

import (
	"github.com/viniqrz/gin-crud/internal/log/dto"
	"github.com/viniqrz/gin-crud/internal/log/entity"
	"github.com/viniqrz/gin-crud/internal/log/repository"
	"github.com/viniqrz/gin-crud/internal/log/transformer"
)

type CreateLogUseCase struct {
	logRepository repository.LogRepository
}

func NewCreateLogUseCase(
	logRepository repository.LogRepository,
) *CreateLogUseCase {
	return &CreateLogUseCase{
		logRepository: logRepository,
	}
}

func (c *CreateLogUseCase) Execute(input *dto.CreateLogInput) (*entity.Log, error) {
	log := transformer.TransformLogCreateInputToDomain(input)

	log, err := c.logRepository.Create(log)
	if (err != nil) {
		return nil, err
	}

	return log, nil
}
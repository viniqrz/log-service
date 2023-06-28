package find_all_logs

import (
	"github.com/viniqrz/gin-crud/internal/log/entity"
	"github.com/viniqrz/gin-crud/internal/log/repository"
)

type FindAllLogsUseCase struct {
	repository repository.LogRepository
}

func NewFindAllLogsUseCase(repository repository.LogRepository) *FindAllLogsUseCase {
	return &FindAllLogsUseCase{
		repository: repository,
	}
}

func (c *FindAllLogsUseCase) Execute() ([]*entity.Log, error) {
	return c.repository.FindAll()
}
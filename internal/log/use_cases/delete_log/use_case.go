package delete_log

import "github.com/viniqrz/gin-crud/internal/log/repository"

type DeleteLogUseCase struct {
	repository repository.LogRepository
}

func NewDeleteLogUseCase(repository repository.LogRepository) *DeleteLogUseCase {
	return &DeleteLogUseCase{
		repository: repository,
	}
}

func (c *DeleteLogUseCase) Execute(id string) error {
	return c.repository.Delete(id)
}
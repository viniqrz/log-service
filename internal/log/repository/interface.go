package repository

import (
	"github.com/viniqrz/gin-crud/internal/log/entity"
)

type LogRepository interface {
	FindAll() ([]*entity.Log, error)
	Create(log *entity.Log) (*entity.Log, error)
	Delete(id string) error
}


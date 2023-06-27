package entity

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	ID    int	`json:"id"`
	Message string	`json:"message"`
	CreatedAt string	`json:"created_at"`
}

func NewLog(message string) *Log {
	return &Log{
		Message: message,
	}
}

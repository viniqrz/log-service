package entity

import (
	"time"

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
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
}

package entity

import (
	"time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type Log struct {
	// gorm.Model
	ID    string	`json:"id"`
	Message string	`json:"message"`
	CreatedAt string	`json:"created_at"`
	Level string	`json:"level"` // ERROR, WARNING, INFO, DEBUG
}

var LogLevels = []string{"ERROR", "WARNING", "INFO", "DEBUG"}

func NewLog(message string, level string) *Log {
	return &Log{
		ID: uuid.New().String(),
		Message: message,
		Level: level,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
}

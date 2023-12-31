package dto

type CreateLogInput struct {
	Message string `json:"message" binding:"required"`
	Level string `json:"level" binding:"required"`
}

type CreateLogOutput struct {
	ID      string   `json:"id"`
	Message string `json:"message"`
	CreatedAt string `json:"created_at"`
	Level string `json:"level"`
}

package main

import (
	"github.com/viniqrz/gin-crud/internal/log/database"
	"github.com/viniqrz/gin-crud/internal/log/routes"
)

func main() {
	db := database.Connect()
	routes.Run(db)
}

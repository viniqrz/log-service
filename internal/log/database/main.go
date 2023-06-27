package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	
	return db
}
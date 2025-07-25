package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Advice struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Message  string `json:"message"`
	Name     string `json:"name"`
	Relation string `json:"relation"`
}

var DB *gorm.DB

func ConnectDB() {
	// Connection string
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	// Open database connection
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate models to form tables in database
	err = DB.AutoMigrate(&Advice{})
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}

	fmt.Println(" connected to database successfully")
}

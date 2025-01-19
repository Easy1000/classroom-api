package main

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	// Database connection string
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("Environment variable DB_DSN is not set")
	}

	// dsn := "root:password@tcp(127.0.0.1:3306)/classroom_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the database
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate the models (creates tables if they don't exist)
	err = db.AutoMigrate(&Student{}, &Class{}, &Enrollment{}, &Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

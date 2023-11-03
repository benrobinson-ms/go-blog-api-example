// database/database.go
package database

import (
	"fmt"

	"github.com/benrobinson-ms/go-blog-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	var err error
	connectionString := "host=localhost port=5432 user=ben dbname=go_blog_api sslmode=disable password="
	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Database connection error:", err)
		panic("Failed to connect to the database.")
	}

	// Migration to create tables for Article struct
	DB.AutoMigrate(&models.Article{})
}

func CloseDB() {
	DB.Close()
}

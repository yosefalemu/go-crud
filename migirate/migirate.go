package main

import (
	"log"
	"os"

	"github.com/yosefalemu/go-crud/initializers"
	"github.com/yosefalemu/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
}

func main() {
	if initializers.DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	initializers.DB.Debug().AutoMigrate(&models.Person{})
}

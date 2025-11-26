package main

import (
	"log"
	"os"
	"task_manager/data"
	"task_manager/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize MongoDB
	err = data.InitMongoDB(mongoURI, "task_manager")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	r := router.SetupRouter()
	r.Run(":" + port)
}

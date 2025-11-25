package main

import (
	"log"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	// Initialize MongoDB
	err := data.InitMongoDB("mongodb://127.0.0.1:27017", "task_manager")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	r := router.SetupRouter()
	r.Run(":8080")
}

package models

import "time"

type Task struct {
	ID          string    `json:"id" bson:"_id"`
	Title       string    `json:"title" binding:"required" bson:"title"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status"`
}

type TaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

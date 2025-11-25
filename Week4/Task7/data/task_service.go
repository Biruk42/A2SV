package data

import (
	"context"
	"errors"
	"time"

	"task_manager/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNotFound = errors.New("task not found")
	taskCollection *mongo.Collection
	userCollection *mongo.Collection
)

// InitMongoDB initializes the MongoDB connection
func InitMongoDB(uri, dbName string) error {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	db := client.Database(dbName)
	taskCollection = db.Collection("tasks")
	userCollection = db.Collection("users")
	return nil
}

// ListTasks retrieves all tasks from MongoDB
func ListTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	
	if tasks == nil {
		tasks = []models.Task{}
	}

	return tasks, nil
}

// GetTask retrieves a specific task by ID
func GetTask(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	filter := bson.M{"_id": id}
	err := taskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, ErrNotFound
		}
		return models.Task{}, err
	}
	return task, nil
}

// CreateTask creates a new task
func CreateTask(ctx context.Context, input models.TaskInput) (models.Task, error) {
	var due time.Time
	var err error
	if input.DueDate != "" {
		due, err = time.Parse(time.RFC3339, input.DueDate)
		if err != nil {
			return models.Task{}, err
		}
	}

	id := uuid.New().String()
	task := models.Task{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		DueDate:     due,
		Status:      input.Status,
	}
	if task.Status == "" {
		task.Status = "pending"
	}

	_, err = taskCollection.InsertOne(ctx, task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

// UpdateTask updates an existing task
func UpdateTask(ctx context.Context, id string, input models.TaskInput) (models.Task, error) {
	// First check if task exists
	_, err := GetTask(ctx, id)
	if err != nil {
		return models.Task{}, err
	}

	update := bson.M{}
	if input.Title != "" {
		update["title"] = input.Title
	}
	if input.Description != "" {
		update["description"] = input.Description
	}
	if input.Status != "" {
		update["status"] = input.Status
	}
	if input.DueDate != "" {
		due, err := time.Parse(time.RFC3339, input.DueDate)
		if err != nil {
			return models.Task{}, err
		}
		update["due_date"] = due
	}

	if len(update) == 0 {
		return GetTask(ctx, id)
	}

	filter := bson.M{"_id": id}
	updateDoc := bson.M{"$set": update}
	
	// Return the updated document
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedTask models.Task
	err = taskCollection.FindOneAndUpdate(ctx, filter, updateDoc, opts).Decode(&updatedTask)
	if err != nil {
		return models.Task{}, err
	}

	return updatedTask, nil
}

// DeleteTask deletes a task
func DeleteTask(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	result, err := taskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrNotFound
	}
	return nil
}

package data

import (
	"errors"
	"sync"
	"time"

	"task_manager/models"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("task not found")
)

type taskStore struct {
	mu    sync.RWMutex
	tasks map[string]models.Task
}

var store = &taskStore{
	tasks: make(map[string]models.Task),
}

// ListTasks
func ListTasks() []models.Task {
	store.mu.RLock()
	defer store.mu.RUnlock()

	res := make([]models.Task, 0, len(store.tasks))
	for _, t := range store.tasks {
		res = append(res, t)
	}
	return res
}

// GetTask
func GetTask(id string) (models.Task, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	t, ok := store.tasks[id]
	if !ok {
		return models.Task{}, ErrNotFound
	}
	return t, nil
}

// CreateTask
func CreateTask(input models.TaskInput) (models.Task, error) {
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

	store.mu.Lock()
	store.tasks[id] = task
	store.mu.Unlock()

	return task, nil
}

// UpdateTask
func UpdateTask(id string, input models.TaskInput) (models.Task, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	task, ok := store.tasks[id]
	if !ok {
		return models.Task{}, ErrNotFound
	}

	if input.Title != "" {
		task.Title = input.Title
	}
	if input.Description != "" {
		task.Description = input.Description
	}
	if input.Status != "" {
		task.Status = input.Status
	}
	if input.DueDate != "" {
		due, err := time.Parse(time.RFC3339, input.DueDate)
		if err != nil {
			return models.Task{}, err
		}
		task.DueDate = due
	}

	store.tasks[id] = task
	return task, nil
}

// DeleteTask
func DeleteTask(id string) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, ok := store.tasks[id]; !ok {
		return ErrNotFound
	}

	delete(store.tasks, id)
	return nil
}

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
mu sync.RWMutex
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
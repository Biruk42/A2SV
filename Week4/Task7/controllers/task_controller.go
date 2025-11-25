package controllers

import (
	"net/http"
	"time"

	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func ListTasksHandler(c *gin.Context) {
	tasks, err := data.ListTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTaskHandler(c *gin.Context) {
	id := c.Param("id")
	t, err := data.GetTask(c, id)
	if err != nil {
		if err == data.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)
}

// CreateTaskHandler POST /tasks
func CreateTaskHandler(c *gin.Context) {
	var input models.TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := data.CreateTask(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// UpdateTaskHandler PUT /tasks/:id
func UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := data.UpdateTask(c, id, input)
	if err != nil {
		if err == data.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// DeleteTaskHandler DELETE /tasks/:id
func DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")
	if err := data.DeleteTask(c, id); err != nil {
		if err == data.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// helper: format time to RFC3339 when zero value should be omitted (not used directly in handlers)
func formatIfNotZero(t time.Time) *string {
	if t.IsZero() {
		return nil
	}
	s := t.Format(time.RFC3339)
	return &s
}

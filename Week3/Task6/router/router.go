package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
)

// SetupRouter configures routes and returns *gin.Engine
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/tasks")
	{
		v1.GET("", controllers.ListTasksHandler)
		v1.POST("", controllers.CreateTaskHandler)
		v1.GET(":id", controllers.GetTaskHandler)
		v1.PUT(":id", controllers.UpdateTaskHandler)
		v1.DELETE(":id", controllers.DeleteTaskHandler)
	}

	return r
}

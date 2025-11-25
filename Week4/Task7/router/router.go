package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures routes and returns *gin.Engine
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/tasks", controllers.ListTasksHandler)
		protected.GET("/tasks/:id", controllers.GetTaskHandler)

		admin := protected.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/tasks", controllers.CreateTaskHandler)
			admin.PUT("/tasks/:id", controllers.UpdateTaskHandler)
			admin.DELETE("/tasks/:id", controllers.DeleteTaskHandler)
			admin.POST("/promote", controllers.PromoteUserHandler)
		}
	}

	return r
}

package main

import (
	"Ampl_project2/database"
	"Ampl_project2/handlers"
	"Ampl_project2/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	public := r.Group("/public")
	{
		public.GET("/tasks", handlers.GetTasks)
	}

	protected := r.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())
	protected.Use(middleware.RateLimiterMiddleware())
	{
		protected.GET(":id", handlers.GetTaskByID)
		protected.POST("", handlers.CreateTask)
		protected.PUT(":id", handlers.UpdateTask)
		protected.DELETE(":id", handlers.DeleteTask)
	}

	r.Run(":8080")
}

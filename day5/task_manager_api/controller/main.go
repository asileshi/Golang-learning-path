package main

import (
	"github.com/asileshi/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	repository.InitDB()

	// Create a new Gin router
	router := gin.Default()
	
	// Define the routes and their handlers
	router.GET("/tasks", GetTasksHandler)
	router.GET("/tasks/:id", GetTaskByIDHandler)
	router.POST("/tasks", CreateTaskHandler)
	router.PUT("/tasks/:id", UpdateTaskHandler)
	router.DELETE("/tasks/:id", DeleteTaskHandler)

	// Start the server on port 8080
	router.Run(":8080")
}

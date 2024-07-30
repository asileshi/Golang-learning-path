package main

import (
	"log"

	"github.com/asileshi/task_manager_api/repository"
	"github.com/gin-gonic/gin"
)

func main(){
	err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
		
	}


	router := gin.Default()
	router.POST("/register", RegistrationHandler)
	router.POST("/login", LoginHandler)
  
	router.GET("/tasks", GetTasksHandler)
	router.GET("/tasks/:id", GetTaskByIDHandler)
  
	router.PATCH("/tasks/:id", AuthMiddleware, AdminMidleware, UpdateTaskHandler)
	router.POST("/tasks", AuthMiddleware, AdminMidleware, CreateTaskHandler)
	router.DELETE("/tasks/:id", AuthMiddleware, AdminMidleware, DeleteTaskHandler)
  
	router.Run("localhost:8080")
  

}
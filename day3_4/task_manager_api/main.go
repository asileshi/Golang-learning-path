package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},	

}


func main() {
	
	router:=gin.Default()
	router.GET("/ping",func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message":"pong",
		})
	})
	
	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.IndentedJSON(200,tasks)
	})

	router.GET("/tasks/:id", func(ctx *gin.Context) {
		id:=ctx.Param("id")
		for _,task:=range tasks{
			if task.ID == id{
				ctx.IndentedJSON(200,task)
				return
			}
		}
		ctx.IndentedJSON(404,gin.H{"message":"task not found"})
	})
	router.PUT("/tasks/:id", func(ctx *gin.Context){

		id:=ctx.Param("id")
		var updatedTask Task
		if err:=ctx.BindJSON(&updatedTask); err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		for _,task:= range tasks{
			if task.ID == id{
				if updatedTask.Title != ""{
					task.Title = updatedTask.Title
				}
				if updatedTask.Description != ""{
					task.Description = updatedTask.Description
				}
				ctx.JSON(http.StatusOK,gin.H{"message":"task updated"})
				
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message":"task not found"})
	})

	router.DELETE("/tasks/:id", func(ctx *gin.Context){
		id:=ctx.Param("id")
		for i,task := range tasks{
			if task.ID == id{
				tasks = append(tasks[:i], tasks[i+1:]...)
				ctx.JSON(http.StatusOK,gin.H{"message":"task deleted"})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message":"task not found"})
	})

	router.POST("/tasks", func(ctx *gin.Context){
		var newTask Task
		if err:=ctx.BindJSON(&newTask); err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		tasks = append(tasks,newTask)
		ctx.JSON(http.StatusCreated,gin.H{"message":"task created"})

	})
	router.Run()

}
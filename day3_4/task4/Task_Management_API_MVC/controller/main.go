package main

import (
	"net/http"

	model "github.com/asileshi/task_manager_api/Model"
	"github.com/asileshi/task_manager_api/data"
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	router.GET("/tasks",func(ctx *gin.Context) {
		tasks:=data.GetTasks()
		ctx.IndentedJSON(http.StatusOK, tasks)
	})

	router.GET("/tasks/:id",func(ctx *gin.Context) {
		id:=ctx.Param("id")
		task,err:=data.GetTaskById(id)
		if err != nil{
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message":"task not found"})
			return 
		}
		ctx.IndentedJSON(http.StatusOK,task)
			
	})

	router.POST("/tasks",func(ctx *gin.Context) {
		var newTask model.Task
		if err := ctx.BindJSON(&newTask); err != nil{
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		data.CreateTask(newTask)
		ctx.IndentedJSON(http.StatusOK,gin.H{"message":"task created"})
	})

	router.PUT("/tasks/:id",func(ctx *gin.Context) {
		id:=ctx.Param("id")
		var updatedTask model.Task

		if err:= ctx.BindJSON(&updatedTask); err!=nil{
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		task,err := data.UpdateTask(id,updatedTask)

		if err != nil{
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message":"task not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK,gin.H{"message":"task update","task":task})
	
	})

	router.DELETE("/tasks/:id",func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := data.DeleteTask(id)
		if err!= nil{
			ctx.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{"message":"task deleted"})
	})

	router.Run(":8080")

}
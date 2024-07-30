package main

import (
	"net/http"

	"github.com/asileshi/task_manager_api/model"
	"github.com/asileshi/task_manager_api/repository"
	"github.com/gin-gonic/gin"
)
func GetTasksHandler(ctx *gin.Context) {
    tasks, err := repository.GetTasks()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByIDHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    task, err := repository.GetTaskByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusOK, task)

}

func CreateTaskHandler(ctx *gin.Context) {
    var task model.Task
    err := ctx.BindJSON(&task)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task, err = repository.CreateTask(task)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusCreated, task)

}

func UpdateTaskHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    var task model.Task
    err := ctx.BindJSON(&task)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task, err = repository.UpdateTask(id, task)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusOK, task)
}

func DeleteTaskHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    err := repository.DeleteTask(id)
    //fmt.Print("err",err)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusOK, gin.H{"message":"task deleted"})

}

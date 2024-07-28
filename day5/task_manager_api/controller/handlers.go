package main

import (
    "net/http"

    "github.com/asileshi/model"
    "github.com/asileshi/repository"
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
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusNoContent, gin.H{"message":"task deleted"})

}

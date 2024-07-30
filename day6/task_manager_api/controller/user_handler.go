package main

import (
	"net/http"

	"github.com/asileshi/task_manager_api/model"
	"github.com/asileshi/task_manager_api/repository"
	"github.com/gin-gonic/gin"
)

func RegistrationHandler(ctx *gin.Context){

	var user model.User
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid request"})
		return
	}
	_, message := repository.CreateUser(user)

	if message != "" {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message":message})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message":"User registered successfully!"})
}

func LoginHandler(ctx *gin.Context){
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"invalid request"})
	}
	response, message := repository.Login(user)
	if message != "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"token":message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
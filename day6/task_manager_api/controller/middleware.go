package main

import (
	"net/http"
	"strings"

	"github.com/asileshi/task_manager_api/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)
func AuthMiddleware(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == ""{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"missing token"})
			return
		}

		authParts := strings.Split(auth, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"invalid authorization token"})
			return

	}
	claims := &model.Claim{}
	token, err := jwt.ParseWithClaims(authParts[1], claims, func(t *jwt.Token) (interface{}, error) {
			return model.Secretkey,nil
		})
	if err != nil || !token.Valid{
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message":"Unauthorized"})
		return
	}
	ctx.Set("name",claims.Email)
	ctx.Set("role", claims.Role)
	ctx.Next()
	}

	

func AdminMidleware(ctx *gin.Context){
	role, exists := ctx.Get("role")
	if !exists || role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message":"Unauthorized"})
		return
	}
	ctx.Next()
}
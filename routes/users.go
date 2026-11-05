package routes

import (
	"lukasrar/go-rest-api-with-db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func getUsers(context *gin.Context) {

	result, err := models.GetUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, result)

}

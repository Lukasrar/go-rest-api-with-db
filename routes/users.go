package routes

import (
	"lukasrar/go-rest-api-with-db/models"
	"lukasrar/go-rest-api-with-db/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func getUsers(context *gin.Context) {

	result, err := models.GetUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, result)

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	errValidateCred := user.ValidateCredentials()

	if errValidateCred != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": errValidateCred.Error()})
		return
	}

	token, errToken := utils.GenerateToken(user.Email, user.ID)

	if errToken != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": errToken.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "sucess", "token": token})
}

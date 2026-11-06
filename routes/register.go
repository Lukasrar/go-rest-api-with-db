package routes

import (
	"lukasrar/go-rest-api-with-db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, errEvent := models.GetEventById(eventId)

	if errEvent != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could fetch event"})
		return
	}

	errRegister := event.Register(userId)

	if errRegister != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "sucess"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	var event models.Event

	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled"})

}

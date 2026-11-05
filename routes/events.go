package routes

import (
	"lukasrar/go-rest-api-with-db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	userId := context.GetInt64("userId")
	event.UserID = userId

	errSave := event.Save()

	if errSave != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEventById(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	event, errGetEvent := models.GetEventById(eventId)

	requestUserId := context.GetInt64("userId")

	if event.UserID != requestUserId {
		context.JSON(http.StatusUnauthorized, err)
		return
	}

	if errGetEvent != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	var updatedEvent models.Event

	errBind := context.ShouldBindJSON(&updatedEvent)

	if errBind != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	updatedEvent.ID = eventId

	errUpdate := updatedEvent.UpdateEvent()

	if errUpdate != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	event, errGetEvent := models.GetEventById(eventId)

	if errGetEvent != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	errDelete := event.DeleteEvent()

	if errDelete != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}

package routes

import (
	"lukasrar/go-rest-api-with-db/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventId", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)
	authenticated.POST("/events/:eventId/register", registerForEvent)
	authenticated.DELETE("/events/:eventId/register", cancelRegistration)

	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.POST("/login", login)
}

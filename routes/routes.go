package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEventById)
	server.POST("/events", createEvent)
	server.PUT("/events/:eventId", updateEvent)
	server.DELETE("/events/:eventId", deleteEvent)
	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.POST("/login", login)
}

package routes

import (
    "example.com/rest-api/middleware"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
    server.GET("/events", getEvents)
    server.GET("/events/:id", getEvent)

    authenticate := server.Group("/")
    authenticate.Use(middleware.Authenticate)
    authenticate.POST("/events", createEvent)
    authenticate.PUT("/events/:id", updateEvent)
    authenticate.DELETE("/events/:id", deleteEvent)

    authenticate.POST("/events/:id/register", registerForEvents)
    authenticate.DELETE("/events/:id/register", cancelRegistration)

    server.POST("/signup", signup)
    server.POST("/login", login)
}
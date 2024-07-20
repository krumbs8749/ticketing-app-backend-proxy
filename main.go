package main

import (
	"go-ticketing-app/controllers"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Adjust this to your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/events", controllers.GetEvents)
	router.GET("/events/:id", controllers.GetEventByID)
	router.POST("/register/:eventId", controllers.RegisterAttendee)

	log.Println("Starting server on :8090")
	router.Run(":8090")
}

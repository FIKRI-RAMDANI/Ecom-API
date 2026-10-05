package main

import (
	"event-app/config"
	"event-app/controller"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	config.ConnectDB()

	server := gin.Default()

	// Route
	api := server.Group("/api")
	{
		api.POST("/events", controller.CreateEvent)
		api.GET("/events", controller.GetEvents)
		api.GET("/events/:id", controller.GetEventById)
		api.PUT("/events/:id", controller.UpdateEvent)
		api.DELETE("/events/:id", controller.DeleteEvent)

		// Register
		api.POST("/auth/register", controller.RegisterUser)
	}

	server.Run(":8080")
}

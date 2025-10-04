package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Anurag/url-shortner/api/database"
	"github.com/Anurag/url-shortner/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err, "Error loading .env file")
	}

	database.Inizialize()

	router := gin.Default()
	setupRouters(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(router.Run(":" + port))
}

func setupRouters(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.PUT("/api/v1/:shortID", routes.EditURL)
	router.DELETE("/api/v1/:shortID", routes.DeleteURL)
	router.POST("/api/v1/tag", routes.AddTag)
	router.GET("/:shortID", routes.RedirectByShortID)
}

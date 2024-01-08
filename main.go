package main

import (
	"daily-quote/internal/controllers"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()

	router.POST("/quote/daily", controllers.SendDailyQuote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

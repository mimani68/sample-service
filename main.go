package main

import (
	"log"
	"net/http"
	"os"

	"app.io/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {

	logger.Log("The server is about to start ... ")
	logger := log.New(os.Stderr, "", 0)

	r := gin.Default()

	// r.GET("/*", func(c *gin.Context) {
	r.NoRoute(func(c *gin.Context) {
		logger.Println("Someone call echo page")
		c.JSON(http.StatusOK, gin.H{
			"message":  "Hello dear users",
			"version":  "1.0.5-alpha",
			"logLevel": os.Getenv("LOG_LEVEL"),
		})
	})

	r.Run(":3000")
}

package main

import (
	"log"
	"net/http"
	"os"

	"app.io/server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {

	logger.Log("Server is about to start")
	logger := log.New(os.Stderr, "", 0)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		logger.Println("Someone call echo page")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
			"version": "1.0.2",
		})
	})

	r.Run(":3000")
}

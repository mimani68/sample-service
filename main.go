package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.New(os.Stderr, "", 0)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		logger.Println("Someone call echo page")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
			"version": "1.0.1",
		})
	})

	r.Run(":3000")
}

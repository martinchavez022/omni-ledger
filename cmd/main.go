package main

import (

	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize the Gin engine
	r := gin.Default()

	// 2. Define a root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "OmniLedger API is live!",
		})
	})

	// Another route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	})

	// 3. Start the server on port 8080
	r.Run(":8080")
}
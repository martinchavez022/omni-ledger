package main

import (
	"omni-ledger/internal/database"
	"omni-ledger/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.Connect()

	// Initialize Router
	r := gin.Default()

	// Routes
	r.POST("/transactions",  handlers.CreateTransaction)
	r.GET("/transactions", handlers.GetTransaction)

	r.Run(":8080")
}
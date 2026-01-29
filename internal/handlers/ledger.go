package handlers

import (
	"net/http"
	"omni-ledger/internal/database"
	"omni-ledger/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

func GetTransaction(c *gin.Context) {
	var list []models.Transaction
	database.DB.Find(&list)
	c.JSON(http.StatusOK, list)
}
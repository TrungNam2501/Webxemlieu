package controllers

import (
	"log"
	"net/http"

	"Webxemlieu/database"
	"Webxemlieu/models"

	"github.com/gin-gonic/gin"
)

// GetItems trả về toàn bộ items.
func GetItems(c *gin.Context) {
	var items []models.Item
	if err := database.DB.Find(&items).Error; err != nil {
		log.Println("GetItems: query failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(http.StatusOK, items)
}

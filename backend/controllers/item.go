package controllers

import (
	"Webxemlieu/database"
	"Webxemlieu/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(200, items)
}
func GetRecipeDefinitions(c *gin.Context) {
	var recipes []models.RecipeProcessDefinition
	if err := database.DB.Limit(100).Find(&recipes).Error; err != nil {
		log.Println("Query failed:", err)
		c.JSON(500, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(200, recipes)
}

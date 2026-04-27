package controllers

import (
	"log"
	"net/http"

	"Webxemlieu/database"
	"Webxemlieu/models"

	"github.com/gin-gonic/gin"
)

// GetRecipeDefinitions trả về tối đa 100 dòng đầu tiên của bảng recipe_process_definition.
func GetRecipeDefinitions(c *gin.Context) {
	var recipes []models.RecipeProcessDefinition
	if err := database.DB.Limit(100).Find(&recipes).Error; err != nil {
		log.Println("GetRecipeDefinitions: query failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

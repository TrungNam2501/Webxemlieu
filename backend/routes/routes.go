package routes

import (
	"Webxemlieu/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/items", controllers.GetItems)
		api.GET("/recipe-definitions", controllers.GetRecipeDefinitions)
	}
}

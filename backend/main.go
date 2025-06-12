package main

import (
	"strings"

	"Webxemlieu/database"
	"Webxemlieu/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	database.Connect()
	routes.SetupRoutes(r)

	// Serve static assets
	r.Static("/app/assets", "../frontend/dist/assets")
	r.StaticFile("/app/favicon.ico", "../frontend/dist/favicon.ico")

	r.GET("/app/", func(c *gin.Context) {
		c.File("../frontend/dist/index.html")
	})
	// ...existing code...
	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/app") {
			c.File("../frontend/dist/index.html")
		} else {
			c.JSON(404, gin.H{"message": "Not Found"})
		}
	})

	r.Run(":8080")
}

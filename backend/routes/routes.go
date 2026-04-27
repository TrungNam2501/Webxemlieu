package routes

import (
	"net/http"
	"path/filepath"
	"strings"

	"Webxemlieu/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAPI đăng ký các route REST dưới prefix /api.
func SetupAPI(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/items", controllers.GetItems)
		api.GET("/recipe-definitions", controllers.GetRecipeDefinitions)
	}
}

// SetupSPA đăng ký các route phục vụ static build của frontend (Vue SPA dưới /app).
// distDir là đường dẫn tới thư mục build (mặc định ../frontend/dist).
func SetupSPA(r *gin.Engine, distDir string) {
	r.Static("/app/assets", filepath.Join(distDir, "assets"))
	r.StaticFile("/app/favicon.ico", filepath.Join(distDir, "favicon.ico"))

	indexFile := filepath.Join(distDir, "index.html")
	r.GET("/app/", func(c *gin.Context) {
		c.File(indexFile)
	})

	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/app") {
			c.File(indexFile)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})
}

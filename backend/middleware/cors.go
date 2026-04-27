package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS trả về một middleware CORS đơn giản, chỉ cho phép các origin trong allowed.
// Sử dụng "*" trong allowed để cho phép mọi origin (chỉ nên dùng trong dev).
func CORS(allowed []string) gin.HandlerFunc {
	allowMap := make(map[string]struct{}, len(allowed))
	allowAll := false
	for _, o := range allowed {
		if o == "*" {
			allowAll = true
			continue
		}
		allowMap[o] = struct{}{}
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			if allowAll {
				c.Header("Access-Control-Allow-Origin", "*")
			} else if _, ok := allowMap[origin]; ok {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Vary", "Origin")
			}
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
			c.Header("Access-Control-Max-Age", "600")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

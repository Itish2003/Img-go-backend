package cors

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define allowed origins (frontend URLs)
		allowedOrigins := []string{
			"https://img-react-frontend.onrender.com", // Production frontend
			"http://localhost:3000",                  // Local development
		}

		origin := c.Request.Header.Get("Origin") // Get the origin of the request
		isAllowedOrigin := false

		// Check if the request origin is in the allowed origins list
		for _, o := range allowedOrigins {
			if origin == o {
				isAllowedOrigin = true
				break
			}
		}

		// Set CORS headers if the origin is allowed
		if isAllowedOrigin {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			log.Printf("Allowed Origin: %s", origin) // Log allowed origins for debugging
		} else {
			log.Printf("Blocked Origin: %s", origin) // Log blocked origins for debugging
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")

		// Preflight (OPTIONS) requests should respond immediately
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No Content
			return
		}

		c.Next()
	}
}

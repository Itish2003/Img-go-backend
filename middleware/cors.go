package cors

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow specific origins in development and production
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // For testing purposes, replace "*" with specific origins for production.
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")

		// Log the CORS headers to check if they are set
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		log.Printf("CORS Headers: %v", c.Writer.Header())

		// If it's a preflight (OPTIONS) request, respond immediately
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No Content
			return
		}
		c.Next()
	}
}

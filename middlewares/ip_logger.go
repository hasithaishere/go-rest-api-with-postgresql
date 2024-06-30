package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func IPLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        ipAddress := c.ClientIP()
        log.Printf("IP Address: %s", ipAddress)

        // Optionally, you can store the IP address in the context
        c.Set("ClientIP", ipAddress)

        // Continue to next middleware/handler
        c.Next()
    }
}
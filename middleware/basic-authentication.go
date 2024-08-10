package middleware

import (
	"encoding/base64"
	"github.com/MESMUR/fixed-term-track-web-server/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var username = config.AppConfig.AppUsername
		var password = config.AppConfig.AppPassword

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the header is in the correct format
		if !strings.HasPrefix(authHeader, "Basic ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Decode the base64 encoded username:password
		payload := strings.TrimPrefix(authHeader, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(payload)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Split the decoded string into username and password
		parts := strings.SplitN(string(decoded), ":", 2)

		if len(parts) != 2 || parts[0] != username || parts[1] != password {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If the username and password are correct, proceed to the next handler
		c.Next()
	}
}

package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	auth "github.com/rosariocannavo/api_gateway/config"
	"github.com/rosariocannavo/api_gateway/internal/nats"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s", time.Now().UTC().Format(time.RFC3339), "middleware/Authenticate", http.StatusUnauthorized, "error: Authorization header is missing")
			nats.NatsConnection.PublishMessage(message)

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return auth.JWTSecretKey, nil
		})

		if err != nil || !token.Valid {
			message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s", time.Now().UTC().Format(time.RFC3339), "middleware/Authenticate", http.StatusUnauthorized, "error: Invalid token")
			nats.NatsConnection.PublishMessage(message)

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			fmt.Println("Invalid token claims")
			c.Abort()
			return
		}

		if !ok {
			fmt.Println("Username is not a string or doesn't exist")
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}

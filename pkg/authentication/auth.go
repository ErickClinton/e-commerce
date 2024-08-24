package authentication

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not defined")
	}
	tokenExpiry := 24 * time.Hour

	tokenService := NewTokenManager(secretKey, tokenExpiry)
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := tokenService.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}

}

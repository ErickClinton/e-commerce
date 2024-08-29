package middleware

import (
	"eccomerce/pkg/authentication"
	"log"
	"net/http"
	"os"
	"strconv"
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

	tokenService := authentication.NewTokenManager(secretKey, tokenExpiry)
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := tokenService.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		userIDUint, err := strconv.ParseUint(claims.UserID, 10, 32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}
		userID := uint(userIDUint)
		c.Set("userID", &userID)
		c.Set("userRole", claims.Role)
		println(c.Get("userID"))
		c.Next()
	}

}

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exist := c.Get("userRole")

		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}

		if userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission for this execution"})
			c.Abort()
			return
		}

		c.Next()
	}

}

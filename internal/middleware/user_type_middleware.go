package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentOnlyMiddleware - middleware untuk endpoint yang hanya bisa diakses student
func StudentOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, exists := c.Get(ContextCurrentUser)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan di konteks"})
			return
		}

		user, ok := currentUser.(ClaimResult)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
			return
		}

		if user.Role != "student" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Access denied",
				"message": "Endpoint ini hanya dapat diakses oleh student",
			})
			return
		}

		c.Next()
	}
}

// TeacherOnlyMiddleware - middleware untuk endpoint yang hanya bisa diakses teacher
func TeacherOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, exists := c.Get(ContextCurrentUser)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan di konteks"})
			return
		}

		user, ok := currentUser.(ClaimResult)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
			return
		}

		if user.Role != "teacher" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Access denied",
				"message": "Endpoint ini hanya dapat diakses oleh teacher",
			})
			return
		}

		c.Next()
	}
}

func UserTypeMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, exists := c.Get(ContextCurrentUser)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan di konteks"})
			return
		}

		user, ok := currentUser.(ClaimResult)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
			return
		}

		// Check if user role is in allowed roles
		for _, role := range allowedRoles {
			if user.Role == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error":   "Access denied",
			"message": "Anda tidak memiliki akses ke endpoint ini",
		})
	}
}

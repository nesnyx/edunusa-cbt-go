package middleware

import (
	"cbt/internal/repository"
	"cbt/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	CookieSessionToken = "session_token"
	ContextCurrentUser = "currentUser"
)

type ClaimResult struct {
	ID      string      `json:"id"`
	Role    string      `json:"role"`
	Profile interface{} `json:"profile"`
}

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	studentRepo := repository.NewStudentRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}
		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": err.Error()})
			return
		}
		validateAndSetUser(c, claims, studentRepo, teacherRepo)
	}
}

func validateAndSetUser(
	c *gin.Context,
	claims *utils.MyCustomClaims,
	studentRepo repository.StudentRepositoryInterface, // Gunakan interface jika ada
	teacherRepo repository.TeacherRepositoryInterface, // Gunakan interface jika ada
) {
	var user interface{}
	var err error

	if claims.Role == "student" {
		user, err = studentRepo.FindByID(claims.ID)
	} else if claims.Role == "teacher" {
		user, err = teacherRepo.GetByID(claims.ID)
	} else {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid user role"})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, user not found"})
		return
	}
	dataUser := ClaimResult{
		ID:      claims.ID,
		Role:    claims.Role,
		Profile: user,
	}
	c.Set("currentUser", dataUser)
	c.Next()
}

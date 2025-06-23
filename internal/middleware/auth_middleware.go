package middleware

import (
	repositoryextention "cbt/extentions/repositoryExtention"
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

// AuthMiddleware adalah middleware untuk otentikasi JWT
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	studentRepo := repositoryextention.NewStudentRepository(db)
	teacherRepo := repositoryextention.NewTeacherRepository(db)
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
	studentRepo repositoryextention.StudentRepositoryInterface, // Gunakan interface jika ada
	teacherRepo repositoryextention.TeacherRepositoryInterface, // Gunakan interface jika ada
) {
	var user interface{}
	var err error
	type ClaimResult struct {
		Profile interface{} `json:"data"`
		Role    string      `json:"role"`
	}

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
		Profile: user,
		Role:    claims.Role,
	}
	c.Set(ContextCurrentUser, dataUser)
	c.Next()
}

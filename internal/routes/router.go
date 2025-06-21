package routes

import (
	repositoryextention "cbt/extentions/repositoryExtention"
	"cbt/internal/handler"
	"cbt/internal/repository"
	"cbt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	authRoutes := routerGroup.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})
	}
}

func SetupExamRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	teacherRepo := repositoryextention.NewTeacherRepository(db)
	examRepo := repository.NewExamRepository(db)
	examService := service.NewExamService(examRepo, teacherRepo)
	examHandler := handler.NewExamHandler(examService)
	examRoutes := routerGroup.Group("/exams")
	{
		examRoutes.POST("/create", examHandler.InsertNewExam)
		examRoutes.GET("/:id", examHandler.FindByID)
		examRoutes.GET("/get-by-teacher-id", examHandler.FindByTeacherID)
		examRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})
	}
}

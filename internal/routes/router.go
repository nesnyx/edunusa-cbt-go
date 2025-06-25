package routes

import (
	repositoryextention "cbt/extentions/repositoryExtention"
	"cbt/internal/handler"
	"cbt/internal/middleware"
	"cbt/internal/repository"
	"cbt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupExamRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	teacherRepo := repositoryextention.NewTeacherRepository(db)
	examRepo := repository.NewExamRepository(db)
	examService := service.NewExamService(examRepo, teacherRepo)
	examHandler := handler.NewExamHandler(examService)
	examRoutes := routerGroup.Group("/exams", middleware.AuthMiddleware(db))
	{
		examRoutes.POST("/create", examHandler.InsertNewExam)
		examRoutes.DELETE("/delete/:id", examHandler.DeleteExam)
		examRoutes.GET("/get-by-id/:id", examHandler.FindByID)
		examRoutes.GET("/get-by-teacher", examHandler.FindByTeacherID)
		examRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})
	}
}

func SetupExamTokenUsageRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	examTokenUsageRepo := repository.NewExamTokenUsageRepository(db)
	examTokenUsageService := service.NewExamTokenUsage(examTokenUsageRepo)
	examTokenUsageHandler := handler.NewExamTokenUsageHandler(examTokenUsageService)
	examTokenUsageRoutes := routerGroup.Group("/exam-token-usage", middleware.AuthMiddleware(db))
	{
		examTokenUsageRoutes.POST("/create", examTokenUsageHandler.Insert)
		examTokenUsageRoutes.DELETE("/delete/:id", examTokenUsageHandler.Delete)
	}
}

func SetupQuestionRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	questionRepo := repository.NewQuestionRepository(db)
	questionService := service.NewQuestionService(questionRepo)
	questionHandler := handler.NewQuestionHandler(questionService)
	questionRoutes := routerGroup.Group("/questions", middleware.AuthMiddleware(db))
	{
		questionRoutes.POST("/create", questionHandler.CreateQuestion)
		questionRoutes.DELETE("/delete/:id", questionHandler.DeleteQuestion)
		questionRoutes.PUT("/update/:id", questionHandler.UpdateQuestion)
		questionRoutes.GET("/get-by-teacher", questionHandler.GetByTeacher)
	}
}

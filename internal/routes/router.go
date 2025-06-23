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

func SetupExamRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	teacherRepo := repositoryextention.NewTeacherRepository(db)
	examRepo := repository.NewExamRepository(db)
	examService := service.NewExamService(examRepo, teacherRepo)
	examHandler := handler.NewExamHandler(examService)
	examRoutes := routerGroup.Group("/exams")
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

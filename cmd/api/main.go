package main

import (
	"cbt/internal/config"
	"cbt/internal/middleware"
	"cbt/internal/repository"
	"cbt/internal/routes"
	"cbt/pkg/logger"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := repository.InitDB(cfg)

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.SimpleLoggingMiddleware())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	apiRoutes := router.Group("/v1")

	// Exam
	routes.SetupExamRoutes(apiRoutes, db)

	// Exam Token Usage
	routes.SetupExamTokenUsageRoutes(apiRoutes, db)

	// Student
	routes.SetupStudentRoutes(apiRoutes, db)

	// Subject
	routes.SetupSubjectAndClassRoutes(apiRoutes, db)

	// Teacher
	routes.SetupTeacherRoutes(apiRoutes, db)

	// Auth
	routes.SetupAuthRoutes(apiRoutes, db)

	// Question
	routes.SetupQuestionRoutes(apiRoutes, db)

	// Question Bank
	routes.SetupQuestionBankRoutes(apiRoutes, db)

	// Exam Question
	routes.SetupExamQuestionRoutes(apiRoutes, db)

	// Student Answer
	routes.SetupStudentAnswerRoutes(apiRoutes, db)

	// Student Exam Attempt
	routes.SetupStudentExamAttemptRoutes(apiRoutes, db)

	logger.Info("Server starting on port " + cfg.APIServerPort)
	if err := router.Run(":" + cfg.APIServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}

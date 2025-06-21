package main

import (
	"cbt/extentions/configExtention"
	repositoryextention "cbt/extentions/repositoryExtention"
	routeextention "cbt/extentions/routeExtention"
	"cbt/internal/config"
	"cbt/internal/repository"
	"cbt/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Main
	cfgExtention, errCfgExtention := configExtention.LoadConfig("./extentions")
	if errCfgExtention != nil {
		log.Fatalf("Failed to load configuration: %v", errCfgExtention)
	}
	dbExtention, errExtention := repositoryextention.InitDBExtention(cfgExtention)
	if errExtention != nil {
		log.Fatalf("Failed to initialize database: %v", errExtention)
	}

	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := repository.InitDB(cfg)

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := gin.Default()
	apiRoutes := router.Group("/v1")

	// Authentication
	routes.SetupAuthRoutes(apiRoutes, db)
	// Exam
	routes.SetupExamRoutes(apiRoutes, db)

	// Exam Questions
	// ####

	// Student
	routeextention.SetupStudentRoutes(apiRoutes, dbExtention)

	// Subject
	routeextention.SetupSubjectAndClassRoutes(apiRoutes, dbExtention)

	// Teacher
	routeextention.SetupTeacherRoutes(apiRoutes, dbExtention)

	if err := router.Run(":" + cfg.APIServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}

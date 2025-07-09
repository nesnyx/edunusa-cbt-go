package routes

import (
	"cbt/internal/handler"
	"cbt/internal/middleware"
	"cbt/internal/repository"
	"cbt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// done
func SetupExamRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	teacherRepo := repository.NewTeacherRepository(db)
	examRepo := repository.NewExamRepository(db)
	examService := service.NewExamService(examRepo, teacherRepo)
	examHandler := handler.NewExamHandler(examService)
	examRoutes := routerGroup.Group("/exams", middleware.AuthMiddleware(db))
	{
		examRoutes.POST("/create", examHandler.InsertNewExam)
		examRoutes.DELETE("/delete/:id", examHandler.DeleteExam)
		examRoutes.GET("/get-by-id/:id", examHandler.FindByID)
		examRoutes.GET("/get-by-teacher", examHandler.FindByTeacherID)
		examRoutes.PUT("/update/:id", examHandler.UpdateExam)
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
		examTokenUsageRoutes.GET("/get-by-student/:id", examTokenUsageHandler.FindByStudent)
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

func SetupExamQuestionRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	examQuestionRepo := repository.NewExamQuestionRepository(db)
	examQuestionService := service.NewExamQuestionService(examQuestionRepo)
	examQuestionHandler := handler.NewExamQuestionHandler(examQuestionService)
	examQuestionRoutes := routerGroup.Group("/exam/questions", middleware.AuthMiddleware(db))
	{
		examQuestionRoutes.POST("/create", examQuestionHandler.CreateExamQuestion)
		examQuestionRoutes.DELETE("/delete/:id", examQuestionHandler.Delete)
		examQuestionRoutes.GET("/get-by-exam/:id", examQuestionHandler.GetByExam)
		examQuestionRoutes.GET("/get-by-subject/:id", examQuestionHandler.GetBySubject)
	}
}

// student session
func SetupStudentExamAttemptRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	examTokenUsageRepo := repository.NewExamTokenUsageRepository(db)
	examTokenUsageService := service.NewExamTokenUsage(examTokenUsageRepo)
	examRepo := repository.NewExamRepository(db)
	studentExamAttemptRepo := repository.NewStudentExamAttemptRepository(db)

	studentExamAttemptService := service.NewStudentExamAttemptService(
		studentExamAttemptRepo,
		examRepo,
		examTokenUsageRepo,
	)

	studentExamAttemptHandler := handler.NewStudentExamAttemptHandler(studentExamAttemptService, examTokenUsageService)

	// Use enhanced middleware
	studentExamAttemptRoutes := routerGroup.Group("/exams/students",
		middleware.AuthMiddleware(db),
		middleware.StudentOnlyMiddleware(),
		middleware.EnhancedExamSessionMiddleware(db),
	)
	{
		studentExamAttemptRoutes.POST("/start", studentExamAttemptHandler.StartExamination)
		studentExamAttemptRoutes.GET("/progress/:attemptId", studentExamAttemptHandler.GetExamProgress)
		studentExamAttemptRoutes.POST("/finish/:attemptId", studentExamAttemptHandler.FinishExam)
	}
}

func SetupStudentAnswerRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	studentAnswerRepo := repository.NewStudentAnswerRepository(db)
	studentAnswerService := service.NewStudentAnswerService(studentAnswerRepo)
	studentAnswerHandler := handler.NewStudentAnswerHandler(studentAnswerService)
	studentAnswerRoutes := routerGroup.Group("/students/exam", middleware.AuthMiddleware(db), middleware.StudentOnlyMiddleware())
	{
		studentAnswerRoutes.POST("/answer", studentAnswerHandler.InsertOrUpdate)
		studentAnswerRoutes.GET("/get-by-answer", studentAnswerHandler.GetByQuestionAndStudentAttempt)
	}
}

func SetupQuestionBankRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	questionBankRepo := repository.NewQuestionBankRepository(db)
	questionBankService := service.NewQuestionBankService(questionBankRepo)
	questionBankHandler := handler.NewQuestionBankHandler(questionBankService)
	questionBankRoutes := routerGroup.Group("/question-banks", middleware.AuthMiddleware(db))
	{
		questionBankRoutes.POST("/create", questionBankHandler.CreateQuestionBank)
		questionBankRoutes.DELETE("/delete/:id", questionBankHandler.DeleteQuestionBank)
		questionBankRoutes.GET("/get-by-teacher", questionBankHandler.GetQuestionBankByTeacher)
		questionBankRoutes.PUT("/update/:id", questionBankHandler.Update)
	}
}

func SetupStudentRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	hasRole := repository.NewHasRoleRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo, hasRole)
	studentHandler := handler.NewStudentHandler(studentService)
	studentRoutes := routerGroup.Group("/students")
	{
		studentRoutes.POST("/create", studentHandler.InsertStudent)
		studentRoutes.GET("/get-by-id/:id", studentHandler.FindByID)
		studentRoutes.GET("/get-by-nis/:nis", studentHandler.FindByNIS)
		studentRoutes.GET("/get-all", studentHandler.FindAllStudent)
		studentRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})
	}

}

func SetupTeacherRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	hasRole := repository.NewHasRoleRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	teacherService := service.NewTeacherService(teacherRepo, hasRole)
	teacherHandler := handler.NewTeacherHandler(teacherService)
	teacherRoutes := routerGroup.Group("/teachers")
	{
		teacherRoutes.POST("/create", teacherHandler.Insert)
		teacherRoutes.GET("/get-all", teacherHandler.FindAll)
		teacherRoutes.GET("/get-by-id/:id", teacherHandler.FindByID)
		teacherRoutes.DELETE("/delete/:id", teacherHandler.Delete)
		teacherRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})

	}
}

func SetupSubjectAndClassRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	classRepo := repository.NewClassRepository(db)
	classService := service.NewClassService(classRepo)
	classHandler := handler.NewClassHandler(classService)

	subjectRepo := repository.NewSubjectRepository(db)
	subjectService := service.NewSubjectService(subjectRepo, classService)
	subjectHandler := handler.NewSubjectHandler(subjectService)
	subjectRoutes := routerGroup.Group("/subjects")
	{
		subjectRoutes.POST("/create", subjectHandler.InsertSubject)
		subjectRoutes.GET("/get-all", subjectHandler.FindAll)
		subjectRoutes.GET("/get-by-id/:id", subjectHandler.FindbyID)
		subjectRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})

	}

	classRoutes := routerGroup.Group("/classes")
	{
		classRoutes.POST("/create", classHandler.Insert)
		classRoutes.GET("/get-all", classHandler.FindAll)
		classRoutes.GET("/get-by-id/:id", classHandler.FindByID)
		classRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})

	}
}

func SetupAuthRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	authRoutes := routerGroup.Group("/auth")

	// Group untuk siswa
	studentRoutes := authRoutes.Group("/students")
	studentRoutes.POST("/login", authHandler.LoginStudent)

	// Group untuk guru
	teacherRoutes := authRoutes.Group("/teachers")
	teacherRoutes.POST("/login", authHandler.LoginTeacher)

	profileRoutes := authRoutes.Group("/profile", middleware.AuthMiddleware(db))
	profileRoutes.GET("/me", func(c *gin.Context) {
		currentUser, exists := c.Get(middleware.ContextCurrentUser)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan di konteks"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": currentUser})
	})
}

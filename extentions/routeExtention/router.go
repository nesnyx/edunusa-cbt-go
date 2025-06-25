package routeextention

import (
	handlerextention "cbt/extentions/handlerExtention"
	repositoryextention "cbt/extentions/repositoryExtention"
	serviceextention "cbt/extentions/serviceExtention"
	"cbt/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupStudentRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	hasRole := repositoryextention.NewHasRoleRepository(db)
	studentRepo := repositoryextention.NewStudentRepository(db)
	studentService := serviceextention.NewStudentService(studentRepo, hasRole)
	studentHandler := handlerextention.NewStudentHandler(studentService)
	studentRoutes := routerGroup.Group("/students")
	{
		studentRoutes.POST("/create", studentHandler.InsertStudent)
		studentRoutes.GET("/get-all", studentHandler.FindAllStudent)
		studentRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})
	}

}

func SetupTeacherRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	hasRole := repositoryextention.NewHasRoleRepository(db)
	teacherRepo := repositoryextention.NewTeacherRepository(db)
	teacherService := serviceextention.NewTeacherService(teacherRepo, hasRole)
	teacherHandler := handlerextention.NewTeacherHandler(teacherService)
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
	classRepo := repositoryextention.NewClassRepository(db)
	classService := serviceextention.NewClassService(classRepo)
	classHandler := handlerextention.NewClassHandler(classService)

	subjectRepo := repositoryextention.NewSubjectRepository(db)
	subjectService := serviceextention.NewSubjectService(subjectRepo, classService)
	subjectHandler := handlerextention.NewSubjectHandler(subjectService)
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
	authRepo := repositoryextention.NewAuthRepository(db)
	authService := serviceextention.NewAuthService(authRepo)
	authHandler := handlerextention.NewAuthHandler(authService)

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

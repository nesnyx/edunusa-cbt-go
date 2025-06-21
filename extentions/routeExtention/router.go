package routeextention

import (
	handlerextention "cbt/extentions/handlerExtention"
	repositoryextention "cbt/extentions/repositoryExtention"
	serviceextention "cbt/extentions/serviceExtention"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var (
// 	db          *gorm.DB
// 	hasRoleRepo = repositoryextention.NewHasRoleRepository(db)
// )

func SetupStudentRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {

	studentRepo := repositoryextention.NewStudentRepository(db)
	studentService := serviceextention.NewStudentService(studentRepo)
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
		teacherRoutes.GET("/:id", teacherHandler.FindByID)
		teacherRoutes.DELETE("/:id", teacherHandler.Delete)
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
		subjectRoutes.GET("/:id", subjectHandler.FindbyID)
		subjectRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})

	}

	classRoutes := routerGroup.Group("/classes")
	{
		classRoutes.POST("/create", classHandler.Insert)
		classRoutes.GET("/get-all", classHandler.FindAll)
		classRoutes.GET("/:id", classHandler.FindByID)
		classRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "testing"})
		})

	}
}

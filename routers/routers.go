package routers

import (
	"github.com/gin-gonic/gin"
	"student-api/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	studentRoutes := router.Group("/students")
	{
		studentRoutes.GET("/", controllers.GetAllStudents)
		studentRoutes.POST("/", controllers.CreateStudent)
		studentRoutes.GET("/:id", controllers.GetStudentByID)
		studentRoutes.PUT("/:id", controllers.UpdateStudent)
		studentRoutes.DELETE("/:id", controllers.DeleteStudent)
	}

	return router
}

package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterEmployeeRoutes(router *gin.RouterGroup, container *app.Container, authMiddleware gin.HandlerFunc) {

	// Setup layers
	empRepo := repo.NewEmployeeRepository(container.DB)
	empService := services.NewEmployeeService(empRepo)
	empController := controllers.NewEmployeeController(empService)
	empRoutes := router.Group("/employees")
	// 🔒 Protected endpoints
	protected := Protected(empRoutes, authMiddleware)
	{
		protected.POST("", empController.CreateEmployee)
		protected.GET("", empController.ListEmployees)
		protected.GET("/:id", empController.GetEmployee)
		protected.PUT("/:id", empController.UpdateEmployee)
		protected.DELETE("/:id", empController.DeleteEmployee)
	}
}

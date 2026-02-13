package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterEmployeeRoutes(router *gin.RouterGroup, container *app.Container) {

	// Setup layers
	empRepo := repo.NewEmployeeRepository(container.DB)
	empService := services.NewEmployeeService(empRepo)
	empController := controllers.NewEmployeeController(empService)
	empRoutes := router.Group("/employees")
	{
		empRoutes.POST("", empController.CreateEmployee)
		empRoutes.GET("", empController.ListEmployees)
		empRoutes.GET("/:id", empController.GetEmployee)
		empRoutes.PUT("/:id", empController.UpdateEmployee)
		empRoutes.DELETE("/:id", empController.DeleteEmployee)
	}
}

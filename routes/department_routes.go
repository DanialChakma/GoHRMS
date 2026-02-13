package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterDepartmentRoutes(router *gin.RouterGroup, container *app.Container) {

	// Setup layers
	repo := repo.NewDepartmentRepository(container.DB)
	service := services.NewDepartmentService(repo)
	deptController := controllers.NewDepartmentController(service)
	// Department
	deptRoutes := router.Group("/departments")
	{
		deptRoutes.POST("", deptController.Create)
		deptRoutes.GET("", deptController.List)
		deptRoutes.GET("/:id", deptController.Get)
		deptRoutes.PUT("/:id", deptController.Update)
		deptRoutes.DELETE("/:id", deptController.Delete)
	}
}

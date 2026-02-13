package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterDepartmentRoutes(router *gin.RouterGroup, container *app.Container, authMiddleware gin.HandlerFunc) {

	// Setup layers
	repo := repo.NewDepartmentRepository(container.DB)
	service := services.NewDepartmentService(repo)
	deptController := controllers.NewDepartmentController(service)
	// Department
	deptRoutes := router.Group("/departments")
	// 🔒 Protected endpoints
	protected := Protected(deptRoutes, authMiddleware)
	{
		protected.POST("", deptController.Create)
		protected.GET("", deptController.List)
		protected.GET("/:id", deptController.Get)
		protected.PUT("/:id", deptController.Update)
		protected.DELETE("/:id", deptController.Delete)
	}
}

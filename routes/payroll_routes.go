package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterPayrollRoutes(router *gin.RouterGroup, container *app.Container, authMiddleware gin.HandlerFunc) {

	// Setup layers
	repo := repo.NewPayrollRepository(container.DB)
	service := services.NewPayrollService(repo)
	payrollController := controllers.NewPayrollController(service)
	payrollRoutes := router.Group("/payrolls")
	// payrollRoutes.Use(authMiddleware)
	// 🔒 Protected endpoints
	protected := Protected(payrollRoutes, authMiddleware)
	{
		protected.POST("", payrollController.CreatePayroll)
		protected.GET("", payrollController.ListPayrolls)
		protected.GET("/:id", payrollController.GetPayroll)
		protected.PUT("/:id", payrollController.UpdatePayroll)
		protected.DELETE("/:id", payrollController.DeletePayroll)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterPayrollRoutes(router *gin.RouterGroup, container *app.Container) {

	// Setup layers
	repo := repo.NewPayrollRepository(container.DB)
	service := services.NewPayrollService(repo)
	payrollController := controllers.NewPayrollController(service)
	payrollRoutes := router.Group("/payrolls")
	{
		payrollRoutes.POST("", payrollController.CreatePayroll)
		payrollRoutes.GET("", payrollController.ListPayrolls)
		payrollRoutes.GET("/:id", payrollController.GetPayroll)
		payrollRoutes.PUT("/:id", payrollController.UpdatePayroll)
		payrollRoutes.DELETE("/:id", payrollController.DeletePayroll)
	}
}

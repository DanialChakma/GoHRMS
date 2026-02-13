package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterLeaveRoutes(router *gin.RouterGroup, container *app.Container) {

	// Setup layers
	leaveRepo := repo.NewLeaveRepository(container.DB)
	leaveService := services.NewLeaveService(leaveRepo)
	leaveController := controllers.NewLeaveController(leaveService)
	leaveRoutes := router.Group("/leaves")
	{
		leaveRoutes.POST("", leaveController.CreateLeave)
		leaveRoutes.GET("", leaveController.ListLeaves)
		leaveRoutes.GET("/:id", leaveController.GetLeave)
		leaveRoutes.PUT("/:id", leaveController.UpdateLeave)
		leaveRoutes.DELETE("/:id", leaveController.DeleteLeave)
		leaveRoutes.PATCH("/:id/approve", leaveController.ApproveLeave)
		leaveRoutes.PATCH("/:id/reject", leaveController.RejectLeave)
	}
}

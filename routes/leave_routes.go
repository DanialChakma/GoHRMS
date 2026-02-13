package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterLeaveRoutes(router *gin.RouterGroup, container *app.Container, authMiddleware gin.HandlerFunc) {

	// Setup layers
	leaveRepo := repo.NewLeaveRepository(container.DB)
	leaveService := services.NewLeaveService(leaveRepo)
	leaveController := controllers.NewLeaveController(leaveService)
	leaveRoutes := router.Group("/leaves")
	// 🔒 Protected endpoints
	protected := Protected(leaveRoutes, authMiddleware)
	{
		protected.POST("", leaveController.CreateLeave)
		protected.GET("", leaveController.ListLeaves)
		protected.GET("/:id", leaveController.GetLeave)
		protected.PUT("/:id", leaveController.UpdateLeave)
		protected.DELETE("/:id", leaveController.DeleteLeave)
		protected.PATCH("/:id/approve", leaveController.ApproveLeave)
		protected.PATCH("/:id/reject", leaveController.RejectLeave)
	}
}

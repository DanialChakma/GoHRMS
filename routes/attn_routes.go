package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterAttendanceRoutes(router *gin.RouterGroup, container *app.Container, authMiddleware gin.HandlerFunc) {

	// Setup layers
	repo := repo.NewAttendanceRepository(container.DB)
	service := services.NewAttendanceService(repo)
	attendanceController := controllers.NewAttendanceController(service)
	attendanceRoutes := router.Group("/attendance")
	// 🔒 Protected endpoints
	protected := Protected(attendanceRoutes, authMiddleware)
	{
		protected.POST("/checkin/:employee_id", attendanceController.CheckIn)
		protected.POST("/checkout/:employee_id", attendanceController.CheckOut)
	}
}

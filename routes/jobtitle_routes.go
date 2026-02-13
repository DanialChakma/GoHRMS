package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterJobtitleRoutes(router *gin.RouterGroup, container *app.Container, authMiddleware gin.HandlerFunc) {

	// Setup layers
	repo := repo.NewJobTitleRepository(container.DB)
	service := services.NewJobTitleService(repo)
	jobTitleController := controllers.NewJobTitleController(service)
	jobTitleRoutes := router.Group("/job-titles")
	// 🔒 Protected endpoints
	protected := Protected(jobTitleRoutes, authMiddleware)
	{
		protected.POST("", jobTitleController.Create)
		protected.GET("", jobTitleController.List)
		protected.GET("/:id", jobTitleController.Get)
		protected.PUT("/:id", jobTitleController.Update)
		protected.DELETE("/:id", jobTitleController.Delete)
	}
}

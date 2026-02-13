package routes

import (
	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/controllers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterJobtitleRoutes(router *gin.RouterGroup, container *app.Container) {

	// Setup layers
	repo := repo.NewJobTitleRepository(container.DB)
	service := services.NewJobTitleService(repo)
	jobTitleController := controllers.NewJobTitleController(service)
	jobTitleRoutes := router.Group("/job-titles")
	{
		jobTitleRoutes.POST("", jobTitleController.Create)
		jobTitleRoutes.GET("", jobTitleController.List)
		jobTitleRoutes.GET("/:id", jobTitleController.Get)
		jobTitleRoutes.PUT("/:id", jobTitleController.Update)
		jobTitleRoutes.DELETE("/:id", jobTitleController.Delete)
	}
}

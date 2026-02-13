package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/auth"
	"go.mod/initializers"
	"go.mod/repo"
	"go.mod/services"
)

func RegisterRoutes(router *gin.Engine, container *app.Container) {
	// Read API version from env
	apiPrefix := os.Getenv("API_PREFIX")
	if apiPrefix == "" {
		apiPrefix = "/api/v1"
	}

	// ✅ Create Auth Middleware ONCE
	authRepo := repo.NewAuthRepository(container.DB)
	tokenService := services.NewTokenService(initializers.JwtKey, initializers.JwtRefreshKey)
	authService := services.NewAuthService(authRepo, tokenService)

	authMiddleware := auth.AuthMiddleware(authService, tokenService)

	// Register Swagger directly on root router
	RegisterSwaggerRoute(router)

	// Global API group
	api := router.Group(apiPrefix)
	api.Use(initializers.RateLimiterMiddleware)
	RegisterAuthRoutes(api, container)
	// employe
	RegisterEmployeeRoutes(api, container, authMiddleware)
	// department
	RegisterDepartmentRoutes(api, container, authMiddleware)
	// jobtitle
	RegisterJobtitleRoutes(api, container, authMiddleware)
	// attendance
	RegisterAttendanceRoutes(api, container, authMiddleware)
	// leave
	RegisterLeaveRoutes(api, container, authMiddleware)
	// payroll
	RegisterPayrollRoutes(api, container, authMiddleware)
}

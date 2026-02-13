package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mod/app"
)

func RegisterRoutes(router *gin.Engine, container *app.Container) {
	// Read API version from env
	apiPrefix := os.Getenv("API_PREFIX")
	if apiPrefix == "" {
		apiPrefix = "/api/v1"
	}

	// Register Swagger directly on root router
	RegisterSwaggerRoute(router)

	// Global API group
	api := router.Group(apiPrefix)

	RegisterAuthRoutes(api, container)
	// employe
	RegisterEmployeeRoutes(api, container)
	// department
	RegisterDepartmentRoutes(api, container)
	// jobtitle
	RegisterJobtitleRoutes(api, container)
	// attendance
	RegisterAttendanceRoutes(api, container)
	// leave
	RegisterLeaveRoutes(api, container)
	// payroll
	RegisterPayrollRoutes(api, container)
}

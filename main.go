package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mod/app"
	"go.mod/routes"
)

// @title HRMS API
// @version 1.0
// @description Human Resource Management System API built with Go (Gin framework), GORM, and MySQL.
// It follows an enterprise-grade architecture with **Controller → Service → Repository pattern**, supports
// Employee management, Attendance, Payroll, Leave Requests, Departments, Job Titles, and User Authentication/Authorization.
// Designed for production-grade usage with modular, extensible, and maintainable code.

// @termsOfService http://hrms.example.com/terms/

// @contact.name HRMS Support
// @contact.url http://hrms.example.com/support
// @contact.email support@hrms.example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// ---------------- OpenAPI 3.0 Settings ----------------
// @openapi: 3.0.3
// @BasePath /api/v1
// @schemes http https

// @servers.local.url http://localhost:8080
// @servers.local.description Local Development Server

// @servers.dev.url https://dev.hrms.example.com
// @servers.dev.description Development Server

// @servers.preprod.url https://preprod.hrms.example.com
// @servers.preprod.description Pre-Production Server

// @servers.prod.url https://hrms.example.com
// @servers.prod.description Production Server

// ---------------- Security Definitions ----------------
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by your JWT token. Example: "Bearer {token}"

// ---------------- Enterprise-Grade Notes ----------------
// 1. Controller layer: Handles HTTP requests, validation, and formatting responses.
// 2. Service layer: Handles business logic and orchestration between controllers and repositories.
// 3. Repository layer: Handles database operations using GORM, supports context and transaction handling.
// 4. Pagination: All list endpoints use standard query parameters `page` and `page_size` with a consistent PaginatedResponse struct.
// 5. JWT Authentication: Protected routes use Bearer token validation via auth middleware.
// 6. Modular routing: Routes are grouped by module (employees, payroll, attendance, leave, departments, job titles).
// 7. Extensibility: Adding new modules only requires creating new Controller → Service → Repository layers and registering routes.
// 8. Rate Limiting: Optional per-route or global rate limit middleware for DDOS protection.

// ---------------- Example Module Tags ----------------
// @tag.name Employees
// @tag.description Operations related to Employee management

// @tag.name Attendance
// @tag.description Operations for Attendance check-in, check-out

// @tag.name Payroll
// @tag.description Payroll creation, listing, and management

// @tag.name Leaves
// @tag.description Leave request creation, approval, rejection

// @tag.name Departments
// @tag.description Department management

// @tag.name JobTitles
// @tag.description Job Title management

// @tag.name Auth
// @tag.description Authentication & Authorization endpoints

func main() {
	container := app.Bootstrap()
	// Read PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // fallback
	}

	router := gin.Default()

	// Register all routes
	routes.RegisterRoutes(router, container)

	log.Printf("Server running on %s\n", port)
	router.Run(port)
}

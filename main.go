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
// @description Human Resource Management System API built with Go (Gin framework), GORM, and MySQL. Supports Employee management, Attendance, Payroll, Leave Requests, Departments, Job Titles, and User Authentication/Authorization.
// @termsOfService http://hrms.example.com/terms/

// @contact.name HRMS Support
// @contact.url http://hrms.example.com/support
// @contact.email support@hrms.example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

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

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by your JWT token. Example: "Bearer {token}"

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

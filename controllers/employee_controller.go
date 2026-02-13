package controllers

import (
	"net/http"
	"strconv"

	"go.mod/models"
	"go.mod/services"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service services.EmployeeService
}

func NewEmployeeController(s services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

// CreateEmployee POST /employees
// CreateEmployee godoc
// @Summary Create a new employee
// @Description Create a new employee with all required fields
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee object"
// @Success 201 {object} models.Employee
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees [post]
func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var employee models.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateEmployee(&employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, employee)
}

// GetEmployee GET /employees/:id
// GetEmployee godoc
// @Summary Get employee by ID
// @Description Get an employee's details by ID
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [get]
func (c *EmployeeController) GetEmployee(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)

	employee, err := c.service.GetEmployee(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// UpdateEmployee PUT /employees/:id
// UpdateEmployee godoc
// @Summary Update an employee
// @Description Update employee details by ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body models.Employee true "Employee object"
// @Success 200 {object} models.Employee
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [put]
func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)

	var employee models.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.ID = id
	if err := c.service.UpdateEmployee(&employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// DeleteEmployee DELETE /employees/:id
// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Delete an employee by ID
// @Tags Employees
// @Param id path int true "Employee ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id} [delete]
func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)

	if err := c.service.DeleteEmployee(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ListEmployees GET /employees?page=1&page_size=10

// ListEmployees godoc
// @Summary List employees with pagination
// @Description Get a paginated list of employees
// @Tags Employees
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} dto.EmployeeListResponse
// @Failure 500 {object} map[string]string
// @Router /employees [get]
func (c *EmployeeController) ListEmployees(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "1")
	pageSizeParam := ctx.DefaultQuery("page_size", "10")

	page, _ := strconv.Atoi(pageParam)
	pageSize, _ := strconv.Atoi(pageSizeParam)

	result, err := c.service.ListEmployees(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return generic response
	ctx.JSON(http.StatusOK, result)
}

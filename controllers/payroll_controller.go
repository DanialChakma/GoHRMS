package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/models"
	"go.mod/services"
)

type PayrollController struct {
	services services.PayrollService
}

func NewPayrollController(s services.PayrollService) *PayrollController {
	return &PayrollController{services: s}
}

// POST /payrolls
// CreatePayroll godoc
// @Summary Create Payroll
// @Description Create a new payroll record for an employee
// @Tags Payroll
// @Accept json
// @Produce json
// @Param payroll body models.Payroll true "Payroll data"
// @Success 201 {object} models.Payroll
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payrolls [post]
func (c *PayrollController) CreatePayroll(ctx *gin.Context) {
	var payroll models.Payroll
	if err := ctx.ShouldBindJSON(&payroll); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.services.CreatePayroll(ctx, &payroll); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, payroll)
}

// GET /payrolls/:id
// GetPayroll godoc
// @Summary Get Payroll
// @Description Get a payroll record by ID
// @Tags Payroll
// @Accept json
// @Produce json
// @Param id path int true "Payroll ID"
// @Success 200 {object} models.Payroll
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payrolls/{id} [get]
func (c *PayrollController) GetPayroll(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	payroll, err := c.services.GetPayroll(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Payroll not found"})
		return
	}

	ctx.JSON(http.StatusOK, payroll)
}

// GET /payrolls?page=1&page_size=10&employee_id=5
// ListPayrolls godoc
// @Summary List Payrolls
// @Description Get paginated list of payrolls, optionally filter by employee_id
// @Tags Payroll
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param employee_id query int false "Filter by employee ID"
// @Success 200 {object} dto.PayrollListResponse
// @Failure 500 {object} map[string]string
// @Router /payrolls [get]
func (c *PayrollController) ListPayrolls(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	var employeeID *uint64
	if empParam := ctx.Query("employee_id"); empParam != "" {
		id, _ := strconv.ParseUint(empParam, 10, 64)
		employeeID = &id
	}

	result, err := c.services.ListPayrolls(ctx, page, pageSize, employeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Returning generic response at runtime
	ctx.JSON(http.StatusOK, result)
}

// PUT /payrolls/:id

// UpdatePayroll godoc
// @Summary Update Payroll
// @Description Update a payroll record by ID
// @Tags Payroll
// @Accept json
// @Produce json
// @Param id path int true "Payroll ID"
// @Param payroll body models.Payroll true "Updated payroll data"
// @Success 200 {object} models.Payroll
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payrolls/{id} [put]
func (c *PayrollController) UpdatePayroll(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	var payroll models.Payroll
	if err := ctx.ShouldBindJSON(&payroll); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payroll.ID = id

	if err := c.services.UpdatePayroll(ctx, &payroll); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payroll)
}

// DELETE /payrolls/:id
// DeletePayroll godoc
// @Summary Delete Payroll
// @Description Delete a payroll record by ID
// @Tags Payroll
// @Accept json
// @Produce json
// @Param id path int true "Payroll ID"
// @Success 204 {string} string "No Content"
// @Failure 500 {object} map[string]string
// @Router /payrolls/{id} [delete]
func (c *PayrollController) DeletePayroll(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err := c.services.DeletePayroll(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/models"
	"go.mod/services"
)

type DepartmentController struct {
	service services.DepartmentService
}

func NewDepartmentController(s services.DepartmentService) *DepartmentController {
	return &DepartmentController{service: s}
}

// Create godoc
// @Summary Create Department
// @Description Create a new department
// @Tags Department
// @Accept json
// @Produce json
// @Param department body models.Department true "Department data"
// @Success 201 {object} models.Department
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /departments [post]
func (c *DepartmentController) Create(ctx *gin.Context) {
	var dept models.Department
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(ctx, &dept); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dept)
}

// Get godoc
// @Summary Get Department
// @Description Get department by ID
// @Tags Department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} models.Department
// @Failure 404 {object} map[string]string
// @Router /departments/{id} [get]
func (c *DepartmentController) Get(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	dept, err := c.service.Get(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	ctx.JSON(http.StatusOK, dept)
}

// List godoc
// @Summary List Departments
// @Description Get paginated list of departments
// @Tags Department
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} dto.DepartmentListResponse
// @Failure 500 {object} map[string]string
// @Router /departments [get]
func (c *DepartmentController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	result, err := c.service.List(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return generic response
	ctx.JSON(http.StatusOK, result)

}

// Update godoc
// @Summary Update Department
// @Description Update department by ID
// @Tags Department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Param department body models.Department true "Updated department data"
// @Success 200 {object} models.Department
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /departments/{id} [put]
func (c *DepartmentController) Update(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	var dept models.Department
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dept.ID = id

	if err := c.service.Update(ctx, &dept); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dept)
}

// Delete godoc
// @Summary Delete Department
// @Description Delete department by ID
// @Tags Department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Success 204 {string} string "No Content"
// @Failure 500 {object} map[string]string
// @Router /departments/{id} [delete]
func (c *DepartmentController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err := c.service.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/models"
	"go.mod/services"
)

type LeaveController struct {
	service services.LeaveService
}

func NewLeaveController(s services.LeaveService) *LeaveController {
	return &LeaveController{service: s}
}

// POST /leaves

// CreateLeave godoc
// @Summary Create a new leave request
// @Description Create a new leave request for an employee
// @Tags Leaves
// @Accept json
// @Produce json
// @Param leave body models.LeaveRequest true "Leave Request object"
// @Success 201 {object} models.LeaveRequest
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /leaves [post]
func (c *LeaveController) CreateLeave(ctx *gin.Context) {
	var leave models.LeaveRequest
	if err := ctx.ShouldBindJSON(&leave); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateLeave(ctx, &leave); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, leave)
}

// GET /leaves/:id

// GetLeave godoc
// @Summary Get leave request by ID
// @Description Retrieve leave request details by ID
// @Tags Leaves
// @Produce json
// @Param id path int true "Leave Request ID"
// @Success 200 {object} models.LeaveRequest
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /leaves/{id} [get]
func (c *LeaveController) GetLeave(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	leave, err := c.service.GetLeave(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Leave not found"})
		return
	}

	ctx.JSON(http.StatusOK, leave)
}

// GET /leaves?page=1&page_size=10
// ListLeaves godoc
// @Summary List leave requests with pagination
// @Description Get paginated list of leave requests
// @Tags Leaves
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} dto.LeaveListResponse
// @Failure 500 {object} map[string]string
// @Router /leaves [get]
func (c *LeaveController) ListLeaves(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	result, err := c.service.ListLeaves(ctx, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Returning generic response at runtime
	ctx.JSON(http.StatusOK, result)
}

// PUT /leaves/:id

// UpdateLeave godoc
// @Summary Update a leave request
// @Description Update leave request details by ID
// @Tags Leaves
// @Accept json
// @Produce json
// @Param id path int true "Leave Request ID"
// @Param leave body models.LeaveRequest true "Leave Request object"
// @Success 200 {object} models.LeaveRequest
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /leaves/{id} [put]
func (c *LeaveController) UpdateLeave(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	var leave models.LeaveRequest
	if err := ctx.ShouldBindJSON(&leave); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leave.ID = id
	if err := c.service.UpdateLeave(ctx, &leave); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, leave)
}

// DELETE /leaves/:id
// DeleteLeave godoc
// @Summary Delete a leave request
// @Description Delete a leave request by ID
// @Tags Leaves
// @Param id path int true "Leave Request ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /leaves/{id} [delete]
func (c *LeaveController) DeleteLeave(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err := c.service.DeleteLeave(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// PATCH /leaves/:id/approve
// ApproveLeave godoc
// @Summary Approve a leave request
// @Description Approve leave request by ID
// @Tags Leaves
// @Param id path int true "Leave Request ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /leaves/{id}/approve [patch]
func (c *LeaveController) ApproveLeave(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err := c.service.ApproveLeave(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Leave approved"})
}

// PATCH /leaves/:id/reject
// RejectLeave godoc
// @Summary Reject a leave request
// @Description Reject leave request by ID
// @Tags Leaves
// @Param id path int true "Leave Request ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /leaves/{id}/reject [patch]

func (c *LeaveController) RejectLeave(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err := c.service.RejectLeave(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Leave rejected"})
}

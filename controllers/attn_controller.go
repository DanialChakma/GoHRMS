package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mod/services"
)

type AttendanceController struct {
	service services.AttendanceService
}

func NewAttendanceController(s services.AttendanceService) *AttendanceController {
	return &AttendanceController{service: s}
}

// POST /attendance/checkin/:employee_id
// CheckIn godoc
// @Summary Employee Check-In
// @Description Record check-in time for an employee
// @Tags Attendance
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Success 200 {object} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/checkin/{employee_id} [post]
func (c *AttendanceController) CheckIn(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Param("employee_id"), 10, 64)

	attendance, err := c.service.CheckIn(ctx, employeeID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, attendance)
}

// POST /attendance/checkout/:employee_id
// CheckOut godoc
// @Summary Employee Check-Out
// @Description Record check-out time for an employee
// @Tags Attendance
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Success 200 {object} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/checkout/{employee_id} [post]
func (c *AttendanceController) CheckOut(ctx *gin.Context) {
	employeeID, _ := strconv.ParseUint(ctx.Param("employee_id"), 10, 64)

	attendance, err := c.service.CheckOut(ctx, employeeID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, attendance)
}

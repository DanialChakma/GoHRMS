package dto

import "go.mod/models"

// EmployeeListResponse is used only for Swagger documentation
type PayrollListResponse struct {
	Data      []models.Payroll `json:"data"`
	Page      int              `json:"page"`
	PageSize  int              `json:"page_size"`
	Total     int64            `json:"total"`
	TotalPage int64            `json:"total_page"`
}

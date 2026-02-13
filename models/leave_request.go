package models

import (
	"time"
)

// LeaveRequest represents a leave request
type LeaveRequest struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID uint64    `gorm:"not null" json:"employee_id"`
	StartDate  time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate    time.Time `gorm:"type:date;not null" json:"end_date"`
	LeaveType  string    `gorm:"size:50;not null" json:"leave_type"`
	Status     string    `gorm:"size:50;not null" json:"status"`
}

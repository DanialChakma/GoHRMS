package models

import (
	"time"
)

// Attendance represents employee attendance
type Attendance struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID uint64     `gorm:"not null" json:"employee_id"`
	Date       time.Time  `gorm:"type:date;not null" json:"date"`
	CheckIn    *time.Time `gorm:"type:time" json:"check_in,omitempty"`
	CheckOut   *time.Time `gorm:"type:time" json:"check_out,omitempty"`
}

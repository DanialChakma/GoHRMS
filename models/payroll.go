package models

import (
	"time"
)

// Payroll represents employee payroll
type Payroll struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID uint64    `gorm:"not null" json:"employee_id"`
	Month      time.Time `gorm:"type:date;not null" json:"month"` // could be first day of month
	BaseSalary float64   `gorm:"not null" json:"base_salary"`
	Bonus      float64   `gorm:"default:0" json:"bonus"`
	Deductions float64   `gorm:"default:0" json:"deductions"`
	NetSalary  float64   `gorm:"not null" json:"net_salary"`
}

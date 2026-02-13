package models

import (
	"time"
)

// Employee represents an employee
type Employee struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeCode string    `gorm:"size:50;not null;unique" json:"employee_code"`
	FirstName    string    `gorm:"size:100;not null" json:"first_name"`
	LastName     string    `gorm:"size:100;not null" json:"last_name"`
	Email        string    `gorm:"size:100;not null;unique" json:"email"`
	Phone        *string   `gorm:"size:20" json:"phone,omitempty"`
	DepartmentID uint64    `gorm:"not null" json:"department_id"`
	JobTitleID   uint64    `gorm:"not null" json:"job_title_id"`
	HireDate     time.Time `gorm:"type:date;not null" json:"hire_date"`
	Status       string    `gorm:"size:50;not null" json:"status"`
}

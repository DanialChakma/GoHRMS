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

// Department represents a company department
type Department struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

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

// JobTitle represents a job title
type JobTitle struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"size:100;not null" json:"title"`
}

// LeaveRequest represents a leave request
type LeaveRequest struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID uint64    `gorm:"not null" json:"employee_id"`
	StartDate  time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate    time.Time `gorm:"type:date;not null" json:"end_date"`
	LeaveType  string    `gorm:"size:50;not null" json:"leave_type"`
	Status     string    `gorm:"size:50;not null" json:"status"`
}

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

// Role enum
type Role uint8

const (
	Admin        Role = 1
	Hr           Role = 2
	EmployeeRole Role = 3
	System       Role = 4
	ApiUser      Role = 5
)

// String converts a Role to a human-readable string
func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Hr:
		return "Hr"
	case EmployeeRole:
		return "Employee"
	case System:
		return "System"
	case ApiUser:
		return "ApiUser"
	default:
		return "Unknown"
	}
}

// FromID converts a numeric role_id to a Role enum
func RoleFromID(id uint8) Role {
	switch id {
	case 1:
		return Admin
	case 2:
		return Hr
	case 3:
		return EmployeeRole
	case 4:
		return System
	case 5:
		return ApiUser
	default:
		return 0 // or a special Unknown role
	}
}

// RoleFromString converts a string to a Role enum
func RoleFromString(name string) Role {
	switch name {
	case "Admin":
		return Admin
	case "Hr":
		return Hr
	case "Employee":
		return EmployeeRole
	case "System":
		return System
	case "ApiUser":
		return ApiUser
	default:
		return 0 // Unknown role
	}
}

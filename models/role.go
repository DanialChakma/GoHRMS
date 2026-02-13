package models

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

package models

import "time"

// type User struct {
// 	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
// 	Username     string             `bson:"username" json:"username"`
// 	Password     string             `bson:"password" json:"-"`
// 	Email        string             `bson:"email" json:"email"`
// 	Role         string             `bson:"role" json:"role"`
// 	TokenVersion int                `bson:"token_version" json:"-"`
// 	IsActive     bool               `bson:"is_active" json:"is_active"`
// 	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
// 	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
// }

// User represents system users
type User struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string    `gorm:"size:50;not null;unique" json:"username"`
	Password     string    `gorm:"size:255;not null" json:"password"`
	Email        string    `gorm:"size:100;not null;unique" json:"email"`
	RoleID       uint8     `gorm:"not null" json:"role_id"`
	TokenVersion int       `gorm:"token_version" json:"-"`
	EmployeeID   *uint64   `json:"employee_id,omitempty"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at" json:"updated_at"`
}

package models

// JobTitle represents a job title
type JobTitle struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"size:100;not null" json:"title"`
}

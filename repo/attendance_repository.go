package repo

import (
	"context"
	"time"

	"go.mod/models"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	FindByEmployeeAndDate(ctx context.Context, employeeID uint64, date time.Time) (*models.Attendance, error)
	Create(ctx context.Context, attendance *models.Attendance) error
	Update(ctx context.Context, attendance *models.Attendance) error
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

func (r *attendanceRepository) FindByEmployeeAndDate(
	ctx context.Context,
	employeeID uint64,
	date time.Time,
) (*models.Attendance, error) {

	var attendance models.Attendance

	err := r.db.WithContext(ctx).
		Where("employee_id = ? AND date = ?", employeeID, date).
		First(&attendance).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &attendance, nil
}

func (r *attendanceRepository) Create(ctx context.Context, attendance *models.Attendance) error {
	return r.db.WithContext(ctx).Create(attendance).Error
}

func (r *attendanceRepository) Update(ctx context.Context, attendance *models.Attendance) error {
	return r.db.WithContext(ctx).Save(attendance).Error
}

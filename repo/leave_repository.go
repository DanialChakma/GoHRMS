package repo

import (
	"context"

	"go.mod/models"
	"gorm.io/gorm"
)

type LeaveRepository interface {
	Create(ctx context.Context, leave *models.LeaveRequest) error
	GetByID(ctx context.Context, id uint64) (*models.LeaveRequest, error)
	Update(ctx context.Context, leave *models.LeaveRequest) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, offset, limit int) ([]models.LeaveRequest, int64, error)
	UpdateStatus(ctx context.Context, id uint64, status string) error
}

type leaveRepository struct {
	db *gorm.DB
}

func NewLeaveRepository(db *gorm.DB) LeaveRepository {
	return &leaveRepository{db: db}
}

func (r *leaveRepository) Create(ctx context.Context, leave *models.LeaveRequest) error {
	return r.db.WithContext(ctx).Create(leave).Error
}

func (r *leaveRepository) GetByID(ctx context.Context, id uint64) (*models.LeaveRequest, error) {
	var leave models.LeaveRequest
	err := r.db.WithContext(ctx).First(&leave, id).Error
	if err != nil {
		return nil, err
	}
	return &leave, nil
}

func (r *leaveRepository) Update(ctx context.Context, leave *models.LeaveRequest) error {
	return r.db.WithContext(ctx).Save(leave).Error
}

func (r *leaveRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&models.LeaveRequest{}, id).Error
}

func (r *leaveRepository) List(ctx context.Context, offset, limit int) ([]models.LeaveRequest, int64, error) {
	var leaves []models.LeaveRequest
	var total int64

	r.db.WithContext(ctx).Model(&models.LeaveRequest{}).Count(&total)

	err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&leaves).Error

	return leaves, total, err
}

func (r *leaveRepository) UpdateStatus(ctx context.Context, id uint64, status string) error {
	return r.db.WithContext(ctx).
		Model(&models.LeaveRequest{}).
		Where("id = ?", id).
		Update("status", status).Error
}

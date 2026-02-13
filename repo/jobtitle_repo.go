package repo

import (
	"context"

	"go.mod/models"
	"gorm.io/gorm"
)

type JobTitleRepository interface {
	Create(ctx context.Context, job *models.JobTitle) error
	GetByID(ctx context.Context, id uint64) (*models.JobTitle, error)
	Update(ctx context.Context, job *models.JobTitle) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, offset, limit int) ([]models.JobTitle, int64, error)
}

type jobTitleRepository struct {
	db *gorm.DB
}

func NewJobTitleRepository(db *gorm.DB) JobTitleRepository {
	return &jobTitleRepository{db: db}
}

func (r *jobTitleRepository) Create(ctx context.Context, job *models.JobTitle) error {
	return r.db.WithContext(ctx).Create(job).Error
}

func (r *jobTitleRepository) GetByID(ctx context.Context, id uint64) (*models.JobTitle, error) {
	var job models.JobTitle
	err := r.db.WithContext(ctx).First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *jobTitleRepository) Update(ctx context.Context, job *models.JobTitle) error {
	return r.db.WithContext(ctx).Save(job).Error
}

func (r *jobTitleRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&models.JobTitle{}, id).Error
}

func (r *jobTitleRepository) List(ctx context.Context, offset, limit int) ([]models.JobTitle, int64, error) {
	var jobs []models.JobTitle
	var total int64

	r.db.WithContext(ctx).Model(&models.JobTitle{}).Count(&total)

	err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&jobs).Error

	return jobs, total, err
}

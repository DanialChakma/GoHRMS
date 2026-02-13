package repo

import (
	"context"

	"go.mod/models"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(ctx context.Context, dept *models.Department) error
	GetByID(ctx context.Context, id uint64) (*models.Department, error)
	Update(ctx context.Context, dept *models.Department) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, offset, limit int) ([]models.Department, int64, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Create(ctx context.Context, dept *models.Department) error {
	return r.db.WithContext(ctx).Create(dept).Error
}

func (r *departmentRepository) GetByID(ctx context.Context, id uint64) (*models.Department, error) {
	var dept models.Department
	err := r.db.WithContext(ctx).First(&dept, id).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *departmentRepository) Update(ctx context.Context, dept *models.Department) error {
	return r.db.WithContext(ctx).Save(dept).Error
}

func (r *departmentRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&models.Department{}, id).Error
}

func (r *departmentRepository) List(ctx context.Context, offset, limit int) ([]models.Department, int64, error) {
	var depts []models.Department
	var total int64

	r.db.WithContext(ctx).Model(&models.Department{}).Count(&total)

	err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&depts).Error

	return depts, total, err
}

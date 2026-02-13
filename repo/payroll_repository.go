package repo

import (
	"context"

	"go.mod/models"
	"gorm.io/gorm"
)

type PayrollRepository interface {
	Create(ctx context.Context, payroll *models.Payroll) error
	GetByID(ctx context.Context, id uint64) (*models.Payroll, error)
	Update(ctx context.Context, payroll *models.Payroll) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, offset, limit int, employeeID *uint64) ([]models.Payroll, int64, error)
}

type payrollRepository struct {
	db *gorm.DB
}

func NewPayrollRepository(db *gorm.DB) PayrollRepository {
	return &payrollRepository{db: db}
}

func (r *payrollRepository) Create(ctx context.Context, payroll *models.Payroll) error {
	return r.db.WithContext(ctx).Create(payroll).Error
}

func (r *payrollRepository) GetByID(ctx context.Context, id uint64) (*models.Payroll, error) {
	var payroll models.Payroll
	err := r.db.WithContext(ctx).First(&payroll, id).Error
	if err != nil {
		return nil, err
	}
	return &payroll, nil
}

func (r *payrollRepository) Update(ctx context.Context, payroll *models.Payroll) error {
	return r.db.WithContext(ctx).Save(payroll).Error
}

func (r *payrollRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&models.Payroll{}, id).Error
}

func (r *payrollRepository) List(ctx context.Context, offset, limit int, employeeID *uint64) ([]models.Payroll, int64, error) {
	var payrolls []models.Payroll
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Payroll{})

	if employeeID != nil {
		query = query.Where("employee_id = ?", *employeeID)
	}

	query.Count(&total)

	err := query.
		Order("month desc").
		Offset(offset).
		Limit(limit).
		Find(&payrolls).Error

	return payrolls, total, err
}

package services

import (
	"context"
	"errors"

	"go.mod/dto"
	"go.mod/models"
	"go.mod/repo"
)

type PayrollService interface {
	CreatePayroll(ctx context.Context, payroll *models.Payroll) error
	GetPayroll(ctx context.Context, id uint64) (*models.Payroll, error)
	UpdatePayroll(ctx context.Context, payroll *models.Payroll) error
	DeletePayroll(ctx context.Context, id uint64) error
	ListPayrolls(ctx context.Context, page, pageSize int, employeeID *uint64) (dto.PaginatedResponse[models.Payroll], error)
}

type payrollService struct {
	repo repo.PayrollRepository
}

func NewPayrollService(repo repo.PayrollRepository) PayrollService {
	return &payrollService{repo: repo}
}

func (s *payrollService) calculateNetSalary(p *models.Payroll) {
	p.NetSalary = p.BaseSalary + p.Bonus - p.Deductions
}

func (s *payrollService) CreatePayroll(ctx context.Context, payroll *models.Payroll) error {
	if payroll.BaseSalary < 0 {
		return errors.New("base salary cannot be negative")
	}

	s.calculateNetSalary(payroll)
	return s.repo.Create(ctx, payroll)
}

func (s *payrollService) GetPayroll(ctx context.Context, id uint64) (*models.Payroll, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *payrollService) UpdatePayroll(ctx context.Context, payroll *models.Payroll) error {
	s.calculateNetSalary(payroll)
	return s.repo.Update(ctx, payroll)
}

func (s *payrollService) DeletePayroll(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *payrollService) ListPayrolls(ctx context.Context, page, pageSize int, employeeID *uint64) (dto.PaginatedResponse[models.Payroll], error) {
	offset := (page - 1) * pageSize
	payrolls, total, err := s.repo.List(ctx, offset, pageSize, employeeID)
	if err != nil {
		return dto.PaginatedResponse[models.Payroll]{}, err
	}

	return dto.PaginatedResponse[models.Payroll]{
		Data:      payrolls,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: (total + int64(pageSize) - 1) / int64(pageSize),
	}, nil

	// return s.repo.List(ctx, offset, pageSize, employeeID)
}

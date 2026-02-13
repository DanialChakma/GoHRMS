package services

import (
	"context"

	"go.mod/dto"
	"go.mod/models"
	"go.mod/repo"
)

type DepartmentService interface {
	Create(ctx context.Context, dept *models.Department) error
	Get(ctx context.Context, id uint64) (*models.Department, error)
	Update(ctx context.Context, dept *models.Department) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, page, pageSize int) (dto.PaginatedResponse[models.Department], error)
}

type departmentService struct {
	repo repo.DepartmentRepository
}

func NewDepartmentService(repo repo.DepartmentRepository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) Create(ctx context.Context, dept *models.Department) error {
	return s.repo.Create(ctx, dept)
}

func (s *departmentService) Get(ctx context.Context, id uint64) (*models.Department, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *departmentService) Update(ctx context.Context, dept *models.Department) error {
	return s.repo.Update(ctx, dept)
}

func (s *departmentService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *departmentService) List(ctx context.Context, page, pageSize int) (dto.PaginatedResponse[models.Department], error) {
	offset := (page - 1) * pageSize
	departments, total, err := s.repo.List(ctx, offset, pageSize)
	if err != nil {
		return dto.PaginatedResponse[models.Department]{}, err
	}

	return dto.PaginatedResponse[models.Department]{
		Data:      departments,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: (total + int64(pageSize) - 1) / int64(pageSize),
	}, nil
}

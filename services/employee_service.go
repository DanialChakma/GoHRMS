package services

import (
	"go.mod/dto"
	"go.mod/models"
	"go.mod/repo"
)

type EmployeeService interface {
	CreateEmployee(employee *models.Employee) error
	GetEmployee(id uint64) (*models.Employee, error)
	UpdateEmployee(employee *models.Employee) error
	DeleteEmployee(id uint64) error
	ListEmployees(page int, pageSize int) (dto.PaginatedResponse[models.Employee], error)
}

type employeeService struct {
	repo repo.EmployeeRepository
}

func NewEmployeeService(repo repo.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(employee *models.Employee) error {
	return s.repo.Create(employee)
}

func (s *employeeService) GetEmployee(id uint64) (*models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) UpdateEmployee(employee *models.Employee) error {
	return s.repo.Update(employee)
}

func (s *employeeService) DeleteEmployee(id uint64) error {
	return s.repo.Delete(id)
}

// func (s *employeeService) ListEmployees(page int, pageSize int) ([]models.Employee, int64, error) {
// 	offset := (page - 1) * pageSize
// 	return s.repo.List(offset, pageSize)
// }

func (s *employeeService) ListEmployees(page, pageSize int) (dto.PaginatedResponse[models.Employee], error) {

	offset := (page - 1) * pageSize

	employees, total, err := s.repo.List(offset, pageSize)
	if err != nil {
		return dto.PaginatedResponse[models.Employee]{}, err
	}

	return dto.PaginatedResponse[models.Employee]{
		Data:      employees,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: (total + int64(pageSize) - 1) / int64(pageSize),
	}, nil
}

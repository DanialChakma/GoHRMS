package repo

import (
	"go.mod/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) error
	GetByID(id uint64) (*models.Employee, error)
	Update(employee *models.Employee) error
	Delete(id uint64) error
	List(offset int, limit int) ([]models.Employee, int64, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) GetByID(id uint64) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepository) Delete(id uint64) error {
	return r.db.Delete(&models.Employee{}, id).Error
}

func (r *employeeRepository) List(offset int, limit int) ([]models.Employee, int64, error) {
	var employees []models.Employee
	var total int64

	r.db.Model(&models.Employee{}).Count(&total)

	err := r.db.Offset(offset).Limit(limit).Find(&employees).Error
	if err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

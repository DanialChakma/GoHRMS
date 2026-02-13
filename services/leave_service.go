package services

import (
	"context"
	"errors"

	"go.mod/dto"
	"go.mod/models"
	"go.mod/repo"
)

type LeaveService interface {
	CreateLeave(ctx context.Context, leave *models.LeaveRequest) error
	GetLeave(ctx context.Context, id uint64) (*models.LeaveRequest, error)
	UpdateLeave(ctx context.Context, leave *models.LeaveRequest) error
	DeleteLeave(ctx context.Context, id uint64) error
	ListLeaves(ctx context.Context, page, pageSize int) (dto.PaginatedResponse[models.LeaveRequest], error)
	ApproveLeave(ctx context.Context, id uint64) error
	RejectLeave(ctx context.Context, id uint64) error
}

type leaveService struct {
	repo repo.LeaveRepository
}

func NewLeaveService(repo repo.LeaveRepository) LeaveService {
	return &leaveService{repo: repo}
}

func (s *leaveService) CreateLeave(ctx context.Context, leave *models.LeaveRequest) error {
	if leave.EndDate.Before(leave.StartDate) {
		return errors.New("end_date cannot be before start_date")
	}

	leave.Status = "pending"
	return s.repo.Create(ctx, leave)
}

func (s *leaveService) GetLeave(ctx context.Context, id uint64) (*models.LeaveRequest, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *leaveService) UpdateLeave(ctx context.Context, leave *models.LeaveRequest) error {
	return s.repo.Update(ctx, leave)
}

func (s *leaveService) DeleteLeave(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

// func (s *leaveService) ListLeaves(ctx context.Context, page, pageSize int) ([]models.LeaveRequest, int64, error) {
// 	offset := (page - 1) * pageSize
// 	return s.repo.List(ctx, offset, pageSize)
// }

func (s *leaveService) ListLeaves(ctx context.Context, page, pageSize int) (dto.PaginatedResponse[models.LeaveRequest], error) {

	offset := (page - 1) * pageSize

	leaves, total, err := s.repo.List(ctx, offset, pageSize)
	if err != nil {
		return dto.PaginatedResponse[models.LeaveRequest]{}, err
	}

	return dto.PaginatedResponse[models.LeaveRequest]{
		Data:      leaves,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: (total + int64(pageSize) - 1) / int64(pageSize),
	}, nil
}

func (s *leaveService) ApproveLeave(ctx context.Context, id uint64) error {
	return s.repo.UpdateStatus(ctx, id, "approved")
}

func (s *leaveService) RejectLeave(ctx context.Context, id uint64) error {
	return s.repo.UpdateStatus(ctx, id, "rejected")
}

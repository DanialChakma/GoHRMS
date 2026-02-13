package services

import (
	"context"

	"go.mod/dto"
	"go.mod/models"
	"go.mod/repo"
)

type JobTitleService interface {
	Create(ctx context.Context, job *models.JobTitle) error
	Get(ctx context.Context, id uint64) (*models.JobTitle, error)
	Update(ctx context.Context, job *models.JobTitle) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, page, pageSize int) (dto.PaginatedResponse[models.JobTitle], error)
}

type jobTitleService struct {
	repo repo.JobTitleRepository
}

func NewJobTitleService(repo repo.JobTitleRepository) JobTitleService {
	return &jobTitleService{repo: repo}
}

func (s *jobTitleService) Create(ctx context.Context, job *models.JobTitle) error {
	return s.repo.Create(ctx, job)
}

func (s *jobTitleService) Get(ctx context.Context, id uint64) (*models.JobTitle, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *jobTitleService) Update(ctx context.Context, job *models.JobTitle) error {
	return s.repo.Update(ctx, job)
}

func (s *jobTitleService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *jobTitleService) List(ctx context.Context, page, pageSize int) (dto.PaginatedResponse[models.JobTitle], error) {
	offset := (page - 1) * pageSize

	datas, total, err := s.repo.List(ctx, offset, pageSize)
	if err != nil {
		return dto.PaginatedResponse[models.JobTitle]{}, err
	}

	return dto.PaginatedResponse[models.JobTitle]{
		Data:      datas,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: (total + int64(pageSize) - 1) / int64(pageSize),
	}, nil

}

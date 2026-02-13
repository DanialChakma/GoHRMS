package services

import (
	"context"
	"errors"
	"time"

	"go.mod/models"
	"go.mod/repo"
)

type AttendanceService interface {
	CheckIn(ctx context.Context, employeeID uint64) (*models.Attendance, error)
	CheckOut(ctx context.Context, employeeID uint64) (*models.Attendance, error)
}

type attendanceService struct {
	repo repo.AttendanceRepository
}

func NewAttendanceService(repo repo.AttendanceRepository) AttendanceService {
	return &attendanceService{repo: repo}
}

// CheckIn logic
func (s *attendanceService) CheckIn(ctx context.Context, employeeID uint64) (*models.Attendance, error) {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)

	existing, err := s.repo.FindByEmployeeAndDate(ctx, employeeID, today)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		if existing.CheckIn != nil {
			return nil, errors.New("already checked in today")
		}
		existing.CheckIn = &now
		err = s.repo.Update(ctx, existing)
		return existing, err
	}

	attendance := &models.Attendance{
		EmployeeID: employeeID,
		Date:       today,
		CheckIn:    &now,
	}

	err = s.repo.Create(ctx, attendance)
	return attendance, err
}

// CheckOut logic
func (s *attendanceService) CheckOut(ctx context.Context, employeeID uint64) (*models.Attendance, error) {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)

	existing, err := s.repo.FindByEmployeeAndDate(ctx, employeeID, today)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		return nil, errors.New("no check-in found for today")
	}

	if existing.CheckIn == nil {
		return nil, errors.New("cannot checkout without check-in")
	}

	if existing.CheckOut != nil {
		return nil, errors.New("already checked out today")
	}

	existing.CheckOut = &now
	err = s.repo.Update(ctx, existing)

	return existing, err
}

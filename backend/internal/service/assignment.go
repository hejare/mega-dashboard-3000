package service

import "github.com/hejare/mega-dashboard-3000/internal/repository"

type AssignmentRepository interface {
	CreateAssignment(*repository.CreateAssignmentData) error
}

type AssignmentService struct {
	repository AssignmentRepository
}

func NewAssignmentService(repo AssignmentRepository) *AssignmentService {
	return &AssignmentService{
		repository: repo,
	}
}

func (s *AssignmentService) CreateAssignment(data *repository.CreateAssignmentData) error {
	return s.repository.CreateAssignment(data)
}

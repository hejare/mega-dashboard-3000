package service

import "github.com/hejare/mega-dashboard-3000/internal/repository"

type LeadRepository interface {
	CreateLead(*repository.CreateLeadData) error
	Search(query string) ([]repository.LeadWithConsultantID, error)
}

type LeadService struct {
	repository LeadRepository
}

func NewLeadService(repo LeadRepository) *LeadService {
	return &LeadService{
		repository: repo,
	}
}

func (s *LeadService) CreateLead(data *repository.CreateLeadData) error {
	return s.repository.CreateLead(data)
}

func (s *LeadService) Search(query string) ([]repository.LeadWithConsultantID, error) {
	return s.repository.Search(query)
}

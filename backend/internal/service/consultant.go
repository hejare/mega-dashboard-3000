package service

import "github.com/hejare/mega-dashboard-3000/internal/repository"

type ConsultantRepository interface {
	GetAllConsultants() ([]repository.Consultant, error)
	GetConsultant(id int) (*repository.Consultant, error)
}

// Consultant + all leads + assignment data (currently not all that, but that's the purpose)
type RichConsultant struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ChangeStatus string `json:"change_status"`
}

type ConsultantService struct {
	repository ConsultantRepository
}

func NewConsultantService(repo ConsultantRepository) *ConsultantService {
	return &ConsultantService{
		repository: repo,
	}
}

func (s *ConsultantService) GetConsultant(id int) (*repository.Consultant, error) {
	consultant, err := s.repository.GetConsultant(id)
	if err != nil {
		return nil, err
	}
	return consultant, nil
}

func (s *ConsultantService) GetAllConsultants() ([]RichConsultant, error) {
	consultants, err := s.repository.GetAllConsultants()
	if err != nil {
		return nil, err
	}
	richConsultants := []RichConsultant{}
	for _, c := range consultants {
		richConsultants = append(richConsultants, RichConsultant{
			Id:           c.Id,
			Name:         c.Name,
			ChangeStatus: c.ChangeStatus,
		})
	}
	return richConsultants, nil
}

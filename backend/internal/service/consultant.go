package service

type ConsultantRepository interface {
	GetAllConsultants()
}

type ConsultantService struct {
	repository ConsultantRepository
}

func NewConsultantService(repo ConsultantRepository) *ConsultantService {
	return &ConsultantService{
		repository: repo,
	}
}

func (s *ConsultantService) GetAllConsultants() {}

package service

type ConsultantService struct{}

func NewConsultantService() *ConsultantService {
	return &ConsultantService{}
}

func (s *ConsultantService) GetAllConsultants() {}

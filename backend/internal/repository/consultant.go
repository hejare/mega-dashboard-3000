package repository

type ConsultantRepository struct{}

func NewConsultantRepository() *ConsultantRepository {
	return &ConsultantRepository{}
}

func (r *ConsultantRepository) GetAllConsultants() {}

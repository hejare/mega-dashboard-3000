package service

type DashboardRepository interface {
}

type DashboardService struct {
	repository DashboardRepository
}

func NewDashboardService(repo DashboardRepository) *DashboardService {
	return &DashboardService{
		repository: repo,
	}
}

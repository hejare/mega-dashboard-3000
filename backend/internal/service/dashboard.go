package service

import (
	"github.com/hejare/mega-dashboard-3000/internal/repository"
)

type DashboardRepository interface {
	GetDashboard() ([]repository.DashboardRow, error)
}

type DashboardService struct {
	repository DashboardRepository
}

type RichDashboardRow struct {
	Consultant *repository.Consultant `json:"consultant"`
	Leads      []repository.Lead      `json:"leads"`
	Assignment *repository.Assignment `json:"assignment,omitempty"`
}

func NewDashboardService(repo DashboardRepository) *DashboardService {
	return &DashboardService{
		repository: repo,
	}
}

func (s *DashboardService) GetDashboard() ([]RichDashboardRow, error) {
	dashboard, err := s.repository.GetDashboard()
	if err != nil {
		return nil, err
	}

	richDashboard := []RichDashboardRow{}
	for _, dRow := range dashboard {
		var assignment *repository.Assignment
		if dRow.AssignmentId != nil {
			assignment = &repository.Assignment{
				Id:           *dRow.AssignmentId,
				Organization: *dRow.AssignmentOrganization,
				PeriodEndAt:  *dRow.AssignmentEndDate,
				Role:         *dRow.AssignmentRole,
			}
		}
		richDashboard = append(richDashboard, RichDashboardRow{
			Consultant: &repository.Consultant{
				Id:                      dRow.ConsultantId,
				Name:                    dRow.ConsultantName,
				ChangeStatus:            dRow.ChangeStatus,
				ProbableExtensionStatus: dRow.ExtensionStatus,
			},
			Assignment: assignment,
			Leads:      dRow.Leads,
		})
	}

	return richDashboard, nil
}

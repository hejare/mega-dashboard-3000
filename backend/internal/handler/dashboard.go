package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

type DashboardService interface{}

type DashboardHandler struct {
	service DashboardService
}

type TmpDummyDashboardRow struct {
	ConsultantName    string    `json:"consultant_name"`
	CurrentAssignment string    `json:"current_assignment"`
	AvailableFrom     time.Time `json:"available_from"`
	ExtensionStatus   string    `json:"extension_status"`
	ChangeStatus      string    `json:"change_status"`
	Leads             []string  `json:"leads"`
}

func NewDashboardHandler(service DashboardService) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}

func (h *DashboardHandler) GetDashboard(ctx *gin.Context) {
	mockDashboard := []TmpDummyDashboardRow{
		{
			ConsultantName:    "Carl Calle L",
			CurrentAssignment: "Relex",
			AvailableFrom:     time.Now(),
			ChangeStatus:      "NOT_LOOKING",
			Leads:             []string{"Avalanche", "Klarna", "Spotify"},
		},

		{
			ConsultantName:    "Calle Carl L",
			CurrentAssignment: "Xeler",
			AvailableFrom:     time.Now(),
			ChangeStatus:      "LOOKING",
			Leads:             []string{"Hej", "Hejare", "Tjenare"},
		},

		{
			ConsultantName:    "L Calle Carl",
			CurrentAssignment: "Rexel",
			AvailableFrom:     time.Now(),
			ChangeStatus:      "NOT_LOOKING",
			Leads:             []string{"Spotify", "AAA", "BBBBBBBBBBB", "CCCCCCCC", "ADLKJFLADKJFJLKDASJF"},
		},
	}
	ctx.JSON(200, mockDashboard)
}

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
	ConsultantName    string    `json:"consultantName"`
	CurrentAssignment string    `json:"currentAssignment"`
	AvailableFrom     time.Time `json:"availableFrom"`
	ExtensionStatus   string    `json:"extensionStatus"`
	ChangeStatus      string    `json:"changeStatus"`
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
			ChangeStatus:      "YELLOW",
			ExtensionStatus:   "",
			Leads:             []string{"Avalanche", "Klarna", "Spotify"},
		},

		{
			ConsultantName:    "Calle Carl L",
			CurrentAssignment: "Xeler",
			AvailableFrom:     time.Now(),
			ChangeStatus:      "RED",
			ExtensionStatus:   "Ja",
			Leads:             []string{"Hej", "Hejare", "Tjenare"},
		},

		{
			ConsultantName:    "L Calle Carl",
			CurrentAssignment: "Rexel",
			AvailableFrom:     time.Now(),
			ChangeStatus:      "YELLOW",
			ExtensionStatus:   "Troligtvis",
			Leads:             []string{"Spotify", "AAA", "BBBBBBBBBBB", "CCCCCCCC", "ADLKJFLADKJFJLKDASJF"},
		},
	}
	ctx.JSON(200, mockDashboard)
}

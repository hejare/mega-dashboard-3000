package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hejare/mega-dashboard-3000/internal/service"
)

type DashboardService interface {
	GetDashboard() ([]service.RichDashboardRow, error)
}

type DashboardHandler struct {
	service DashboardService
}

func NewDashboardHandler(service DashboardService) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}

func (h *DashboardHandler) GetDashboard(ctx *gin.Context) {
	dashboard, err := h.service.GetDashboard()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, dashboard)
}

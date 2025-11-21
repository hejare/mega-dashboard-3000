package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hejare/mega-dashboard-3000/internal/service"
)

type ConsultantService interface {
	GetAllConsultants() ([]service.RichConsultant, error)
}

type ConsultantHandler struct {
	service ConsultantService
}

func NewConsultantHandler(service ConsultantService) *ConsultantHandler {
	return &ConsultantHandler{service: service}
}

func (h *ConsultantHandler) Get(ctx *gin.Context) {
	consultants, err := h.service.GetAllConsultants()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, consultants)
}

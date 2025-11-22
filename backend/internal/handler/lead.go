package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hejare/mega-dashboard-3000/internal/repository"
)

type LeadService interface {
	CreateLead(*repository.CreateLeadData) error
}

type LeadHandler struct {
	service LeadService
}

func NewLeadHandler(service LeadService) *LeadHandler {
	return &LeadHandler{
		service: service,
	}
}

func (h *LeadHandler) Post(ctx *gin.Context) {
	var data repository.CreateLeadData

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(400, err)
		return
	}
	if err := h.service.CreateLead(&data); err != nil {
		ctx.JSON(500, err) // TODO: Make more fine grained handling of errors?
		return
	}
	ctx.Status(200)
}

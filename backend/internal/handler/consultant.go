package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hejare/mega-dashboard-3000/internal/repository"
	"github.com/hejare/mega-dashboard-3000/internal/service"
)

type ConsultantService interface {
	GetAllConsultants() ([]service.RichConsultant, error)
	GetConsultant(id int) (*repository.Consultant, error)
}

type ConsultantHandler struct {
	service ConsultantService
}

func NewConsultantHandler(service ConsultantService) *ConsultantHandler {
	return &ConsultantHandler{service: service}
}

func (h *ConsultantHandler) GetAll(ctx *gin.Context) {
	consultants, err := h.service.GetAllConsultants()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, consultants)
}

func (h *ConsultantHandler) GetOne(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	consultant, err := h.service.GetConsultant(id)
	if err != nil {
		log.Printf("internal server error: %v\n", err)
		ctx.JSON(500, err)
		return
	}

	if consultant == nil {
		ctx.Status(404)
		return
	}
	ctx.JSON(200, consultant)
}

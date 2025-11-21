package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ConsultantService interface {
	GetAllConsultants()
}

type ConsultantHandler struct {
	service ConsultantService
}

func NewConsultantHandler(service ConsultantService) *ConsultantHandler {
	return &ConsultantHandler{service: service}
}

func (h *ConsultantHandler) Get(ctx *gin.Context) {
	fmt.Println("All the consultants")
}

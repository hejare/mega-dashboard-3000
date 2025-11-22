package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hejare/mega-dashboard-3000/internal/repository"
)

type AssignmentService interface {
	CreateAssignment(*repository.CreateAssignmentData) error
}

type AssignmentHandler struct {
	service AssignmentService
}

func NewAssignmentHandler(service AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{
		service: service,
	}
}

func (h *AssignmentHandler) Post(ctx *gin.Context) {
	var data repository.CreateAssignmentData

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(400, err)
		return
	}
	if err := h.service.CreateAssignment(&data); err != nil {
		ctx.JSON(500, err) // TODO: Make more fine grained handling of errors?
		return
	}
	ctx.Status(200)
}

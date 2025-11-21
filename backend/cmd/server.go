package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hejare/mega-dashboard-3000/internal/handler"
	"github.com/hejare/mega-dashboard-3000/internal/repository"
	"github.com/hejare/mega-dashboard-3000/internal/service"
)

func main() {

	consultantRepo := repository.NewConsultantRepository()
	consultantService := service.NewConsultantService(consultantRepo)
	consultantHandler := handler.NewConsultantHandler(consultantService)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin in mega-dashboard-3000!",
		})
	})

	r.GET("/consultant", consultantHandler.Get)

	r.Run(":8080")
}

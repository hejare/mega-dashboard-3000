package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hejare/mega-dashboard-3000/clients"
	"github.com/hejare/mega-dashboard-3000/internal/handler"
	"github.com/hejare/mega-dashboard-3000/internal/repository"
	"github.com/hejare/mega-dashboard-3000/internal/service"
)

func main() {

	db := clients.NewDBClient()
	consultantRepo := repository.NewConsultantRepository(db)
	consultantService := service.NewConsultantService(consultantRepo)
	consultantHandler := handler.NewConsultantHandler(consultantService)

	dashboardRepo := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(dashboardRepo)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin in mega-dashboard-3000!",
		})
	})

	r.GET("/consultant", consultantHandler.GetAll)
	r.GET("/consultant/:id", consultantHandler.GetOne)

	r.GET("/dashboard", dashboardHandler.GetDashboard)

	r.Run(":8080")
}

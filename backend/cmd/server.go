package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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

	leadRepo := repository.NewLeadRepository(db)
	leadService := service.NewLeadService(leadRepo)
	leadHandler := handler.NewLeadHandler(leadService)

	assingmentRepo := repository.NewAssignmentRepository(db)
	assingmentService := service.NewAssignmentService(assingmentRepo)
	assignmentHandler := handler.NewAssignmentHandler(assingmentService)

	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin in mega-dashboard-3000!",
		})
	})

	r.GET("/consultant", consultantHandler.GetAll)
	r.GET("/consultant/:id", consultantHandler.GetOne)

	r.GET("/dashboard", dashboardHandler.GetDashboard)

	r.POST("/lead", leadHandler.Post)

	r.POST("/assignment", assignmentHandler.Post)

	r.Run(":8080")
}

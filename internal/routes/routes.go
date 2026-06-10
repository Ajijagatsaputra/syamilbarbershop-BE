package routes

import (
	"syamilbarbershop-be/internal/handlers"
	"syamilbarbershop-be/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Public routes
	api := r.Group("/api")
	{
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)
		api.GET("/services", handlers.GetServices)
		api.POST("/appointments", handlers.CreateAppointment)
	}

	// Protected routes
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// Services
		auth.POST("/services", handlers.CreateService)
		auth.PUT("/services/:id", handlers.UpdateService)
		auth.DELETE("/services/:id", handlers.DeleteService)

		// Customers
		auth.GET("/customers", handlers.GetCustomers)
		auth.POST("/customers", handlers.CreateCustomer)
		auth.PUT("/customers/:id", handlers.UpdateCustomer)
		auth.DELETE("/customers/:id", handlers.DeleteCustomer)

		// Appointments
		auth.GET("/appointments", handlers.GetAppointments)
		auth.PUT("/appointments/:id", handlers.UpdateAppointment)
		auth.PUT("/appointments/:id/status", handlers.UpdateAppointmentStatus)
		auth.DELETE("/appointments/:id", handlers.DeleteAppointment)
	}

	return r
}

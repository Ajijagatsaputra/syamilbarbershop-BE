package handlers

import (
	"net/http"
	"syamilbarbershop-be/internal/database"
	"syamilbarbershop-be/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment
	if err := database.DB.Preload("Customer").Preload("Service").Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func CreateAppointment(c *gin.Context) {
	var req struct {
		ServiceID     uint   `json:"service_id" binding:"required"`
		CustomerName  string `json:"customer_name"`
		CustomerPhone string `json:"customer_phone" binding:"required"`
		Date          string `json:"date" binding:"required"`
		Time          string `json:"time" binding:"required"`
		PaymentMethod string `json:"payment_method"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Find or Create Customer
	var customer models.Customer
	if err := database.DB.Where("phone = ?", req.CustomerPhone).First(&customer).Error; err != nil {
		customer = models.Customer{
			Name:  req.CustomerName,
			Phone: req.CustomerPhone,
		}
		if err := database.DB.Create(&customer).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
			return
		}
	}

	// 2. Parse Date
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	// 3. Find Service for Price
	var service models.Service
	if err := database.DB.First(&service, req.ServiceID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Service not found"})
		return
	}

	// 4. Create Appointment
	appointment := models.Appointment{
		CustomerID:    customer.ID,
		ServiceID:     service.ID,
		Date:          models.JSONDate(parsedDate),
		Time:          req.Time,
		Status:        models.StatusPending,
		PaymentMethod: req.PaymentMethod,
		Price:         service.Price,
	}

	if err := database.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Preload for response
	database.DB.Preload("Customer").Preload("Service").First(&appointment, appointment.ID)
	c.JSON(http.StatusCreated, appointment)
}

func UpdateAppointmentStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status models.AppointmentStatus `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&models.Appointment{}).Where("id = ?", id).Update("status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment status updated"})
}

func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Appointment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted"})
}

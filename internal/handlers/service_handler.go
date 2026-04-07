package handlers

import (
	"net/http"
	"syamilbarbershop-be/internal/database"
	"syamilbarbershop-be/internal/models"

	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
	var services []models.Service
	if err := database.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, service)
}

func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := database.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&service)
	c.JSON(http.StatusOK, service)
}

func DeleteService(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Service{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
}

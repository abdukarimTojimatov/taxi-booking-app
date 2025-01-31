package controllers

import (
	"net/http"
	"taxi-booking-app/models"
	"taxi-booking-app/services"

	"github.com/gin-gonic/gin"
)

type DriverController struct {
	driverService services.DriverService
}

func NewDriverController(driverService services.DriverService) *DriverController {
	return &DriverController{driverService: driverService}
}

// UpdateDriverLocation - Haydovchining joylashuvini yangilash
func (dc *DriverController) UpdateDriverLocation(c *gin.Context) {
	var location struct {
		DriverID string             `json:"driver_id"`
		Location models.GeoJSON     `json:"location"`
	}

	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Joylashuvni yangilash
	err := dc.driverService.UpdateDriverLocation(location.DriverID, location.Location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Driver location updated successfully"})
}
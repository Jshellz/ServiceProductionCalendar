package controllers

import (
	"ServiceProductionCalendar/initializers"
	"ServiceProductionCalendar/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HolidayCreate(c *gin.Context) {

	//var body struct {
	//	Name string
	//	Date time.Time
	//}

	//Create a data
	//holiday := models.Holiday{Name: body.Name, Date: body.Date}
	holiday := models.Holiday{Name: "test", Date: time.Date(2024, 12, 12, 12, 12, 12, 12, time.UTC)}

	result := initializers.DB.Create(&holiday)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"holiday": holiday,
	})
}

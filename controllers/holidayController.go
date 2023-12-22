package controllers

import (
	"ServiceProductionCalendar/initializers"
	"ServiceProductionCalendar/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HolidayCreate(c *gin.Context) {

	var body struct {
		Name string
		Date time.Time
	}

	//Create a data holiday
	holiday := models.Holiday{Name: body.Name, Date: body.Date}

	//test
	//holiday := models.Holiday{Name: "test1", Date: time.Date(2023, 12, 12, 12, 12, 12, 12, time.UTC)}

	result := initializers.DB.Create(&holiday)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"holiday": holiday,
	})
}

func GetAllHoliday(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Holidays)
}

func GetHoliday(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, id)
}
func UpdateHoliday(c *gin.Context) {}
func DeleteHoliday(c *gin.Context) {}

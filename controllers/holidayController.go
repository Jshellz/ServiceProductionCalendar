package controllers

import (
	"ServiceProductionCalendar/initializers"
	"ServiceProductionCalendar/models"
	"github.com/gin-gonic/gin"
	"log"
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

func UpdateHoliday(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
		Date time.Time
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	var holiday models.Holiday

	initializers.DB.First(&holiday, id)

	//holiday.ID = body.ID
	holiday.Name = body.Name
	holiday.Date = body.Date

	initializers.DB.Save(&holiday)

	c.JSON(http.StatusOK, &holiday)
}

func DeleteHoliday(c *gin.Context) {
	id := c.Param("id")
	var holiday models.Holiday

	if result := initializers.DB.First(&holiday, id); result.Error != nil {
		log.Fatal(c.AbortWithError(http.StatusNotFound, result.Error))
		return
	}

	initializers.DB.Delete(&holiday)

	c.Status(http.StatusOK)
}

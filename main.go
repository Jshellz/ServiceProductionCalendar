package main

import (
	"ServiceProductionCalendar/controllers"
	"ServiceProductionCalendar/initializers"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnToDB()
}

func main() {

	r := gin.Default()
	r.POST("/createHoliday", controllers.HolidayCreate)
	r.GET("/getHoliday")
	r.GET("/getHoliday/:id")
	r.PATCH("/updateHoliday")
	r.DELETE("/deleteHoliday/:id")

	log.Fatal(r.Run())
}

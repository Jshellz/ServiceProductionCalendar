package main

import (
	"ServiceProductionCalendar/spc"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()
	r.GET("/", spc.GetHoliday)
	r.POST("/holiday_create", spc.CreateHoliday)
	r.GET("/holiday/:id", spc.HolidayById)
	r.PATCH("/checkout_holiday", spc.CheckoutHoliday)
	r.DELETE("/holiday/:id", spc.DeleteHoliday)
	log.Fatal(r.Run("localhost:8080"))
}

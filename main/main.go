package main

import (
	"ServiceProductionCalendar/spc"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	errMigrate := db.AutoMigrate(&spc.Holiday{})
	if errMigrate != nil {
		panic("failed to migrate on db")
	}

	// Create inserts value
	db.Create(&spc.Holidays)

	r := gin.Default()
	r.GET("/", spc.GetHoliday)
	r.POST("/holiday_create", spc.CreateHoliday)
	r.GET("/holiday/:id", spc.HolidayById)
	r.PATCH("/checkout_holiday", spc.CheckoutHoliday)
	r.DELETE("/holiday/:id", spc.DeleteHoliday)
	log.Fatal(r.Run("localhost:8080"))
}

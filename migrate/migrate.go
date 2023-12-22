package main

import (
	"ServiceProductionCalendar/initializers"
	"ServiceProductionCalendar/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnToDB()
}

func main() {
	log.Fatal(initializers.DB.AutoMigrate(&models.Holiday{}))
}

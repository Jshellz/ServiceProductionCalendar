package main

import (
	"ServiceProductionCalendar/spc"
	pb "ServiceProductionCalendar/spc/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net"
	"time"
)

const (
	network       = "tcp"
	address       = ":50051"
	addressClient = "localhost:50051"
	defaultName   = "День к новому году"
	defaultDate   = "31.12.2023"
)

func main() {

	//runJSON()
	serverRPC()
	clientRPC()
	runBD()
}

func runBD() {
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
}

func runJSON() {
	r := gin.Default()
	//r.GET("/", spc.GetHoliday)
	//r.POST("/holiday_create", spc.CreateHoliday)
	//r.GET("/holiday/:id", spc.HolidayById)
	//r.PATCH("/checkout_holiday", spc.CheckoutHoliday)
	//r.DELETE("/holiday/:id", spc.DeleteHoliday)
	log.Fatal(r.Run("localhost:8080"))
}

func serverRPC() {
	lis, err := net.Listen(network, address)
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterSPCServer(s, &spc.Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}

func clientRPC() {
	conn, err := grpc.Dial(addressClient, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewSPCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	holidayName := defaultName
	holidayDate := defaultDate
	r, err := c.AddHoliday(ctx, &pb.AddHolidayRequest{Name: holidayName, Data: holidayDate})

	if err != nil {
		log.Fatalf("Error when calling AddHoliday: %v", err)
	}

	fmt.Printf("Response from server: %s and %s\n", r.HolidayName, r.HolidayData)
}

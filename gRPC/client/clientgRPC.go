package main

import (
	"ServiceProductionCalendar/gRPC/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

const serverAddr = "localhost:50051"

// Client gRPC
func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	client := proto.NewServiceProductionCalendarClient(conn)

	var body struct {
		ID   int32
		Name string
		Date string
	}

	addHoliday := &proto.AddHolidayRequest{Id: body.ID, Name: body.Name, Data: body.Date}
	create, err := client.HolidayCreate(context.Background(), addHoliday)

	if err != nil {
		log.Fatalf("failed to add holiday: %v", err)
	}
	log.Printf("AddHolidayRequest: %v %v", create.Name, create.Data)
	log.Println(create)

	getAllHoliday := &proto.GetAllHolidayRequest{}
	getAll, err := client.GetAllHoliday(context.Background(), getAllHoliday)

	if err != nil {
		log.Fatalf("failed to get all holidays: %v", err)
	}
	log.Println(getAll)

	getHoliday := &proto.GetHolidayRequest{}
	getHolidays, err := client.GetHoliday(context.Background(), getHoliday)

	if err != nil {
		log.Fatalf("failed to get holiday: %v", err)
	}
	log.Println(getHolidays)

	editHoliday := &proto.EditHolidayRequest{Id: body.ID, Name: body.Name, Data: body.Date}
	update, err := client.UpdateHoliday(context.Background(), editHoliday)

	if err != nil {
		log.Fatalf("failed to edit holiday: %v", err)
	}
	log.Printf("EditHolidayRequest: %v %v %v", update.Id, update.Name, update.Data)
	log.Println(update)

	deleteHoliday := &proto.DeleteHolidayRequest{Id: body.ID, Name: body.Name}
	deletes, err := client.DeleteHoliday(context.Background(), deleteHoliday)

	if err != nil {
		log.Fatalf("failed to delete holiday: %v", err)
	}
	log.Printf("DeleteHolidayRequest: %v %v", deletes.Id, deletes.Name)
	log.Println(deletes)

	if conn != nil {
		log.Fatal(conn.Close())
	}
}

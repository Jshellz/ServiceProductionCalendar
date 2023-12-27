package main

import (
	"ServiceProductionCalendar/gRPC/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

const serverAddr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer log.Fatal(conn.Close())

	client := proto.NewServiceProductionCalendarClient(conn)

	var body struct {
		ID   int32
		Name string
		Date string
	}

	reqTest := &proto.AddHolidayRequest{Id: body.ID, Name: body.Name, Data: body.Date}
	res, err := client.HolidayCreate(context.Background(), reqTest)
	if err != nil {
		log.Fatalf("failed to call AddHolidayRequest: %v", err)
	}

	log.Printf("AddHolidayRequest: %v %v", res.Name, res.Data)
}

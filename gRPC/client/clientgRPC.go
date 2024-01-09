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

	client := proto.NewServiceProductionCalendarClient(conn)

	//var body struct {
	//	ID   int32
	//	Name string
	//	Date string
	//}

	reqTest := &proto.AddHolidayRequest{Id: 1, Name: "test", Data: "10.10.1010"}
	//reqTest := &proto.GetAllHolidayRequest{}
	respTest, err := client.HolidayCreate(context.Background(), reqTest)
	//respTest, err := client.GetAllHoliday(context.Background(), reqTest)
	if err != nil {
		log.Fatalf("failed to add holiday: %v", err)
	}
	log.Println(respTest)

	res, err := client.HolidayCreate(context.Background(), reqTest)
	//res, err := client.GetAllHoliday(context.Background(), reqTest)
	if err != nil {
		log.Fatalf("failed to call AddHolidayRequest: %v", err)
	}

	log.Printf("AddHolidayRequest: %v %v", res.Name, res.Data)

	if conn != nil {
		log.Fatal(conn.Close())
	}
}

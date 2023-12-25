package main

import (
	pb "ServiceProductionCalendar/rpc/proto"
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

	client := pb.NewServiceProductionCalendarClient(conn)

	reqTest := &pb.AddHolidayRequest{Id: 1, Name: "test", Data: "2023:12:31"}
	res, err := client.HolidayCreate(context.Background(), reqTest)
	if err != nil {
		log.Fatalf("failed to call AddHolidayRequest: %v", err)
	}

	log.Printf("AddHolidayRequest: %v %v", res.Name, res.Data)
}

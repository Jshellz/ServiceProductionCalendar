package main

import (
	"ServiceProductionCalendar/gRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	network = "tcp"
	address = ":50051"
)

var spcHoliday = proto.SPC{}

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterServiceProductionCalendarServer(s, &spcHoliday)

	log.Printf("Server listening on %s", lis)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

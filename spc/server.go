package spc

import (
	pb "ServiceProductionCalendar/spc/proto"
	"context"
	"fmt"
)

type Server struct {
	pb.UnimplementedSPCServer
}

func (c CalendarService) AddHoliday(ctx context.Context, r *pb.AddHolidayRequest) (*pb.AddHolidayResponse, error) {
	fmt.Printf("Received request to add holiday: %s and %s\n", r.GetName(), r.GetData())
	return &pb.AddHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

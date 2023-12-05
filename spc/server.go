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
	ctx.Err().Error()
	return &pb.AddHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

func (c CalendarService) GetHoliday(ctx context.Context, r *pb.GetHolidayRequest) (*pb.GetHolidayResponse, error) {
	fmt.Printf("Received request to add holiday: %s, %s and %s\n", r.GetId(), r.GetName(), r.GetData())
	ctx.Err().Error()
	return &pb.GetHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

func (c CalendarService) EditHoliday(ctx context.Context, r *pb.EditHolidayRequest) (*pb.EditHolidayResponse, error) {
	fmt.Printf("Received request to add holiday: %s and %s\n", r.GetName(), r.GetData())
	ctx.Err().Error()
	return &pb.EditHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

func (c CalendarService) DeleteHoliday(ctx context.Context, r *pb.DeleteHolidayRequest) (*pb.DeleteHolidayResponse, error) {
	fmt.Printf("Received request to add holiday: %s, %s and %s\n", r.GetId(), r.GetName(), r.GetData())
	ctx.Err().Error()
	return &pb.DeleteHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

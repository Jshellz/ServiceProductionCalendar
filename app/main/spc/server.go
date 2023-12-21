package spc

import (
	proto2 "ServiceProductionCalendar/app/main/spc/proto"
	"context"
	"fmt"
)

type Server struct {
	proto2.UnimplementedSPCServer
}

func (c CalendarService) AddHoliday(ctx context.Context, r *proto2.AddHolidayRequest) (*proto2.AddHolidayResponse, error) {
	fmt.Printf("Received request to add holiday: %s and %s\n", r.GetName(), r.GetData())
	ctx.Err().Error()
	return &proto2.AddHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

func (c CalendarService) GetHoliday(ctx context.Context, r *proto2.GetHolidayRequest) (*proto2.GetHolidayResponse, error) {
	fmt.Printf("Received request to add holiday: %s, %s and %s\n", r.GetId(), r.GetName(), r.GetData())
	ctx.Err().Error()
	return &proto2.GetHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

func (c CalendarService) EditHoliday(ctx context.Context, r *proto2.EditHolidayRequest) (*proto2.EditHolidayResponse, error) {
	fmt.Printf("Received request to add holiday: %s and %s\n", r.GetName(), r.GetData())
	ctx.Err().Error()
	return &proto2.EditHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

func (c CalendarService) DeleteHoliday(ctx context.Context, r *proto2.DeleteHolidayRequest) (*proto2.DeleteHolidayResponse, error) {
	fmt.Printf("Received request to add holiday: %s, %s and %s\n", r.GetId(), r.GetName(), r.GetData())
	ctx.Err().Error()
	return &proto2.DeleteHolidayResponse{HolidayName: "День к новому году", HolidayData: "31.12.2023"}, nil
}

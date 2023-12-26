package main

import (
	"ServiceProductionCalendar/rpc/proto"
	"context"
)

const (
	network = "tcp"
	address = ":50051"
)

type sPC struct{}

func (spc *sPC) GetAllHoliday(ctx context.Context, req *proto.GetAllHolidayRequest) (*proto.GetAllHolidayResponse, error) {
	return &proto.GetAllHolidayResponse{}, nil
}

func (spc *sPC) GetHoliday(ctx context.Context, req *proto.GetHolidayRequest) (*proto.GetHolidayResponse, error) {
	return &proto.GetHolidayResponse{}, nil
}

func (spc *sPC) HolidayCreate(ctx context.Context, req *proto.AddHolidayRequest) (*proto.AddHolidayResponse, error) {
	return &proto.AddHolidayResponse{}, nil
}

func (spc *sPC) UpdateHoliday(ctx context.Context, req *proto.EditHolidayRequest) (*proto.EditHolidayResponse, error) {
	return &proto.EditHolidayResponse{}, nil
}

func (spc *sPC) DeleteHoliday(ctx context.Context, req *proto.DeleteHolidayRequest) (*proto.DeleteHolidayResponse, error) {
	return &proto.DeleteHolidayResponse{}, nil
}

func (spc *sPC) mustEmbedUnimplementedServiceProductionCalendarServer() {

}

func main() {
	//lis, err := net.Listen(network, address)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//s := grpc.NewServer()
	//proto.RegisterServiceProductionCalendarServer(s, &sPC{})
	//
	//log.Printf("Server listening on %s", lis)
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}

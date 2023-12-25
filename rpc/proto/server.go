package proto

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	network = "tcp"
	address = ":50051"
)

type sPC struct{}

func (spc *sPC) GetAllHoliday(ctx context.Context, req *GetAllHolidayRequest) (*GetAllHolidayResponse, error) {
	return &GetAllHolidayResponse{}, nil
}

func (spc *sPC) GetHoliday(ctx context.Context, req *GetHolidayRequest) (*GetHolidayResponse, error) {
	return &GetHolidayResponse{}, nil
}

func (spc *sPC) HolidayCreate(ctx context.Context, req *AddHolidayRequest) (*AddHolidayResponse, error) {
	return &AddHolidayResponse{}, nil
}

func (spc *sPC) UpdateHoliday(ctx context.Context, req *EditHolidayRequest) (*EditHolidayResponse, error) {
	return &EditHolidayResponse{}, nil
}

func (spc *sPC) DeleteHoliday(ctx context.Context, req *DeleteHolidayRequest) (*DeleteHolidayResponse, error) {
	return &DeleteHolidayResponse{}, nil
}

func (spc *sPC) mustEmbedUnimplementedServiceProductionCalendarServer() {

}

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterServiceProductionCalendarServer(s, &sPC{})

	log.Printf("Server listening on %s", lis)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package proto

import (
	"context"
	"fmt"
)

// SPC Methods

type SPC struct {
	ServiceProductionCalendarServer
}

func (s *SPC) GetAllHoliday(ctx context.Context, req *GetAllHolidayRequest) (*GetAllHolidayResponse, error) {
	var body struct {
		Id   int32
		Name string
		Data string
	}

	return &GetAllHolidayResponse{
		Id:   body.Id,
		Name: body.Name,
		Data: body.Data,
	}, nil
}

func (s *SPC) GetHoliday(ctx context.Context, req *GetHolidayRequest) (*GetHolidayResponse, error) {
	//var body struct {
	//	Id   int32
	//	Name string
	//	Data string
	//}
	//return &GetHolidayResponse{
	//	Id:   body.Id,
	//	Name: body.Name,
	//	Data: body.Data,
	//}, nil

	// Test
	return &GetHolidayResponse{
		Name: fmt.Sprintf("Test value %s", req.Name),
		Data: fmt.Sprintf("Test value %s", req.Data),
	}, nil
}

func (s *SPC) HolidayCreate(ctx context.Context, req *AddHolidayRequest) (*AddHolidayResponse, error) {
	//var body struct {
	//	Name string
	//	Data string
	//}
	//return &AddHolidayResponse{
	//	Name: body.Name,
	//	Data: body.Data,
	//}, nil

	// Test
	return &AddHolidayResponse{
		Name: fmt.Sprintf("Test value %s", req.Name),
		Data: fmt.Sprintf("Test value %s", req.Data),
	}, nil
}

func (s *SPC) UpdateHoliday(ctx context.Context, req *EditHolidayRequest) (*EditHolidayResponse, error) {
	var body struct {
		Id   int32
		Name string
		Data string
	}
	return &EditHolidayResponse{
		Id:   body.Id,
		Name: body.Name,
		Data: body.Data,
	}, nil
}

func (s *SPC) DeleteHoliday(ctx context.Context, req *DeleteHolidayRequest) (*DeleteHolidayResponse, error) {
	var body struct {
		Id   int32
		Name string
	}
	return &DeleteHolidayResponse{
		Id:   body.Id,
		Name: body.Name,
	}, nil
}

func (s *SPC) mustEmbedUnimplementedServiceProductionCalendarServer() {}

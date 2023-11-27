// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/spc.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetHolidayRequest) Reset() {
	*x = GetHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHolidayRequest) ProtoMessage() {}

func (x *GetHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHolidayRequest.ProtoReflect.Descriptor instead.
func (*GetHolidayRequest) Descriptor() ([]byte, []int) {
	return file_proto_spc_proto_rawDescGZIP(), []int{0}
}

func (x *GetHolidayRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetHolidayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetHolidayRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetHolidayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HolidayId   string `protobuf:"bytes,1,opt,name=holiday_id,json=holidayId,proto3" json:"holiday_id,omitempty"`
	HolidayName string `protobuf:"bytes,2,opt,name=holiday_name,json=holidayName,proto3" json:"holiday_name,omitempty"`
	HolidayData string `protobuf:"bytes,3,opt,name=holiday_data,json=holidayData,proto3" json:"holiday_data,omitempty"`
}

func (x *GetHolidayResponse) Reset() {
	*x = GetHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHolidayResponse) ProtoMessage() {}

func (x *GetHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHolidayResponse.ProtoReflect.Descriptor instead.
func (*GetHolidayResponse) Descriptor() ([]byte, []int) {
	return file_proto_spc_proto_rawDescGZIP(), []int{1}
}

func (x *GetHolidayResponse) GetHolidayId() string {
	if x != nil {
		return x.HolidayId
	}
	return ""
}

func (x *GetHolidayResponse) GetHolidayName() string {
	if x != nil {
		return x.HolidayName
	}
	return ""
}

func (x *GetHolidayResponse) GetHolidayData() string {
	if x != nil {
		return x.HolidayData
	}
	return ""
}

type AddHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddHolidayRequest) Reset() {
	*x = AddHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHolidayRequest) ProtoMessage() {}

func (x *AddHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHolidayRequest.ProtoReflect.Descriptor instead.
func (*AddHolidayRequest) Descriptor() ([]byte, []int) {
	return file_proto_spc_proto_rawDescGZIP(), []int{2}
}

func (x *AddHolidayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddHolidayRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type AddHolidayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HolidayName string `protobuf:"bytes,1,opt,name=holiday_name,json=holidayName,proto3" json:"holiday_name,omitempty"`
	HolidayData string `protobuf:"bytes,2,opt,name=holiday_data,json=holidayData,proto3" json:"holiday_data,omitempty"`
}

func (x *AddHolidayResponse) Reset() {
	*x = AddHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHolidayResponse) ProtoMessage() {}

func (x *AddHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHolidayResponse.ProtoReflect.Descriptor instead.
func (*AddHolidayResponse) Descriptor() ([]byte, []int) {
	return file_proto_spc_proto_rawDescGZIP(), []int{3}
}

func (x *AddHolidayResponse) GetHolidayName() string {
	if x != nil {
		return x.HolidayName
	}
	return ""
}

func (x *AddHolidayResponse) GetHolidayData() string {
	if x != nil {
		return x.HolidayData
	}
	return ""
}

type EditHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EditHolidayRequest) Reset() {
	*x = EditHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditHolidayRequest) ProtoMessage() {}

func (x *EditHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditHolidayRequest.ProtoReflect.Descriptor instead.
func (*EditHolidayRequest) Descriptor() ([]byte, []int) {
	return file_proto_spc_proto_rawDescGZIP(), []int{4}
}

func (x *EditHolidayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EditHolidayRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EditHolidayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HolidayName string `protobuf:"bytes,1,opt,name=holiday_name,json=holidayName,proto3" json:"holiday_name,omitempty"`
	HolidayData string `protobuf:"bytes,2,opt,name=holiday_data,json=holidayData,proto3" json:"holiday_data,omitempty"`
}

func (x *EditHolidayResponse) Reset() {
	*x = EditHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditHolidayResponse) ProtoMessage() {}

func (x *EditHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditHolidayResponse.ProtoReflect.Descriptor instead.
func (*EditHolidayResponse) Descriptor() ([]byte, []int) {
	return file_proto_spc_proto_rawDescGZIP(), []int{5}
}

func (x *EditHolidayResponse) GetHolidayName() string {
	if x != nil {
		return x.HolidayName
	}
	return ""
}

func (x *EditHolidayResponse) GetHolidayData() string {
	if x != nil {
		return x.HolidayData
	}
	return ""
}

type DeleteHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DeleteHolidayRequest) Reset() {
	*x = DeleteHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHolidayRequest) ProtoMessage() {}

func (x *DeleteHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHolidayRequest.ProtoReflect.Descriptor instead.
func (*DeleteHolidayRequest) Descriptor() ([]byte, []int) {
	return file_proto_spc_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteHolidayRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteHolidayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteHolidayRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DeleteHolidayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HolidayId   string `protobuf:"bytes,1,opt,name=holiday_id,json=holidayId,proto3" json:"holiday_id,omitempty"`
	HolidayName string `protobuf:"bytes,2,opt,name=holiday_name,json=holidayName,proto3" json:"holiday_name,omitempty"`
	HolidayData string `protobuf:"bytes,3,opt,name=holiday_data,json=holidayData,proto3" json:"holiday_data,omitempty"`
}

func (x *DeleteHolidayResponse) Reset() {
	*x = DeleteHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHolidayResponse) ProtoMessage() {}

func (x *DeleteHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHolidayResponse.ProtoReflect.Descriptor instead.
func (*DeleteHolidayResponse) Descriptor() ([]byte, []int) {
	return file_proto_spc_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteHolidayResponse) GetHolidayId() string {
	if x != nil {
		return x.HolidayId
	}
	return ""
}

func (x *DeleteHolidayResponse) GetHolidayName() string {
	if x != nil {
		return x.HolidayName
	}
	return ""
}

func (x *DeleteHolidayResponse) GetHolidayData() string {
	if x != nil {
		return x.HolidayData
	}
	return ""
}

var File_proto_spc_proto protoreflect.FileDescriptor

var file_proto_spc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x79, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x3b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x12,
	0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x12, 0x45, 0x64, 0x69, 0x74,
	0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x7c, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x68,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x32, 0x95, 0x02, 0x0a, 0x03, 0x53, 0x50, 0x43, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x69,
	0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x45,
	0x64, 0x69, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_spc_proto_rawDescOnce sync.Once
	file_proto_spc_proto_rawDescData = file_proto_spc_proto_rawDesc
)

func file_proto_spc_proto_rawDescGZIP() []byte {
	file_proto_spc_proto_rawDescOnce.Do(func() {
		file_proto_spc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_spc_proto_rawDescData)
	})
	return file_proto_spc_proto_rawDescData
}

var file_proto_spc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_spc_proto_goTypes = []interface{}{
	(*GetHolidayRequest)(nil),     // 0: main.GetHolidayRequest
	(*GetHolidayResponse)(nil),    // 1: main.GetHolidayResponse
	(*AddHolidayRequest)(nil),     // 2: main.AddHolidayRequest
	(*AddHolidayResponse)(nil),    // 3: main.AddHolidayResponse
	(*EditHolidayRequest)(nil),    // 4: main.EditHolidayRequest
	(*EditHolidayResponse)(nil),   // 5: main.EditHolidayResponse
	(*DeleteHolidayRequest)(nil),  // 6: main.DeleteHolidayRequest
	(*DeleteHolidayResponse)(nil), // 7: main.DeleteHolidayResponse
}
var file_proto_spc_proto_depIdxs = []int32{
	0, // 0: main.SPC.GetHoliday:input_type -> main.GetHolidayRequest
	2, // 1: main.SPC.AddHoliday:input_type -> main.AddHolidayRequest
	4, // 2: main.SPC.EditHoliday:input_type -> main.EditHolidayRequest
	6, // 3: main.SPC.DeleteHoliday:input_type -> main.DeleteHolidayRequest
	1, // 4: main.SPC.GetHoliday:output_type -> main.GetHolidayResponse
	3, // 5: main.SPC.AddHoliday:output_type -> main.AddHolidayResponse
	5, // 6: main.SPC.EditHoliday:output_type -> main.EditHolidayResponse
	7, // 7: main.SPC.DeleteHoliday:output_type -> main.DeleteHolidayResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_spc_proto_init() }
func file_proto_spc_proto_init() {
	if File_proto_spc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_spc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHolidayRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_spc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHolidayResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_spc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHolidayRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_spc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHolidayResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_spc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditHolidayRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_spc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditHolidayResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_spc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHolidayRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_spc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHolidayResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_spc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_spc_proto_goTypes,
		DependencyIndexes: file_proto_spc_proto_depIdxs,
		MessageInfos:      file_proto_spc_proto_msgTypes,
	}.Build()
	File_proto_spc_proto = out.File
	file_proto_spc_proto_rawDesc = nil
	file_proto_spc_proto_goTypes = nil
	file_proto_spc_proto_depIdxs = nil
}

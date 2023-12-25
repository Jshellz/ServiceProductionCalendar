// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: rpc/spc.proto

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

type GetAllHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllHolidayRequest) Reset() {
	*x = GetAllHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllHolidayRequest) ProtoMessage() {}

func (x *GetAllHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllHolidayRequest.ProtoReflect.Descriptor instead.
func (*GetAllHolidayRequest) Descriptor() ([]byte, []int) {
	return file_rpc_spc_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllHolidayRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAllHolidayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllHolidayRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetAllHolidayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []int32  `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Name []string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Data []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllHolidayResponse) Reset() {
	*x = GetAllHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllHolidayResponse) ProtoMessage() {}

func (x *GetAllHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllHolidayResponse.ProtoReflect.Descriptor instead.
func (*GetAllHolidayResponse) Descriptor() ([]byte, []int) {
	return file_rpc_spc_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllHolidayResponse) GetId() []int32 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetAllHolidayResponse) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *GetAllHolidayResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetHolidayRequest) Reset() {
	*x = GetHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHolidayRequest) ProtoMessage() {}

func (x *GetHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[2]
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
	return file_rpc_spc_proto_rawDescGZIP(), []int{2}
}

func (x *GetHolidayRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
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

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetHolidayResponse) Reset() {
	*x = GetHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHolidayResponse) ProtoMessage() {}

func (x *GetHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[3]
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
	return file_rpc_spc_proto_rawDescGZIP(), []int{3}
}

func (x *GetHolidayResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetHolidayResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetHolidayResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type AddHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddHolidayRequest) Reset() {
	*x = AddHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHolidayRequest) ProtoMessage() {}

func (x *AddHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[4]
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
	return file_rpc_spc_proto_rawDescGZIP(), []int{4}
}

func (x *AddHolidayRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
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

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddHolidayResponse) Reset() {
	*x = AddHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHolidayResponse) ProtoMessage() {}

func (x *AddHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[5]
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
	return file_rpc_spc_proto_rawDescGZIP(), []int{5}
}

func (x *AddHolidayResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddHolidayResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EditHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EditHolidayRequest) Reset() {
	*x = EditHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditHolidayRequest) ProtoMessage() {}

func (x *EditHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[6]
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
	return file_rpc_spc_proto_rawDescGZIP(), []int{6}
}

func (x *EditHolidayRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
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

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EditHolidayResponse) Reset() {
	*x = EditHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditHolidayResponse) ProtoMessage() {}

func (x *EditHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[7]
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
	return file_rpc_spc_proto_rawDescGZIP(), []int{7}
}

func (x *EditHolidayResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditHolidayResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EditHolidayResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DeleteHolidayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteHolidayRequest) Reset() {
	*x = DeleteHolidayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHolidayRequest) ProtoMessage() {}

func (x *DeleteHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[8]
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
	return file_rpc_spc_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteHolidayRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteHolidayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteHolidayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteHolidayResponse) Reset() {
	*x = DeleteHolidayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_spc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHolidayResponse) ProtoMessage() {}

func (x *DeleteHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_spc_proto_msgTypes[9]
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
	return file_rpc_spc_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteHolidayResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteHolidayResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_rpc_spc_proto protoreflect.FileDescriptor

var file_rpc_spc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x4e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x4b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c,
	0x0a, 0x12, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x12,
	0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a, 0x13, 0x45, 0x64,
	0x69, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0xfa, 0x02, 0x0a, 0x19, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61,
	0x79, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69,
	0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64,
	0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x12, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x6c, 0x69,
	0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_spc_proto_rawDescOnce sync.Once
	file_rpc_spc_proto_rawDescData = file_rpc_spc_proto_rawDesc
)

func file_rpc_spc_proto_rawDescGZIP() []byte {
	file_rpc_spc_proto_rawDescOnce.Do(func() {
		file_rpc_spc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_spc_proto_rawDescData)
	})
	return file_rpc_spc_proto_rawDescData
}

var file_rpc_spc_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rpc_spc_proto_goTypes = []interface{}{
	(*GetAllHolidayRequest)(nil),  // 0: main.GetAllHolidayRequest
	(*GetAllHolidayResponse)(nil), // 1: main.GetAllHolidayResponse
	(*GetHolidayRequest)(nil),     // 2: main.GetHolidayRequest
	(*GetHolidayResponse)(nil),    // 3: main.GetHolidayResponse
	(*AddHolidayRequest)(nil),     // 4: main.AddHolidayRequest
	(*AddHolidayResponse)(nil),    // 5: main.AddHolidayResponse
	(*EditHolidayRequest)(nil),    // 6: main.EditHolidayRequest
	(*EditHolidayResponse)(nil),   // 7: main.EditHolidayResponse
	(*DeleteHolidayRequest)(nil),  // 8: main.DeleteHolidayRequest
	(*DeleteHolidayResponse)(nil), // 9: main.DeleteHolidayResponse
}
var file_rpc_spc_proto_depIdxs = []int32{
	0, // 0: main.ServiceProductionCalendar.GetAllHoliday:input_type -> main.GetAllHolidayRequest
	2, // 1: main.ServiceProductionCalendar.GetHoliday:input_type -> main.GetHolidayRequest
	4, // 2: main.ServiceProductionCalendar.HolidayCreate:input_type -> main.AddHolidayRequest
	6, // 3: main.ServiceProductionCalendar.UpdateHoliday:input_type -> main.EditHolidayRequest
	8, // 4: main.ServiceProductionCalendar.DeleteHoliday:input_type -> main.DeleteHolidayRequest
	1, // 5: main.ServiceProductionCalendar.GetAllHoliday:output_type -> main.GetAllHolidayResponse
	3, // 6: main.ServiceProductionCalendar.GetHoliday:output_type -> main.GetHolidayResponse
	5, // 7: main.ServiceProductionCalendar.HolidayCreate:output_type -> main.AddHolidayResponse
	7, // 8: main.ServiceProductionCalendar.UpdateHoliday:output_type -> main.EditHolidayResponse
	9, // 9: main.ServiceProductionCalendar.DeleteHoliday:output_type -> main.DeleteHolidayResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_spc_proto_init() }
func file_rpc_spc_proto_init() {
	if File_rpc_spc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_spc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllHolidayRequest); i {
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
		file_rpc_spc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllHolidayResponse); i {
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
		file_rpc_spc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_spc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_spc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_spc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_spc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_spc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_spc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_spc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_rpc_spc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_spc_proto_goTypes,
		DependencyIndexes: file_rpc_spc_proto_depIdxs,
		MessageInfos:      file_rpc_spc_proto_msgTypes,
	}.Build()
	File_rpc_spc_proto = out.File
	file_rpc_spc_proto_rawDesc = nil
	file_rpc_spc_proto_goTypes = nil
	file_rpc_spc_proto_depIdxs = nil
}

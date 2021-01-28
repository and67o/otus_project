// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: BannerRotationService.proto

package server

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AddBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId   uint64 `protobuf:"varint,1,opt,name=SlotId,proto3" json:"SlotId,omitempty"`
	BannerId uint64 `protobuf:"varint,2,opt,name=BannerId,proto3" json:"BannerId,omitempty"`
}

func (x *AddBannerRequest) Reset() {
	*x = AddBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BannerRotationService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerRequest) ProtoMessage() {}

func (x *AddBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BannerRotationService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerRequest.ProtoReflect.Descriptor instead.
func (*AddBannerRequest) Descriptor() ([]byte, []int) {
	return file_BannerRotationService_proto_rawDescGZIP(), []int{0}
}

func (x *AddBannerRequest) GetSlotId() uint64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *AddBannerRequest) GetBannerId() uint64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

type AddBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddBannerResponse) Reset() {
	*x = AddBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BannerRotationService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerResponse) ProtoMessage() {}

func (x *AddBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BannerRotationService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerResponse.ProtoReflect.Descriptor instead.
func (*AddBannerResponse) Descriptor() ([]byte, []int) {
	return file_BannerRotationService_proto_rawDescGZIP(), []int{1}
}

type DeleteBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId uint64 `protobuf:"varint,1,opt,name=BannerId,proto3" json:"BannerId,omitempty"`
	SlotId   int64  `protobuf:"varint,2,opt,name=SlotId,proto3" json:"SlotId,omitempty"`
}

func (x *DeleteBannerRequest) Reset() {
	*x = DeleteBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BannerRotationService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBannerRequest) ProtoMessage() {}

func (x *DeleteBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BannerRotationService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBannerRequest.ProtoReflect.Descriptor instead.
func (*DeleteBannerRequest) Descriptor() ([]byte, []int) {
	return file_BannerRotationService_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBannerRequest) GetBannerId() uint64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *DeleteBannerRequest) GetSlotId() int64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

type DeleteBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBannerResponse) Reset() {
	*x = DeleteBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BannerRotationService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBannerResponse) ProtoMessage() {}

func (x *DeleteBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BannerRotationService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBannerResponse.ProtoReflect.Descriptor instead.
func (*DeleteBannerResponse) Descriptor() ([]byte, []int) {
	return file_BannerRotationService_proto_rawDescGZIP(), []int{3}
}

type ClickBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId   int64  `protobuf:"varint,1,opt,name=SlotId,proto3" json:"SlotId,omitempty"`
	BannerId uint64 `protobuf:"varint,2,opt,name=BannerId,proto3" json:"BannerId,omitempty"`
	GroupId  int64  `protobuf:"varint,3,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
}

func (x *ClickBannerRequest) Reset() {
	*x = ClickBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BannerRotationService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickBannerRequest) ProtoMessage() {}

func (x *ClickBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BannerRotationService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickBannerRequest.ProtoReflect.Descriptor instead.
func (*ClickBannerRequest) Descriptor() ([]byte, []int) {
	return file_BannerRotationService_proto_rawDescGZIP(), []int{4}
}

func (x *ClickBannerRequest) GetSlotId() int64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *ClickBannerRequest) GetBannerId() uint64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *ClickBannerRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type ClickBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClickBannerResponse) Reset() {
	*x = ClickBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BannerRotationService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickBannerResponse) ProtoMessage() {}

func (x *ClickBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BannerRotationService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickBannerResponse.ProtoReflect.Descriptor instead.
func (*ClickBannerResponse) Descriptor() ([]byte, []int) {
	return file_BannerRotationService_proto_rawDescGZIP(), []int{5}
}

type ShowBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId  int64 `protobuf:"varint,1,opt,name=SlotId,proto3" json:"SlotId,omitempty"`
	GroupId int64 `protobuf:"varint,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
}

func (x *ShowBannerRequest) Reset() {
	*x = ShowBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BannerRotationService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowBannerRequest) ProtoMessage() {}

func (x *ShowBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BannerRotationService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowBannerRequest.ProtoReflect.Descriptor instead.
func (*ShowBannerRequest) Descriptor() ([]byte, []int) {
	return file_BannerRotationService_proto_rawDescGZIP(), []int{6}
}

func (x *ShowBannerRequest) GetSlotId() int64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *ShowBannerRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type ShowBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId int64 `protobuf:"varint,1,opt,name=BannerId,proto3" json:"BannerId,omitempty"`
}

func (x *ShowBannerResponse) Reset() {
	*x = ShowBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BannerRotationService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowBannerResponse) ProtoMessage() {}

func (x *ShowBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BannerRotationService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowBannerResponse.ProtoReflect.Descriptor instead.
func (*ShowBannerResponse) Descriptor() ([]byte, []int) {
	return file_BannerRotationService_proto_rawDescGZIP(), []int{7}
}

func (x *ShowBannerResponse) GetBannerId() int64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

var File_BannerRotationService_proto protoreflect.FileDescriptor

var file_BannerRotationService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a,
	0x10, 0x61, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53,
	0x6c, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62, 0x0a,
	0x12, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x15, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x77,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53,
	0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0x30, 0x0a, 0x12, 0x73, 0x68, 0x6f, 0x77, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x32, 0xf2, 0x02, 0x0a, 0x0e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x20, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x2e,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x68, 0x6f, 0x77, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BannerRotationService_proto_rawDescOnce sync.Once
	file_BannerRotationService_proto_rawDescData = file_BannerRotationService_proto_rawDesc
)

func file_BannerRotationService_proto_rawDescGZIP() []byte {
	file_BannerRotationService_proto_rawDescOnce.Do(func() {
		file_BannerRotationService_proto_rawDescData = protoimpl.X.CompressGZIP(file_BannerRotationService_proto_rawDescData)
	})
	return file_BannerRotationService_proto_rawDescData
}

var file_BannerRotationService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_BannerRotationService_proto_goTypes = []interface{}{
	(*AddBannerRequest)(nil),     // 0: bannerRotation.addBannerRequest
	(*AddBannerResponse)(nil),    // 1: bannerRotation.addBannerResponse
	(*DeleteBannerRequest)(nil),  // 2: bannerRotation.deleteBannerRequest
	(*DeleteBannerResponse)(nil), // 3: bannerRotation.deleteBannerResponse
	(*ClickBannerRequest)(nil),   // 4: bannerRotation.clickBannerRequest
	(*ClickBannerResponse)(nil),  // 5: bannerRotation.clickBannerResponse
	(*ShowBannerRequest)(nil),    // 6: bannerRotation.showBannerRequest
	(*ShowBannerResponse)(nil),   // 7: bannerRotation.showBannerResponse
}
var file_BannerRotationService_proto_depIdxs = []int32{
	0, // 0: bannerRotation.bannerRotation.addBanner:input_type -> bannerRotation.addBannerRequest
	2, // 1: bannerRotation.bannerRotation.deleteBanner:input_type -> bannerRotation.deleteBannerRequest
	4, // 2: bannerRotation.bannerRotation.clickBanner:input_type -> bannerRotation.clickBannerRequest
	6, // 3: bannerRotation.bannerRotation.showBanner:input_type -> bannerRotation.showBannerRequest
	1, // 4: bannerRotation.bannerRotation.addBanner:output_type -> bannerRotation.addBannerResponse
	3, // 5: bannerRotation.bannerRotation.deleteBanner:output_type -> bannerRotation.deleteBannerResponse
	5, // 6: bannerRotation.bannerRotation.clickBanner:output_type -> bannerRotation.clickBannerResponse
	7, // 7: bannerRotation.bannerRotation.showBanner:output_type -> bannerRotation.showBannerResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BannerRotationService_proto_init() }
func file_BannerRotationService_proto_init() {
	if File_BannerRotationService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BannerRotationService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerRequest); i {
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
		file_BannerRotationService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerResponse); i {
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
		file_BannerRotationService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBannerRequest); i {
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
		file_BannerRotationService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBannerResponse); i {
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
		file_BannerRotationService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickBannerRequest); i {
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
		file_BannerRotationService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickBannerResponse); i {
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
		file_BannerRotationService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowBannerRequest); i {
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
		file_BannerRotationService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowBannerResponse); i {
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
			RawDescriptor: file_BannerRotationService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_BannerRotationService_proto_goTypes,
		DependencyIndexes: file_BannerRotationService_proto_depIdxs,
		MessageInfos:      file_BannerRotationService_proto_msgTypes,
	}.Build()
	File_BannerRotationService_proto = out.File
	file_BannerRotationService_proto_rawDesc = nil
	file_BannerRotationService_proto_goTypes = nil
	file_BannerRotationService_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: app/app.proto

package app

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

type AppOpRequest_Operation int32

const (
	AppOpRequest_Install   AppOpRequest_Operation = 0
	AppOpRequest_Uninstall AppOpRequest_Operation = 1
	AppOpRequest_Upgrade   AppOpRequest_Operation = 2
)

// Enum value maps for AppOpRequest_Operation.
var (
	AppOpRequest_Operation_name = map[int32]string{
		0: "Install",
		1: "Uninstall",
		2: "Upgrade",
	}
	AppOpRequest_Operation_value = map[string]int32{
		"Install":   0,
		"Uninstall": 1,
		"Upgrade":   2,
	}
)

func (x AppOpRequest_Operation) Enum() *AppOpRequest_Operation {
	p := new(AppOpRequest_Operation)
	*p = x
	return p
}

func (x AppOpRequest_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppOpRequest_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_app_app_proto_enumTypes[0].Descriptor()
}

func (AppOpRequest_Operation) Type() protoreflect.EnumType {
	return &file_app_app_proto_enumTypes[0]
}

func (x AppOpRequest_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppOpRequest_Operation.Descriptor instead.
func (AppOpRequest_Operation) EnumDescriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{3, 0}
}

type AppStatusRequest_AppStatus int32

const (
	AppStatusRequest_Running   AppStatusRequest_AppStatus = 0
	AppStatusRequest_Stopped   AppStatusRequest_AppStatus = 1
	AppStatusRequest_Upgrading AppStatusRequest_AppStatus = 2
)

// Enum value maps for AppStatusRequest_AppStatus.
var (
	AppStatusRequest_AppStatus_name = map[int32]string{
		0: "Running",
		1: "Stopped",
		2: "Upgrading",
	}
	AppStatusRequest_AppStatus_value = map[string]int32{
		"Running":   0,
		"Stopped":   1,
		"Upgrading": 2,
	}
)

func (x AppStatusRequest_AppStatus) Enum() *AppStatusRequest_AppStatus {
	p := new(AppStatusRequest_AppStatus)
	*p = x
	return p
}

func (x AppStatusRequest_AppStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppStatusRequest_AppStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_app_app_proto_enumTypes[1].Descriptor()
}

func (AppStatusRequest_AppStatus) Type() protoreflect.EnumType {
	return &file_app_app_proto_enumTypes[1]
}

func (x AppStatusRequest_AppStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppStatusRequest_AppStatus.Descriptor instead.
func (AppStatusRequest_AppStatus) EnumDescriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{5, 0}
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Intro   string  `protobuf:"bytes,3,opt,name=intro,proto3" json:"intro,omitempty"`
	Image   string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Port    []int32 `protobuf:"varint,5,rep,packed,name=port,proto3" json:"port,omitempty"`
	Config  string  `protobuf:"bytes,6,opt,name=config,proto3" json:"config,omitempty"`
	Depends string  `protobuf:"bytes,7,opt,name=depends,proto3" json:"depends,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *Application) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Application) GetPort() []int32 {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *Application) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Application) GetDepends() string {
	if x != nil {
		return x.Depends
	}
	return ""
}

type AppListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AppListRequest) Reset() {
	*x = AppListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppListRequest) ProtoMessage() {}

func (x *AppListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppListRequest.ProtoReflect.Descriptor instead.
func (*AppListRequest) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{1}
}

func (x *AppListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AppList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Application `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *AppList) Reset() {
	*x = AppList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppList) ProtoMessage() {}

func (x *AppList) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppList.ProtoReflect.Descriptor instead.
func (*AppList) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{2}
}

func (x *AppList) GetItems() []*Application {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AppList) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AppOpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Op AppOpRequest_Operation `protobuf:"varint,4,opt,name=op,proto3,enum=app.AppOpRequest_Operation" json:"op,omitempty"`
}

func (x *AppOpRequest) Reset() {
	*x = AppOpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppOpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppOpRequest) ProtoMessage() {}

func (x *AppOpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppOpRequest.ProtoReflect.Descriptor instead.
func (*AppOpRequest) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{3}
}

func (x *AppOpRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppOpRequest) GetOp() AppOpRequest_Operation {
	if x != nil {
		return x.Op
	}
	return AppOpRequest_Install
}

type AppOpReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppOpReply) Reset() {
	*x = AppOpReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppOpReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppOpReply) ProtoMessage() {}

func (x *AppOpReply) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppOpReply.ProtoReflect.Descriptor instead.
func (*AppOpReply) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{4}
}

type AppStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status AppStatusRequest_AppStatus `protobuf:"varint,2,opt,name=status,proto3,enum=app.AppStatusRequest_AppStatus" json:"status,omitempty"`
}

func (x *AppStatusRequest) Reset() {
	*x = AppStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppStatusRequest) ProtoMessage() {}

func (x *AppStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppStatusRequest.ProtoReflect.Descriptor instead.
func (*AppStatusRequest) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{5}
}

func (x *AppStatusRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppStatusRequest) GetStatus() AppStatusRequest_AppStatus {
	if x != nil {
		return x.Status
	}
	return AppStatusRequest_Running
}

type AppStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppStatusReply) Reset() {
	*x = AppStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppStatusReply) ProtoMessage() {}

func (x *AppStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppStatusReply.ProtoReflect.Descriptor instead.
func (*AppStatusReply) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{6}
}

var File_app_app_proto protoreflect.FileDescriptor

var file_app_app_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x70, 0x22, 0xa5, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x74, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x24, 0x0a, 0x0e,
	0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x47, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x81, 0x01, 0x0a, 0x0c,
	0x41, 0x70, 0x70, 0x4f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x02,
	0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41,
	0x70, 0x70, 0x4f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x22, 0x34, 0x0a, 0x09, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x10, 0x02, 0x22,
	0x0c, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x4f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x91, 0x01,
	0x0a, 0x10, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x34, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x10,
	0x02, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x83, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41,
	0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x4f, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70,
	0x4f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x55, 0x6e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70,
	0x4f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x41, 0x70, 0x70, 0x4f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x15, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6b, 0x6f, 0x73, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x6b, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_app_proto_rawDescOnce sync.Once
	file_app_app_proto_rawDescData = file_app_app_proto_rawDesc
)

func file_app_app_proto_rawDescGZIP() []byte {
	file_app_app_proto_rawDescOnce.Do(func() {
		file_app_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_app_proto_rawDescData)
	})
	return file_app_app_proto_rawDescData
}

var file_app_app_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_app_app_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_app_app_proto_goTypes = []interface{}{
	(AppOpRequest_Operation)(0),     // 0: app.AppOpRequest.Operation
	(AppStatusRequest_AppStatus)(0), // 1: app.AppStatusRequest.AppStatus
	(*Application)(nil),             // 2: app.Application
	(*AppListRequest)(nil),          // 3: app.AppListRequest
	(*AppList)(nil),                 // 4: app.AppList
	(*AppOpRequest)(nil),            // 5: app.AppOpRequest
	(*AppOpReply)(nil),              // 6: app.AppOpReply
	(*AppStatusRequest)(nil),        // 7: app.AppStatusRequest
	(*AppStatusReply)(nil),          // 8: app.AppStatusReply
}
var file_app_app_proto_depIdxs = []int32{
	2, // 0: app.AppList.items:type_name -> app.Application
	0, // 1: app.AppOpRequest.op:type_name -> app.AppOpRequest.Operation
	1, // 2: app.AppStatusRequest.status:type_name -> app.AppStatusRequest.AppStatus
	3, // 3: app.App.List:input_type -> app.AppListRequest
	5, // 4: app.App.Install:input_type -> app.AppOpRequest
	5, // 5: app.App.Uninstall:input_type -> app.AppOpRequest
	7, // 6: app.App.Start:input_type -> app.AppStatusRequest
	7, // 7: app.App.Stop:input_type -> app.AppStatusRequest
	4, // 8: app.App.List:output_type -> app.AppList
	6, // 9: app.App.Install:output_type -> app.AppOpReply
	6, // 10: app.App.Uninstall:output_type -> app.AppOpReply
	8, // 11: app.App.Start:output_type -> app.AppStatusReply
	8, // 12: app.App.Stop:output_type -> app.AppStatusReply
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_app_proto_init() }
func file_app_app_proto_init() {
	if File_app_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_app_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppListRequest); i {
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
		file_app_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppList); i {
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
		file_app_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppOpRequest); i {
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
		file_app_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppOpReply); i {
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
		file_app_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppStatusRequest); i {
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
		file_app_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppStatusReply); i {
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
			RawDescriptor: file_app_app_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_app_proto_goTypes,
		DependencyIndexes: file_app_app_proto_depIdxs,
		EnumInfos:         file_app_app_proto_enumTypes,
		MessageInfos:      file_app_app_proto_msgTypes,
	}.Build()
	File_app_app_proto = out.File
	file_app_app_proto_rawDesc = nil
	file_app_app_proto_goTypes = nil
	file_app_app_proto_depIdxs = nil
}

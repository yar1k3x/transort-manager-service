// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/transport.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTransportRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Number          string                 `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	TypeId          int32                  `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	IsActive        int32                  `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CurrentDriverId int32                  `protobuf:"varint,4,opt,name=current_driver_id,json=currentDriverId,proto3" json:"current_driver_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateTransportRequest) Reset() {
	*x = CreateTransportRequest{}
	mi := &file_proto_transport_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTransportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransportRequest) ProtoMessage() {}

func (x *CreateTransportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransportRequest.ProtoReflect.Descriptor instead.
func (*CreateTransportRequest) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTransportRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CreateTransportRequest) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *CreateTransportRequest) GetIsActive() int32 {
	if x != nil {
		return x.IsActive
	}
	return 0
}

func (x *CreateTransportRequest) GetCurrentDriverId() int32 {
	if x != nil {
		return x.CurrentDriverId
	}
	return 0
}

type CreateTransportResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TransportId   int64                  `protobuf:"varint,2,opt,name=transport_id,json=transportId,proto3" json:"transport_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTransportResponse) Reset() {
	*x = CreateTransportResponse{}
	mi := &file_proto_transport_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTransportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransportResponse) ProtoMessage() {}

func (x *CreateTransportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransportResponse.ProtoReflect.Descriptor instead.
func (*CreateTransportResponse) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTransportResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateTransportResponse) GetTransportId() int64 {
	if x != nil {
		return x.TransportId
	}
	return 0
}

type UpdateTransportRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TransportId     *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=transport_id,json=transportId,proto3" json:"transport_id,omitempty"`
	IsActive        *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CurrentDriverId *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=current_driver_id,json=currentDriverId,proto3" json:"current_driver_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UpdateTransportRequest) Reset() {
	*x = UpdateTransportRequest{}
	mi := &file_proto_transport_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTransportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTransportRequest) ProtoMessage() {}

func (x *UpdateTransportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTransportRequest.ProtoReflect.Descriptor instead.
func (*UpdateTransportRequest) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTransportRequest) GetTransportId() *wrapperspb.Int32Value {
	if x != nil {
		return x.TransportId
	}
	return nil
}

func (x *UpdateTransportRequest) GetIsActive() *wrapperspb.Int32Value {
	if x != nil {
		return x.IsActive
	}
	return nil
}

func (x *UpdateTransportRequest) GetCurrentDriverId() *wrapperspb.Int32Value {
	if x != nil {
		return x.CurrentDriverId
	}
	return nil
}

type UpdateTransportResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTransportResponse) Reset() {
	*x = UpdateTransportResponse{}
	mi := &file_proto_transport_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTransportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTransportResponse) ProtoMessage() {}

func (x *UpdateTransportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTransportResponse.ProtoReflect.Descriptor instead.
func (*UpdateTransportResponse) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTransportResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetTransportInfoRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TransportId     *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=transport_id,json=transportId,proto3" json:"transport_id,omitempty"`
	IsActive        *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CurrentDriverId *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=current_driver_id,json=currentDriverId,proto3" json:"current_driver_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetTransportInfoRequest) Reset() {
	*x = GetTransportInfoRequest{}
	mi := &file_proto_transport_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransportInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransportInfoRequest) ProtoMessage() {}

func (x *GetTransportInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransportInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTransportInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{4}
}

func (x *GetTransportInfoRequest) GetTransportId() *wrapperspb.Int32Value {
	if x != nil {
		return x.TransportId
	}
	return nil
}

func (x *GetTransportInfoRequest) GetIsActive() *wrapperspb.Int32Value {
	if x != nil {
		return x.IsActive
	}
	return nil
}

func (x *GetTransportInfoRequest) GetCurrentDriverId() *wrapperspb.Int32Value {
	if x != nil {
		return x.CurrentDriverId
	}
	return nil
}

type GetTransportInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Transports    []*TransportInfo       `protobuf:"bytes,1,rep,name=transports,proto3" json:"transports,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransportInfoResponse) Reset() {
	*x = GetTransportInfoResponse{}
	mi := &file_proto_transport_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransportInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransportInfoResponse) ProtoMessage() {}

func (x *GetTransportInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransportInfoResponse.ProtoReflect.Descriptor instead.
func (*GetTransportInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{5}
}

func (x *GetTransportInfoResponse) GetTransports() []*TransportInfo {
	if x != nil {
		return x.Transports
	}
	return nil
}

type TransportInfo struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TransportId     int32                  `protobuf:"varint,1,opt,name=transport_id,json=transportId,proto3" json:"transport_id,omitempty"`
	Number          string                 `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	TransportType   string                 `protobuf:"bytes,3,opt,name=transport_type,json=transportType,proto3" json:"transport_type,omitempty"`
	IsActive        int32                  `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CurrentDriverId int32                  `protobuf:"varint,5,opt,name=current_driver_id,json=currentDriverId,proto3" json:"current_driver_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *TransportInfo) Reset() {
	*x = TransportInfo{}
	mi := &file_proto_transport_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransportInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportInfo) ProtoMessage() {}

func (x *TransportInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportInfo.ProtoReflect.Descriptor instead.
func (*TransportInfo) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{6}
}

func (x *TransportInfo) GetTransportId() int32 {
	if x != nil {
		return x.TransportId
	}
	return 0
}

func (x *TransportInfo) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *TransportInfo) GetTransportType() string {
	if x != nil {
		return x.TransportType
	}
	return ""
}

func (x *TransportInfo) GetIsActive() int32 {
	if x != nil {
		return x.IsActive
	}
	return 0
}

func (x *TransportInfo) GetCurrentDriverId() int32 {
	if x != nil {
		return x.CurrentDriverId
	}
	return 0
}

type CreateTransportLogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TransportId   int32                  `protobuf:"varint,2,opt,name=transport_id,json=transportId,proto3" json:"transport_id,omitempty"`
	ServiceTypeId int32                  `protobuf:"varint,3,opt,name=service_type_id,json=serviceTypeId,proto3" json:"service_type_id,omitempty"`
	ServiceDate   string                 `protobuf:"bytes,4,opt,name=service_date,json=serviceDate,proto3" json:"service_date,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Mileage       int32                  `protobuf:"varint,6,opt,name=mileage,proto3" json:"mileage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTransportLogRequest) Reset() {
	*x = CreateTransportLogRequest{}
	mi := &file_proto_transport_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTransportLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransportLogRequest) ProtoMessage() {}

func (x *CreateTransportLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransportLogRequest.ProtoReflect.Descriptor instead.
func (*CreateTransportLogRequest) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{7}
}

func (x *CreateTransportLogRequest) GetTransportId() int32 {
	if x != nil {
		return x.TransportId
	}
	return 0
}

func (x *CreateTransportLogRequest) GetServiceTypeId() int32 {
	if x != nil {
		return x.ServiceTypeId
	}
	return 0
}

func (x *CreateTransportLogRequest) GetServiceDate() string {
	if x != nil {
		return x.ServiceDate
	}
	return ""
}

func (x *CreateTransportLogRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTransportLogRequest) GetMileage() int32 {
	if x != nil {
		return x.Mileage
	}
	return 0
}

type CreateTransportLogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTransportLogResponse) Reset() {
	*x = CreateTransportLogResponse{}
	mi := &file_proto_transport_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTransportLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransportLogResponse) ProtoMessage() {}

func (x *CreateTransportLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransportLogResponse.ProtoReflect.Descriptor instead.
func (*CreateTransportLogResponse) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{8}
}

func (x *CreateTransportLogResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type TransportLogInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TransportId   int32                  `protobuf:"varint,2,opt,name=transport_id,json=transportId,proto3" json:"transport_id,omitempty"`
	ServiceType   string                 `protobuf:"bytes,3,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Mileage       int32                  `protobuf:"varint,5,opt,name=mileage,proto3" json:"mileage,omitempty"`
	ServiceDate   string                 `protobuf:"bytes,6,opt,name=service_date,json=serviceDate,proto3" json:"service_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransportLogInfo) Reset() {
	*x = TransportLogInfo{}
	mi := &file_proto_transport_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransportLogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportLogInfo) ProtoMessage() {}

func (x *TransportLogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportLogInfo.ProtoReflect.Descriptor instead.
func (*TransportLogInfo) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{9}
}

func (x *TransportLogInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TransportLogInfo) GetTransportId() int32 {
	if x != nil {
		return x.TransportId
	}
	return 0
}

func (x *TransportLogInfo) GetServiceType() string {
	if x != nil {
		return x.ServiceType
	}
	return ""
}

func (x *TransportLogInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TransportLogInfo) GetMileage() int32 {
	if x != nil {
		return x.Mileage
	}
	return 0
}

func (x *TransportLogInfo) GetServiceDate() string {
	if x != nil {
		return x.ServiceDate
	}
	return ""
}

type GetTransportLogsInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TransportId   *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=transport_id,json=transportId,proto3" json:"transport_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransportLogsInfoRequest) Reset() {
	*x = GetTransportLogsInfoRequest{}
	mi := &file_proto_transport_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransportLogsInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransportLogsInfoRequest) ProtoMessage() {}

func (x *GetTransportLogsInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransportLogsInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTransportLogsInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{10}
}

func (x *GetTransportLogsInfoRequest) GetTransportId() *wrapperspb.Int32Value {
	if x != nil {
		return x.TransportId
	}
	return nil
}

type GetTransportLogsInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TransportLogs []*TransportLogInfo    `protobuf:"bytes,1,rep,name=transport_logs,json=transportLogs,proto3" json:"transport_logs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransportLogsInfoResponse) Reset() {
	*x = GetTransportLogsInfoResponse{}
	mi := &file_proto_transport_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransportLogsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransportLogsInfoResponse) ProtoMessage() {}

func (x *GetTransportLogsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransportLogsInfoResponse.ProtoReflect.Descriptor instead.
func (*GetTransportLogsInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{11}
}

func (x *GetTransportLogsInfoResponse) GetTransportLogs() []*TransportLogInfo {
	if x != nil {
		return x.TransportLogs
	}
	return nil
}

var File_proto_transport_proto protoreflect.FileDescriptor

const file_proto_transport_proto_rawDesc = "" +
	"\n" +
	"\x15proto/transport.proto\x12\x05proto\x1a\x1egoogle/protobuf/wrappers.proto\"\x92\x01\n" +
	"\x16CreateTransportRequest\x12\x16\n" +
	"\x06number\x18\x01 \x01(\tR\x06number\x12\x17\n" +
	"\atype_id\x18\x02 \x01(\x05R\x06typeId\x12\x1b\n" +
	"\tis_active\x18\x03 \x01(\x05R\bisActive\x12*\n" +
	"\x11current_driver_id\x18\x04 \x01(\x05R\x0fcurrentDriverId\"V\n" +
	"\x17CreateTransportResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12!\n" +
	"\ftransport_id\x18\x02 \x01(\x03R\vtransportId\"\xdb\x01\n" +
	"\x16UpdateTransportRequest\x12>\n" +
	"\ftransport_id\x18\x01 \x01(\v2\x1b.google.protobuf.Int32ValueR\vtransportId\x128\n" +
	"\tis_active\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueR\bisActive\x12G\n" +
	"\x11current_driver_id\x18\x03 \x01(\v2\x1b.google.protobuf.Int32ValueR\x0fcurrentDriverId\"3\n" +
	"\x17UpdateTransportResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\xdc\x01\n" +
	"\x17GetTransportInfoRequest\x12>\n" +
	"\ftransport_id\x18\x01 \x01(\v2\x1b.google.protobuf.Int32ValueR\vtransportId\x128\n" +
	"\tis_active\x18\x02 \x01(\v2\x1b.google.protobuf.Int32ValueR\bisActive\x12G\n" +
	"\x11current_driver_id\x18\x03 \x01(\v2\x1b.google.protobuf.Int32ValueR\x0fcurrentDriverId\"P\n" +
	"\x18GetTransportInfoResponse\x124\n" +
	"\n" +
	"transports\x18\x01 \x03(\v2\x14.proto.TransportInfoR\n" +
	"transports\"\xba\x01\n" +
	"\rTransportInfo\x12!\n" +
	"\ftransport_id\x18\x01 \x01(\x05R\vtransportId\x12\x16\n" +
	"\x06number\x18\x02 \x01(\tR\x06number\x12%\n" +
	"\x0etransport_type\x18\x03 \x01(\tR\rtransportType\x12\x1b\n" +
	"\tis_active\x18\x04 \x01(\x05R\bisActive\x12*\n" +
	"\x11current_driver_id\x18\x05 \x01(\x05R\x0fcurrentDriverId\"\xc5\x01\n" +
	"\x19CreateTransportLogRequest\x12!\n" +
	"\ftransport_id\x18\x02 \x01(\x05R\vtransportId\x12&\n" +
	"\x0fservice_type_id\x18\x03 \x01(\x05R\rserviceTypeId\x12!\n" +
	"\fservice_date\x18\x04 \x01(\tR\vserviceDate\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12\x18\n" +
	"\amileage\x18\x06 \x01(\x05R\amileage\"6\n" +
	"\x1aCreateTransportLogResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\xc7\x01\n" +
	"\x10TransportLogInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12!\n" +
	"\ftransport_id\x18\x02 \x01(\x05R\vtransportId\x12!\n" +
	"\fservice_type\x18\x03 \x01(\tR\vserviceType\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x18\n" +
	"\amileage\x18\x05 \x01(\x05R\amileage\x12!\n" +
	"\fservice_date\x18\x06 \x01(\tR\vserviceDate\"]\n" +
	"\x1bGetTransportLogsInfoRequest\x12>\n" +
	"\ftransport_id\x18\x01 \x01(\v2\x1b.google.protobuf.Int32ValueR\vtransportId\"^\n" +
	"\x1cGetTransportLogsInfoResponse\x12>\n" +
	"\x0etransport_logs\x18\x01 \x03(\v2\x17.proto.TransportLogInfoR\rtransportLogs2\xc7\x03\n" +
	"\x10TransportService\x12P\n" +
	"\x0fCreateTransport\x12\x1d.proto.CreateTransportRequest\x1a\x1e.proto.CreateTransportResponse\x12P\n" +
	"\x0fUpdateTransport\x12\x1d.proto.UpdateTransportRequest\x1a\x1e.proto.UpdateTransportResponse\x12S\n" +
	"\x10GetTransportInfo\x12\x1e.proto.GetTransportInfoRequest\x1a\x1f.proto.GetTransportInfoResponse\x12Y\n" +
	"\x12CreateTransportLog\x12 .proto.CreateTransportLogRequest\x1a!.proto.CreateTransportLogResponse\x12_\n" +
	"\x14GetTransportLogsInfo\x12\".proto.GetTransportLogsInfoRequest\x1a#.proto.GetTransportLogsInfoResponseB(Z&TransportManagementService/proto;protob\x06proto3"

var (
	file_proto_transport_proto_rawDescOnce sync.Once
	file_proto_transport_proto_rawDescData []byte
)

func file_proto_transport_proto_rawDescGZIP() []byte {
	file_proto_transport_proto_rawDescOnce.Do(func() {
		file_proto_transport_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_transport_proto_rawDesc), len(file_proto_transport_proto_rawDesc)))
	})
	return file_proto_transport_proto_rawDescData
}

var file_proto_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_transport_proto_goTypes = []any{
	(*CreateTransportRequest)(nil),       // 0: proto.CreateTransportRequest
	(*CreateTransportResponse)(nil),      // 1: proto.CreateTransportResponse
	(*UpdateTransportRequest)(nil),       // 2: proto.UpdateTransportRequest
	(*UpdateTransportResponse)(nil),      // 3: proto.UpdateTransportResponse
	(*GetTransportInfoRequest)(nil),      // 4: proto.GetTransportInfoRequest
	(*GetTransportInfoResponse)(nil),     // 5: proto.GetTransportInfoResponse
	(*TransportInfo)(nil),                // 6: proto.TransportInfo
	(*CreateTransportLogRequest)(nil),    // 7: proto.CreateTransportLogRequest
	(*CreateTransportLogResponse)(nil),   // 8: proto.CreateTransportLogResponse
	(*TransportLogInfo)(nil),             // 9: proto.TransportLogInfo
	(*GetTransportLogsInfoRequest)(nil),  // 10: proto.GetTransportLogsInfoRequest
	(*GetTransportLogsInfoResponse)(nil), // 11: proto.GetTransportLogsInfoResponse
	(*wrapperspb.Int32Value)(nil),        // 12: google.protobuf.Int32Value
}
var file_proto_transport_proto_depIdxs = []int32{
	12, // 0: proto.UpdateTransportRequest.transport_id:type_name -> google.protobuf.Int32Value
	12, // 1: proto.UpdateTransportRequest.is_active:type_name -> google.protobuf.Int32Value
	12, // 2: proto.UpdateTransportRequest.current_driver_id:type_name -> google.protobuf.Int32Value
	12, // 3: proto.GetTransportInfoRequest.transport_id:type_name -> google.protobuf.Int32Value
	12, // 4: proto.GetTransportInfoRequest.is_active:type_name -> google.protobuf.Int32Value
	12, // 5: proto.GetTransportInfoRequest.current_driver_id:type_name -> google.protobuf.Int32Value
	6,  // 6: proto.GetTransportInfoResponse.transports:type_name -> proto.TransportInfo
	12, // 7: proto.GetTransportLogsInfoRequest.transport_id:type_name -> google.protobuf.Int32Value
	9,  // 8: proto.GetTransportLogsInfoResponse.transport_logs:type_name -> proto.TransportLogInfo
	0,  // 9: proto.TransportService.CreateTransport:input_type -> proto.CreateTransportRequest
	2,  // 10: proto.TransportService.UpdateTransport:input_type -> proto.UpdateTransportRequest
	4,  // 11: proto.TransportService.GetTransportInfo:input_type -> proto.GetTransportInfoRequest
	7,  // 12: proto.TransportService.CreateTransportLog:input_type -> proto.CreateTransportLogRequest
	10, // 13: proto.TransportService.GetTransportLogsInfo:input_type -> proto.GetTransportLogsInfoRequest
	1,  // 14: proto.TransportService.CreateTransport:output_type -> proto.CreateTransportResponse
	3,  // 15: proto.TransportService.UpdateTransport:output_type -> proto.UpdateTransportResponse
	5,  // 16: proto.TransportService.GetTransportInfo:output_type -> proto.GetTransportInfoResponse
	8,  // 17: proto.TransportService.CreateTransportLog:output_type -> proto.CreateTransportLogResponse
	11, // 18: proto.TransportService.GetTransportLogsInfo:output_type -> proto.GetTransportLogsInfoResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_transport_proto_init() }
func file_proto_transport_proto_init() {
	if File_proto_transport_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_transport_proto_rawDesc), len(file_proto_transport_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_transport_proto_goTypes,
		DependencyIndexes: file_proto_transport_proto_depIdxs,
		MessageInfos:      file_proto_transport_proto_msgTypes,
	}.Build()
	File_proto_transport_proto = out.File
	file_proto_transport_proto_goTypes = nil
	file_proto_transport_proto_depIdxs = nil
}

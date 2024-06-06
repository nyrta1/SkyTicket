// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: flight.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From                 string                 `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   string                 `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	DepartureDate        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=departure_date,json=departureDate,proto3" json:"departure_date,omitempty"`
	ArrivalDate          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=arrival_date,json=arrivalDate,proto3" json:"arrival_date,omitempty"`
	AvailableFirstSlot   int64                  `protobuf:"varint,7,opt,name=available_first_slot,json=availableFirstSlot,proto3" json:"available_first_slot,omitempty"`
	AvailableEconomySlot int64                  `protobuf:"varint,8,opt,name=available_economy_slot,json=availableEconomySlot,proto3" json:"available_economy_slot,omitempty"`
	Status               string                 `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{0}
}

func (x *Flight) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Flight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flight) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Flight) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Flight) GetDepartureDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DepartureDate
	}
	return nil
}

func (x *Flight) GetArrivalDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ArrivalDate
	}
	return nil
}

func (x *Flight) GetAvailableFirstSlot() int64 {
	if x != nil {
		return x.AvailableFirstSlot
	}
	return 0
}

func (x *Flight) GetAvailableEconomySlot() int64 {
	if x != nil {
		return x.AvailableEconomySlot
	}
	return 0
}

func (x *Flight) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Flight) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Flight) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFlightRequest) Reset() {
	*x = GetFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightRequest) ProtoMessage() {}

func (x *GetFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightRequest.ProtoReflect.Descriptor instead.
func (*GetFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{1}
}

func (x *GetFlightRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFlightRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightList           []*Flight `protobuf:"bytes,1,rep,name=flight_list,json=flightList,proto3" json:"flight_list,omitempty"`
	AvailableFirstSlot   int64     `protobuf:"varint,2,opt,name=available_first_slot,json=availableFirstSlot,proto3" json:"available_first_slot,omitempty"`
	AvailableEconomySlot int64     `protobuf:"varint,3,opt,name=available_economy_slot,json=availableEconomySlot,proto3" json:"available_economy_slot,omitempty"`
}

func (x *GetFlightResponse) Reset() {
	*x = GetFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightResponse) ProtoMessage() {}

func (x *GetFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightResponse.ProtoReflect.Descriptor instead.
func (*GetFlightResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{2}
}

func (x *GetFlightResponse) GetFlightList() []*Flight {
	if x != nil {
		return x.FlightList
	}
	return nil
}

func (x *GetFlightResponse) GetAvailableFirstSlot() int64 {
	if x != nil {
		return x.AvailableFirstSlot
	}
	return 0
}

func (x *GetFlightResponse) GetAvailableEconomySlot() int64 {
	if x != nil {
		return x.AvailableEconomySlot
	}
	return 0
}

type DeleteFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFlightRequest) Reset() {
	*x = DeleteFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlightRequest) ProtoMessage() {}

func (x *DeleteFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlightRequest.ProtoReflect.Descriptor instead.
func (*DeleteFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteFlightRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListFlightRequest) Reset() {
	*x = ListFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlightRequest) ProtoMessage() {}

func (x *ListFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlightRequest.ProtoReflect.Descriptor instead.
func (*ListFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{4}
}

func (x *ListFlightRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListFlightRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightList []*Flight `protobuf:"bytes,1,rep,name=flight_list,json=flightList,proto3" json:"flight_list,omitempty"`
	Total      int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page       int64     `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListFlightResponse) Reset() {
	*x = ListFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlightResponse) ProtoMessage() {}

func (x *ListFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlightResponse.ProtoReflect.Descriptor instead.
func (*ListFlightResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{5}
}

func (x *ListFlightResponse) GetFlightList() []*Flight {
	if x != nil {
		return x.FlightList
	}
	return nil
}

func (x *ListFlightResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListFlightResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type UpdateFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateFlightRequest) Reset() {
	*x = UpdateFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlightRequest) ProtoMessage() {}

func (x *UpdateFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlightRequest.ProtoReflect.Descriptor instead.
func (*UpdateFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFlightRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateFlightSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TicketType int64 `protobuf:"varint,2,opt,name=ticket_type,json=ticketType,proto3" json:"ticket_type,omitempty"`
}

func (x *UpdateFlightSlotRequest) Reset() {
	*x = UpdateFlightSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlightSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlightSlotRequest) ProtoMessage() {}

func (x *UpdateFlightSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlightSlotRequest.ProtoReflect.Descriptor instead.
func (*UpdateFlightSlotRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateFlightSlotRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFlightSlotRequest) GetTicketType() int64 {
	if x != nil {
		return x.TicketType
	}
	return 0
}

type SearchFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	From          string                 `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	DepartureDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=departure_date,json=departureDate,proto3" json:"departure_date,omitempty"`
	ArrivalDate   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=arrival_date,json=arrivalDate,proto3" json:"arrival_date,omitempty"`
}

func (x *SearchFlightRequest) Reset() {
	*x = SearchFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlightRequest) ProtoMessage() {}

func (x *SearchFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlightRequest.ProtoReflect.Descriptor instead.
func (*SearchFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{8}
}

func (x *SearchFlightRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchFlightRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchFlightRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SearchFlightRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SearchFlightRequest) GetDepartureDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DepartureDate
	}
	return nil
}

func (x *SearchFlightRequest) GetArrivalDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ArrivalDate
	}
	return nil
}

type SearchFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightList []*Flight `protobuf:"bytes,1,rep,name=flight_list,json=flightList,proto3" json:"flight_list,omitempty"`
	Total      int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page       int64     `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchFlightResponse) Reset() {
	*x = SearchFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlightResponse) ProtoMessage() {}

func (x *SearchFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlightResponse.ProtoReflect.Descriptor instead.
func (*SearchFlightResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{9}
}

func (x *SearchFlightResponse) GetFlightList() []*Flight {
	if x != nil {
		return x.FlightList
	}
	return nil
}

func (x *SearchFlightResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchFlightResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

var File_flight_proto protoreflect.FileDescriptor

var file_flight_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x03, 0x0a, 0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x41, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x61,
	0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x61,
	0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x34, 0x0a, 0x16,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x6c,
	0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x0b, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x0a, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5f,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x6c, 0x6f, 0x74,
	0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x68, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0b,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x0a, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x22, 0x25, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x41, 0x0a, 0x0e, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x22, 0x6a, 0x0a, 0x14, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x0a, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0xb2, 0x04, 0x0a, 0x0d, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x07, 0x2e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x1a, 0x07, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x07, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x07, 0x2e, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a,
	0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x56, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x58, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x18,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x73,
	0x6c, 0x6f, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x57, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x14, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x4f, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flight_proto_rawDescOnce sync.Once
	file_flight_proto_rawDescData = file_flight_proto_rawDesc
)

func file_flight_proto_rawDescGZIP() []byte {
	file_flight_proto_rawDescOnce.Do(func() {
		file_flight_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_proto_rawDescData)
	})
	return file_flight_proto_rawDescData
}

var file_flight_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_flight_proto_goTypes = []interface{}{
	(*Flight)(nil),                  // 0: Flight
	(*GetFlightRequest)(nil),        // 1: GetFlightRequest
	(*GetFlightResponse)(nil),       // 2: GetFlightResponse
	(*DeleteFlightRequest)(nil),     // 3: DeleteFlightRequest
	(*ListFlightRequest)(nil),       // 4: ListFlightRequest
	(*ListFlightResponse)(nil),      // 5: ListFlightResponse
	(*UpdateFlightRequest)(nil),     // 6: UpdateFlightRequest
	(*UpdateFlightSlotRequest)(nil), // 7: UpdateFlightSlotRequest
	(*SearchFlightRequest)(nil),     // 8: SearchFlightRequest
	(*SearchFlightResponse)(nil),    // 9: SearchFlightResponse
	(*timestamppb.Timestamp)(nil),   // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),           // 11: google.protobuf.Empty
}
var file_flight_proto_depIdxs = []int32{
	10, // 0: Flight.departure_date:type_name -> google.protobuf.Timestamp
	10, // 1: Flight.arrival_date:type_name -> google.protobuf.Timestamp
	10, // 2: Flight.created_at:type_name -> google.protobuf.Timestamp
	10, // 3: Flight.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 4: GetFlightResponse.flight_list:type_name -> Flight
	0,  // 5: ListFlightResponse.flight_list:type_name -> Flight
	10, // 6: SearchFlightRequest.departure_date:type_name -> google.protobuf.Timestamp
	10, // 7: SearchFlightRequest.arrival_date:type_name -> google.protobuf.Timestamp
	0,  // 8: SearchFlightResponse.flight_list:type_name -> Flight
	0,  // 9: FlightManager.CreateFlight:input_type -> Flight
	0,  // 10: FlightManager.UpdateFlight:input_type -> Flight
	3,  // 11: FlightManager.DeleteFlight:input_type -> DeleteFlightRequest
	7,  // 12: FlightManager.UpdateFlightSlot:input_type -> UpdateFlightSlotRequest
	1,  // 13: FlightManager.GetFlight:input_type -> GetFlightRequest
	8,  // 14: FlightManager.SearchFlight:input_type -> SearchFlightRequest
	4,  // 15: FlightManager.ListFlight:input_type -> ListFlightRequest
	0,  // 16: FlightManager.CreateFlight:output_type -> Flight
	0,  // 17: FlightManager.UpdateFlight:output_type -> Flight
	11, // 18: FlightManager.DeleteFlight:output_type -> google.protobuf.Empty
	0,  // 19: FlightManager.UpdateFlightSlot:output_type -> Flight
	2,  // 20: FlightManager.GetFlight:output_type -> GetFlightResponse
	9,  // 21: FlightManager.SearchFlight:output_type -> SearchFlightResponse
	5,  // 22: FlightManager.ListFlight:output_type -> ListFlightResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_flight_proto_init() }
func file_flight_proto_init() {
	if File_flight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
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
		file_flight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightRequest); i {
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
		file_flight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightResponse); i {
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
		file_flight_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlightRequest); i {
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
		file_flight_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlightRequest); i {
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
		file_flight_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlightResponse); i {
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
		file_flight_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlightRequest); i {
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
		file_flight_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlightSlotRequest); i {
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
		file_flight_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlightRequest); i {
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
		file_flight_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlightResponse); i {
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
			RawDescriptor: file_flight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flight_proto_goTypes,
		DependencyIndexes: file_flight_proto_depIdxs,
		MessageInfos:      file_flight_proto_msgTypes,
	}.Build()
	File_flight_proto = out.File
	file_flight_proto_rawDesc = nil
	file_flight_proto_goTypes = nil
	file_flight_proto_depIdxs = nil
}

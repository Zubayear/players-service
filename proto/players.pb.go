// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: proto/players.proto

package players

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

type Position int32

const (
	Position_DEFAULT  Position = 0
	Position_MIDFIELD Position = 1
	Position_FORWARD  Position = 2
	Position_DEFENDER Position = 3
)

// Enum value maps for Position.
var (
	Position_name = map[int32]string{
		0: "DEFAULT",
		1: "MIDFIELD",
		2: "FORWARD",
		3: "DEFENDER",
	}
	Position_value = map[string]int32{
		"DEFAULT":  0,
		"MIDFIELD": 1,
		"FORWARD":  2,
		"DEFENDER": 3,
	}
)

func (x Position) Enum() *Position {
	p := new(Position)
	*p = x
	return p
}

func (x Position) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Position) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_players_proto_enumTypes[0].Descriptor()
}

func (Position) Type() protoreflect.EnumType {
	return &file_proto_players_proto_enumTypes[0]
}

func (x Position) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Position.Descriptor instead.
func (Position) EnumDescriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{0}
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position Position `protobuf:"varint,3,opt,name=position,proto3,enum=players.Position" json:"position,omitempty"`
	ClubName string   `protobuf:"bytes,4,opt,name=clubName,proto3" json:"clubName,omitempty"`
	Age      int32    `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetPosition() Position {
	if x != nil {
		return x.Position
	}
	return Position_DEFAULT
}

func (x *UpdateRequest) GetClubName() string {
	if x != nil {
		return x.ClubName
	}
	return ""
}

func (x *UpdateRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type FilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterKey   string `protobuf:"bytes,1,opt,name=filterKey,proto3" json:"filterKey,omitempty"`
	FilterValue string `protobuf:"bytes,2,opt,name=filterValue,proto3" json:"filterValue,omitempty"`
}

func (x *FilterRequest) Reset() {
	*x = FilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterRequest) ProtoMessage() {}

func (x *FilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterRequest.ProtoReflect.Descriptor instead.
func (*FilterRequest) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{2}
}

func (x *FilterRequest) GetFilterKey() string {
	if x != nil {
		return x.FilterKey
	}
	return ""
}

func (x *FilterRequest) GetFilterValue() string {
	if x != nil {
		return x.FilterValue
	}
	return ""
}

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Position Position `protobuf:"varint,2,opt,name=position,proto3,enum=players.Position" json:"position,omitempty"`
	ClubName string   `protobuf:"bytes,3,opt,name=clubName,proto3" json:"clubName,omitempty"`
	Age      int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{3}
}

func (x *SaveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaveRequest) GetPosition() Position {
	if x != nil {
		return x.Position
	}
	return Position_DEFAULT
}

func (x *SaveRequest) GetClubName() string {
	if x != nil {
		return x.ClubName
	}
	return ""
}

func (x *SaveRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type SaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{4}
}

func (x *SaveResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position Position `protobuf:"varint,3,opt,name=position,proto3,enum=players.Position" json:"position,omitempty"`
	ClubName string   `protobuf:"bytes,4,opt,name=clubName,proto3" json:"clubName,omitempty"`
	Age      int32    `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{5}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetPosition() Position {
	if x != nil {
		return x.Position
	}
	return Position_DEFAULT
}

func (x *Player) GetClubName() string {
	if x != nil {
		return x.ClubName
	}
	return ""
}

func (x *Player) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type PlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PlayerRequest) Reset() {
	*x = PlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRequest) ProtoMessage() {}

func (x *PlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRequest.ProtoReflect.Descriptor instead.
func (*PlayerRequest) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{6}
}

func (x *PlayerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *PlayerResponse) Reset() {
	*x = PlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerResponse) ProtoMessage() {}

func (x *PlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerResponse.ProtoReflect.Descriptor instead.
func (*PlayerResponse) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{7}
}

func (x *PlayerResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type PlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayersRequest) Reset() {
	*x = PlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayersRequest) ProtoMessage() {}

func (x *PlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayersRequest.ProtoReflect.Descriptor instead.
func (*PlayersRequest) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{8}
}

type PlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *PlayersResponse) Reset() {
	*x = PlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayersResponse) ProtoMessage() {}

func (x *PlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayersResponse.ProtoReflect.Descriptor instead.
func (*PlayersResponse) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{9}
}

func (x *PlayersResponse) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_players_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_players_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_players_proto_rawDescGZIP(), []int{11}
}

var File_proto_players_proto protoreflect.FileDescriptor

var file_proto_players_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x90,
	0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x39, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x0d,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7e, 0x0a,
	0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a,
	0x0c, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x89, 0x01,
	0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x0e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x40, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x49, 0x44, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x46, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x03, 0x32, 0xf1, 0x02, 0x0a, 0x07, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x14,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11,
	0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_players_proto_rawDescOnce sync.Once
	file_proto_players_proto_rawDescData = file_proto_players_proto_rawDesc
)

func file_proto_players_proto_rawDescGZIP() []byte {
	file_proto_players_proto_rawDescOnce.Do(func() {
		file_proto_players_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_players_proto_rawDescData)
	})
	return file_proto_players_proto_rawDescData
}

var file_proto_players_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_players_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_players_proto_goTypes = []interface{}{
	(Position)(0),           // 0: players.Position
	(*UpdateRequest)(nil),   // 1: players.UpdateRequest
	(*UpdateResponse)(nil),  // 2: players.UpdateResponse
	(*FilterRequest)(nil),   // 3: players.FilterRequest
	(*SaveRequest)(nil),     // 4: players.SaveRequest
	(*SaveResponse)(nil),    // 5: players.SaveResponse
	(*Player)(nil),          // 6: players.Player
	(*PlayerRequest)(nil),   // 7: players.PlayerRequest
	(*PlayerResponse)(nil),  // 8: players.PlayerResponse
	(*PlayersRequest)(nil),  // 9: players.PlayersRequest
	(*PlayersResponse)(nil), // 10: players.PlayersResponse
	(*DeleteRequest)(nil),   // 11: players.DeleteRequest
	(*DeleteResponse)(nil),  // 12: players.DeleteResponse
}
var file_proto_players_proto_depIdxs = []int32{
	0,  // 0: players.UpdateRequest.position:type_name -> players.Position
	6,  // 1: players.UpdateResponse.player:type_name -> players.Player
	0,  // 2: players.SaveRequest.position:type_name -> players.Position
	0,  // 3: players.Player.position:type_name -> players.Position
	6,  // 4: players.PlayerResponse.player:type_name -> players.Player
	6,  // 5: players.PlayersResponse.players:type_name -> players.Player
	4,  // 6: players.Players.Save:input_type -> players.SaveRequest
	7,  // 7: players.Players.Get:input_type -> players.PlayerRequest
	9,  // 8: players.Players.GetAll:input_type -> players.PlayersRequest
	11, // 9: players.Players.Delete:input_type -> players.DeleteRequest
	3,  // 10: players.Players.Filter:input_type -> players.FilterRequest
	1,  // 11: players.Players.Update:input_type -> players.UpdateRequest
	5,  // 12: players.Players.Save:output_type -> players.SaveResponse
	8,  // 13: players.Players.Get:output_type -> players.PlayerResponse
	10, // 14: players.Players.GetAll:output_type -> players.PlayersResponse
	12, // 15: players.Players.Delete:output_type -> players.DeleteResponse
	10, // 16: players.Players.Filter:output_type -> players.PlayersResponse
	2,  // 17: players.Players.Update:output_type -> players.UpdateResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_players_proto_init() }
func file_proto_players_proto_init() {
	if File_proto_players_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_players_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_players_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_proto_players_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterRequest); i {
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
		file_proto_players_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
		file_proto_players_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResponse); i {
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
		file_proto_players_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_players_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRequest); i {
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
		file_proto_players_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerResponse); i {
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
		file_proto_players_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayersRequest); i {
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
		file_proto_players_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayersResponse); i {
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
		file_proto_players_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_players_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_proto_players_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_players_proto_goTypes,
		DependencyIndexes: file_proto_players_proto_depIdxs,
		EnumInfos:         file_proto_players_proto_enumTypes,
		MessageInfos:      file_proto_players_proto_msgTypes,
	}.Build()
	File_proto_players_proto = out.File
	file_proto_players_proto_rawDesc = nil
	file_proto_players_proto_goTypes = nil
	file_proto_players_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/minecraft.proto

package minecraft

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

type PlayerSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique id of player
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// name of the player
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PlayerSample) Reset() {
	*x = PlayerSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minecraft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerSample) ProtoMessage() {}

func (x *PlayerSample) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minecraft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerSample.ProtoReflect.Descriptor instead.
func (*PlayerSample) Descriptor() ([]byte, []int) {
	return file_proto_minecraft_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerSample) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PlayerSample) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Ping a minecraft server
type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address of the server
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minecraft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minecraft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_proto_minecraft_proto_rawDescGZIP(), []int{1}
}

func (x *PingRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Latency (ms) between us and the server (EU)
	Latency uint32 `protobuf:"varint,1,opt,name=latency,proto3" json:"latency,omitempty"`
	// Number of players online
	Players int32 `protobuf:"varint,2,opt,name=players,proto3" json:"players,omitempty"`
	// Max players ever
	MaxPlayers int32 `protobuf:"varint,3,opt,name=max_players,json=maxPlayers,proto3" json:"max_players,omitempty"`
	// Protocol number of the server
	Protocol int32 `protobuf:"varint,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Favicon in base64
	Favicon string `protobuf:"bytes,5,opt,name=favicon,proto3" json:"favicon,omitempty"`
	// Message of the day
	Motd string `protobuf:"bytes,6,opt,name=motd,proto3" json:"motd,omitempty"`
	// Version of the server
	Version string `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	// List of connected players
	Sample []*PlayerSample `protobuf:"bytes,8,rep,name=sample,proto3" json:"sample,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_minecraft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_minecraft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_proto_minecraft_proto_rawDescGZIP(), []int{2}
}

func (x *PingResponse) GetLatency() uint32 {
	if x != nil {
		return x.Latency
	}
	return 0
}

func (x *PingResponse) GetPlayers() int32 {
	if x != nil {
		return x.Players
	}
	return 0
}

func (x *PingResponse) GetMaxPlayers() int32 {
	if x != nil {
		return x.MaxPlayers
	}
	return 0
}

func (x *PingResponse) GetProtocol() int32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *PingResponse) GetFavicon() string {
	if x != nil {
		return x.Favicon
	}
	return ""
}

func (x *PingResponse) GetMotd() string {
	if x != nil {
		return x.Motd
	}
	return ""
}

func (x *PingResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PingResponse) GetSample() []*PlayerSample {
	if x != nil {
		return x.Sample
	}
	return nil
}

var File_proto_minecraft_proto protoreflect.FileDescriptor

var file_proto_minecraft_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61,
	0x66, 0x74, 0x22, 0x36, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0xf8, 0x01, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x74, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x6f, 0x74, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a,
	0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x32, 0x46,
	0x0a, 0x09, 0x4d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x69,
	0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_minecraft_proto_rawDescOnce sync.Once
	file_proto_minecraft_proto_rawDescData = file_proto_minecraft_proto_rawDesc
)

func file_proto_minecraft_proto_rawDescGZIP() []byte {
	file_proto_minecraft_proto_rawDescOnce.Do(func() {
		file_proto_minecraft_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_minecraft_proto_rawDescData)
	})
	return file_proto_minecraft_proto_rawDescData
}

var file_proto_minecraft_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_minecraft_proto_goTypes = []interface{}{
	(*PlayerSample)(nil), // 0: minecraft.PlayerSample
	(*PingRequest)(nil),  // 1: minecraft.PingRequest
	(*PingResponse)(nil), // 2: minecraft.PingResponse
}
var file_proto_minecraft_proto_depIdxs = []int32{
	0, // 0: minecraft.PingResponse.sample:type_name -> minecraft.PlayerSample
	1, // 1: minecraft.Minecraft.Ping:input_type -> minecraft.PingRequest
	2, // 2: minecraft.Minecraft.Ping:output_type -> minecraft.PingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_minecraft_proto_init() }
func file_proto_minecraft_proto_init() {
	if File_proto_minecraft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_minecraft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerSample); i {
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
		file_proto_minecraft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_proto_minecraft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_proto_minecraft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_minecraft_proto_goTypes,
		DependencyIndexes: file_proto_minecraft_proto_depIdxs,
		MessageInfos:      file_proto_minecraft_proto_msgTypes,
	}.Build()
	File_proto_minecraft_proto = out.File
	file_proto_minecraft_proto_rawDesc = nil
	file_proto_minecraft_proto_goTypes = nil
	file_proto_minecraft_proto_depIdxs = nil
}

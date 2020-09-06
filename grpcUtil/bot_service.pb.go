// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: bot_service.proto

package grpcUtil

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type WordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WordCount int32 `protobuf:"varint,1,opt,name=word_count,json=wordCount,proto3" json:"word_count,omitempty"`
}

func (x *WordRequest) Reset() {
	*x = WordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordRequest) ProtoMessage() {}

func (x *WordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordRequest.ProtoReflect.Descriptor instead.
func (*WordRequest) Descriptor() ([]byte, []int) {
	return file_bot_service_proto_rawDescGZIP(), []int{0}
}

func (x *WordRequest) GetWordCount() int32 {
	if x != nil {
		return x.WordCount
	}
	return 0
}

type WordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word         string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	PartOfSpeech string   `protobuf:"bytes,2,opt,name=part_of_speech,json=partOfSpeech,proto3" json:"part_of_speech,omitempty"`
	Definition   []string `protobuf:"bytes,3,rep,name=definition,proto3" json:"definition,omitempty"`
	Example      []string `protobuf:"bytes,4,rep,name=example,proto3" json:"example,omitempty"`
}

func (x *WordResponse) Reset() {
	*x = WordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordResponse) ProtoMessage() {}

func (x *WordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordResponse.ProtoReflect.Descriptor instead.
func (*WordResponse) Descriptor() ([]byte, []int) {
	return file_bot_service_proto_rawDescGZIP(), []int{1}
}

func (x *WordResponse) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *WordResponse) GetPartOfSpeech() string {
	if x != nil {
		return x.PartOfSpeech
	}
	return ""
}

func (x *WordResponse) GetDefinition() []string {
	if x != nil {
		return x.Definition
	}
	return nil
}

func (x *WordResponse) GetExample() []string {
	if x != nil {
		return x.Example
	}
	return nil
}

var File_bot_service_proto protoreflect.FileDescriptor

var file_bot_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x62, 0x6f, 0x74, 0x22, 0x2c, 0x0a, 0x0b, 0x57, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x64,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x6f,
	0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x53, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x32, 0x63, 0x0a, 0x0a,
	0x42, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x62, 0x6f,
	0x74, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x62, 0x6f, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x55, 0x74, 0x69, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bot_service_proto_rawDescOnce sync.Once
	file_bot_service_proto_rawDescData = file_bot_service_proto_rawDesc
)

func file_bot_service_proto_rawDescGZIP() []byte {
	file_bot_service_proto_rawDescOnce.Do(func() {
		file_bot_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bot_service_proto_rawDescData)
	})
	return file_bot_service_proto_rawDescData
}

var file_bot_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bot_service_proto_goTypes = []interface{}{
	(*WordRequest)(nil),  // 0: com.grpc.languagebot.WordRequest
	(*WordResponse)(nil), // 1: com.grpc.languagebot.WordResponse
}
var file_bot_service_proto_depIdxs = []int32{
	0, // 0: com.grpc.languagebot.BotService.GenerateWord:input_type -> com.grpc.languagebot.WordRequest
	1, // 1: com.grpc.languagebot.BotService.GenerateWord:output_type -> com.grpc.languagebot.WordResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bot_service_proto_init() }
func file_bot_service_proto_init() {
	if File_bot_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bot_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordRequest); i {
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
		file_bot_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordResponse); i {
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
			RawDescriptor: file_bot_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bot_service_proto_goTypes,
		DependencyIndexes: file_bot_service_proto_depIdxs,
		MessageInfos:      file_bot_service_proto_msgTypes,
	}.Build()
	File_bot_service_proto = out.File
	file_bot_service_proto_rawDesc = nil
	file_bot_service_proto_goTypes = nil
	file_bot_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BotServiceClient is the client API for BotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BotServiceClient interface {
	GenerateWord(ctx context.Context, in *WordRequest, opts ...grpc.CallOption) (*WordResponse, error)
}

type botServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBotServiceClient(cc grpc.ClientConnInterface) BotServiceClient {
	return &botServiceClient{cc}
}

func (c *botServiceClient) GenerateWord(ctx context.Context, in *WordRequest, opts ...grpc.CallOption) (*WordResponse, error) {
	out := new(WordResponse)
	err := c.cc.Invoke(ctx, "/com.grpc.languagebot.BotService/GenerateWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServiceServer is the server API for BotService service.
type BotServiceServer interface {
	GenerateWord(context.Context, *WordRequest) (*WordResponse, error)
}

// UnimplementedBotServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBotServiceServer struct {
}

func (*UnimplementedBotServiceServer) GenerateWord(context.Context, *WordRequest) (*WordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateWord not implemented")
}

func RegisterBotServiceServer(s *grpc.Server, srv BotServiceServer) {
	s.RegisterService(&_BotService_serviceDesc, srv)
}

func _BotService_GenerateWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).GenerateWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.grpc.languagebot.BotService/GenerateWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).GenerateWord(ctx, req.(*WordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.grpc.languagebot.BotService",
	HandlerType: (*BotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateWord",
			Handler:    _BotService_GenerateWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot_service.proto",
}
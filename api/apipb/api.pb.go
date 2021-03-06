// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        (unknown)
// source: github.com/Go-gRPC-Server-Client/api/apipb/api.proto

package apipb

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

type ApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ApiRequest) Reset() {
	*x = ApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiRequest) ProtoMessage() {}

func (x *ApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiRequest.ProtoReflect.Descriptor instead.
func (*ApiRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescGZIP(), []int{0}
}

func (x *ApiRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *ApiResponse) Reset() {
	*x = ApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponse) ProtoMessage() {}

func (x *ApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponse.ProtoReflect.Descriptor instead.
func (*ApiResponse) Descriptor() ([]byte, []int) {
	return file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescGZIP(), []int{1}
}

func (x *ApiResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

var File_github_com_Go_gRPC_Server_Client_api_apipb_api_proto protoreflect.FileDescriptor

var file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x2d,
	0x67, 0x52, 0x50, 0x43, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x20, 0x0a, 0x0a, 0x61,
	0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a,
	0x0b, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x32, 0x38, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x0f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x61, 0x70, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescOnce sync.Once
	file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescData = file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDesc
)

func file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescGZIP() []byte {
	file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescOnce.Do(func() {
		file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescData)
	})
	return file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDescData
}

var file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_goTypes = []interface{}{
	(*ApiRequest)(nil),  // 0: api.apiRequest
	(*ApiResponse)(nil), // 1: api.apiResponse
}
var file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_depIdxs = []int32{
	0, // 0: api.apiService.Api:input_type -> api.apiRequest
	1, // 1: api.apiService.Api:output_type -> api.apiResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_init() }
func file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_init() {
	if File_github_com_Go_gRPC_Server_Client_api_apipb_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiRequest); i {
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
		file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResponse); i {
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
			RawDescriptor: file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_goTypes,
		DependencyIndexes: file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_depIdxs,
		MessageInfos:      file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_msgTypes,
	}.Build()
	File_github_com_Go_gRPC_Server_Client_api_apipb_api_proto = out.File
	file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_rawDesc = nil
	file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_goTypes = nil
	file_github_com_Go_gRPC_Server_Client_api_apipb_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiServiceClient interface {
	Api(ctx context.Context, in *ApiRequest, opts ...grpc.CallOption) (*ApiResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) Api(ctx context.Context, in *ApiRequest, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := c.cc.Invoke(ctx, "/api.apiService/Api", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
type ApiServiceServer interface {
	Api(context.Context, *ApiRequest) (*ApiResponse, error)
}

// UnimplementedApiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (*UnimplementedApiServiceServer) Api(context.Context, *ApiRequest) (*ApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Api not implemented")
}

func RegisterApiServiceServer(s *grpc.Server, srv ApiServiceServer) {
	s.RegisterService(&_ApiService_serviceDesc, srv)
}

func _ApiService_Api_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).Api(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.apiService/Api",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).Api(ctx, req.(*ApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.apiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Api",
			Handler:    _ApiService_Api_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/Go-gRPC-Server-Client/api/apipb/api.proto",
}

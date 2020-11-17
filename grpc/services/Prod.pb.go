// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: Prod.proto

package services

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

type ProAreas int32

const (
	ProAreas_A ProAreas = 0
	ProAreas_B ProAreas = 1
	ProAreas_C ProAreas = 2
)

// Enum value maps for ProAreas.
var (
	ProAreas_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
	}
	ProAreas_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
	}
)

func (x ProAreas) Enum() *ProAreas {
	p := new(ProAreas)
	*p = x
	return p
}

func (x ProAreas) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProAreas) Descriptor() protoreflect.EnumDescriptor {
	return file_Prod_proto_enumTypes[0].Descriptor()
}

func (ProAreas) Type() protoreflect.EnumType {
	return &file_Prod_proto_enumTypes[0]
}

func (x ProAreas) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProAreas.Descriptor instead.
func (ProAreas) EnumDescriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId   int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdArea ProAreas `protobuf:"varint,2,opt,name=prod_area,json=prodArea,proto3,enum=services.ProAreas" json:"prod_area,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *Request) GetProdArea() ProAreas {
	if x != nil {
		return x.ProdArea
	}
	return ProAreas_A
}

type QuerySize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *QuerySize) Reset() {
	*x = QuerySize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySize) ProtoMessage() {}

func (x *QuerySize) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySize.ProtoReflect.Descriptor instead.
func (*QuerySize) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{1}
}

func (x *QuerySize) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdStock int32 `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetProdStock() int32 {
	if x != nil {
		return x.ProdStock
	}
	return 0
}

type ResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prodreq []*Response `protobuf:"bytes,1,rep,name=prodreq,proto3" json:"prodreq,omitempty"`
}

func (x *ResponseList) Reset() {
	*x = ResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseList) ProtoMessage() {}

func (x *ResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseList.ProtoReflect.Descriptor instead.
func (*ResponseList) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseList) GetProdreq() []*Response {
	if x != nil {
		return x.Prodreq
	}
	return nil
}

var File_Prod_proto protoreflect.FileDescriptor

var file_Prod_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x41, 0x72, 0x65, 0x61, 0x73, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x41, 0x72, 0x65, 0x61, 0x22, 0x1f, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x29, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x3c, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x72, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x72, 0x65, 0x71, 0x2a, 0x1f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12,
	0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x01, 0x12, 0x05, 0x0a,
	0x01, 0x43, 0x10, 0x02, 0x32, 0xbf, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x69, 0x7a, 0x65, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Prod_proto_rawDescOnce sync.Once
	file_Prod_proto_rawDescData = file_Prod_proto_rawDesc
)

func file_Prod_proto_rawDescGZIP() []byte {
	file_Prod_proto_rawDescOnce.Do(func() {
		file_Prod_proto_rawDescData = protoimpl.X.CompressGZIP(file_Prod_proto_rawDescData)
	})
	return file_Prod_proto_rawDescData
}

var file_Prod_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Prod_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Prod_proto_goTypes = []interface{}{
	(ProAreas)(0),        // 0: services.ProAreas
	(*Request)(nil),      // 1: services.Request
	(*QuerySize)(nil),    // 2: services.QuerySize
	(*Response)(nil),     // 3: services.Response
	(*ResponseList)(nil), // 4: services.ResponseList
	(*ProdModel)(nil),    // 5: services.ProdModel
}
var file_Prod_proto_depIdxs = []int32{
	0, // 0: services.Request.prod_area:type_name -> services.ProAreas
	3, // 1: services.ResponseList.prodreq:type_name -> services.Response
	1, // 2: services.ProdService.GetProdStock:input_type -> services.Request
	2, // 3: services.ProdService.GetProdStocks:input_type -> services.QuerySize
	1, // 4: services.ProdService.GetProdInfo:input_type -> services.Request
	3, // 5: services.ProdService.GetProdStock:output_type -> services.Response
	4, // 6: services.ProdService.GetProdStocks:output_type -> services.ResponseList
	5, // 7: services.ProdService.GetProdInfo:output_type -> services.ProdModel
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Prod_proto_init() }
func file_Prod_proto_init() {
	if File_Prod_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Prod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Prod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySize); i {
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
		file_Prod_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_Prod_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseList); i {
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
			RawDescriptor: file_Prod_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Prod_proto_goTypes,
		DependencyIndexes: file_Prod_proto_depIdxs,
		EnumInfos:         file_Prod_proto_enumTypes,
		MessageInfos:      file_Prod_proto_msgTypes,
	}.Build()
	File_Prod_proto = out.File
	file_Prod_proto_rawDesc = nil
	file_Prod_proto_goTypes = nil
	file_Prod_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ResponseList, error)
	GetProdInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ProdModel, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ResponseList, error) {
	out := new(ResponseList)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ProdModel, error) {
	out := new(ProdModel)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *Request) (*Response, error)
	GetProdStocks(context.Context, *QuerySize) (*ResponseList, error)
	GetProdInfo(context.Context, *Request) (*ProdModel, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}
func (*UnimplementedProdServiceServer) GetProdStocks(context.Context, *QuerySize) (*ResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStocks not implemented")
}
func (*UnimplementedProdServiceServer) GetProdInfo(context.Context, *Request) (*ProdModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdInfo not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStocks(ctx, req.(*QuerySize))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
		{
			MethodName: "GetProdStocks",
			Handler:    _ProdService_GetProdStocks_Handler,
		},
		{
			MethodName: "GetProdInfo",
			Handler:    _ProdService_GetProdInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Prod.proto",
}

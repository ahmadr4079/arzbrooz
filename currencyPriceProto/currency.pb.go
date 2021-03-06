// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/currency.proto

package currencyPriceProto

import (
	context "context"
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

type RequestCurrencyPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestCurrencyPrice) Reset() {
	*x = RequestCurrencyPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCurrencyPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCurrencyPrice) ProtoMessage() {}

func (x *RequestCurrencyPrice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCurrencyPrice.ProtoReflect.Descriptor instead.
func (*RequestCurrencyPrice) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{0}
}

type CurrencyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Market string `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	Price  string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CurrencyInfo) Reset() {
	*x = CurrencyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyInfo) ProtoMessage() {}

func (x *CurrencyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyInfo.ProtoReflect.Descriptor instead.
func (*CurrencyInfo) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyInfo) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

func (x *CurrencyInfo) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type ResponseCurrencyPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyPriceList []*CurrencyInfo `protobuf:"bytes,1,rep,name=currencyPriceList,proto3" json:"currencyPriceList,omitempty"`
}

func (x *ResponseCurrencyPrice) Reset() {
	*x = ResponseCurrencyPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCurrencyPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCurrencyPrice) ProtoMessage() {}

func (x *ResponseCurrencyPrice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCurrencyPrice.ProtoReflect.Descriptor instead.
func (*ResponseCurrencyPrice) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseCurrencyPrice) GetCurrencyPriceList() []*CurrencyInfo {
	if x != nil {
		return x.CurrencyPriceList
	}
	return nil
}

var File_proto_currency_proto protoreflect.FileDescriptor

var file_proto_currency_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x3c,
	0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x15,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0x56, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x16,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_currency_proto_rawDescOnce sync.Once
	file_proto_currency_proto_rawDescData = file_proto_currency_proto_rawDesc
)

func file_proto_currency_proto_rawDescGZIP() []byte {
	file_proto_currency_proto_rawDescOnce.Do(func() {
		file_proto_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currency_proto_rawDescData)
	})
	return file_proto_currency_proto_rawDescData
}

var file_proto_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_currency_proto_goTypes = []interface{}{
	(*RequestCurrencyPrice)(nil),  // 0: RequestCurrencyPrice
	(*CurrencyInfo)(nil),          // 1: CurrencyInfo
	(*ResponseCurrencyPrice)(nil), // 2: ResponseCurrencyPrice
}
var file_proto_currency_proto_depIdxs = []int32{
	1, // 0: ResponseCurrencyPrice.currencyPriceList:type_name -> CurrencyInfo
	0, // 1: currencyPrice.getCurrencyPrice:input_type -> RequestCurrencyPrice
	2, // 2: currencyPrice.getCurrencyPrice:output_type -> ResponseCurrencyPrice
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_currency_proto_init() }
func file_proto_currency_proto_init() {
	if File_proto_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCurrencyPrice); i {
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
		file_proto_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyInfo); i {
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
		file_proto_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCurrencyPrice); i {
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
			RawDescriptor: file_proto_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currency_proto_goTypes,
		DependencyIndexes: file_proto_currency_proto_depIdxs,
		MessageInfos:      file_proto_currency_proto_msgTypes,
	}.Build()
	File_proto_currency_proto = out.File
	file_proto_currency_proto_rawDesc = nil
	file_proto_currency_proto_goTypes = nil
	file_proto_currency_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyPriceClient is the client API for CurrencyPrice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyPriceClient interface {
	GetCurrencyPrice(ctx context.Context, in *RequestCurrencyPrice, opts ...grpc.CallOption) (CurrencyPrice_GetCurrencyPriceClient, error)
}

type currencyPriceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyPriceClient(cc grpc.ClientConnInterface) CurrencyPriceClient {
	return &currencyPriceClient{cc}
}

func (c *currencyPriceClient) GetCurrencyPrice(ctx context.Context, in *RequestCurrencyPrice, opts ...grpc.CallOption) (CurrencyPrice_GetCurrencyPriceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CurrencyPrice_serviceDesc.Streams[0], "/currencyPrice/getCurrencyPrice", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyPriceGetCurrencyPriceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyPrice_GetCurrencyPriceClient interface {
	Recv() (*ResponseCurrencyPrice, error)
	grpc.ClientStream
}

type currencyPriceGetCurrencyPriceClient struct {
	grpc.ClientStream
}

func (x *currencyPriceGetCurrencyPriceClient) Recv() (*ResponseCurrencyPrice, error) {
	m := new(ResponseCurrencyPrice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyPriceServer is the server API for CurrencyPrice service.
type CurrencyPriceServer interface {
	GetCurrencyPrice(*RequestCurrencyPrice, CurrencyPrice_GetCurrencyPriceServer) error
}

// UnimplementedCurrencyPriceServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyPriceServer struct {
}

func (*UnimplementedCurrencyPriceServer) GetCurrencyPrice(*RequestCurrencyPrice, CurrencyPrice_GetCurrencyPriceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCurrencyPrice not implemented")
}

func RegisterCurrencyPriceServer(s *grpc.Server, srv CurrencyPriceServer) {
	s.RegisterService(&_CurrencyPrice_serviceDesc, srv)
}

func _CurrencyPrice_GetCurrencyPrice_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestCurrencyPrice)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyPriceServer).GetCurrencyPrice(m, &currencyPriceGetCurrencyPriceServer{stream})
}

type CurrencyPrice_GetCurrencyPriceServer interface {
	Send(*ResponseCurrencyPrice) error
	grpc.ServerStream
}

type currencyPriceGetCurrencyPriceServer struct {
	grpc.ServerStream
}

func (x *currencyPriceGetCurrencyPriceServer) Send(m *ResponseCurrencyPrice) error {
	return x.ServerStream.SendMsg(m)
}

var _CurrencyPrice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "currencyPrice",
	HandlerType: (*CurrencyPriceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getCurrencyPrice",
			Handler:       _CurrencyPrice_GetCurrencyPrice_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/currency.proto",
}

// Code generated by protoc-gen-go.
// source: tradeCheck.proto
// DO NOT EDIT!

/*
Package tradeCheck is a generated protocol buffer package.

It is generated from these files:
	tradeCheck.proto

It has these top-level messages:
	TradeHandle
	TradeCheckInfo
	TokenName
	TradeExpect
*/
package tradeCheck

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TradeHandle struct {
	Handler int32 `protobuf:"varint,1,opt,name=handler" json:"handler,omitempty"`
}

func (m *TradeHandle) Reset()                    { *m = TradeHandle{} }
func (m *TradeHandle) String() string            { return proto.CompactTextString(m) }
func (*TradeHandle) ProtoMessage()               {}
func (*TradeHandle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TradeCheckInfo struct {
	Info string `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *TradeCheckInfo) Reset()                    { *m = TradeCheckInfo{} }
func (m *TradeCheckInfo) String() string            { return proto.CompactTextString(m) }
func (*TradeCheckInfo) ProtoMessage()               {}
func (*TradeCheckInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TokenName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TokenName) Reset()                    { *m = TokenName{} }
func (m *TokenName) String() string            { return proto.CompactTextString(m) }
func (*TokenName) ProtoMessage()               {}
func (*TokenName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type TradeExpect struct {
	Pairs     string  `protobuf:"bytes,1,opt,name=pairs" json:"pairs,omitempty"`
	LimitBuy  float32 `protobuf:"fixed32,2,opt,name=limitBuy" json:"limitBuy,omitempty"`
	LimitSeel float32 `protobuf:"fixed32,3,opt,name=limitSeel" json:"limitSeel,omitempty"`
}

func (m *TradeExpect) Reset()                    { *m = TradeExpect{} }
func (m *TradeExpect) String() string            { return proto.CompactTextString(m) }
func (*TradeExpect) ProtoMessage()               {}
func (*TradeExpect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*TradeHandle)(nil), "tradeCheck.TradeHandle")
	proto.RegisterType((*TradeCheckInfo)(nil), "tradeCheck.TradeCheckInfo")
	proto.RegisterType((*TokenName)(nil), "tradeCheck.TokenName")
	proto.RegisterType((*TradeExpect)(nil), "tradeCheck.TradeExpect")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for TradeCheckService service

type TradeCheckServiceClient interface {
	// 返回当前买卖策略说明
	GetInfo(ctx context.Context, in *TradeHandle, opts ...grpc.CallOption) (*TradeCheckInfo, error)
	// 根据现有走势数组判断买卖点
	TradeCheck(ctx context.Context, in *TokenName, opts ...grpc.CallOption) (*TradeExpect, error)
}

type tradeCheckServiceClient struct {
	cc *grpc.ClientConn
}

func NewTradeCheckServiceClient(cc *grpc.ClientConn) TradeCheckServiceClient {
	return &tradeCheckServiceClient{cc}
}

func (c *tradeCheckServiceClient) GetInfo(ctx context.Context, in *TradeHandle, opts ...grpc.CallOption) (*TradeCheckInfo, error) {
	out := new(TradeCheckInfo)
	err := grpc.Invoke(ctx, "/tradeCheck.TradeCheckService/getInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeCheckServiceClient) TradeCheck(ctx context.Context, in *TokenName, opts ...grpc.CallOption) (*TradeExpect, error) {
	out := new(TradeExpect)
	err := grpc.Invoke(ctx, "/tradeCheck.TradeCheckService/tradeCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TradeCheckService service

type TradeCheckServiceServer interface {
	// 返回当前买卖策略说明
	GetInfo(context.Context, *TradeHandle) (*TradeCheckInfo, error)
	// 根据现有走势数组判断买卖点
	TradeCheck(context.Context, *TokenName) (*TradeExpect, error)
}

func RegisterTradeCheckServiceServer(s *grpc.Server, srv TradeCheckServiceServer) {
	s.RegisterService(&_TradeCheckService_serviceDesc, srv)
}

func _TradeCheckService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeHandle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeCheckServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tradeCheck.TradeCheckService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeCheckServiceServer).GetInfo(ctx, req.(*TradeHandle))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeCheckService_TradeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeCheckServiceServer).TradeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tradeCheck.TradeCheckService/TradeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeCheckServiceServer).TradeCheck(ctx, req.(*TokenName))
	}
	return interceptor(ctx, in, info, handler)
}

var _TradeCheckService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tradeCheck.TradeCheckService",
	HandlerType: (*TradeCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getInfo",
			Handler:    _TradeCheckService_GetInfo_Handler,
		},
		{
			MethodName: "tradeCheck",
			Handler:    _TradeCheckService_TradeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("tradeCheck.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x29, 0x4a, 0x4c,
	0x49, 0x75, 0xce, 0x48, 0x4d, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xa9, 0x73, 0x71, 0x87, 0x80, 0x78, 0x1e, 0x89, 0x79, 0x29, 0x39, 0xa9, 0x42, 0x12, 0x5c,
	0xec, 0x19, 0x60, 0x56, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x8c, 0xab, 0xa4, 0xc2,
	0xc5, 0x17, 0x02, 0xd7, 0xe6, 0x99, 0x97, 0x96, 0x2f, 0x24, 0xc4, 0xc5, 0x92, 0x99, 0x97, 0x96,
	0x0f, 0x56, 0xc8, 0x19, 0x04, 0x66, 0x2b, 0xc9, 0x73, 0x71, 0x86, 0xe4, 0x67, 0xa7, 0xe6, 0xf9,
	0x25, 0xe6, 0xa6, 0x82, 0x14, 0xe4, 0x25, 0xe6, 0xa6, 0xc2, 0x14, 0x80, 0xd8, 0x4a, 0xb1, 0x50,
	0xfb, 0x5c, 0x2b, 0x0a, 0x52, 0x93, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x0b, 0x12, 0x33, 0x8b, 0x8a,
	0xa1, 0x6a, 0x20, 0x1c, 0x21, 0x29, 0x2e, 0x8e, 0x9c, 0xcc, 0xdc, 0xcc, 0x12, 0xa7, 0xd2, 0x4a,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xa6, 0x20, 0x38, 0x5f, 0x48, 0x86, 0x8b, 0x13, 0xcc, 0x0e, 0x4e,
	0x4d, 0xcd, 0x91, 0x60, 0x06, 0x4b, 0x22, 0x04, 0x8c, 0xa6, 0x32, 0x72, 0x09, 0x22, 0x9c, 0x19,
	0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc0, 0xc5, 0x9e, 0x9e, 0x5a, 0x02, 0x76, 0xb4,
	0xb8, 0x1e, 0x52, 0x70, 0x20, 0xf9, 0x5c, 0x4a, 0x0a, 0x43, 0x02, 0xee, 0x53, 0x25, 0x06, 0x21,
	0x3b, 0x2e, 0xa4, 0x40, 0x13, 0x12, 0x45, 0x51, 0x0b, 0xf3, 0xaf, 0x14, 0xa6, 0xd9, 0x10, 0x5f,
	0x2a, 0x31, 0x38, 0x99, 0x70, 0x09, 0x65, 0xe6, 0xeb, 0xa5, 0x17, 0x15, 0x24, 0x23, 0xa9, 0x71,
	0xe2, 0x47, 0xb0, 0x03, 0x40, 0x31, 0x13, 0xc0, 0xb8, 0x88, 0x09, 0xc9, 0x9e, 0x24, 0x36, 0x70,
	0x7c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x14, 0x97, 0xa6, 0xf5, 0xc3, 0x01, 0x00, 0x00,
}
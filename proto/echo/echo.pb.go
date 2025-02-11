// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/echo/echo.proto

package echo

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

type RepeatRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepeatRequest) Reset()         { *m = RepeatRequest{} }
func (m *RepeatRequest) String() string { return proto.CompactTextString(m) }
func (*RepeatRequest) ProtoMessage()    {}
func (*RepeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{0}
}
func (m *RepeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatRequest.Unmarshal(m, b)
}
func (m *RepeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatRequest.Marshal(b, m, deterministic)
}
func (dst *RepeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatRequest.Merge(dst, src)
}
func (m *RepeatRequest) XXX_Size() int {
	return xxx_messageInfo_RepeatRequest.Size(m)
}
func (m *RepeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatRequest proto.InternalMessageInfo

func (m *RepeatRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type RepeatResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepeatResponse) Reset()         { *m = RepeatResponse{} }
func (m *RepeatResponse) String() string { return proto.CompactTextString(m) }
func (*RepeatResponse) ProtoMessage()    {}
func (*RepeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{1}
}
func (m *RepeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatResponse.Unmarshal(m, b)
}
func (m *RepeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatResponse.Marshal(b, m, deterministic)
}
func (dst *RepeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatResponse.Merge(dst, src)
}
func (m *RepeatResponse) XXX_Size() int {
	return xxx_messageInfo_RepeatResponse.Size(m)
}
func (m *RepeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatResponse proto.InternalMessageInfo

func (m *RepeatResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ScreamRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScreamRequest) Reset()         { *m = ScreamRequest{} }
func (m *ScreamRequest) String() string { return proto.CompactTextString(m) }
func (*ScreamRequest) ProtoMessage()    {}
func (*ScreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{2}
}
func (m *ScreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScreamRequest.Unmarshal(m, b)
}
func (m *ScreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScreamRequest.Marshal(b, m, deterministic)
}
func (dst *ScreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScreamRequest.Merge(dst, src)
}
func (m *ScreamRequest) XXX_Size() int {
	return xxx_messageInfo_ScreamRequest.Size(m)
}
func (m *ScreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScreamRequest proto.InternalMessageInfo

func (m *ScreamRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ScreamResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScreamResponse) Reset()         { *m = ScreamResponse{} }
func (m *ScreamResponse) String() string { return proto.CompactTextString(m) }
func (*ScreamResponse) ProtoMessage()    {}
func (*ScreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{3}
}
func (m *ScreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScreamResponse.Unmarshal(m, b)
}
func (m *ScreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScreamResponse.Marshal(b, m, deterministic)
}
func (dst *ScreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScreamResponse.Merge(dst, src)
}
func (m *ScreamResponse) XXX_Size() int {
	return xxx_messageInfo_ScreamResponse.Size(m)
}
func (m *ScreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScreamResponse proto.InternalMessageInfo

func (m *ScreamResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AskGoogleRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskGoogleRequest) Reset()         { *m = AskGoogleRequest{} }
func (m *AskGoogleRequest) String() string { return proto.CompactTextString(m) }
func (*AskGoogleRequest) ProtoMessage()    {}
func (*AskGoogleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{4}
}
func (m *AskGoogleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskGoogleRequest.Unmarshal(m, b)
}
func (m *AskGoogleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskGoogleRequest.Marshal(b, m, deterministic)
}
func (dst *AskGoogleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskGoogleRequest.Merge(dst, src)
}
func (m *AskGoogleRequest) XXX_Size() int {
	return xxx_messageInfo_AskGoogleRequest.Size(m)
}
func (m *AskGoogleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AskGoogleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AskGoogleRequest proto.InternalMessageInfo

func (m *AskGoogleRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AskGoogleResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskGoogleResponse) Reset()         { *m = AskGoogleResponse{} }
func (m *AskGoogleResponse) String() string { return proto.CompactTextString(m) }
func (*AskGoogleResponse) ProtoMessage()    {}
func (*AskGoogleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{5}
}
func (m *AskGoogleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskGoogleResponse.Unmarshal(m, b)
}
func (m *AskGoogleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskGoogleResponse.Marshal(b, m, deterministic)
}
func (dst *AskGoogleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskGoogleResponse.Merge(dst, src)
}
func (m *AskGoogleResponse) XXX_Size() int {
	return xxx_messageInfo_AskGoogleResponse.Size(m)
}
func (m *AskGoogleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AskGoogleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AskGoogleResponse proto.InternalMessageInfo

func (m *AskGoogleResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AskDBRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskDBRequest) Reset()         { *m = AskDBRequest{} }
func (m *AskDBRequest) String() string { return proto.CompactTextString(m) }
func (*AskDBRequest) ProtoMessage()    {}
func (*AskDBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{6}
}
func (m *AskDBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskDBRequest.Unmarshal(m, b)
}
func (m *AskDBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskDBRequest.Marshal(b, m, deterministic)
}
func (dst *AskDBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskDBRequest.Merge(dst, src)
}
func (m *AskDBRequest) XXX_Size() int {
	return xxx_messageInfo_AskDBRequest.Size(m)
}
func (m *AskDBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AskDBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AskDBRequest proto.InternalMessageInfo

func (m *AskDBRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AskDBResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskDBResponse) Reset()         { *m = AskDBResponse{} }
func (m *AskDBResponse) String() string { return proto.CompactTextString(m) }
func (*AskDBResponse) ProtoMessage()    {}
func (*AskDBResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{7}
}
func (m *AskDBResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskDBResponse.Unmarshal(m, b)
}
func (m *AskDBResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskDBResponse.Marshal(b, m, deterministic)
}
func (dst *AskDBResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskDBResponse.Merge(dst, src)
}
func (m *AskDBResponse) XXX_Size() int {
	return xxx_messageInfo_AskDBResponse.Size(m)
}
func (m *AskDBResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AskDBResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AskDBResponse proto.InternalMessageInfo

func (m *AskDBResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AskRedisRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskRedisRequest) Reset()         { *m = AskRedisRequest{} }
func (m *AskRedisRequest) String() string { return proto.CompactTextString(m) }
func (*AskRedisRequest) ProtoMessage()    {}
func (*AskRedisRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{8}
}
func (m *AskRedisRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskRedisRequest.Unmarshal(m, b)
}
func (m *AskRedisRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskRedisRequest.Marshal(b, m, deterministic)
}
func (dst *AskRedisRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskRedisRequest.Merge(dst, src)
}
func (m *AskRedisRequest) XXX_Size() int {
	return xxx_messageInfo_AskRedisRequest.Size(m)
}
func (m *AskRedisRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AskRedisRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AskRedisRequest proto.InternalMessageInfo

func (m *AskRedisRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AskRedisResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskRedisResponse) Reset()         { *m = AskRedisResponse{} }
func (m *AskRedisResponse) String() string { return proto.CompactTextString(m) }
func (*AskRedisResponse) ProtoMessage()    {}
func (*AskRedisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{9}
}
func (m *AskRedisResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskRedisResponse.Unmarshal(m, b)
}
func (m *AskRedisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskRedisResponse.Marshal(b, m, deterministic)
}
func (dst *AskRedisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskRedisResponse.Merge(dst, src)
}
func (m *AskRedisResponse) XXX_Size() int {
	return xxx_messageInfo_AskRedisResponse.Size(m)
}
func (m *AskRedisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AskRedisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AskRedisResponse proto.InternalMessageInfo

func (m *AskRedisResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AskOracleRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskOracleRequest) Reset()         { *m = AskOracleRequest{} }
func (m *AskOracleRequest) String() string { return proto.CompactTextString(m) }
func (*AskOracleRequest) ProtoMessage()    {}
func (*AskOracleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{10}
}
func (m *AskOracleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskOracleRequest.Unmarshal(m, b)
}
func (m *AskOracleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskOracleRequest.Marshal(b, m, deterministic)
}
func (dst *AskOracleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskOracleRequest.Merge(dst, src)
}
func (m *AskOracleRequest) XXX_Size() int {
	return xxx_messageInfo_AskOracleRequest.Size(m)
}
func (m *AskOracleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AskOracleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AskOracleRequest proto.InternalMessageInfo

func (m *AskOracleRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AskOracleResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AskOracleResponse) Reset()         { *m = AskOracleResponse{} }
func (m *AskOracleResponse) String() string { return proto.CompactTextString(m) }
func (*AskOracleResponse) ProtoMessage()    {}
func (*AskOracleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_2bb4d59af0570099, []int{11}
}
func (m *AskOracleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskOracleResponse.Unmarshal(m, b)
}
func (m *AskOracleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskOracleResponse.Marshal(b, m, deterministic)
}
func (dst *AskOracleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskOracleResponse.Merge(dst, src)
}
func (m *AskOracleResponse) XXX_Size() int {
	return xxx_messageInfo_AskOracleResponse.Size(m)
}
func (m *AskOracleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AskOracleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AskOracleResponse proto.InternalMessageInfo

func (m *AskOracleResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*RepeatRequest)(nil), "echo.RepeatRequest")
	proto.RegisterType((*RepeatResponse)(nil), "echo.RepeatResponse")
	proto.RegisterType((*ScreamRequest)(nil), "echo.ScreamRequest")
	proto.RegisterType((*ScreamResponse)(nil), "echo.ScreamResponse")
	proto.RegisterType((*AskGoogleRequest)(nil), "echo.AskGoogleRequest")
	proto.RegisterType((*AskGoogleResponse)(nil), "echo.AskGoogleResponse")
	proto.RegisterType((*AskDBRequest)(nil), "echo.AskDBRequest")
	proto.RegisterType((*AskDBResponse)(nil), "echo.AskDBResponse")
	proto.RegisterType((*AskRedisRequest)(nil), "echo.AskRedisRequest")
	proto.RegisterType((*AskRedisResponse)(nil), "echo.AskRedisResponse")
	proto.RegisterType((*AskOracleRequest)(nil), "echo.AskOracleRequest")
	proto.RegisterType((*AskOracleResponse)(nil), "echo.AskOracleResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	// First test endpoint
	Repeat(ctx context.Context, in *RepeatRequest, opts ...grpc.CallOption) (*RepeatResponse, error)
	// Second test endpoint
	Scream(ctx context.Context, in *ScreamRequest, opts ...grpc.CallOption) (*ScreamResponse, error)
	// Check integration with http endpoint
	AskGoogle(ctx context.Context, in *AskGoogleRequest, opts ...grpc.CallOption) (*AskGoogleResponse, error)
	// Check integartion with Mysql database
	AskDB(ctx context.Context, in *AskDBRequest, opts ...grpc.CallOption) (*AskDBResponse, error)
	// Check integartion with Redis cache
	AskRedis(ctx context.Context, in *AskRedisRequest, opts ...grpc.CallOption) (*AskRedisResponse, error)
	// Check integartion with second grpc service
	AskOracle(ctx context.Context, in *AskOracleRequest, opts ...grpc.CallOption) (*AskOracleResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Repeat(ctx context.Context, in *RepeatRequest, opts ...grpc.CallOption) (*RepeatResponse, error) {
	out := new(RepeatResponse)
	err := c.cc.Invoke(ctx, "/echo.Echo/Repeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Scream(ctx context.Context, in *ScreamRequest, opts ...grpc.CallOption) (*ScreamResponse, error) {
	out := new(ScreamResponse)
	err := c.cc.Invoke(ctx, "/echo.Echo/Scream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) AskGoogle(ctx context.Context, in *AskGoogleRequest, opts ...grpc.CallOption) (*AskGoogleResponse, error) {
	out := new(AskGoogleResponse)
	err := c.cc.Invoke(ctx, "/echo.Echo/AskGoogle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) AskDB(ctx context.Context, in *AskDBRequest, opts ...grpc.CallOption) (*AskDBResponse, error) {
	out := new(AskDBResponse)
	err := c.cc.Invoke(ctx, "/echo.Echo/AskDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) AskRedis(ctx context.Context, in *AskRedisRequest, opts ...grpc.CallOption) (*AskRedisResponse, error) {
	out := new(AskRedisResponse)
	err := c.cc.Invoke(ctx, "/echo.Echo/AskRedis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) AskOracle(ctx context.Context, in *AskOracleRequest, opts ...grpc.CallOption) (*AskOracleResponse, error) {
	out := new(AskOracleResponse)
	err := c.cc.Invoke(ctx, "/echo.Echo/AskOracle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	// First test endpoint
	Repeat(context.Context, *RepeatRequest) (*RepeatResponse, error)
	// Second test endpoint
	Scream(context.Context, *ScreamRequest) (*ScreamResponse, error)
	// Check integration with http endpoint
	AskGoogle(context.Context, *AskGoogleRequest) (*AskGoogleResponse, error)
	// Check integartion with Mysql database
	AskDB(context.Context, *AskDBRequest) (*AskDBResponse, error)
	// Check integartion with Redis cache
	AskRedis(context.Context, *AskRedisRequest) (*AskRedisResponse, error)
	// Check integartion with second grpc service
	AskOracle(context.Context, *AskOracleRequest) (*AskOracleResponse, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Repeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Repeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Repeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Repeat(ctx, req.(*RepeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Scream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Scream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Scream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Scream(ctx, req.(*ScreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_AskGoogle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskGoogleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).AskGoogle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/AskGoogle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).AskGoogle(ctx, req.(*AskGoogleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_AskDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).AskDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/AskDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).AskDB(ctx, req.(*AskDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_AskRedis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskRedisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).AskRedis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/AskRedis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).AskRedis(ctx, req.(*AskRedisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_AskOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskOracleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).AskOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/AskOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).AskOracle(ctx, req.(*AskOracleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Repeat",
			Handler:    _Echo_Repeat_Handler,
		},
		{
			MethodName: "Scream",
			Handler:    _Echo_Scream_Handler,
		},
		{
			MethodName: "AskGoogle",
			Handler:    _Echo_AskGoogle_Handler,
		},
		{
			MethodName: "AskDB",
			Handler:    _Echo_AskDB_Handler,
		},
		{
			MethodName: "AskRedis",
			Handler:    _Echo_AskRedis_Handler,
		},
		{
			MethodName: "AskOracle",
			Handler:    _Echo_AskOracle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/echo/echo.proto",
}

func init() { proto.RegisterFile("proto/echo/echo.proto", fileDescriptor_echo_2bb4d59af0570099) }

var fileDescriptor_echo_2bb4d59af0570099 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x53, 0x89, 0xc5, 0x0e, 0xc6, 0x8f, 0xa9, 0xad, 0x92, 0x93, 0x6c, 0xfc, 0x3a, 0x55,
	0xa8, 0x78, 0x12, 0x84, 0x88, 0xe2, 0x51, 0x88, 0x4f, 0xb0, 0xa6, 0x83, 0x95, 0xa8, 0x1b, 0xb3,
	0xf1, 0x91, 0x7d, 0x0f, 0xe9, 0x6e, 0x77, 0xd9, 0x8d, 0xba, 0xf5, 0x52, 0xb6, 0xc3, 0x6f, 0xfe,
	0xcc, 0xcc, 0x8f, 0xc0, 0xa8, 0x6e, 0x44, 0x2b, 0xce, 0xa9, 0x9c, 0xeb, 0x9f, 0x89, 0xfa, 0x8f,
	0xf1, 0xe2, 0xcd, 0x32, 0x48, 0x0a, 0xaa, 0x89, 0xb7, 0x05, 0x7d, 0x7c, 0x92, 0x6c, 0x11, 0x21,
	0x9e, 0xf1, 0x96, 0x1f, 0xf4, 0x0e, 0x7b, 0x67, 0x83, 0x42, 0xbd, 0xd9, 0x11, 0x6c, 0x19, 0x48,
	0xd6, 0xe2, 0x5d, 0xd2, 0xaf, 0x54, 0x06, 0xc9, 0x63, 0xd9, 0x10, 0x7f, 0x5b, 0x11, 0x65, 0xa0,
	0x40, 0xd4, 0x09, 0xec, 0xe4, 0xb2, 0xba, 0x17, 0xe2, 0xf9, 0x95, 0x42, 0x69, 0xa7, 0xb0, 0xeb,
	0x70, 0x81, 0x40, 0x06, 0x9b, 0xb9, 0xac, 0x6e, 0x6f, 0x42, 0x61, 0x19, 0x24, 0x4b, 0x26, 0x10,
	0x74, 0x0c, 0xdb, 0xb9, 0xac, 0x0a, 0x9a, 0xbd, 0xc8, 0x50, 0x96, 0x5e, 0x60, 0x89, 0xad, 0x5c,
	0xf4, 0xa1, 0xe1, 0xe5, 0x7f, 0x16, 0x35, 0xdc, 0xdf, 0x81, 0xd3, 0xaf, 0x35, 0x88, 0xef, 0xca,
	0xb9, 0xc0, 0x4b, 0xe8, 0x6b, 0x67, 0x38, 0x9c, 0x28, 0xeb, 0x9e, 0xe6, 0x74, 0xcf, 0x2f, 0xea,
	0x44, 0x16, 0x2d, 0xda, 0xb4, 0x1f, 0xd3, 0xe6, 0x29, 0x35, 0x6d, 0xbe, 0x42, 0x16, 0xe1, 0x35,
	0x0c, 0xac, 0x08, 0x1c, 0x6b, 0xa8, 0x6b, 0x30, 0xdd, 0xff, 0x51, 0xb7, 0xfd, 0x53, 0x58, 0x57,
	0xb7, 0x47, 0xb4, 0x8c, 0x95, 0x95, 0x0e, 0xbd, 0x9a, 0xed, 0xb9, 0x82, 0x0d, 0x73, 0x63, 0x1c,
	0x59, 0xc4, 0x55, 0x93, 0x8e, 0xbb, 0xe5, 0xce, 0xc0, 0xfa, 0xa0, 0xce, 0xc0, 0x9e, 0x09, 0x67,
	0x60, 0xff, 0xf2, 0x2c, 0x7a, 0xea, 0xab, 0x8f, 0xe8, 0xe2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x62,
	0xf1, 0x3d, 0xbe, 0x5d, 0x03, 0x00, 0x00,
}

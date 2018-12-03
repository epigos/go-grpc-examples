// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc/calcpb/calc.proto

package calcpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Calculator struct {
	OperandOne           int32    `protobuf:"varint,1,opt,name=operand_one,json=operandOne,proto3" json:"operand_one,omitempty"`
	OperandTwo           int32    `protobuf:"varint,2,opt,name=operand_two,json=operandTwo,proto3" json:"operand_two,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculator) Reset()         { *m = Calculator{} }
func (m *Calculator) String() string { return proto.CompactTextString(m) }
func (*Calculator) ProtoMessage()    {}
func (*Calculator) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e5b7e388c7a849, []int{0}
}

func (m *Calculator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculator.Unmarshal(m, b)
}
func (m *Calculator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculator.Marshal(b, m, deterministic)
}
func (m *Calculator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculator.Merge(m, src)
}
func (m *Calculator) XXX_Size() int {
	return xxx_messageInfo_Calculator.Size(m)
}
func (m *Calculator) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculator.DiscardUnknown(m)
}

var xxx_messageInfo_Calculator proto.InternalMessageInfo

func (m *Calculator) GetOperandOne() int32 {
	if m != nil {
		return m.OperandOne
	}
	return 0
}

func (m *Calculator) GetOperandTwo() int32 {
	if m != nil {
		return m.OperandTwo
	}
	return 0
}

type CalculatorRequest struct {
	Calculate            *Calculator `protobuf:"bytes,1,opt,name=calculate,proto3" json:"calculate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e5b7e388c7a849, []int{1}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetCalculate() *Calculator {
	if m != nil {
		return m.Calculate
	}
	return nil
}

type CalculatorResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e5b7e388c7a849, []int{2}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Calculator)(nil), "calc.Calculator")
	proto.RegisterType((*CalculatorRequest)(nil), "calc.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calc.CalculatorResponse")
}

func init() { proto.RegisterFile("calc/calcpb/calc.proto", fileDescriptor_53e5b7e388c7a849) }

var fileDescriptor_53e5b7e388c7a849 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4e, 0xcc, 0x49,
	0xd6, 0x07, 0x11, 0x05, 0x49, 0x60, 0x4a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x05, 0xc4,
	0x56, 0xf2, 0xe3, 0xe2, 0x72, 0x4e, 0xcc, 0x49, 0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0x12, 0x92,
	0xe7, 0xe2, 0xce, 0x2f, 0x48, 0x2d, 0x4a, 0xcc, 0x4b, 0x89, 0xcf, 0xcf, 0x4b, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x0d, 0xe2, 0x82, 0x0a, 0xf9, 0xe7, 0xa5, 0x22, 0x2b, 0x28, 0x29, 0xcf, 0x97,
	0x60, 0x42, 0x51, 0x10, 0x52, 0x9e, 0xaf, 0xe4, 0xcc, 0x25, 0x88, 0x30, 0x2f, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8f, 0x8b, 0x33, 0x19, 0x2a, 0x08, 0x31, 0x94, 0xdb, 0x48, 0x40,
	0x0f, 0xec, 0x14, 0x24, 0xb5, 0x08, 0x25, 0x4a, 0x3a, 0x5c, 0x42, 0xc8, 0x86, 0x14, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x40, 0xdd, 0x05,
	0xe5, 0x19, 0xf9, 0x23, 0x5b, 0x19, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xc5, 0xc5,
	0x1c, 0x5c, 0x9a, 0x2b, 0x24, 0x8e, 0x61, 0x0d, 0xc4, 0x49, 0x52, 0x12, 0x98, 0x12, 0x10, 0x6b,
	0x94, 0x18, 0x9c, 0x38, 0xa2, 0xd8, 0x20, 0xc1, 0x95, 0xc4, 0x06, 0x0e, 0x2a, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x25, 0x71, 0xfa, 0x26, 0x44, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calc.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc/calcpb/calc.proto",
}
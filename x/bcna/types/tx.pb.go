// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bcna/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # proto/tx/message
type MsgCreateBitcannaid struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Bcnaid  string `protobuf:"bytes,2,opt,name=bcnaid,proto3" json:"bcnaid,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgCreateBitcannaid) Reset()         { *m = MsgCreateBitcannaid{} }
func (m *MsgCreateBitcannaid) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBitcannaid) ProtoMessage()    {}
func (*MsgCreateBitcannaid) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d2bc8ae8552a21, []int{0}
}
func (m *MsgCreateBitcannaid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBitcannaid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBitcannaid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBitcannaid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBitcannaid.Merge(m, src)
}
func (m *MsgCreateBitcannaid) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBitcannaid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBitcannaid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBitcannaid proto.InternalMessageInfo

func (m *MsgCreateBitcannaid) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateBitcannaid) GetBcnaid() string {
	if m != nil {
		return m.Bcnaid
	}
	return ""
}

func (m *MsgCreateBitcannaid) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgCreateBitcannaidResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateBitcannaidResponse) Reset()         { *m = MsgCreateBitcannaidResponse{} }
func (m *MsgCreateBitcannaidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBitcannaidResponse) ProtoMessage()    {}
func (*MsgCreateBitcannaidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d2bc8ae8552a21, []int{1}
}
func (m *MsgCreateBitcannaidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBitcannaidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBitcannaidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBitcannaidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBitcannaidResponse.Merge(m, src)
}
func (m *MsgCreateBitcannaidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBitcannaidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBitcannaidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBitcannaidResponse proto.InternalMessageInfo

func (m *MsgCreateBitcannaidResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgUpdateBitcannaid struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Bcnaid  string `protobuf:"bytes,3,opt,name=bcnaid,proto3" json:"bcnaid,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgUpdateBitcannaid) Reset()         { *m = MsgUpdateBitcannaid{} }
func (m *MsgUpdateBitcannaid) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateBitcannaid) ProtoMessage()    {}
func (*MsgUpdateBitcannaid) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d2bc8ae8552a21, []int{2}
}
func (m *MsgUpdateBitcannaid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateBitcannaid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateBitcannaid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateBitcannaid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateBitcannaid.Merge(m, src)
}
func (m *MsgUpdateBitcannaid) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateBitcannaid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateBitcannaid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateBitcannaid proto.InternalMessageInfo

func (m *MsgUpdateBitcannaid) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateBitcannaid) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgUpdateBitcannaid) GetBcnaid() string {
	if m != nil {
		return m.Bcnaid
	}
	return ""
}

func (m *MsgUpdateBitcannaid) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgUpdateBitcannaidResponse struct {
}

func (m *MsgUpdateBitcannaidResponse) Reset()         { *m = MsgUpdateBitcannaidResponse{} }
func (m *MsgUpdateBitcannaidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateBitcannaidResponse) ProtoMessage()    {}
func (*MsgUpdateBitcannaidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d2bc8ae8552a21, []int{3}
}
func (m *MsgUpdateBitcannaidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateBitcannaidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateBitcannaidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateBitcannaidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateBitcannaidResponse.Merge(m, src)
}
func (m *MsgUpdateBitcannaidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateBitcannaidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateBitcannaidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateBitcannaidResponse proto.InternalMessageInfo

type MsgDeleteBitcannaid struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeleteBitcannaid) Reset()         { *m = MsgDeleteBitcannaid{} }
func (m *MsgDeleteBitcannaid) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteBitcannaid) ProtoMessage()    {}
func (*MsgDeleteBitcannaid) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d2bc8ae8552a21, []int{4}
}
func (m *MsgDeleteBitcannaid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteBitcannaid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteBitcannaid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteBitcannaid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteBitcannaid.Merge(m, src)
}
func (m *MsgDeleteBitcannaid) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteBitcannaid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteBitcannaid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteBitcannaid proto.InternalMessageInfo

func (m *MsgDeleteBitcannaid) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteBitcannaid) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgDeleteBitcannaidResponse struct {
}

func (m *MsgDeleteBitcannaidResponse) Reset()         { *m = MsgDeleteBitcannaidResponse{} }
func (m *MsgDeleteBitcannaidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteBitcannaidResponse) ProtoMessage()    {}
func (*MsgDeleteBitcannaidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d2bc8ae8552a21, []int{5}
}
func (m *MsgDeleteBitcannaidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteBitcannaidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteBitcannaidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteBitcannaidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteBitcannaidResponse.Merge(m, src)
}
func (m *MsgDeleteBitcannaidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteBitcannaidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteBitcannaidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteBitcannaidResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateBitcannaid)(nil), "BitCannaGlobal.bcna.bcna.MsgCreateBitcannaid")
	proto.RegisterType((*MsgCreateBitcannaidResponse)(nil), "BitCannaGlobal.bcna.bcna.MsgCreateBitcannaidResponse")
	proto.RegisterType((*MsgUpdateBitcannaid)(nil), "BitCannaGlobal.bcna.bcna.MsgUpdateBitcannaid")
	proto.RegisterType((*MsgUpdateBitcannaidResponse)(nil), "BitCannaGlobal.bcna.bcna.MsgUpdateBitcannaidResponse")
	proto.RegisterType((*MsgDeleteBitcannaid)(nil), "BitCannaGlobal.bcna.bcna.MsgDeleteBitcannaid")
	proto.RegisterType((*MsgDeleteBitcannaidResponse)(nil), "BitCannaGlobal.bcna.bcna.MsgDeleteBitcannaidResponse")
}

func init() { proto.RegisterFile("bcna/tx.proto", fileDescriptor_12d2bc8ae8552a21) }

var fileDescriptor_12d2bc8ae8552a21 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0xce, 0x4b,
	0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x70, 0xca, 0x2c, 0x71, 0x4e,
	0xcc, 0xcb, 0x4b, 0x74, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0xd1, 0x03, 0xc9, 0x82, 0x09, 0xa5, 0x44,
	0x2e, 0x61, 0xdf, 0xe2, 0x74, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0xa7, 0xcc, 0x92, 0x64, 0x90,
	0xa2, 0xcc, 0x14, 0x21, 0x09, 0x2e, 0xf6, 0x64, 0x90, 0x58, 0x7e, 0x91, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc6, 0xc5, 0x06, 0xd2, 0x98, 0x99, 0x22, 0xc1, 0x04, 0x96,
	0x80, 0xf2, 0x40, 0x3a, 0x12, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0x21, 0x3a, 0xa0,
	0x5c, 0x25, 0x5d, 0x2e, 0x69, 0x2c, 0x56, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x80, 0x6d, 0x61, 0x09, 0x62, 0xca, 0x4c, 0x51, 0x2a, 0x04, 0xbb,
	0x28, 0xb4, 0x20, 0x85, 0x58, 0x17, 0x41, 0x0c, 0x60, 0x82, 0x19, 0x80, 0xe4, 0x42, 0x66, 0x5c,
	0x2e, 0x64, 0x41, 0x75, 0xa1, 0x2c, 0xd8, 0x85, 0xe8, 0x56, 0xc2, 0x5c, 0xa8, 0x64, 0x0f, 0x76,
	0x91, 0x4b, 0x6a, 0x4e, 0x2a, 0x79, 0x2e, 0x82, 0x9a, 0x8f, 0x6e, 0x00, 0xcc, 0x7c, 0xa3, 0xcf,
	0x4c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x15, 0x5c, 0x02, 0x18, 0x11, 0xa1, 0xab, 0x87, 0x2b,
	0xea, 0xf4, 0xb0, 0x04, 0xaa, 0x94, 0x29, 0x49, 0xca, 0xe1, 0x71, 0x50, 0xc1, 0x25, 0x80, 0x11,
	0xe0, 0xf8, 0x6d, 0x46, 0x57, 0x4e, 0xc0, 0x66, 0x5c, 0x61, 0x0b, 0xb2, 0x19, 0x23, 0x60, 0xf1,
	0xdb, 0x8c, 0xae, 0x9c, 0x80, 0xcd, 0xb8, 0x42, 0xdd, 0xc9, 0xf5, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0x51, 0x8d, 0xd6, 0x07, 0x67, 0xab, 0x0a, 0x08, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0xce, 0x61, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x3f, 0x35, 0x65, 0x72, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateBitcannaid(ctx context.Context, in *MsgCreateBitcannaid, opts ...grpc.CallOption) (*MsgCreateBitcannaidResponse, error)
	UpdateBitcannaid(ctx context.Context, in *MsgUpdateBitcannaid, opts ...grpc.CallOption) (*MsgUpdateBitcannaidResponse, error)
	DeleteBitcannaid(ctx context.Context, in *MsgDeleteBitcannaid, opts ...grpc.CallOption) (*MsgDeleteBitcannaidResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBitcannaid(ctx context.Context, in *MsgCreateBitcannaid, opts ...grpc.CallOption) (*MsgCreateBitcannaidResponse, error) {
	out := new(MsgCreateBitcannaidResponse)
	err := c.cc.Invoke(ctx, "/BitCannaGlobal.bcna.bcna.Msg/CreateBitcannaid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateBitcannaid(ctx context.Context, in *MsgUpdateBitcannaid, opts ...grpc.CallOption) (*MsgUpdateBitcannaidResponse, error) {
	out := new(MsgUpdateBitcannaidResponse)
	err := c.cc.Invoke(ctx, "/BitCannaGlobal.bcna.bcna.Msg/UpdateBitcannaid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteBitcannaid(ctx context.Context, in *MsgDeleteBitcannaid, opts ...grpc.CallOption) (*MsgDeleteBitcannaidResponse, error) {
	out := new(MsgDeleteBitcannaidResponse)
	err := c.cc.Invoke(ctx, "/BitCannaGlobal.bcna.bcna.Msg/DeleteBitcannaid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateBitcannaid(context.Context, *MsgCreateBitcannaid) (*MsgCreateBitcannaidResponse, error)
	UpdateBitcannaid(context.Context, *MsgUpdateBitcannaid) (*MsgUpdateBitcannaidResponse, error)
	DeleteBitcannaid(context.Context, *MsgDeleteBitcannaid) (*MsgDeleteBitcannaidResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBitcannaid(ctx context.Context, req *MsgCreateBitcannaid) (*MsgCreateBitcannaidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBitcannaid not implemented")
}
func (*UnimplementedMsgServer) UpdateBitcannaid(ctx context.Context, req *MsgUpdateBitcannaid) (*MsgUpdateBitcannaidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBitcannaid not implemented")
}
func (*UnimplementedMsgServer) DeleteBitcannaid(ctx context.Context, req *MsgDeleteBitcannaid) (*MsgDeleteBitcannaidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBitcannaid not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBitcannaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBitcannaid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBitcannaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BitCannaGlobal.bcna.bcna.Msg/CreateBitcannaid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBitcannaid(ctx, req.(*MsgCreateBitcannaid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateBitcannaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateBitcannaid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateBitcannaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BitCannaGlobal.bcna.bcna.Msg/UpdateBitcannaid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateBitcannaid(ctx, req.(*MsgUpdateBitcannaid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteBitcannaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteBitcannaid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteBitcannaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BitCannaGlobal.bcna.bcna.Msg/DeleteBitcannaid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteBitcannaid(ctx, req.(*MsgDeleteBitcannaid))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BitCannaGlobal.bcna.bcna.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBitcannaid",
			Handler:    _Msg_CreateBitcannaid_Handler,
		},
		{
			MethodName: "UpdateBitcannaid",
			Handler:    _Msg_UpdateBitcannaid_Handler,
		},
		{
			MethodName: "DeleteBitcannaid",
			Handler:    _Msg_DeleteBitcannaid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bcna/tx.proto",
}

func (m *MsgCreateBitcannaid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBitcannaid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBitcannaid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Bcnaid) > 0 {
		i -= len(m.Bcnaid)
		copy(dAtA[i:], m.Bcnaid)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Bcnaid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateBitcannaidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBitcannaidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBitcannaidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateBitcannaid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateBitcannaid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateBitcannaid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bcnaid) > 0 {
		i -= len(m.Bcnaid)
		copy(dAtA[i:], m.Bcnaid)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Bcnaid)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateBitcannaidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateBitcannaidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateBitcannaidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteBitcannaid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteBitcannaid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteBitcannaid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteBitcannaidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteBitcannaidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteBitcannaidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateBitcannaid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Bcnaid)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateBitcannaidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgUpdateBitcannaid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Bcnaid)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateBitcannaidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteBitcannaid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgDeleteBitcannaidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateBitcannaid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateBitcannaid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBitcannaid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bcnaid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bcnaid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateBitcannaidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateBitcannaidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBitcannaidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateBitcannaid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateBitcannaid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateBitcannaid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bcnaid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bcnaid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateBitcannaidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateBitcannaidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateBitcannaidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteBitcannaid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteBitcannaid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteBitcannaid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteBitcannaidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteBitcannaidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteBitcannaidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

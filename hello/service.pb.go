// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Data struct {
	Chunk                string   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetChunk() string {
	if m != nil {
		return m.Chunk
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "hello.Data")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0x48, 0xcd, 0xc9, 0xc9,
	0x97, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x26, 0x95, 0xa6, 0xe9, 0xa7,
	0xe6, 0x16, 0x94, 0x54, 0x42, 0xd4, 0x28, 0xc9, 0x70, 0xb1, 0xb8, 0x24, 0x96, 0x24, 0x0a, 0x89,
	0x70, 0xb1, 0x26, 0x67, 0x94, 0xe6, 0x65, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38,
	0x46, 0x56, 0x5c, 0x9c, 0xce, 0xf9, 0x79, 0x79, 0xa9, 0xc9, 0x25, 0xf9, 0x45, 0x42, 0xba, 0x5c,
	0x4c, 0x1e, 0x99, 0x42, 0xdc, 0x7a, 0x60, 0x53, 0xf5, 0x40, 0xba, 0xa4, 0xc4, 0xf4, 0x20, 0x66,
	0xeb, 0xc1, 0xcc, 0xd6, 0x73, 0x05, 0x99, 0xad, 0xc4, 0xa0, 0xc1, 0xe8, 0xc4, 0x1e, 0x05, 0xb1,
	0x3f, 0x89, 0x0d, 0x2c, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xd1, 0xa5, 0xc1, 0x9e,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConnectorClient is the client API for Connector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectorClient interface {
	Hi(ctx context.Context, opts ...grpc.CallOption) (Connector_HiClient, error)
}

type connectorClient struct {
	cc *grpc.ClientConn
}

func NewConnectorClient(cc *grpc.ClientConn) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) Hi(ctx context.Context, opts ...grpc.CallOption) (Connector_HiClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Connector_serviceDesc.Streams[0], "/hello.Connector/Hi", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectorHiClient{stream}
	return x, nil
}

type Connector_HiClient interface {
	Send(*Data) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type connectorHiClient struct {
	grpc.ClientStream
}

func (x *connectorHiClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectorHiClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectorServer is the server API for Connector service.
type ConnectorServer interface {
	Hi(Connector_HiServer) error
}

// UnimplementedConnectorServer can be embedded to have forward compatible implementations.
type UnimplementedConnectorServer struct {
}

func (*UnimplementedConnectorServer) Hi(srv Connector_HiServer) error {
	return status.Errorf(codes.Unimplemented, "method Hi not implemented")
}

func RegisterConnectorServer(s *grpc.Server, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_Hi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectorServer).Hi(&connectorHiServer{stream})
}

type Connector_HiServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type connectorHiServer struct {
	grpc.ServerStream
}

func (x *connectorHiServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectorHiServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Hi",
			Handler:       _Connector_Hi_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

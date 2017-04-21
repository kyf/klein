// Code generated by protoc-gen-go.
// source: connector.proto
// DO NOT EDIT!

/*
Package connector is a generated protocol buffer package.

It is generated from these files:
	connector.proto

It has these top-level messages:
	MessageRequest
	Reply
*/
package connector

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

type MessageRequest struct {
	MsgType    int32  `protobuf:"varint,1,opt,name=MsgType,json=msgType" json:"MsgType,omitempty"`
	SequenceId string `protobuf:"bytes,2,opt,name=SequenceId,json=sequenceId" json:"SequenceId,omitempty"`
	ConnId     string `protobuf:"bytes,3,opt,name=ConnId,json=connId" json:"ConnId,omitempty"`
	Sender     string `protobuf:"bytes,4,opt,name=Sender,json=sender" json:"Sender,omitempty"`
	Receiver   string `protobuf:"bytes,5,opt,name=Receiver,json=receiver" json:"Receiver,omitempty"`
	Body       string `protobuf:"bytes,6,opt,name=Body,json=body" json:"Body,omitempty"`
}

func (m *MessageRequest) Reset()                    { *m = MessageRequest{} }
func (m *MessageRequest) String() string            { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()               {}
func (*MessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Reply struct {
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*MessageRequest)(nil), "connector.MessageRequest")
	proto.RegisterType((*Reply)(nil), "connector.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Connector service

type ConnectorClient interface {
	ReceiveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*Reply, error)
}

type connectorClient struct {
	cc *grpc.ClientConn
}

func NewConnectorClient(cc *grpc.ClientConn) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) ReceiveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/connector.Connector/ReceiveMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Connector service

type ConnectorServer interface {
	ReceiveMessage(context.Context, *MessageRequest) (*Reply, error)
}

func RegisterConnectorServer(s *grpc.Server, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_ReceiveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).ReceiveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connector.Connector/ReceiveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).ReceiveMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connector.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveMessage",
			Handler:    _Connector_ReceiveMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("connector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0x85, 0x30,
	0x14, 0x86, 0xad, 0x42, 0xb9, 0x9c, 0xe1, 0x6a, 0xce, 0x60, 0xea, 0x1d, 0xcc, 0x0d, 0x13, 0x13,
	0x83, 0x3e, 0x80, 0x89, 0x4c, 0x24, 0xb2, 0x14, 0x5f, 0x40, 0xda, 0x13, 0x62, 0xa2, 0x2d, 0xb6,
	0x68, 0xd2, 0x57, 0xf2, 0x29, 0x0d, 0x05, 0xd1, 0x3b, 0x7e, 0xdf, 0xdf, 0x9c, 0xe4, 0x2b, 0x5c,
	0x2a, 0x6b, 0x0c, 0xa9, 0xc9, 0xba, 0x6a, 0x74, 0x76, 0xb2, 0x98, 0x6f, 0xa2, 0xf8, 0x66, 0xb0,
	0x6f, 0xc9, 0xfb, 0x97, 0x81, 0x24, 0x7d, 0x7c, 0x92, 0x9f, 0x50, 0x40, 0xd6, 0xfa, 0xe1, 0x39,
	0x8c, 0x24, 0xd8, 0x91, 0x95, 0xa9, 0xcc, 0xde, 0x17, 0xc4, 0x5b, 0x80, 0x6e, 0x7e, 0x64, 0x14,
	0x35, 0x5a, 0x9c, 0x1f, 0x59, 0x99, 0x4b, 0xf0, 0x9b, 0xc1, 0x6b, 0xe0, 0xb5, 0x35, 0xa6, 0xd1,
	0xe2, 0x22, 0x6e, 0x5c, 0x45, 0x9a, 0x7d, 0x47, 0x46, 0x93, 0x13, 0xc9, 0xe2, 0x7d, 0x24, 0x3c,
	0xc0, 0x4e, 0x92, 0xa2, 0xd7, 0x2f, 0x72, 0x22, 0x8d, 0xcb, 0xce, 0xad, 0x8c, 0x08, 0xc9, 0xa3,
	0xd5, 0x41, 0xf0, 0xe8, 0x93, 0xde, 0xea, 0x50, 0x64, 0x90, 0x4a, 0x1a, 0xdf, 0xc2, 0xdd, 0x13,
	0xe4, 0xf5, 0x6f, 0x02, 0x3e, 0xc0, 0x7e, 0xbd, 0xb2, 0x86, 0xe0, 0x4d, 0xf5, 0x57, 0x7c, 0x1a,
	0x77, 0xb8, 0xfa, 0x37, 0xc5, 0x5b, 0xc5, 0x59, 0xcf, 0xe3, 0xaf, 0xdc, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x11, 0x65, 0x77, 0xce, 0x28, 0x01, 0x00, 0x00,
}
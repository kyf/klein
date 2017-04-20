// Code generated by protoc-gen-go.
// source: connector.proto
// DO NOT EDIT!

/*
Package connector is a generated protocol buffer package.

It is generated from these files:
	connector.proto

It has these top-level messages:
	MessageRequest
	ReceiveReply
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

type ReceiveReply struct {
	Status  bool   `protobuf:"varint,1,opt,name=Status,json=status" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,json=message" json:"Message,omitempty"`
}

func (m *ReceiveReply) Reset()                    { *m = ReceiveReply{} }
func (m *ReceiveReply) String() string            { return proto.CompactTextString(m) }
func (*ReceiveReply) ProtoMessage()               {}
func (*ReceiveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*MessageRequest)(nil), "connector.MessageRequest")
	proto.RegisterType((*ReceiveReply)(nil), "connector.ReceiveReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Connector service

type ConnectorClient interface {
	ReceiveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*ReceiveReply, error)
}

type connectorClient struct {
	cc *grpc.ClientConn
}

func NewConnectorClient(cc *grpc.ClientConn) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) ReceiveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*ReceiveReply, error) {
	out := new(ReceiveReply)
	err := grpc.Invoke(ctx, "/connector.Connector/ReceiveMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Connector service

type ConnectorServer interface {
	ReceiveMessage(context.Context, *MessageRequest) (*ReceiveReply, error)
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
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x5d, 0xdd, 0x66, 0x77, 0x07, 0xa9, 0x30, 0x07, 0x8d, 0x3d, 0x48, 0xd9, 0x53, 0x4f,
	0x3d, 0xe8, 0x1f, 0x10, 0x0b, 0x42, 0x0f, 0xbd, 0x64, 0xfd, 0x03, 0x6d, 0x32, 0x14, 0xc1, 0x26,
	0x6b, 0x92, 0x0a, 0xfb, 0x97, 0xfc, 0x95, 0xd2, 0x49, 0xac, 0xeb, 0xf1, 0x7b, 0xf3, 0x18, 0xde,
	0x7b, 0x70, 0xa3, 0x9d, 0xb5, 0xa4, 0xa3, 0xf3, 0xcb, 0xde, 0xbb, 0xe8, 0xb0, 0x39, 0x0b, 0xed,
	0x77, 0x01, 0xd3, 0x0d, 0x85, 0xb0, 0xdd, 0x93, 0xa2, 0xcf, 0x23, 0x85, 0x88, 0x12, 0xaa, 0x4d,
	0xd8, 0xbf, 0x0d, 0x3d, 0xc9, 0x62, 0x5e, 0x2c, 0x26, 0xaa, 0x3a, 0x24, 0xc4, 0x07, 0x80, 0xee,
	0x64, 0xb2, 0x9a, 0xd6, 0x46, 0x5e, 0xce, 0x8b, 0x45, 0xa3, 0x20, 0x9c, 0x15, 0xbc, 0x05, 0xb1,
	0x72, 0xd6, 0xae, 0x8d, 0xbc, 0xe2, 0x9b, 0xd0, 0x4c, 0x27, 0xbd, 0x23, 0x6b, 0xc8, 0xcb, 0x32,
	0xe9, 0x81, 0x09, 0x67, 0x50, 0x2b, 0xd2, 0xf4, 0xfe, 0x45, 0x5e, 0x4e, 0xf8, 0x52, 0xfb, 0xcc,
	0x88, 0x50, 0xbe, 0x38, 0x33, 0x48, 0xc1, 0x7a, 0xb9, 0x73, 0x66, 0x68, 0x9f, 0xe1, 0x3a, 0xfb,
	0x15, 0xf5, 0x1f, 0x03, 0xff, 0x8d, 0xdb, 0x78, 0x0c, 0x1c, 0xb4, 0x56, 0x22, 0x30, 0x71, 0x83,
	0xd4, 0x29, 0x87, 0xac, 0x0e, 0x09, 0x1f, 0x3b, 0x68, 0x56, 0xbf, 0xdd, 0xf1, 0x15, 0xa6, 0xf9,
	0x5d, 0x76, 0xe3, 0xfd, 0xf2, 0x6f, 0xaa, 0xff, 0xab, 0xcc, 0xee, 0x46, 0xa7, 0x71, 0x88, 0xf6,
	0x62, 0x27, 0x78, 0xd5, 0xa7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x83, 0x23, 0x8f, 0x68,
	0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Client struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{0}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Client) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Client) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ListClient struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListClient) Reset()         { *m = ListClient{} }
func (m *ListClient) String() string { return proto.CompactTextString(m) }
func (*ListClient) ProtoMessage()    {}
func (*ListClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{1}
}

func (m *ListClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClient.Unmarshal(m, b)
}
func (m *ListClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClient.Marshal(b, m, deterministic)
}
func (m *ListClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClient.Merge(m, src)
}
func (m *ListClient) XXX_Size() int {
	return xxx_messageInfo_ListClient.Size(m)
}
func (m *ListClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClient.DiscardUnknown(m)
}

var xxx_messageInfo_ListClient proto.InternalMessageInfo

type ClientResponse struct {
	Clients              []*Client `protobuf:"bytes,1,rep,name=Clients,proto3" json:"Clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ClientResponse) Reset()         { *m = ClientResponse{} }
func (m *ClientResponse) String() string { return proto.CompactTextString(m) }
func (*ClientResponse) ProtoMessage()    {}
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{2}
}

func (m *ClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientResponse.Unmarshal(m, b)
}
func (m *ClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientResponse.Marshal(b, m, deterministic)
}
func (m *ClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientResponse.Merge(m, src)
}
func (m *ClientResponse) XXX_Size() int {
	return xxx_messageInfo_ClientResponse.Size(m)
}
func (m *ClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientResponse proto.InternalMessageInfo

func (m *ClientResponse) GetClients() []*Client {
	if m != nil {
		return m.Clients
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "pb.Client")
	proto.RegisterType((*ListClient)(nil), "pb.ListClient")
	proto.RegisterType((*ClientResponse)(nil), "pb.ClientResponse")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_014de31d7ac8c57c) }

var fileDescriptor_014de31d7ac8c57c = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd7, 0x56, 0x3b, 0x7c, 0x9b, 0x43, 0xde, 0x29, 0x6c, 0x07, 0x47, 0xf0, 0xb0, 0x53,
	0xc0, 0x09, 0xde, 0x37, 0x85, 0x21, 0x78, 0x18, 0x93, 0x7d, 0x80, 0xb6, 0x79, 0xcc, 0x40, 0xd6,
	0x84, 0x24, 0x13, 0x3f, 0x9d, 0x9f, 0x4d, 0x9a, 0x76, 0x58, 0x10, 0xbd, 0xfd, 0xf3, 0xfb, 0x3d,
	0x92, 0xfc, 0x79, 0x30, 0xae, 0xb4, 0xa2, 0x3a, 0x08, 0xeb, 0x4c, 0x30, 0x98, 0xda, 0x72, 0x3a,
	0x3e, 0x68, 0x53, 0x16, 0xba, 0x25, 0xfc, 0x1d, 0xf2, 0xa7, 0x38, 0x81, 0x13, 0x48, 0x95, 0x64,
	0xc9, 0x3c, 0x59, 0x5c, 0xed, 0x52, 0x25, 0x71, 0x0e, 0x23, 0x49, 0xbe, 0x72, 0xca, 0x06, 0x65,
	0x6a, 0x96, 0x46, 0xd1, 0x47, 0x88, 0x70, 0x51, 0x17, 0x47, 0x62, 0x59, 0x54, 0x31, 0x23, 0x83,
	0x61, 0x21, 0xa5, 0x23, 0xef, 0xd9, 0x65, 0xc4, 0xe7, 0x23, 0x1f, 0x03, 0xbc, 0x2a, 0x1f, 0xda,
	0xd7, 0xf8, 0x23, 0x4c, 0xda, 0xb4, 0x23, 0x6f, 0x4d, 0xed, 0x09, 0xef, 0x60, 0xd8, 0x12, 0xcf,
	0x92, 0x79, 0xb6, 0x18, 0x2d, 0x41, 0xd8, 0x52, 0x74, 0x43, 0x67, 0xb5, 0xfc, 0x4a, 0xe0, 0xba,
	0xcd, 0x6f, 0xe4, 0x3e, 0x54, 0x45, 0x78, 0x0b, 0xd9, 0x4a, 0x4a, 0xec, 0x4d, 0x4f, 0x7b, 0x99,
	0x0f, 0x70, 0x06, 0xd9, 0x86, 0x02, 0xe6, 0x0d, 0x7c, 0x91, 0xbf, 0xe5, 0x33, 0xe9, 0x3f, 0x24,
	0x87, 0x7c, 0x6f, 0x65, 0x11, 0xe8, 0x9f, 0xdb, 0x05, 0xe4, 0x1b, 0x0a, 0x2b, 0xad, 0x71, 0xd2,
	0xf0, 0x9f, 0x8a, 0x53, 0xec, 0xfd, 0xbf, 0x2b, 0xc9, 0x07, 0xeb, 0x7b, 0x98, 0x29, 0x23, 0x0e,
	0xce, 0x56, 0x82, 0x3e, 0x8b, 0xa3, 0xd5, 0xe4, 0x85, 0x33, 0xa7, 0x40, 0x87, 0x93, 0x92, 0xb4,
	0xbe, 0xd9, 0x7b, 0x72, 0x5d, 0xb5, 0x6d, 0xb3, 0xa1, 0x6d, 0x52, 0xe6, 0x71, 0x55, 0x0f, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0xcc, 0x10, 0x7e, 0xcc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	Add(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Client, error)
	Del(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Client, error)
	Update(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error)
	GetAll(ctx context.Context, in *ListClient, opts ...grpc.CallOption) (*ClientResponse, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Add(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/pb.ClientService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/pb.ClientService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Del(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/pb.ClientService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Update(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/pb.ClientService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetAll(ctx context.Context, in *ListClient, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	Add(context.Context, *Client) (*Client, error)
	Get(context.Context, *Id) (*Client, error)
	Del(context.Context, *Id) (*Client, error)
	Update(context.Context, *Client) (*Client, error)
	GetAll(context.Context, *ListClient) (*ClientResponse, error)
}

// UnimplementedClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (*UnimplementedClientServiceServer) Add(ctx context.Context, req *Client) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedClientServiceServer) Get(ctx context.Context, req *Id) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedClientServiceServer) Del(ctx context.Context, req *Id) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (*UnimplementedClientServiceServer) Update(ctx context.Context, req *Client) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedClientServiceServer) GetAll(ctx context.Context, req *ListClient) (*ClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Add(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Del(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Update(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetAll(ctx, req.(*ListClient))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ClientService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ClientService_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _ClientService_Del_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClientService_Update_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ClientService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}

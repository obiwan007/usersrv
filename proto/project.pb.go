// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

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

type Project struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Client               string   `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	Tags                 string   `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Team                 string   `protobuf:"bytes,11,opt,name=team,proto3" json:"team,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *Project) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *Project) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Project) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListProject struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProject) Reset()         { *m = ListProject{} }
func (m *ListProject) String() string { return proto.CompactTextString(m) }
func (*ListProject) ProtoMessage()    {}
func (*ListProject) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{1}
}

func (m *ListProject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProject.Unmarshal(m, b)
}
func (m *ListProject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProject.Marshal(b, m, deterministic)
}
func (m *ListProject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProject.Merge(m, src)
}
func (m *ListProject) XXX_Size() int {
	return xxx_messageInfo_ListProject.Size(m)
}
func (m *ListProject) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProject.DiscardUnknown(m)
}

var xxx_messageInfo_ListProject proto.InternalMessageInfo

type ProjectResponse struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=Projects,proto3" json:"Projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProjectResponse) Reset()         { *m = ProjectResponse{} }
func (m *ProjectResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectResponse) ProtoMessage()    {}
func (*ProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{2}
}

func (m *ProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectResponse.Unmarshal(m, b)
}
func (m *ProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectResponse.Marshal(b, m, deterministic)
}
func (m *ProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectResponse.Merge(m, src)
}
func (m *ProjectResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectResponse.Size(m)
}
func (m *ProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectResponse proto.InternalMessageInfo

func (m *ProjectResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*Project)(nil), "pb.Project")
	proto.RegisterType((*ListProject)(nil), "pb.ListProject")
	proto.RegisterType((*ProjectResponse)(nil), "pb.ProjectResponse")
}

func init() { proto.RegisterFile("project.proto", fileDescriptor_8340e6318dfdfac2) }

var fileDescriptor_8340e6318dfdfac2 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x18, 0x85, 0xd7, 0x56, 0x3b, 0xfd, 0xeb, 0x36, 0x89, 0x20, 0x61, 0x7a, 0x31, 0x8b, 0xe0, 0xae,
	0x8a, 0xce, 0x3b, 0xef, 0x36, 0x84, 0x21, 0x78, 0x31, 0x26, 0x7b, 0x80, 0xb4, 0xf9, 0x29, 0x91,
	0xb4, 0x09, 0x49, 0x26, 0xbe, 0x91, 0x8f, 0xe2, 0x6b, 0x49, 0xdb, 0x28, 0xbd, 0x50, 0xef, 0xce,
	0xf9, 0xce, 0x29, 0xcd, 0xe1, 0x87, 0x91, 0x36, 0xea, 0x15, 0x0b, 0x97, 0x69, 0xa3, 0x9c, 0x22,
	0xa1, 0xce, 0xa7, 0x27, 0xa5, 0x54, 0x39, 0x93, 0x1d, 0x49, 0x3f, 0x02, 0x18, 0x6e, 0xba, 0x0e,
	0x19, 0x43, 0x28, 0x38, 0x0d, 0x66, 0xc1, 0xfc, 0x78, 0x1b, 0x0a, 0x4e, 0x66, 0x90, 0x70, 0xb4,
	0x85, 0x11, 0xda, 0x09, 0x55, 0xd3, 0xb0, 0x0d, 0xfa, 0x88, 0x9c, 0x43, 0x5c, 0x48, 0x81, 0xb5,
	0xa3, 0x51, 0x1b, 0x7a, 0x47, 0x08, 0x1c, 0x38, 0x56, 0x5a, 0x7a, 0xd8, 0xd2, 0x56, 0x37, 0x5d,
	0xeb, 0x98, 0xdb, 0x5b, 0x1a, 0x77, 0xdd, 0xce, 0xb5, 0x5d, 0x64, 0x15, 0x4d, 0x7c, 0x17, 0x59,
	0xd5, 0xb0, 0x9a, 0x55, 0x48, 0x87, 0x1d, 0x6b, 0x74, 0x3a, 0x82, 0xe4, 0x59, 0x58, 0xe7, 0x1f,
	0x9b, 0x3e, 0xc0, 0xc4, 0xcb, 0x2d, 0x5a, 0xad, 0x6a, 0x8b, 0xe4, 0x06, 0x8e, 0x3c, 0xb2, 0x34,
	0x98, 0x45, 0xf3, 0x64, 0x91, 0x64, 0x3a, 0xcf, 0xbe, 0x6b, 0x3f, 0xe1, 0xe2, 0x33, 0x80, 0xb1,
	0x37, 0x2f, 0x68, 0xde, 0x44, 0x81, 0xe4, 0x0a, 0xa2, 0x25, 0xe7, 0xa4, 0xff, 0xc1, 0xb4, 0x6f,
	0xd2, 0x01, 0xb9, 0x84, 0x68, 0x8d, 0x8e, 0xc4, 0x0d, 0x7d, 0xe2, 0xbf, 0xa4, 0x8f, 0x28, 0xff,
	0x4a, 0xaf, 0x21, 0xde, 0x69, 0xce, 0x1c, 0xfe, 0xfb, 0x87, 0x5b, 0x88, 0xd7, 0xe8, 0x96, 0x52,
	0x92, 0x49, 0x13, 0xf4, 0xe6, 0x4e, 0xcf, 0xfa, 0x4b, 0xfc, 0xe0, 0x74, 0xb0, 0xba, 0x83, 0x0b,
	0xa1, 0xb2, 0xd2, 0xe8, 0x22, 0xc3, 0x77, 0x56, 0x69, 0x89, 0x36, 0x33, 0x6a, 0xef, 0xb0, 0xdc,
	0x0b, 0x8e, 0xab, 0xd3, 0x9d, 0x45, 0xe3, 0x27, 0x6e, 0x9a, 0x7b, 0x6f, 0x82, 0x3c, 0x6e, 0x0f,
	0x7f, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xff, 0x41, 0x05, 0x1b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectServiceClient interface {
	Add(ctx context.Context, in *Project, opts ...grpc.CallOption) (*Project, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Project, error)
	Del(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Project, error)
	Update(ctx context.Context, in *Project, opts ...grpc.CallOption) (*Project, error)
	GetAll(ctx context.Context, in *ListProject, opts ...grpc.CallOption) (*ProjectResponse, error)
}

type projectServiceClient struct {
	cc *grpc.ClientConn
}

func NewProjectServiceClient(cc *grpc.ClientConn) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Add(ctx context.Context, in *Project, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/pb.ProjectService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/pb.ProjectService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Del(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/pb.ProjectService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Update(ctx context.Context, in *Project, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/pb.ProjectService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetAll(ctx context.Context, in *ListProject, opts ...grpc.CallOption) (*ProjectResponse, error) {
	out := new(ProjectResponse)
	err := c.cc.Invoke(ctx, "/pb.ProjectService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
type ProjectServiceServer interface {
	Add(context.Context, *Project) (*Project, error)
	Get(context.Context, *Id) (*Project, error)
	Del(context.Context, *Id) (*Project, error)
	Update(context.Context, *Project) (*Project, error)
	GetAll(context.Context, *ListProject) (*ProjectResponse, error)
}

// UnimplementedProjectServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (*UnimplementedProjectServiceServer) Add(ctx context.Context, req *Project) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedProjectServiceServer) Get(ctx context.Context, req *Id) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedProjectServiceServer) Del(ctx context.Context, req *Id) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (*UnimplementedProjectServiceServer) Update(ctx context.Context, req *Project) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedProjectServiceServer) GetAll(ctx context.Context, req *ListProject) (*ProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterProjectServiceServer(s *grpc.Server, srv ProjectServiceServer) {
	s.RegisterService(&_ProjectService_serviceDesc, srv)
}

func _ProjectService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Project)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProjectService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Add(ctx, req.(*Project))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProjectService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProjectService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Del(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Project)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProjectService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Update(ctx, req.(*Project))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProjectService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetAll(ctx, req.(*ListProject))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ProjectService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProjectService_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _ProjectService_Del_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProjectService_Update_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ProjectService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project.proto",
}
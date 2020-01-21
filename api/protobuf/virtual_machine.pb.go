// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtual_machine.proto

package protobuf

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

type MachineType int32

const (
	MachineType_STANDARD MachineType = 0
	MachineType_HIGHMEM  MachineType = 1
	MachineType_HIGHCPU  MachineType = 2
)

var MachineType_name = map[int32]string{
	0: "STANDARD",
	1: "HIGHMEM",
	2: "HIGHCPU",
}

var MachineType_value = map[string]int32{
	"STANDARD": 0,
	"HIGHMEM":  1,
	"HIGHCPU":  2,
}

func (x MachineType) String() string {
	return proto.EnumName(MachineType_name, int32(x))
}

func (MachineType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0312cad0ca872dab, []int{0}
}

type VirtualMachine struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MachineType          MachineType `protobuf:"varint,3,opt,name=machine_type,json=machineType,proto3,enum=MachineType" json:"machine_type,omitempty"`
	Ip                   string      `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VirtualMachine) Reset()         { *m = VirtualMachine{} }
func (m *VirtualMachine) String() string { return proto.CompactTextString(m) }
func (*VirtualMachine) ProtoMessage()    {}
func (*VirtualMachine) Descriptor() ([]byte, []int) {
	return fileDescriptor_0312cad0ca872dab, []int{0}
}

func (m *VirtualMachine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachine.Unmarshal(m, b)
}
func (m *VirtualMachine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachine.Marshal(b, m, deterministic)
}
func (m *VirtualMachine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachine.Merge(m, src)
}
func (m *VirtualMachine) XXX_Size() int {
	return xxx_messageInfo_VirtualMachine.Size(m)
}
func (m *VirtualMachine) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachine.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachine proto.InternalMessageInfo

func (m *VirtualMachine) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualMachine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMachine) GetMachineType() MachineType {
	if m != nil {
		return m.MachineType
	}
	return MachineType_STANDARD
}

func (m *VirtualMachine) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func init() {
	proto.RegisterEnum("MachineType", MachineType_name, MachineType_value)
	proto.RegisterType((*VirtualMachine)(nil), "VirtualMachine")
}

func init() { proto.RegisterFile("virtual_machine.proto", fileDescriptor_0312cad0ca872dab) }

var fileDescriptor_0312cad0ca872dab = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xdf, 0xf6, 0x1d, 0x53, 0xff, 0x2b, 0xb5, 0x04, 0x06, 0xc5, 0xd3, 0xf0, 0x34, 0x06,
	0x6b, 0x61, 0xe2, 0x55, 0xdc, 0x56, 0x75, 0x1e, 0x3a, 0xa4, 0x9b, 0x1e, 0xbc, 0x8c, 0xb4, 0xfd,
	0xcf, 0x85, 0xad, 0x49, 0xc8, 0x92, 0x41, 0x3f, 0x8d, 0x5f, 0x55, 0xd6, 0x55, 0x74, 0x5e, 0xf4,
	0xf6, 0x3c, 0x4f, 0x7e, 0xfc, 0x48, 0x08, 0xb4, 0x77, 0x4c, 0x69, 0x43, 0x37, 0x8b, 0x82, 0x66,
	0x2b, 0xc6, 0x31, 0x90, 0x4a, 0x68, 0x71, 0x69, 0xc0, 0x7d, 0x39, 0x1c, 0xc4, 0x87, 0x9d, 0xb8,
	0x60, 0xb3, 0xdc, 0xb7, 0x3a, 0x56, 0xf7, 0x2c, 0xb1, 0x59, 0x4e, 0x08, 0x34, 0x38, 0x2d, 0xd0,
	0xb7, 0xab, 0xa5, 0xca, 0x24, 0x04, 0xa7, 0xd6, 0x2c, 0x74, 0x29, 0xd1, 0xff, 0xdf, 0xb1, 0xba,
	0xee, 0xc0, 0x09, 0x6a, 0xc7, 0xbc, 0x94, 0x98, 0xb4, 0x8a, 0xaf, 0x52, 0x49, 0xa5, 0xdf, 0xa8,
	0xa5, 0xb2, 0x77, 0x0d, 0xad, 0x6f, 0x2c, 0x71, 0xe0, 0x74, 0x36, 0x1f, 0x4e, 0xa3, 0x61, 0x12,
	0x79, 0xff, 0x48, 0x0b, 0x4e, 0x26, 0x8f, 0x0f, 0x93, 0xf8, 0x2e, 0xf6, 0xac, 0xcf, 0x32, 0x7e,
	0x7a, 0xf6, 0xec, 0xc1, 0xbb, 0x05, 0xed, 0xe3, 0xeb, 0xce, 0x50, 0xed, 0x58, 0x86, 0xa4, 0x07,
	0xcd, 0xb1, 0x42, 0xaa, 0x91, 0x9c, 0x07, 0xc7, 0xc4, 0xc5, 0xcf, 0x61, 0xcf, 0x46, 0xb8, 0xc1,
	0x3f, 0xb1, 0x01, 0xc0, 0x3d, 0xe3, 0xf9, 0xa8, 0x9c, 0xee, 0xdf, 0xfd, 0x2b, 0x3f, 0xba, 0x7d,
	0xbd, 0x79, 0x63, 0x7a, 0x65, 0xd2, 0x20, 0x13, 0x45, 0xc8, 0xc5, 0x9a, 0x16, 0x42, 0x8b, 0x70,
	0x6d, 0x52, 0x54, 0x1c, 0x35, 0x6e, 0xfb, 0x99, 0xca, 0xfb, 0x8c, 0x2f, 0x15, 0xdd, 0x6a, 0x65,
	0x32, 0x6d, 0x14, 0x86, 0x54, 0xb2, 0xb0, 0xfa, 0x8e, 0xd4, 0x2c, 0xd3, 0x66, 0x95, 0xae, 0x3e,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xd5, 0x45, 0x59, 0xb1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualMachineServiceClient is the client API for VirtualMachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMachineServiceClient interface {
	Create(ctx context.Context, in *VirtualMachine, opts ...grpc.CallOption) (*VirtualMachine, error)
	Delete(ctx context.Context, in *VirtualMachine, opts ...grpc.CallOption) (*VirtualMachine, error)
	FindByName(ctx context.Context, in *VirtualMachine, opts ...grpc.CallOption) (*VirtualMachine, error)
}

type virtualMachineServiceClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMachineServiceClient(cc *grpc.ClientConn) VirtualMachineServiceClient {
	return &virtualMachineServiceClient{cc}
}

func (c *virtualMachineServiceClient) Create(ctx context.Context, in *VirtualMachine, opts ...grpc.CallOption) (*VirtualMachine, error) {
	out := new(VirtualMachine)
	err := c.cc.Invoke(ctx, "/VirtualMachineService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineServiceClient) Delete(ctx context.Context, in *VirtualMachine, opts ...grpc.CallOption) (*VirtualMachine, error) {
	out := new(VirtualMachine)
	err := c.cc.Invoke(ctx, "/VirtualMachineService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineServiceClient) FindByName(ctx context.Context, in *VirtualMachine, opts ...grpc.CallOption) (*VirtualMachine, error) {
	out := new(VirtualMachine)
	err := c.cc.Invoke(ctx, "/VirtualMachineService/FindByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineServiceServer is the server API for VirtualMachineService service.
type VirtualMachineServiceServer interface {
	Create(context.Context, *VirtualMachine) (*VirtualMachine, error)
	Delete(context.Context, *VirtualMachine) (*VirtualMachine, error)
	FindByName(context.Context, *VirtualMachine) (*VirtualMachine, error)
}

// UnimplementedVirtualMachineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineServiceServer struct {
}

func (*UnimplementedVirtualMachineServiceServer) Create(ctx context.Context, req *VirtualMachine) (*VirtualMachine, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedVirtualMachineServiceServer) Delete(ctx context.Context, req *VirtualMachine) (*VirtualMachine, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedVirtualMachineServiceServer) FindByName(ctx context.Context, req *VirtualMachine) (*VirtualMachine, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByName not implemented")
}

func RegisterVirtualMachineServiceServer(s *grpc.Server, srv VirtualMachineServiceServer) {
	s.RegisterService(&_VirtualMachineService_serviceDesc, srv)
}

func _VirtualMachineService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VirtualMachineService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineServiceServer).Create(ctx, req.(*VirtualMachine))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VirtualMachineService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineServiceServer).Delete(ctx, req.(*VirtualMachine))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineService_FindByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineServiceServer).FindByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VirtualMachineService/FindByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineServiceServer).FindByName(ctx, req.(*VirtualMachine))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "VirtualMachineService",
	HandlerType: (*VirtualMachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VirtualMachineService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VirtualMachineService_Delete_Handler,
		},
		{
			MethodName: "FindByName",
			Handler:    _VirtualMachineService_FindByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtual_machine.proto",
}

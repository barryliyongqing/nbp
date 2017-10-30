// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dock.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	dock.proto

It has these top-level messages:
	CreateVolumeOpts
	DeleteVolumeOpts
	CreateVolumeSnapshotOpts
	DeleteVolumeSnapshotOpts
	CreateAttachmentOpts
	HostInfo
	GenericResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// CreateVolumeOpts is a structure which indicates all required properties
// for creating a volume.
type CreateVolumeOpts struct {
	// The uuid of the volume, optional when creating.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The name of the volume, required.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// The requested capacity of the volume, required.
	Size int64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	// The description of the volume, optional.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// When create volume from snapshot, this field is required.
	SnapshotId string `protobuf:"bytes,5,opt,name=snapshotId" json:"snapshotId,omitempty"`
	// The locality that volume belongs to, required.
	AvailabilityZone string `protobuf:"bytes,6,opt,name=availabilityZone" json:"availabilityZone,omitempty"`
	// The service level that volume belongs to, required.
	ProfileId string `protobuf:"bytes,7,opt,name=profileId" json:"profileId,omitempty"`
	// The uuid of the pool on which volume will be created, required.
	PoolId string `protobuf:"bytes,8,opt,name=poolId" json:"poolId,omitempty"`
	// The name of the pool on which volume will be created, required.
	PoolName string `protobuf:"bytes,9,opt,name=poolName" json:"poolName,omitempty"`
	// The metadata of the volume, optional.
	Metadata map[string]string `protobuf:"bytes,10,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The dock infomation on which the request will be executed
	DockId string `protobuf:"bytes,11,opt,name=dockId" json:"dockId,omitempty"`
	// The storage driver type.
	DriverName string `protobuf:"bytes,12,opt,name=driverName" json:"driverName,omitempty"`
}

func (m *CreateVolumeOpts) Reset()                    { *m = CreateVolumeOpts{} }
func (m *CreateVolumeOpts) String() string            { return proto1.CompactTextString(m) }
func (*CreateVolumeOpts) ProtoMessage()               {}
func (*CreateVolumeOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateVolumeOpts) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateVolumeOpts) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateVolumeOpts) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CreateVolumeOpts) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateVolumeOpts) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

func (m *CreateVolumeOpts) GetAvailabilityZone() string {
	if m != nil {
		return m.AvailabilityZone
	}
	return ""
}

func (m *CreateVolumeOpts) GetProfileId() string {
	if m != nil {
		return m.ProfileId
	}
	return ""
}

func (m *CreateVolumeOpts) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *CreateVolumeOpts) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *CreateVolumeOpts) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateVolumeOpts) GetDockId() string {
	if m != nil {
		return m.DockId
	}
	return ""
}

func (m *CreateVolumeOpts) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

// DeleteVolumeOpts is a structure which indicates all required properties
// for deleting a volume.
type DeleteVolumeOpts struct {
	// The uuid of the volume, required.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The dock infomation on which the request will be executed
	DockId string `protobuf:"bytes,2,opt,name=dockId" json:"dockId,omitempty"`
	// The storage driver type.
	DriverName string `protobuf:"bytes,3,opt,name=driverName" json:"driverName,omitempty"`
}

func (m *DeleteVolumeOpts) Reset()                    { *m = DeleteVolumeOpts{} }
func (m *DeleteVolumeOpts) String() string            { return proto1.CompactTextString(m) }
func (*DeleteVolumeOpts) ProtoMessage()               {}
func (*DeleteVolumeOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeleteVolumeOpts) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteVolumeOpts) GetDockId() string {
	if m != nil {
		return m.DockId
	}
	return ""
}

func (m *DeleteVolumeOpts) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

// CreateVolumeSnapshotOpts is a structure which indicates all required
// properties for creating a volume snapshot.
type CreateVolumeSnapshotOpts struct {
	// The uuid of the volume snapshot, optional.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The name of the volume snapshot, required.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// The size of the volume that snapshot belongs to, required.
	Size int64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	// The description of the volume snapshot, optional.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// The uuid of the volume that snapshot belongs to, required.
	VolumeId string `protobuf:"bytes,5,opt,name=volumeId" json:"volumeId,omitempty"`
	// The dock infomation on which the request will be executed
	DockId string `protobuf:"bytes,6,opt,name=dockId" json:"dockId,omitempty"`
	// The storage driver type.
	DriverName string `protobuf:"bytes,7,opt,name=driverName" json:"driverName,omitempty"`
}

func (m *CreateVolumeSnapshotOpts) Reset()                    { *m = CreateVolumeSnapshotOpts{} }
func (m *CreateVolumeSnapshotOpts) String() string            { return proto1.CompactTextString(m) }
func (*CreateVolumeSnapshotOpts) ProtoMessage()               {}
func (*CreateVolumeSnapshotOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateVolumeSnapshotOpts) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateVolumeSnapshotOpts) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateVolumeSnapshotOpts) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CreateVolumeSnapshotOpts) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateVolumeSnapshotOpts) GetVolumeId() string {
	if m != nil {
		return m.VolumeId
	}
	return ""
}

func (m *CreateVolumeSnapshotOpts) GetDockId() string {
	if m != nil {
		return m.DockId
	}
	return ""
}

func (m *CreateVolumeSnapshotOpts) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

// DeleteVolumeSnapshotOpts is a structure which indicates all required
// properties for deleting a volume snapshot.
type DeleteVolumeSnapshotOpts struct {
	// The uuid of the volume snapshot, required.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The dock infomation on which the request will be executed
	DockId string `protobuf:"bytes,2,opt,name=dockId" json:"dockId,omitempty"`
	// The storage driver type.
	DriverName string `protobuf:"bytes,3,opt,name=driverName" json:"driverName,omitempty"`
}

func (m *DeleteVolumeSnapshotOpts) Reset()                    { *m = DeleteVolumeSnapshotOpts{} }
func (m *DeleteVolumeSnapshotOpts) String() string            { return proto1.CompactTextString(m) }
func (*DeleteVolumeSnapshotOpts) ProtoMessage()               {}
func (*DeleteVolumeSnapshotOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteVolumeSnapshotOpts) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteVolumeSnapshotOpts) GetDockId() string {
	if m != nil {
		return m.DockId
	}
	return ""
}

func (m *DeleteVolumeSnapshotOpts) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

// CreateAttachmentOpts is a structure which indicates all required
// properties for creating a volume attachment.
type CreateAttachmentOpts struct {
	// The uuid of the volume attachment, optional.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The uuid of the volume, required.
	VolumeId string `protobuf:"bytes,2,opt,name=volumeId" json:"volumeId,omitempty"`
	// This field indicates if the volume is attached locally, optional.
	DoLocalAttach bool `protobuf:"varint,3,opt,name=doLocalAttach" json:"doLocalAttach,omitempty"`
	// This field indicates if the volume is attached multiple times, optional.
	MultiPath bool `protobuf:"varint,4,opt,name=multiPath" json:"multiPath,omitempty"`
	// The infomation of the host node on which the volume will be attached.
	HostInfo *HostInfo `protobuf:"bytes,5,opt,name=hostInfo" json:"hostInfo,omitempty"`
	// The dock infomation on which the request will be executed
	DockId string `protobuf:"bytes,6,opt,name=dockId" json:"dockId,omitempty"`
	// The storage driver type.
	DriverName string `protobuf:"bytes,7,opt,name=driverName" json:"driverName,omitempty"`
}

func (m *CreateAttachmentOpts) Reset()                    { *m = CreateAttachmentOpts{} }
func (m *CreateAttachmentOpts) String() string            { return proto1.CompactTextString(m) }
func (*CreateAttachmentOpts) ProtoMessage()               {}
func (*CreateAttachmentOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateAttachmentOpts) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateAttachmentOpts) GetVolumeId() string {
	if m != nil {
		return m.VolumeId
	}
	return ""
}

func (m *CreateAttachmentOpts) GetDoLocalAttach() bool {
	if m != nil {
		return m.DoLocalAttach
	}
	return false
}

func (m *CreateAttachmentOpts) GetMultiPath() bool {
	if m != nil {
		return m.MultiPath
	}
	return false
}

func (m *CreateAttachmentOpts) GetHostInfo() *HostInfo {
	if m != nil {
		return m.HostInfo
	}
	return nil
}

func (m *CreateAttachmentOpts) GetDockId() string {
	if m != nil {
		return m.DockId
	}
	return ""
}

func (m *CreateAttachmentOpts) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

type HostInfo struct {
	// The platform of the host, such as "x86_64"
	Platform string `protobuf:"bytes,1,opt,name=platform" json:"platform,omitempty"`
	// The type of OS, such as "linux","windows", etc.
	OsType string `protobuf:"bytes,2,opt,name=osType" json:"osType,omitempty"`
	// The name of the host
	Host string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	// The ip address of the host
	Ip string `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
	// The initiator infomation, such as: "iqn.2017.com.redhat:e08039b48d5c"
	Initiator string `protobuf:"bytes,5,opt,name=initiator" json:"initiator,omitempty"`
}

func (m *HostInfo) Reset()                    { *m = HostInfo{} }
func (m *HostInfo) String() string            { return proto1.CompactTextString(m) }
func (*HostInfo) ProtoMessage()               {}
func (*HostInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HostInfo) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *HostInfo) GetOsType() string {
	if m != nil {
		return m.OsType
	}
	return ""
}

func (m *HostInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HostInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *HostInfo) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

// Generic response, it return:
// 1. Return result with message when create/update resource successfully.
// 2. Return result without message when delete resource successfully.
// 3. Return Error with error code and message when operate unsuccessfully.
type GenericResponse struct {
	// Types that are valid to be assigned to Reply:
	//	*GenericResponse_Result_
	//	*GenericResponse_Error_
	Reply isGenericResponse_Reply `protobuf_oneof:"reply"`
}

func (m *GenericResponse) Reset()                    { *m = GenericResponse{} }
func (m *GenericResponse) String() string            { return proto1.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()               {}
func (*GenericResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isGenericResponse_Reply interface {
	isGenericResponse_Reply()
}

type GenericResponse_Result_ struct {
	Result *GenericResponse_Result `protobuf:"bytes,1,opt,name=result,oneof"`
}
type GenericResponse_Error_ struct {
	Error *GenericResponse_Error `protobuf:"bytes,2,opt,name=error,oneof"`
}

func (*GenericResponse_Result_) isGenericResponse_Reply() {}
func (*GenericResponse_Error_) isGenericResponse_Reply()  {}

func (m *GenericResponse) GetReply() isGenericResponse_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *GenericResponse) GetResult() *GenericResponse_Result {
	if x, ok := m.GetReply().(*GenericResponse_Result_); ok {
		return x.Result
	}
	return nil
}

func (m *GenericResponse) GetError() *GenericResponse_Error {
	if x, ok := m.GetReply().(*GenericResponse_Error_); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GenericResponse) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _GenericResponse_OneofMarshaler, _GenericResponse_OneofUnmarshaler, _GenericResponse_OneofSizer, []interface{}{
		(*GenericResponse_Result_)(nil),
		(*GenericResponse_Error_)(nil),
	}
}

func _GenericResponse_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*GenericResponse)
	// reply
	switch x := m.Reply.(type) {
	case *GenericResponse_Result_:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Result); err != nil {
			return err
		}
	case *GenericResponse_Error_:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GenericResponse.Reply has unexpected type %T", x)
	}
	return nil
}

func _GenericResponse_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*GenericResponse)
	switch tag {
	case 1: // reply.result
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(GenericResponse_Result)
		err := b.DecodeMessage(msg)
		m.Reply = &GenericResponse_Result_{msg}
		return true, err
	case 2: // reply.error
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(GenericResponse_Error)
		err := b.DecodeMessage(msg)
		m.Reply = &GenericResponse_Error_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GenericResponse_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*GenericResponse)
	// reply
	switch x := m.Reply.(type) {
	case *GenericResponse_Result_:
		s := proto1.Size(x.Result)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *GenericResponse_Error_:
		s := proto1.Size(x.Error)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GenericResponse_Result struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *GenericResponse_Result) Reset()                    { *m = GenericResponse_Result{} }
func (m *GenericResponse_Result) String() string            { return proto1.CompactTextString(m) }
func (*GenericResponse_Result) ProtoMessage()               {}
func (*GenericResponse_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *GenericResponse_Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GenericResponse_Error struct {
	Code        string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *GenericResponse_Error) Reset()                    { *m = GenericResponse_Error{} }
func (m *GenericResponse_Error) String() string            { return proto1.CompactTextString(m) }
func (*GenericResponse_Error) ProtoMessage()               {}
func (*GenericResponse_Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 1} }

func (m *GenericResponse_Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GenericResponse_Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto1.RegisterType((*CreateVolumeOpts)(nil), "proto.CreateVolumeOpts")
	proto1.RegisterType((*DeleteVolumeOpts)(nil), "proto.DeleteVolumeOpts")
	proto1.RegisterType((*CreateVolumeSnapshotOpts)(nil), "proto.CreateVolumeSnapshotOpts")
	proto1.RegisterType((*DeleteVolumeSnapshotOpts)(nil), "proto.DeleteVolumeSnapshotOpts")
	proto1.RegisterType((*CreateAttachmentOpts)(nil), "proto.CreateAttachmentOpts")
	proto1.RegisterType((*HostInfo)(nil), "proto.HostInfo")
	proto1.RegisterType((*GenericResponse)(nil), "proto.GenericResponse")
	proto1.RegisterType((*GenericResponse_Result)(nil), "proto.GenericResponse.Result")
	proto1.RegisterType((*GenericResponse_Error)(nil), "proto.GenericResponse.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Dock service

type DockClient interface {
	// Create a volume
	CreateVolume(ctx context.Context, in *CreateVolumeOpts, opts ...grpc.CallOption) (*GenericResponse, error)
	// Delete a volume
	DeleteVolume(ctx context.Context, in *DeleteVolumeOpts, opts ...grpc.CallOption) (*GenericResponse, error)
	// Create a volume snapshot
	CreateVolumeSnapshot(ctx context.Context, in *CreateVolumeSnapshotOpts, opts ...grpc.CallOption) (*GenericResponse, error)
	// Delete a volume snapshot
	DeleteVolumeSnapshot(ctx context.Context, in *DeleteVolumeSnapshotOpts, opts ...grpc.CallOption) (*GenericResponse, error)
	// Create a volume attachment
	CreateAttachment(ctx context.Context, in *CreateAttachmentOpts, opts ...grpc.CallOption) (*GenericResponse, error)
}

type dockClient struct {
	cc *grpc.ClientConn
}

func NewDockClient(cc *grpc.ClientConn) DockClient {
	return &dockClient{cc}
}

func (c *dockClient) CreateVolume(ctx context.Context, in *CreateVolumeOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := grpc.Invoke(ctx, "/proto.Dock/CreateVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) DeleteVolume(ctx context.Context, in *DeleteVolumeOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := grpc.Invoke(ctx, "/proto.Dock/DeleteVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) CreateVolumeSnapshot(ctx context.Context, in *CreateVolumeSnapshotOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := grpc.Invoke(ctx, "/proto.Dock/CreateVolumeSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) DeleteVolumeSnapshot(ctx context.Context, in *DeleteVolumeSnapshotOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := grpc.Invoke(ctx, "/proto.Dock/DeleteVolumeSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockClient) CreateAttachment(ctx context.Context, in *CreateAttachmentOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := grpc.Invoke(ctx, "/proto.Dock/CreateAttachment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dock service

type DockServer interface {
	// Create a volume
	CreateVolume(context.Context, *CreateVolumeOpts) (*GenericResponse, error)
	// Delete a volume
	DeleteVolume(context.Context, *DeleteVolumeOpts) (*GenericResponse, error)
	// Create a volume snapshot
	CreateVolumeSnapshot(context.Context, *CreateVolumeSnapshotOpts) (*GenericResponse, error)
	// Delete a volume snapshot
	DeleteVolumeSnapshot(context.Context, *DeleteVolumeSnapshotOpts) (*GenericResponse, error)
	// Create a volume attachment
	CreateAttachment(context.Context, *CreateAttachmentOpts) (*GenericResponse, error)
}

func RegisterDockServer(s *grpc.Server, srv DockServer) {
	s.RegisterService(&_Dock_serviceDesc, srv)
}

func _Dock_CreateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).CreateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dock/CreateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).CreateVolume(ctx, req.(*CreateVolumeOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_DeleteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVolumeOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).DeleteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dock/DeleteVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).DeleteVolume(ctx, req.(*DeleteVolumeOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_CreateVolumeSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeSnapshotOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).CreateVolumeSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dock/CreateVolumeSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).CreateVolumeSnapshot(ctx, req.(*CreateVolumeSnapshotOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_DeleteVolumeSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVolumeSnapshotOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).DeleteVolumeSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dock/DeleteVolumeSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).DeleteVolumeSnapshot(ctx, req.(*DeleteVolumeSnapshotOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dock_CreateAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttachmentOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockServer).CreateAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dock/CreateAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockServer).CreateAttachment(ctx, req.(*CreateAttachmentOpts))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Dock",
	HandlerType: (*DockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVolume",
			Handler:    _Dock_CreateVolume_Handler,
		},
		{
			MethodName: "DeleteVolume",
			Handler:    _Dock_DeleteVolume_Handler,
		},
		{
			MethodName: "CreateVolumeSnapshot",
			Handler:    _Dock_CreateVolumeSnapshot_Handler,
		},
		{
			MethodName: "DeleteVolumeSnapshot",
			Handler:    _Dock_DeleteVolumeSnapshot_Handler,
		},
		{
			MethodName: "CreateAttachment",
			Handler:    _Dock_CreateAttachment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dock.proto",
}

func init() { proto1.RegisterFile("dock.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xae, 0x9d, 0x26, 0x71, 0x27, 0xed, 0xaf, 0xd1, 0xaa, 0xea, 0x6f, 0x15, 0x0a, 0x44, 0x11,
	0x48, 0x15, 0x48, 0x39, 0x04, 0x24, 0x10, 0x88, 0x43, 0xa1, 0x15, 0x8d, 0xc4, 0x5f, 0x83, 0x38,
	0xf4, 0xb6, 0xb5, 0xb7, 0x64, 0x55, 0xdb, 0x6b, 0xed, 0x6e, 0x22, 0x85, 0x13, 0xcf, 0xc6, 0x89,
	0x77, 0x81, 0x57, 0x40, 0x42, 0xbb, 0x5e, 0x3b, 0x8e, 0x93, 0xf8, 0x50, 0x89, 0x53, 0x66, 0x66,
	0x67, 0x67, 0xe6, 0xfb, 0x76, 0x3e, 0x07, 0x20, 0xe4, 0xc1, 0xf5, 0x30, 0x15, 0x5c, 0x71, 0xd4,
	0x34, 0x3f, 0x83, 0x1f, 0x0d, 0xe8, 0xbe, 0x12, 0x94, 0x28, 0xfa, 0x85, 0x47, 0xd3, 0x98, 0xbe,
	0x4f, 0x95, 0x44, 0xff, 0x81, 0xcb, 0x42, 0xec, 0xf4, 0x9d, 0xe3, 0x1d, 0xdf, 0x65, 0x21, 0x42,
	0xb0, 0x9d, 0x90, 0x98, 0x62, 0xd7, 0x44, 0x8c, 0xad, 0x63, 0x92, 0x7d, 0xa3, 0xb8, 0xd1, 0x77,
	0x8e, 0x1b, 0xbe, 0xb1, 0x51, 0x1f, 0x3a, 0x21, 0x95, 0x81, 0x60, 0xa9, 0x62, 0x3c, 0xc1, 0xdb,
	0x26, 0xbd, 0x1c, 0x42, 0x77, 0x00, 0x64, 0x42, 0x52, 0x39, 0xe1, 0x6a, 0x1c, 0xe2, 0xa6, 0x49,
	0x28, 0x45, 0xd0, 0x03, 0xe8, 0x92, 0x19, 0x61, 0x11, 0xb9, 0x64, 0x11, 0x53, 0xf3, 0x0b, 0x9e,
	0x50, 0xdc, 0x32, 0x59, 0x2b, 0x71, 0x74, 0x04, 0x3b, 0xa9, 0xe0, 0x57, 0x2c, 0xa2, 0xe3, 0x10,
	0xb7, 0x4d, 0xd2, 0x22, 0x80, 0x0e, 0xa1, 0x95, 0x72, 0x1e, 0x8d, 0x43, 0xec, 0x99, 0x23, 0xeb,
	0xa1, 0x1e, 0x78, 0xda, 0x7a, 0xa7, 0xf1, 0xec, 0x98, 0x93, 0xc2, 0x47, 0x27, 0xe0, 0xc5, 0x54,
	0x91, 0x90, 0x28, 0x82, 0xa1, 0xdf, 0x38, 0xee, 0x8c, 0xee, 0x67, 0x6c, 0x0d, 0xab, 0x14, 0x0d,
	0xdf, 0xda, 0xbc, 0xb3, 0x44, 0x89, 0xb9, 0x5f, 0x5c, 0xd3, 0x6d, 0x35, 0xc9, 0xe3, 0x10, 0x77,
	0xb2, 0xb6, 0x99, 0xa7, 0x81, 0x87, 0x82, 0xcd, 0xa8, 0x30, 0x8d, 0x77, 0x33, 0xe0, 0x8b, 0x48,
	0xef, 0x39, 0xec, 0x2d, 0x95, 0x44, 0x5d, 0x68, 0x5c, 0xd3, 0xb9, 0x7d, 0x04, 0x6d, 0xa2, 0x03,
	0x68, 0xce, 0x48, 0x34, 0xcd, 0x9f, 0x21, 0x73, 0x9e, 0xb9, 0x4f, 0x9d, 0xc1, 0x05, 0x74, 0x4f,
	0x69, 0x44, 0x6b, 0xdf, 0x70, 0x31, 0x98, 0x5b, 0x33, 0x58, 0xa3, 0x3a, 0xd8, 0xe0, 0xa7, 0x03,
	0xb8, 0x8c, 0xfe, 0x93, 0x7d, 0xac, 0x7f, 0xbc, 0x28, 0x3d, 0xf0, 0x66, 0xa6, 0x5f, 0xb1, 0x26,
	0x85, 0x5f, 0x82, 0xd2, 0xaa, 0x81, 0xd2, 0x5e, 0x81, 0x72, 0x09, 0xb8, 0x4c, 0x53, 0x2d, 0x92,
	0x9b, 0xd2, 0xf5, 0xcb, 0x81, 0x83, 0x8c, 0xae, 0x13, 0xa5, 0x48, 0x30, 0x89, 0x69, 0xb2, 0xbe,
	0x41, 0x19, 0xa0, 0x5b, 0x01, 0x78, 0x0f, 0xf6, 0x42, 0xfe, 0x86, 0x07, 0x24, 0xca, 0x8a, 0x98,
	0x3e, 0x9e, 0xbf, 0x1c, 0xd4, 0xfb, 0x1f, 0x4f, 0x23, 0xc5, 0x3e, 0x10, 0x35, 0x31, 0x14, 0x7a,
	0xfe, 0x22, 0x80, 0x1e, 0x82, 0x37, 0xe1, 0x52, 0x8d, 0x93, 0x2b, 0x6e, 0x08, 0xec, 0x8c, 0xf6,
	0xed, 0x2e, 0x9f, 0xdb, 0xb0, 0x5f, 0x24, 0xdc, 0x98, 0xd1, 0xef, 0x0e, 0x78, 0x79, 0x39, 0xa3,
	0xac, 0x88, 0xa8, 0x2b, 0x2e, 0x62, 0x8b, 0xb3, 0xf0, 0x75, 0x03, 0x2e, 0x3f, 0xcf, 0xd3, 0x7c,
	0x35, 0xac, 0xa7, 0x97, 0x43, 0x0f, 0x61, 0x89, 0x34, 0xb6, 0x61, 0x2a, 0xb5, 0x3b, 0xe1, 0xb2,
	0x54, 0xe3, 0x64, 0x09, 0x53, 0x8c, 0x28, 0x2e, 0xec, 0x2e, 0x2c, 0x02, 0x83, 0xdf, 0x0e, 0xec,
	0xbf, 0xa6, 0x09, 0x15, 0x2c, 0xf0, 0xa9, 0x4c, 0x79, 0x22, 0x29, 0x7a, 0x02, 0x2d, 0x41, 0xe5,
	0x34, 0x52, 0x66, 0x8e, 0xce, 0xe8, 0xb6, 0x45, 0x5e, 0xc9, 0x1b, 0xfa, 0x26, 0xe9, 0x7c, 0xcb,
	0xb7, 0xe9, 0xe8, 0x31, 0x34, 0xa9, 0x10, 0x5c, 0x98, 0x29, 0x3b, 0xa3, 0xa3, 0x0d, 0xf7, 0xce,
	0x74, 0xce, 0xf9, 0x96, 0x9f, 0x25, 0xf7, 0x06, 0xd0, 0xca, 0x2a, 0x21, 0x0c, 0xed, 0x98, 0x4a,
	0x49, 0xbe, 0x52, 0xcb, 0x40, 0xee, 0xf6, 0x5e, 0x40, 0xd3, 0xdc, 0xd2, 0x88, 0x03, 0x1e, 0xe6,
	0xe7, 0xc6, 0xae, 0xca, 0xc1, 0x5d, 0x91, 0xc3, 0xcb, 0x36, 0x34, 0x05, 0x4d, 0xa3, 0xf9, 0xe8,
	0x8f, 0x0b, 0xdb, 0xa7, 0x3c, 0xb8, 0x46, 0x27, 0xb0, 0x5b, 0x96, 0x25, 0xfa, 0x7f, 0xc3, 0x97,
	0xaa, 0x77, 0xb8, 0x1e, 0xc4, 0x60, 0x4b, 0x97, 0x28, 0xeb, 0xa1, 0x28, 0x51, 0xfd, 0x96, 0xd4,
	0x94, 0xf8, 0x98, 0x6f, 0xfb, 0xb2, 0xa4, 0xd0, 0xdd, 0x35, 0xd3, 0x94, 0xf5, 0x56, 0x5f, 0x72,
	0x9d, 0x4a, 0x8b, 0x92, 0x9b, 0x24, 0x5c, 0x53, 0x72, 0x9c, 0xff, 0xc7, 0x2d, 0x34, 0x89, 0x6e,
	0x2d, 0x4d, 0xb8, 0x2c, 0xd6, 0xcd, 0xa5, 0x2e, 0x5b, 0xe6, 0xe0, 0xd1, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0x21, 0x7b, 0x0c, 0x4b, 0x07, 0x00, 0x00,
}

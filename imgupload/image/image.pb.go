// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

/*
Package image is a generated protocol buffer package.

It is generated from these files:
	image.proto

It has these top-level messages:
	Void
	File
	ID
	IDs
	Files
	Count
*/
package image

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

type Void struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Void) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type File struct {
	// Unix Timestamp from Created at
	UnixCreatedAt int64 `protobuf:"varint,1,opt,name=unixCreatedAt" json:"unixCreatedAt,omitempty"`
	// Image in bytes
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Is already processed
	Processed bool `protobuf:"varint,3,opt,name=processed" json:"processed,omitempty"`
	// extension of file
	Extension string `protobuf:"bytes,4,opt,name=extension" json:"extension,omitempty"`
	Filename  string `protobuf:"bytes,5,opt,name=filename" json:"filename,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *File) GetUnixCreatedAt() int64 {
	if m != nil {
		return m.UnixCreatedAt
	}
	return 0
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *File) GetProcessed() bool {
	if m != nil {
		return m.Processed
	}
	return false
}

func (m *File) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func (m *File) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type ID struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type IDs struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *IDs) Reset()                    { *m = IDs{} }
func (m *IDs) String() string            { return proto.CompactTextString(m) }
func (*IDs) ProtoMessage()               {}
func (*IDs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IDs) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Files struct {
	Files []*File `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
}

func (m *Files) Reset()                    { *m = Files{} }
func (m *Files) String() string            { return proto.CompactTextString(m) }
func (*Files) ProtoMessage()               {}
func (*Files) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Files) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

type Count struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *Count) Reset()                    { *m = Count{} }
func (m *Count) String() string            { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()               {}
func (*Count) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Count) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Void)(nil), "image.Void")
	proto.RegisterType((*File)(nil), "image.File")
	proto.RegisterType((*ID)(nil), "image.ID")
	proto.RegisterType((*IDs)(nil), "image.IDs")
	proto.RegisterType((*Files)(nil), "image.Files")
	proto.RegisterType((*Count)(nil), "image.Count")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FileHandler service

type FileHandlerClient interface {
	GetNewImagesIDs(ctx context.Context, in *Void, opts ...grpc.CallOption) (*IDs, error)
	GetImage(ctx context.Context, in *ID, opts ...grpc.CallOption) (*File, error)
	GetCount(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Count, error)
	PostImage(ctx context.Context, in *File, opts ...grpc.CallOption) (*Void, error)
	MarkProcessed(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Void, error)
}

type fileHandlerClient struct {
	cc *grpc.ClientConn
}

func NewFileHandlerClient(cc *grpc.ClientConn) FileHandlerClient {
	return &fileHandlerClient{cc}
}

func (c *fileHandlerClient) GetNewImagesIDs(ctx context.Context, in *Void, opts ...grpc.CallOption) (*IDs, error) {
	out := new(IDs)
	err := grpc.Invoke(ctx, "/image.FileHandler/GetNewImagesIDs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileHandlerClient) GetImage(ctx context.Context, in *ID, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := grpc.Invoke(ctx, "/image.FileHandler/GetImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileHandlerClient) GetCount(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := grpc.Invoke(ctx, "/image.FileHandler/GetCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileHandlerClient) PostImage(ctx context.Context, in *File, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/image.FileHandler/PostImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileHandlerClient) MarkProcessed(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/image.FileHandler/MarkProcessed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileHandler service

type FileHandlerServer interface {
	GetNewImagesIDs(context.Context, *Void) (*IDs, error)
	GetImage(context.Context, *ID) (*File, error)
	GetCount(context.Context, *Void) (*Count, error)
	PostImage(context.Context, *File) (*Void, error)
	MarkProcessed(context.Context, *ID) (*Void, error)
}

func RegisterFileHandlerServer(s *grpc.Server, srv FileHandlerServer) {
	s.RegisterService(&_FileHandler_serviceDesc, srv)
}

func _FileHandler_GetNewImagesIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileHandlerServer).GetNewImagesIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image.FileHandler/GetNewImagesIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileHandlerServer).GetNewImagesIDs(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileHandler_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileHandlerServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image.FileHandler/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileHandlerServer).GetImage(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileHandler_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileHandlerServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image.FileHandler/GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileHandlerServer).GetCount(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileHandler_PostImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileHandlerServer).PostImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image.FileHandler/PostImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileHandlerServer).PostImage(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileHandler_MarkProcessed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileHandlerServer).MarkProcessed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image.FileHandler/MarkProcessed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileHandlerServer).MarkProcessed(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "image.FileHandler",
	HandlerType: (*FileHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewImagesIDs",
			Handler:    _FileHandler_GetNewImagesIDs_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _FileHandler_GetImage_Handler,
		},
		{
			MethodName: "GetCount",
			Handler:    _FileHandler_GetCount_Handler,
		},
		{
			MethodName: "PostImage",
			Handler:    _FileHandler_PostImage_Handler,
		},
		{
			MethodName: "MarkProcessed",
			Handler:    _FileHandler_MarkProcessed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image.proto",
}

func init() { proto.RegisterFile("image.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0x9a, 0x44, 0x9a, 0x49, 0xab, 0x32, 0x14, 0x0c, 0x41, 0x21, 0x2e, 0x05, 0xab, 0x48,
	0x0f, 0xf5, 0x17, 0x48, 0x83, 0x35, 0x07, 0xa5, 0xe4, 0xe0, 0x3d, 0x76, 0xc7, 0xb2, 0xd8, 0x66,
	0x4b, 0x76, 0x8b, 0xfd, 0x29, 0xfe, 0x3a, 0x7f, 0x8b, 0xec, 0x26, 0xf6, 0x03, 0xbc, 0x84, 0x79,
	0xf3, 0x5e, 0xde, 0xcc, 0x1b, 0x16, 0x42, 0xb1, 0x2a, 0x16, 0x34, 0x5a, 0x57, 0x52, 0x4b, 0xf4,
	0x2d, 0x60, 0x11, 0x78, 0x6f, 0x52, 0x70, 0x3c, 0x07, 0x77, 0xa5, 0x16, 0x91, 0x93, 0x38, 0xc3,
	0x20, 0x37, 0x25, 0xfb, 0x76, 0xc0, 0x7b, 0x12, 0x4b, 0xc2, 0x01, 0xf4, 0x36, 0xa5, 0xd8, 0x4e,
	0x2a, 0x2a, 0x34, 0xf1, 0x47, 0x6d, 0x45, 0x6e, 0x7e, 0xdc, 0x44, 0x04, 0x8f, 0x17, 0xba, 0x88,
	0xda, 0x89, 0x33, 0xec, 0xe6, 0xb6, 0xc6, 0x4b, 0x08, 0xd6, 0x95, 0x9c, 0x93, 0x52, 0xc4, 0x23,
	0x37, 0x71, 0x86, 0x9d, 0x7c, 0xdf, 0x30, 0x2c, 0x6d, 0x35, 0x95, 0x4a, 0xc8, 0x32, 0xf2, 0xec,
	0xe0, 0x7d, 0x03, 0x63, 0xe8, 0x7c, 0x88, 0x25, 0x95, 0xc5, 0x8a, 0x22, 0xdf, 0x92, 0x3b, 0xcc,
	0xfa, 0xd0, 0xce, 0x52, 0x3c, 0x35, 0xdf, 0x66, 0xe3, 0x76, 0x96, 0xb2, 0x0b, 0x70, 0xb3, 0x54,
	0x99, 0x24, 0x82, 0xab, 0xc8, 0x49, 0x5c, 0x93, 0x44, 0x70, 0xc5, 0xee, 0xc0, 0x37, 0x41, 0x14,
	0x5e, 0x83, 0x6f, 0x3c, 0x6a, 0x32, 0x1c, 0x87, 0xa3, 0xfa, 0x20, 0x86, 0xcc, 0x6b, 0x86, 0x5d,
	0x81, 0x3f, 0x91, 0x9b, 0x52, 0x63, 0x1f, 0xfc, 0xb9, 0x29, 0x9a, 0xb4, 0x35, 0x18, 0xff, 0x38,
	0x10, 0x1a, 0xf9, 0x73, 0x51, 0xf2, 0x25, 0x55, 0x78, 0x0f, 0x67, 0x53, 0xd2, 0xaf, 0xf4, 0x95,
	0x19, 0x27, 0x65, 0xe6, 0xff, 0xb9, 0x9a, 0xb3, 0xc6, 0xd0, 0x80, 0x2c, 0x55, 0xac, 0x85, 0x03,
	0xe8, 0x4c, 0x49, 0x5b, 0x29, 0x06, 0x3b, 0x26, 0x3e, 0xdc, 0x83, 0xb5, 0xf0, 0xc6, 0xaa, 0xea,
	0x2d, 0x8e, 0xcc, 0xba, 0x0d, 0xb0, 0x94, 0x15, 0x06, 0x33, 0xa9, 0x1a, 0xbf, 0x43, 0x93, 0xf8,
	0xf0, 0x37, 0xd6, 0xc2, 0x5b, 0xe8, 0xbd, 0x14, 0xd5, 0xe7, 0x6c, 0x77, 0xfa, 0x7f, 0x86, 0xd7,
	0xd2, 0xf7, 0x13, 0xfb, 0x3a, 0x1e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe8, 0x01, 0x30,
	0x2c, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go.
// source: plugins/plugin.proto
// DO NOT EDIT!

/*
Package openapi_plugin_v1 is a generated protocol buffer package.

It is generated from these files:
	plugins/plugin.proto

It has these top-level messages:
	Version
	Parameter
	Request
	Response
	File
	Wrapper
*/
package openapi_plugin_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The version number of OpenAPI compiler.
type Version struct {
	Major int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch int32 `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix string `protobuf:"bytes,4,opt,name=suffix" json:"suffix,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Version) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *Version) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *Version) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *Version) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

// A parameter passed to the plugin from (or through) the OpenAPI compiler.
type Parameter struct {
	// The name of the parameter as specified in the option string
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The parameter value as specified in the option string
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Parameter) Reset()                    { *m = Parameter{} }
func (m *Parameter) String() string            { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()               {}
func (*Parameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// An encoded Request is written to the plugin's stdin.
type Request struct {
	// A wrapped OpenAPI document to process.
	Wrapper *Wrapper `protobuf:"bytes,1,opt,name=wrapper" json:"wrapper,omitempty"`
	// Output path specified in the plugin invocation.
	OutputPath string `protobuf:"bytes,2,opt,name=output_path,json=outputPath" json:"output_path,omitempty"`
	// Plugin parameters parsed from the invocation string.
	Parameters []*Parameter `protobuf:"bytes,3,rep,name=parameters" json:"parameters,omitempty"`
	// The version number of openapi compiler.
	CompilerVersion *Version `protobuf:"bytes,4,opt,name=compiler_version,json=compilerVersion" json:"compiler_version,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Request) GetWrapper() *Wrapper {
	if m != nil {
		return m.Wrapper
	}
	return nil
}

func (m *Request) GetOutputPath() string {
	if m != nil {
		return m.OutputPath
	}
	return ""
}

func (m *Request) GetParameters() []*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Request) GetCompilerVersion() *Version {
	if m != nil {
		return m.CompilerVersion
	}
	return nil
}

// The plugin writes an encoded Response to stdout.
type Response struct {
	// Error message.  If non-empty, the plugin failed.
	// The plugin process should exit with status code zero
	// even if it reports an error in this way.
	//
	// This should be used to indicate errors which prevent the plugin from
	// operating as intended.  Errors which indicate a problem in openapic
	// itself -- such as the input Document being unparseable -- should be
	// reported by writing a message to stderr and exiting with a non-zero
	// status code.
	Errors []string `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty"`
	// file output, each file will be written by openapic to an appropriate location.
	Files []*File `protobuf:"bytes,2,rep,name=files" json:"files,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

// File describes a file generated by a plugin.
type File struct {
	// name of the file
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// data to be written to the file
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Wrapper wraps an OpenAPI document with its version.
type Wrapper struct {
	// filename or URL of the wrapped document
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// version of the OpenAPI specification that is used by the wrapped document
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// valid serialized protocol buffer of the named OpenAPI specification version
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Wrapper) Reset()                    { *m = Wrapper{} }
func (m *Wrapper) String() string            { return proto.CompactTextString(m) }
func (*Wrapper) ProtoMessage()               {}
func (*Wrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Wrapper) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Wrapper) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Wrapper) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Version)(nil), "openapi.plugin.v1.Version")
	proto.RegisterType((*Parameter)(nil), "openapi.plugin.v1.Parameter")
	proto.RegisterType((*Request)(nil), "openapi.plugin.v1.Request")
	proto.RegisterType((*Response)(nil), "openapi.plugin.v1.Response")
	proto.RegisterType((*File)(nil), "openapi.plugin.v1.File")
	proto.RegisterType((*Wrapper)(nil), "openapi.plugin.v1.Wrapper")
}

func init() { proto.RegisterFile("plugins/plugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8f, 0xd3, 0x30,
	0x10, 0x85, 0x95, 0x26, 0xdd, 0x90, 0x29, 0xd2, 0x82, 0xb5, 0x80, 0x85, 0x90, 0x88, 0x72, 0xea,
	0x85, 0xc0, 0x2e, 0x70, 0xe3, 0xd2, 0x56, 0x20, 0x71, 0x40, 0x0d, 0x3e, 0xc0, 0xb1, 0x32, 0x61,
	0xda, 0x1a, 0x25, 0xb1, 0xb1, 0x9d, 0xc0, 0xef, 0xe1, 0xbf, 0xf1, 0x3f, 0x50, 0x6c, 0x07, 0x90,
	0x36, 0xa7, 0xfa, 0x7b, 0x1a, 0x3f, 0xbf, 0x79, 0x0d, 0x5c, 0xa9, 0xa6, 0x3f, 0x89, 0xce, 0x3c,
	0xf7, 0xbf, 0xa5, 0xd2, 0xd2, 0x4a, 0x72, 0x5f, 0x2a, 0xec, 0xb8, 0x12, 0x65, 0x50, 0x87, 0xeb,
	0xa2, 0x86, 0xf4, 0x13, 0x6a, 0x23, 0x64, 0x47, 0xae, 0x60, 0xd9, 0xf2, 0x6f, 0x52, 0xd3, 0x28,
	0x8f, 0xd6, 0x4b, 0xe6, 0xc1, 0xa9, 0xa2, 0x93, 0x9a, 0x2e, 0x82, 0x3a, 0xc2, 0xa8, 0x2a, 0x6e,
	0xeb, 0x33, 0x8d, 0xbd, 0xea, 0x80, 0x3c, 0x84, 0x0b, 0xd3, 0x1f, 0x8f, 0xe2, 0x27, 0x4d, 0xf2,
	0x68, 0x9d, 0xb1, 0x40, 0xc5, 0x6b, 0xc8, 0x2a, 0xae, 0x79, 0x8b, 0x16, 0x35, 0x21, 0x90, 0x74,
	0xbc, 0x45, 0xf7, 0x4a, 0xc6, 0xdc, 0x79, 0xb4, 0x1b, 0x78, 0xd3, 0xa3, 0x7b, 0x24, 0x63, 0x1e,
	0x8a, 0xdf, 0x11, 0xa4, 0x0c, 0xbf, 0xf7, 0x68, 0x2c, 0x79, 0x05, 0xe9, 0x0f, 0xcd, 0x95, 0x42,
	0x1f, 0x6f, 0x75, 0xf3, 0xb8, 0xbc, 0xb5, 0x4c, 0xf9, 0xd9, 0x4f, 0xb0, 0x69, 0x94, 0x3c, 0x85,
	0x95, 0xec, 0xad, 0xea, 0xed, 0x41, 0x71, 0x7b, 0x0e, 0xee, 0xe0, 0xa5, 0x8a, 0xdb, 0x33, 0x79,
	0x03, 0xa0, 0xa6, 0x64, 0x86, 0xc6, 0x79, 0xbc, 0x5e, 0xdd, 0x3c, 0x99, 0x71, 0xfe, 0x1b, 0x9f,
	0xfd, 0x37, 0x4f, 0xde, 0xc2, 0xbd, 0x5a, 0xb6, 0x4a, 0x34, 0xa8, 0x0f, 0x83, 0x6f, 0xd1, 0x6d,
	0x3e, 0x9f, 0x2e, 0xf4, 0xcc, 0x2e, 0xa7, 0x3b, 0x41, 0x28, 0x3e, 0xc2, 0x1d, 0x86, 0x46, 0xc9,
	0xce, 0xe0, 0x58, 0x21, 0x6a, 0x2d, 0xb5, 0xa1, 0x51, 0x1e, 0x8f, 0x15, 0x7a, 0x22, 0xcf, 0x60,
	0x79, 0x14, 0x0d, 0x1a, 0xba, 0x70, 0x19, 0x1f, 0xcd, 0xf8, 0xbf, 0x13, 0x0d, 0x32, 0x3f, 0x55,
	0x94, 0x90, 0x8c, 0x38, 0x5b, 0x36, 0x81, 0xe4, 0x2b, 0xb7, 0xdc, 0xb5, 0x71, 0x97, 0xb9, 0x73,
	0xf1, 0x01, 0xd2, 0x50, 0xde, 0xec, 0x15, 0x0a, 0xe9, 0xb4, 0x9f, 0xef, 0x70, 0xc2, 0x7f, 0xff,
	0x5c, 0xec, 0xdc, 0x3c, 0x6c, 0x5f, 0xc0, 0xa5, 0xd4, 0xa7, 0x29, 0x63, 0x5d, 0x0e, 0xd7, 0xdb,
	0x07, 0x7b, 0x85, 0xdd, 0xa6, 0x7a, 0xbf, 0x0b, 0xcb, 0x57, 0x2e, 0x77, 0x15, 0xfd, 0x5a, 0xc4,
	0xfb, 0xcd, 0xee, 0xcb, 0x85, 0xfb, 0x42, 0x5f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x51, 0x4d,
	0xec, 0x80, 0xb9, 0x02, 0x00, 0x00,
}

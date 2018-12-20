// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: virtual_service.proto

package v1 // import "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=vs
// @solo-kit:resource.plural_name=virtual_services
// @solo-kit:resource.resource_groups=api.gateway.solo.io
// A virtual service describes the set of routes to match for a set of domains.
// Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).
type VirtualService struct {
	VirtualHost *v1.VirtualHost `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost" json:"virtual_host,omitempty"`
	// If provided, the Gateway will serve TLS/SSL traffic for this set of routes
	SslConfig *v1.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig" json:"ssl_config,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VirtualService) Reset()         { *m = VirtualService{} }
func (m *VirtualService) String() string { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()    {}
func (*VirtualService) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_service_6dd0fca6e05ec630, []int{0}
}
func (m *VirtualService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualService.Unmarshal(m, b)
}
func (m *VirtualService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualService.Marshal(b, m, deterministic)
}
func (dst *VirtualService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualService.Merge(dst, src)
}
func (m *VirtualService) XXX_Size() int {
	return xxx_messageInfo_VirtualService.Size(m)
}
func (m *VirtualService) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualService.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualService proto.InternalMessageInfo

func (m *VirtualService) GetVirtualHost() *v1.VirtualHost {
	if m != nil {
		return m.VirtualHost
	}
	return nil
}

func (m *VirtualService) GetSslConfig() *v1.SslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualService) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *VirtualService) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*VirtualService)(nil), "gateway.solo.io.VirtualService")
}
func (this *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.VirtualHost.Equal(that1.VirtualHost) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("virtual_service.proto", fileDescriptor_virtual_service_6dd0fca6e05ec630)
}

var fileDescriptor_virtual_service_6dd0fca6e05ec630 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0x2a, 0x31,
	0x14, 0x86, 0x2f, 0xe4, 0x86, 0x7b, 0x2d, 0x46, 0xe3, 0x04, 0x75, 0x60, 0x21, 0x86, 0x95, 0x1b,
	0xdb, 0x20, 0x89, 0x41, 0xe3, 0x0a, 0x17, 0xba, 0x71, 0x33, 0x24, 0x2e, 0xdc, 0x90, 0x32, 0x94,
	0x52, 0x19, 0x38, 0x93, 0xf6, 0x30, 0xca, 0x23, 0xf8, 0x26, 0x3e, 0x8a, 0x4f, 0xc1, 0xc2, 0x47,
	0xf0, 0x09, 0xcc, 0xb4, 0x1d, 0x09, 0x89, 0x89, 0xae, 0x7a, 0x4e, 0xfe, 0xf3, 0xfd, 0xc9, 0xff,
	0x97, 0xec, 0x67, 0x4a, 0xe3, 0x82, 0x27, 0x03, 0x23, 0x74, 0xa6, 0x62, 0x41, 0x53, 0x0d, 0x08,
	0xc1, 0xae, 0xe4, 0x28, 0x9e, 0xf8, 0x92, 0x1a, 0x48, 0x80, 0x2a, 0x68, 0xd4, 0x24, 0x48, 0xb0,
	0x1a, 0xcb, 0x27, 0x77, 0xd6, 0x68, 0x4b, 0x85, 0x93, 0xc5, 0x90, 0xc6, 0x30, 0x63, 0xf9, 0xe5,
	0xa9, 0x02, 0xf7, 0x4e, 0x15, 0x32, 0x9e, 0x2a, 0x96, 0xb5, 0xd9, 0x4c, 0x20, 0x1f, 0x71, 0xe4,
	0x1e, 0x61, 0xbf, 0x40, 0x0c, 0x72, 0x5c, 0x18, 0x0f, 0x74, 0xbf, 0x01, 0x64, 0x02, 0xc0, 0x52,
	0x0d, 0x8f, 0x22, 0x46, 0xe3, 0x36, 0x8f, 0xa6, 0x1a, 0x9e, 0x97, 0x8e, 0x6c, 0xbd, 0x94, 0xc9,
	0xce, 0xbd, 0x8b, 0xd7, 0x77, 0xe9, 0x82, 0x2b, 0xb2, 0x5d, 0x04, 0x9e, 0x80, 0xc1, 0xb0, 0x74,
	0x5c, 0x3a, 0xa9, 0x9e, 0xd5, 0x69, 0x6e, 0x51, 0x64, 0xa5, 0x9e, 0xb9, 0x05, 0x83, 0x51, 0x35,
	0x5b, 0x2f, 0xc1, 0x39, 0x21, 0xc6, 0x24, 0x83, 0x18, 0xe6, 0x63, 0x25, 0xc3, 0xb2, 0x65, 0x0f,
	0x37, 0xd9, 0xbe, 0x49, 0xae, 0xad, 0x1c, 0x6d, 0x99, 0x62, 0x0c, 0x6e, 0x48, 0xc5, 0x45, 0x0a,
	0x2b, 0x96, 0xa9, 0xd1, 0x18, 0xb4, 0x58, 0x33, 0x56, 0xeb, 0xd5, 0xdf, 0x56, 0xcd, 0x3f, 0x1f,
	0xab, 0xe6, 0x1e, 0x0a, 0x83, 0x23, 0x35, 0x1e, 0x5f, 0xb6, 0x94, 0x9c, 0x83, 0x16, 0xad, 0xc8,
	0xe3, 0x41, 0x97, 0xfc, 0x2f, 0xea, 0x0c, 0xff, 0x59, 0xab, 0x83, 0x4d, 0xab, 0x3b, 0xaf, 0xf6,
	0xfe, 0xe6, 0x66, 0xd1, 0xd7, 0x75, 0xef, 0xe2, 0xf5, 0xfd, 0xa8, 0xf4, 0xd0, 0xf9, 0xb9, 0x4b,
	0xf7, 0xef, 0x2c, 0x9d, 0x4a, 0x5f, 0xe9, 0xb0, 0x62, 0xdb, 0xec, 0x7c, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x27, 0x79, 0xb5, 0x36, 0x2b, 0x02, 0x00, 0x00,
}
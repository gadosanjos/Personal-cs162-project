// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1alpha/finding_type_stats.proto

package websecurityscanner

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// A FindingTypeStats resource represents stats regarding a specific FindingType
// of Findings under a given ScanRun.
type FindingTypeStats struct {
	// The finding type associated with the stats.
	FindingType Finding_FindingType `protobuf:"varint,1,opt,name=finding_type,json=findingType,proto3,enum=google.cloud.websecurityscanner.v1alpha.Finding_FindingType" json:"finding_type,omitempty"`
	// The count of findings belonging to this finding type.
	FindingCount         int32    `protobuf:"varint,2,opt,name=finding_count,json=findingCount,proto3" json:"finding_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindingTypeStats) Reset()         { *m = FindingTypeStats{} }
func (m *FindingTypeStats) String() string { return proto.CompactTextString(m) }
func (*FindingTypeStats) ProtoMessage()    {}
func (*FindingTypeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_91da39ac488bf6ea, []int{0}
}

func (m *FindingTypeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindingTypeStats.Unmarshal(m, b)
}
func (m *FindingTypeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindingTypeStats.Marshal(b, m, deterministic)
}
func (m *FindingTypeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindingTypeStats.Merge(m, src)
}
func (m *FindingTypeStats) XXX_Size() int {
	return xxx_messageInfo_FindingTypeStats.Size(m)
}
func (m *FindingTypeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_FindingTypeStats.DiscardUnknown(m)
}

var xxx_messageInfo_FindingTypeStats proto.InternalMessageInfo

func (m *FindingTypeStats) GetFindingType() Finding_FindingType {
	if m != nil {
		return m.FindingType
	}
	return Finding_FINDING_TYPE_UNSPECIFIED
}

func (m *FindingTypeStats) GetFindingCount() int32 {
	if m != nil {
		return m.FindingCount
	}
	return 0
}

func init() {
	proto.RegisterType((*FindingTypeStats)(nil), "google.cloud.websecurityscanner.v1alpha.FindingTypeStats")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1alpha/finding_type_stats.proto", fileDescriptor_91da39ac488bf6ea)
}

var fileDescriptor_91da39ac488bf6ea = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x2f, 0x4f, 0x4d, 0x2a, 0x4e, 0x4d, 0x2e,
	0x2d, 0xca, 0x2c, 0xa9, 0x2c, 0x4e, 0x4e, 0xcc, 0xcb, 0x4b, 0x2d, 0xd2, 0x2f, 0x33, 0x4c, 0xcc,
	0x29, 0xc8, 0x48, 0xd4, 0x4f, 0xcb, 0xcc, 0x4b, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0xa9, 0x2c, 0x48,
	0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x87, 0x98,
	0xa0, 0x07, 0x36, 0x41, 0x0f, 0xd3, 0x04, 0x3d, 0xa8, 0x09, 0x52, 0xa6, 0x24, 0x5a, 0x05, 0x31,
	0x5f, 0x69, 0x06, 0x23, 0x97, 0x80, 0x1b, 0x44, 0x24, 0xa4, 0xb2, 0x20, 0x35, 0x18, 0x64, 0xb5,
	0x50, 0x3c, 0x17, 0x0f, 0xb2, 0x83, 0x24, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x6c, 0xf4, 0x88,
	0x74, 0x8b, 0x1e, 0xd4, 0x40, 0x3d, 0x24, 0x83, 0x83, 0xb8, 0xd3, 0x10, 0x1c, 0x21, 0x65, 0x2e,
	0x5e, 0x98, 0x05, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x30,
	0x5b, 0x9d, 0x41, 0x62, 0x4e, 0x0b, 0x19, 0xb9, 0xb4, 0x93, 0xf3, 0x73, 0x89, 0xb5, 0xd5, 0x49,
	0x14, 0xdd, 0x1f, 0x01, 0x20, 0x1f, 0x06, 0x30, 0x46, 0x45, 0x42, 0x4d, 0x48, 0xcf, 0xcf, 0x49,
	0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f, 0xcd, 0x03, 0xfb, 0x5f, 0x1f, 0x22, 0x95,
	0x58, 0x90, 0x59, 0x4c, 0x30, 0xe4, 0xac, 0x31, 0xa5, 0x92, 0xd8, 0xc0, 0xa6, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x3e, 0x11, 0xae, 0xcf, 0xe9, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/mutate_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible mutate errors.
type MutateErrorEnum_MutateError int32

const (
	// Enum unspecified.
	MutateErrorEnum_UNSPECIFIED MutateErrorEnum_MutateError = 0
	// The received error code is not known in this version.
	MutateErrorEnum_UNKNOWN MutateErrorEnum_MutateError = 1
	// Requested resource was not found.
	MutateErrorEnum_RESOURCE_NOT_FOUND MutateErrorEnum_MutateError = 3
	// Cannot mutate the same resource twice in one request.
	MutateErrorEnum_ID_EXISTS_IN_MULTIPLE_MUTATES MutateErrorEnum_MutateError = 7
	// The field's contents don't match another field that represents the same
	// data.
	MutateErrorEnum_INCONSISTENT_FIELD_VALUES MutateErrorEnum_MutateError = 8
	// Mutates are not allowed for the requested resource.
	MutateErrorEnum_MUTATE_NOT_ALLOWED MutateErrorEnum_MutateError = 9
	// The resource isn't in Google Ads. It belongs to another ads system.
	MutateErrorEnum_RESOURCE_NOT_IN_GOOGLE_ADS MutateErrorEnum_MutateError = 10
)

var MutateErrorEnum_MutateError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "RESOURCE_NOT_FOUND",
	7:  "ID_EXISTS_IN_MULTIPLE_MUTATES",
	8:  "INCONSISTENT_FIELD_VALUES",
	9:  "MUTATE_NOT_ALLOWED",
	10: "RESOURCE_NOT_IN_GOOGLE_ADS",
}
var MutateErrorEnum_MutateError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"RESOURCE_NOT_FOUND":            3,
	"ID_EXISTS_IN_MULTIPLE_MUTATES": 7,
	"INCONSISTENT_FIELD_VALUES":     8,
	"MUTATE_NOT_ALLOWED":            9,
	"RESOURCE_NOT_IN_GOOGLE_ADS":    10,
}

func (x MutateErrorEnum_MutateError) String() string {
	return proto.EnumName(MutateErrorEnum_MutateError_name, int32(x))
}
func (MutateErrorEnum_MutateError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mutate_error_52bfa33b7773be5c, []int{0, 0}
}

// Container for enum describing possible mutate errors.
type MutateErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateErrorEnum) Reset()         { *m = MutateErrorEnum{} }
func (m *MutateErrorEnum) String() string { return proto.CompactTextString(m) }
func (*MutateErrorEnum) ProtoMessage()    {}
func (*MutateErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_mutate_error_52bfa33b7773be5c, []int{0}
}
func (m *MutateErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateErrorEnum.Unmarshal(m, b)
}
func (m *MutateErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateErrorEnum.Marshal(b, m, deterministic)
}
func (dst *MutateErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateErrorEnum.Merge(dst, src)
}
func (m *MutateErrorEnum) XXX_Size() int {
	return xxx_messageInfo_MutateErrorEnum.Size(m)
}
func (m *MutateErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MutateErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MutateErrorEnum)(nil), "google.ads.googleads.v1.errors.MutateErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.MutateErrorEnum_MutateError", MutateErrorEnum_MutateError_name, MutateErrorEnum_MutateError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/mutate_error.proto", fileDescriptor_mutate_error_52bfa33b7773be5c)
}

var fileDescriptor_mutate_error_52bfa33b7773be5c = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x49, 0x47, 0x62, 0xc0, 0x5d, 0x8c, 0xe5, 0x05, 0x12, 0x23, 0xa6, 0x12, 0x39, 0x80,
	0xa3, 0x88, 0x9d, 0x59, 0x79, 0x1a, 0x37, 0xb2, 0x48, 0xed, 0x68, 0x9c, 0x64, 0x10, 0x8a, 0x64,
	0x05, 0x12, 0x45, 0x95, 0xda, 0xb8, 0x8a, 0xd3, 0x1e, 0x88, 0x25, 0x77, 0xe0, 0x02, 0xec, 0xb8,
	0x06, 0x5c, 0x02, 0x25, 0xa6, 0x55, 0x59, 0x30, 0x2b, 0xff, 0x7e, 0xfa, 0xfe, 0xff, 0xd9, 0xef,
	0x81, 0xb0, 0x35, 0xa6, 0xdd, 0x36, 0x41, 0x55, 0xdb, 0xc0, 0xc9, 0x51, 0x1d, 0xc3, 0xa0, 0xe9,
	0x7b, 0xd3, 0xdb, 0x60, 0x77, 0x18, 0xaa, 0xa1, 0xd1, 0xd3, 0x0d, 0xef, 0x7b, 0x33, 0x18, 0xb4,
	0x70, 0x1c, 0xae, 0x6a, 0x8b, 0xcf, 0x16, 0x7c, 0x0c, 0xb1, 0xb3, 0xdc, 0xbe, 0x39, 0x45, 0xee,
	0x37, 0x41, 0xd5, 0x75, 0x66, 0xa8, 0x86, 0x8d, 0xe9, 0xac, 0x73, 0xfb, 0x3f, 0x3d, 0x70, 0xb3,
	0x9e, 0x42, 0xd9, 0x88, 0xb3, 0xee, 0xb0, 0xf3, 0xbf, 0x7b, 0x60, 0x7e, 0x51, 0x43, 0x37, 0x60,
	0x9e, 0x0b, 0x95, 0xb2, 0x25, 0x5f, 0x71, 0x16, 0xc1, 0x67, 0x68, 0x0e, 0xae, 0x73, 0xf1, 0x41,
	0xc8, 0x47, 0x01, 0x3d, 0xf4, 0x0a, 0xa0, 0x07, 0xa6, 0x64, 0xfe, 0xb0, 0x64, 0x5a, 0xc8, 0x4c,
	0xaf, 0x64, 0x2e, 0x22, 0x78, 0x85, 0xde, 0x82, 0x3b, 0x1e, 0x69, 0xf6, 0x91, 0xab, 0x4c, 0x69,
	0x2e, 0xf4, 0x3a, 0x4f, 0x32, 0x9e, 0x26, 0x4c, 0xaf, 0xf3, 0x8c, 0x66, 0x4c, 0xc1, 0x6b, 0x74,
	0x07, 0x5e, 0x73, 0xb1, 0x94, 0x42, 0x71, 0x95, 0x31, 0x91, 0xe9, 0x15, 0x67, 0x49, 0xa4, 0x0b,
	0x9a, 0xe4, 0x4c, 0xc1, 0x17, 0x63, 0xb2, 0x63, 0xa7, 0x5c, 0x9a, 0x24, 0xf2, 0x91, 0x45, 0xf0,
	0x25, 0x5a, 0x80, 0xdb, 0x7f, 0x3a, 0x72, 0xa1, 0x63, 0x29, 0xe3, 0x84, 0x69, 0x1a, 0x29, 0x08,
	0xee, 0x7f, 0x7b, 0xc0, 0xff, 0x62, 0x76, 0xf8, 0xe9, 0xc1, 0xdc, 0xc3, 0x8b, 0x3f, 0xa6, 0xe3,
	0x30, 0x52, 0xef, 0x53, 0xf4, 0xd7, 0xd3, 0x9a, 0x6d, 0xd5, 0xb5, 0xd8, 0xf4, 0x6d, 0xd0, 0x36,
	0xdd, 0x34, 0xaa, 0xd3, 0x3e, 0xf6, 0x1b, 0xfb, 0xbf, 0xf5, 0xbc, 0x77, 0xc7, 0xd7, 0xd9, 0x55,
	0x4c, 0xe9, 0xb7, 0xd9, 0x22, 0x76, 0x61, 0xb4, 0xb6, 0xd8, 0xc9, 0x51, 0x15, 0x21, 0x9e, 0x5a,
	0xda, 0x1f, 0x27, 0xa0, 0xa4, 0xb5, 0x2d, 0xcf, 0x40, 0x59, 0x84, 0xa5, 0x03, 0x7e, 0xcd, 0x7c,
	0x57, 0x25, 0x84, 0xd6, 0x96, 0x90, 0x33, 0x42, 0x48, 0x11, 0x12, 0xe2, 0xa0, 0xcf, 0xcf, 0xa7,
	0xd7, 0xbd, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xef, 0xa9, 0x17, 0xad, 0x3b, 0x02, 0x00, 0x00,
}

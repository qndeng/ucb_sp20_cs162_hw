// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/ad_group_criterion_approval_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enumerates AdGroupCriterion approval statuses.
type AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus int32

const (
	// Not specified.
	AdGroupCriterionApprovalStatusEnum_UNSPECIFIED AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus = 0
	// The value is unknown in this version.
	AdGroupCriterionApprovalStatusEnum_UNKNOWN AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus = 1
	// Approved.
	AdGroupCriterionApprovalStatusEnum_APPROVED AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus = 2
	// Disapproved.
	AdGroupCriterionApprovalStatusEnum_DISAPPROVED AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus = 3
	// Pending Review.
	AdGroupCriterionApprovalStatusEnum_PENDING_REVIEW AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus = 4
	// Under review.
	AdGroupCriterionApprovalStatusEnum_UNDER_REVIEW AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus = 5
)

var AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "APPROVED",
	3: "DISAPPROVED",
	4: "PENDING_REVIEW",
	5: "UNDER_REVIEW",
}

var AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"APPROVED":       2,
	"DISAPPROVED":    3,
	"PENDING_REVIEW": 4,
	"UNDER_REVIEW":   5,
}

func (x AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus) String() string {
	return proto.EnumName(AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus_name, int32(x))
}

func (AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_073bf627ba32857b, []int{0, 0}
}

// Container for enum describing possible AdGroupCriterion approval statuses.
type AdGroupCriterionApprovalStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupCriterionApprovalStatusEnum) Reset()         { *m = AdGroupCriterionApprovalStatusEnum{} }
func (m *AdGroupCriterionApprovalStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionApprovalStatusEnum) ProtoMessage()    {}
func (*AdGroupCriterionApprovalStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_073bf627ba32857b, []int{0}
}

func (m *AdGroupCriterionApprovalStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionApprovalStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupCriterionApprovalStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionApprovalStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionApprovalStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionApprovalStatusEnum.Merge(m, src)
}
func (m *AdGroupCriterionApprovalStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionApprovalStatusEnum.Size(m)
}
func (m *AdGroupCriterionApprovalStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionApprovalStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionApprovalStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus", AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus_name, AdGroupCriterionApprovalStatusEnum_AdGroupCriterionApprovalStatus_value)
	proto.RegisterType((*AdGroupCriterionApprovalStatusEnum)(nil), "google.ads.googleads.v2.enums.AdGroupCriterionApprovalStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/ad_group_criterion_approval_status.proto", fileDescriptor_073bf627ba32857b)
}

var fileDescriptor_073bf627ba32857b = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcd, 0x6a, 0xab, 0x40,
	0x18, 0xbd, 0x9a, 0x7b, 0x6f, 0xcb, 0x24, 0xb4, 0xe2, 0xb2, 0x34, 0x85, 0xa4, 0xfb, 0x11, 0xec,
	0x6e, 0xba, 0x9a, 0xc4, 0x89, 0x48, 0xc1, 0x48, 0x42, 0x0c, 0x14, 0x41, 0xa6, 0x19, 0x19, 0x84,
	0x64, 0x46, 0x1c, 0xcd, 0x0b, 0xf4, 0x4d, 0xba, 0xcc, 0xa3, 0xf4, 0x51, 0xfa, 0x04, 0x5d, 0x16,
	0xc7, 0xe8, 0xae, 0xd9, 0x0c, 0x87, 0xef, 0x9c, 0x39, 0xe7, 0xfb, 0x01, 0x0b, 0x2e, 0x25, 0xdf,
	0x67, 0x0e, 0x65, 0xca, 0x69, 0x61, 0x83, 0x8e, 0xae, 0x93, 0x89, 0xfa, 0xa0, 0x1c, 0xca, 0x52,
	0x5e, 0xca, 0xba, 0x48, 0x77, 0x65, 0x5e, 0x65, 0x65, 0x2e, 0x45, 0x4a, 0x8b, 0xa2, 0x94, 0x47,
	0xba, 0x4f, 0x55, 0x45, 0xab, 0x5a, 0xc1, 0xa2, 0x94, 0x95, 0xb4, 0xc7, 0xed, 0x67, 0x48, 0x99,
	0x82, 0xbd, 0x0f, 0x3c, 0xba, 0x50, 0xfb, 0xdc, 0xdd, 0x77, 0x31, 0x45, 0xee, 0x50, 0x21, 0x64,
	0x45, 0xab, 0x5c, 0x8a, 0xf3, 0xe7, 0xe9, 0xc9, 0x00, 0x53, 0xcc, 0xfc, 0x26, 0x68, 0xde, 0xe5,
	0xe0, 0x73, 0xcc, 0x5a, 0xa7, 0x10, 0x51, 0x1f, 0xa6, 0xef, 0x06, 0x78, 0xb8, 0x2c, 0xb3, 0x6f,
	0xc1, 0x70, 0x13, 0xae, 0x23, 0x32, 0x0f, 0x16, 0x01, 0xf1, 0xac, 0x3f, 0xf6, 0x10, 0x5c, 0x6d,
	0xc2, 0x97, 0x70, 0xb9, 0x0d, 0x2d, 0xc3, 0x1e, 0x81, 0x6b, 0x1c, 0x45, 0xab, 0x65, 0x4c, 0x3c,
	0xcb, 0x6c, 0xb4, 0x5e, 0xb0, 0xee, 0x0b, 0x03, 0xdb, 0x06, 0x37, 0x11, 0x09, 0xbd, 0x20, 0xf4,
	0xd3, 0x15, 0x89, 0x03, 0xb2, 0xb5, 0xfe, 0xda, 0x16, 0x18, 0x6d, 0x42, 0x8f, 0xac, 0xba, 0xca,
	0xbf, 0xd9, 0xb7, 0x01, 0x26, 0x3b, 0x79, 0x80, 0x17, 0x07, 0x9e, 0x3d, 0x5e, 0x6e, 0x34, 0x6a,
	0xe6, 0x8e, 0x8c, 0xd7, 0xd9, 0xd9, 0x85, 0xcb, 0x3d, 0x15, 0x1c, 0xca, 0x92, 0x3b, 0x3c, 0x13,
	0x7a, 0x2b, 0xdd, 0x39, 0x8a, 0x5c, 0xfd, 0x72, 0x9d, 0x67, 0xfd, 0x7e, 0x98, 0x03, 0x1f, 0xe3,
	0x93, 0x39, 0xf6, 0x5b, 0x2b, 0xcc, 0x14, 0x6c, 0x61, 0x83, 0x62, 0x17, 0x36, 0xbb, 0x53, 0x9f,
	0x1d, 0x9f, 0x60, 0xa6, 0x92, 0x9e, 0x4f, 0x62, 0x37, 0xd1, 0xfc, 0x97, 0x39, 0x69, 0x8b, 0x08,
	0x61, 0xa6, 0x10, 0xea, 0x15, 0x08, 0xc5, 0x2e, 0x42, 0x5a, 0xf3, 0xf6, 0x5f, 0x37, 0xf6, 0xf4,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xde, 0xc8, 0xce, 0xad, 0x35, 0x02, 0x00, 0x00,
}
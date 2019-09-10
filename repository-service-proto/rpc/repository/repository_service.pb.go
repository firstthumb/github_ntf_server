// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/repository_service.proto

package repository

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffa4c3ed6eaea42, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SearchResponse struct {
	Results              []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffa4c3ed6eaea42, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type Result struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffa4c3ed6eaea42, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "com.ekocaman.githubntf.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "com.ekocaman.githubntf.SearchResponse")
	proto.RegisterType((*Result)(nil), "com.ekocaman.githubntf.Result")
}

func init() { proto.RegisterFile("proto/repository_service.proto", fileDescriptor_7ffa4c3ed6eaea42) }

var fileDescriptor_7ffa4c3ed6eaea42 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0x2f, 0xaa, 0x8c, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x4b, 0x08, 0x89, 0x25, 0xe7, 0xe7, 0xea, 0xa5, 0x66, 0xe7,
	0x27, 0x27, 0xe6, 0x26, 0xe6, 0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe5, 0x95, 0xa4, 0x29,
	0xa9, 0x72, 0xf1, 0x06, 0xa7, 0x26, 0x16, 0x25, 0x67, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x89, 0x70, 0xb1, 0x16, 0x96, 0xa6, 0x16, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x41, 0x38, 0x4a, 0x5e, 0x5c, 0x7c, 0x30, 0x65, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x16,
	0x5c, 0xec, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46,
	0x72, 0x7a, 0xd8, 0xad, 0xd0, 0x0b, 0x02, 0x2b, 0x0b, 0x82, 0x29, 0x57, 0x92, 0xe1, 0x62, 0x83,
	0x08, 0x09, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x42, 0xad, 0x02, 0xb3, 0x8d, 0xb2, 0xb9,
	0xd8, 0xdc, 0xc1, 0x5a, 0x85, 0x12, 0xb9, 0x04, 0x60, 0x76, 0xc2, 0x3c, 0x25, 0xa4, 0x8a, 0xcb,
	0x12, 0x14, 0x4f, 0x48, 0xa9, 0x11, 0x52, 0x06, 0xf1, 0x84, 0x13, 0x77, 0x14, 0x27, 0x5c, 0x2e,
	0x89, 0x0d, 0x1c, 0x52, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xa8, 0xd1, 0x92, 0x4b,
	0x01, 0x00, 0x00,
}
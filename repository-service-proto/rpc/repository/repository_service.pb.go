// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repository_service.proto

package githubntf

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
	return fileDescriptor_a7e699c8334bc484, []int{0}
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
	return fileDescriptor_a7e699c8334bc484, []int{1}
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
	return fileDescriptor_a7e699c8334bc484, []int{2}
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

func init() { proto.RegisterFile("repository_service.proto", fileDescriptor_a7e699c8334bc484) }

var fileDescriptor_a7e699c8334bc484 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x4a, 0x2d, 0xc8,
	0x2f, 0xce, 0x2c, 0xc9, 0x2f, 0xaa, 0x8c, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4b, 0xce, 0xcf, 0xd5, 0x4b, 0xcd, 0xce, 0x4f, 0x4e, 0xcc,
	0x4d, 0xcc, 0xd3, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xca, 0x2b, 0x49, 0x53, 0x52, 0xe5, 0xe2,
	0x0d, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1,
	0x62, 0x2d, 0x2c, 0x4d, 0x2d, 0xaa, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94,
	0xbc, 0xb8, 0xf8, 0x60, 0xca, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x2c, 0xb8, 0xd8, 0x8b,
	0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xe4, 0xf4, 0xb0,
	0x5b, 0xa1, 0x17, 0x04, 0x56, 0x16, 0x04, 0x53, 0xae, 0x24, 0xc3, 0xc5, 0x06, 0x11, 0x12, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x85, 0x5a, 0x05, 0x66, 0x1b, 0x65, 0x73, 0xb1, 0xb9, 0x83,
	0xb5, 0x0a, 0x25, 0x72, 0x09, 0xc0, 0xec, 0x84, 0x79, 0x4a, 0x48, 0x15, 0x97, 0x25, 0x28, 0x9e,
	0x90, 0x52, 0x23, 0xa4, 0x0c, 0xe2, 0x09, 0x27, 0xee, 0x28, 0x4e, 0xb8, 0x5c, 0x12, 0x1b, 0x38,
	0xa4, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x2f, 0xef, 0xf5, 0x45, 0x01, 0x00, 0x00,
}

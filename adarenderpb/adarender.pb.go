// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adarender.proto

package adarender

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

// MarkdownData - markdown data
type MarkdownData struct {
	// strData - markdown string
	StrData string `protobuf:"bytes,1,opt,name=strData,proto3" json:"strData,omitempty"`
	// binaryData - binary data, it's like images
	BinaryData map[string][]byte `protobuf:"bytes,2,rep,name=binaryData,proto3" json:"binaryData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// templateName - template Name
	TemplateName string `protobuf:"bytes,10,opt,name=templateName,proto3" json:"templateName,omitempty"`
	// templateData - template data
	TemplateData         string   `protobuf:"bytes,11,opt,name=templateData,proto3" json:"templateData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkdownData) Reset()         { *m = MarkdownData{} }
func (m *MarkdownData) String() string { return proto.CompactTextString(m) }
func (*MarkdownData) ProtoMessage()    {}
func (*MarkdownData) Descriptor() ([]byte, []int) {
	return fileDescriptor_adarender_909440fee60fb195, []int{0}
}
func (m *MarkdownData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkdownData.Unmarshal(m, b)
}
func (m *MarkdownData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkdownData.Marshal(b, m, deterministic)
}
func (dst *MarkdownData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkdownData.Merge(dst, src)
}
func (m *MarkdownData) XXX_Size() int {
	return xxx_messageInfo_MarkdownData.Size(m)
}
func (m *MarkdownData) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkdownData.DiscardUnknown(m)
}

var xxx_messageInfo_MarkdownData proto.InternalMessageInfo

func (m *MarkdownData) GetStrData() string {
	if m != nil {
		return m.StrData
	}
	return ""
}

func (m *MarkdownData) GetBinaryData() map[string][]byte {
	if m != nil {
		return m.BinaryData
	}
	return nil
}

func (m *MarkdownData) GetTemplateName() string {
	if m != nil {
		return m.TemplateName
	}
	return ""
}

func (m *MarkdownData) GetTemplateData() string {
	if m != nil {
		return m.TemplateData
	}
	return ""
}

// MarkdownStream - markdown stream data
type MarkdownStream struct {
	// totalLength - If the message is too long, it will send data in multiple msg, this is the total length.
	TotalLength int32 `protobuf:"varint,1,opt,name=totalLength,proto3" json:"totalLength,omitempty"`
	// curStart - The starting point of the current data (in bytes).
	CurStart int32 `protobuf:"varint,2,opt,name=curStart,proto3" json:"curStart,omitempty"`
	// curLength - The length of the current data (in bytes).
	CurLength int32 `protobuf:"varint,3,opt,name=curLength,proto3" json:"curLength,omitempty"`
	// hashData - This is the hash of each paragraph.
	HashData string `protobuf:"bytes,4,opt,name=hashData,proto3" json:"hashData,omitempty"`
	// totalHashData - If multiple messages return data, this is the hash value of all data, only sent in the last message.
	TotalHashData string `protobuf:"bytes,5,opt,name=totalHashData,proto3" json:"totalHashData,omitempty"`
	// data - binary data
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// error - error string
	Error string `protobuf:"bytes,100,opt,name=error,proto3" json:"error,omitempty"`
	// markdownData - If the data does not exceed 4mb, this is the data that is directly available.
	MarkdownData *MarkdownData `protobuf:"bytes,200,opt,name=markdownData,proto3" json:"markdownData,omitempty"`
	// token - API token
	Token                string   `protobuf:"bytes,300,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkdownStream) Reset()         { *m = MarkdownStream{} }
func (m *MarkdownStream) String() string { return proto.CompactTextString(m) }
func (*MarkdownStream) ProtoMessage()    {}
func (*MarkdownStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_adarender_909440fee60fb195, []int{1}
}
func (m *MarkdownStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkdownStream.Unmarshal(m, b)
}
func (m *MarkdownStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkdownStream.Marshal(b, m, deterministic)
}
func (dst *MarkdownStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkdownStream.Merge(dst, src)
}
func (m *MarkdownStream) XXX_Size() int {
	return xxx_messageInfo_MarkdownStream.Size(m)
}
func (m *MarkdownStream) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkdownStream.DiscardUnknown(m)
}

var xxx_messageInfo_MarkdownStream proto.InternalMessageInfo

func (m *MarkdownStream) GetTotalLength() int32 {
	if m != nil {
		return m.TotalLength
	}
	return 0
}

func (m *MarkdownStream) GetCurStart() int32 {
	if m != nil {
		return m.CurStart
	}
	return 0
}

func (m *MarkdownStream) GetCurLength() int32 {
	if m != nil {
		return m.CurLength
	}
	return 0
}

func (m *MarkdownStream) GetHashData() string {
	if m != nil {
		return m.HashData
	}
	return ""
}

func (m *MarkdownStream) GetTotalHashData() string {
	if m != nil {
		return m.TotalHashData
	}
	return ""
}

func (m *MarkdownStream) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MarkdownStream) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *MarkdownStream) GetMarkdownData() *MarkdownData {
	if m != nil {
		return m.MarkdownData
	}
	return nil
}

func (m *MarkdownStream) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// HTMLData - HTML data
type HTMLData struct {
	// strData - HTML string
	StrData string `protobuf:"bytes,1,opt,name=strData,proto3" json:"strData,omitempty"`
	// binaryData - binary data, it's like images, css file
	BinaryData           map[string][]byte `protobuf:"bytes,2,rep,name=binaryData,proto3" json:"binaryData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HTMLData) Reset()         { *m = HTMLData{} }
func (m *HTMLData) String() string { return proto.CompactTextString(m) }
func (*HTMLData) ProtoMessage()    {}
func (*HTMLData) Descriptor() ([]byte, []int) {
	return fileDescriptor_adarender_909440fee60fb195, []int{2}
}
func (m *HTMLData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTMLData.Unmarshal(m, b)
}
func (m *HTMLData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTMLData.Marshal(b, m, deterministic)
}
func (dst *HTMLData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTMLData.Merge(dst, src)
}
func (m *HTMLData) XXX_Size() int {
	return xxx_messageInfo_HTMLData.Size(m)
}
func (m *HTMLData) XXX_DiscardUnknown() {
	xxx_messageInfo_HTMLData.DiscardUnknown(m)
}

var xxx_messageInfo_HTMLData proto.InternalMessageInfo

func (m *HTMLData) GetStrData() string {
	if m != nil {
		return m.StrData
	}
	return ""
}

func (m *HTMLData) GetBinaryData() map[string][]byte {
	if m != nil {
		return m.BinaryData
	}
	return nil
}

// HTMLStream - HTML data stream
type HTMLStream struct {
	// totalLength - If the message is too long, it will send data in multiple msg, this is the total length.
	TotalLength int32 `protobuf:"varint,1,opt,name=totalLength,proto3" json:"totalLength,omitempty"`
	// curStart - The starting point of the current data (in bytes).
	CurStart int32 `protobuf:"varint,2,opt,name=curStart,proto3" json:"curStart,omitempty"`
	// curLength - The length of the current data (in bytes).
	CurLength int32 `protobuf:"varint,3,opt,name=curLength,proto3" json:"curLength,omitempty"`
	// hashData - This is the hash of each paragraph.
	HashData string `protobuf:"bytes,4,opt,name=hashData,proto3" json:"hashData,omitempty"`
	// totalHashData - If multiple messages return data, this is the hash value of all data, only sent in the last message.
	TotalHashData string `protobuf:"bytes,5,opt,name=totalHashData,proto3" json:"totalHashData,omitempty"`
	// data - binary data
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// error - error string
	Error string `protobuf:"bytes,100,opt,name=error,proto3" json:"error,omitempty"`
	// markdownData - If the data does not exceed 4mb, this is the data that is directly available.
	HtmlData             *HTMLData `protobuf:"bytes,200,opt,name=htmlData,proto3" json:"htmlData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *HTMLStream) Reset()         { *m = HTMLStream{} }
func (m *HTMLStream) String() string { return proto.CompactTextString(m) }
func (*HTMLStream) ProtoMessage()    {}
func (*HTMLStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_adarender_909440fee60fb195, []int{3}
}
func (m *HTMLStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTMLStream.Unmarshal(m, b)
}
func (m *HTMLStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTMLStream.Marshal(b, m, deterministic)
}
func (dst *HTMLStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTMLStream.Merge(dst, src)
}
func (m *HTMLStream) XXX_Size() int {
	return xxx_messageInfo_HTMLStream.Size(m)
}
func (m *HTMLStream) XXX_DiscardUnknown() {
	xxx_messageInfo_HTMLStream.DiscardUnknown(m)
}

var xxx_messageInfo_HTMLStream proto.InternalMessageInfo

func (m *HTMLStream) GetTotalLength() int32 {
	if m != nil {
		return m.TotalLength
	}
	return 0
}

func (m *HTMLStream) GetCurStart() int32 {
	if m != nil {
		return m.CurStart
	}
	return 0
}

func (m *HTMLStream) GetCurLength() int32 {
	if m != nil {
		return m.CurLength
	}
	return 0
}

func (m *HTMLStream) GetHashData() string {
	if m != nil {
		return m.HashData
	}
	return ""
}

func (m *HTMLStream) GetTotalHashData() string {
	if m != nil {
		return m.TotalHashData
	}
	return ""
}

func (m *HTMLStream) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *HTMLStream) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *HTMLStream) GetHtmlData() *HTMLData {
	if m != nil {
		return m.HtmlData
	}
	return nil
}

func init() {
	proto.RegisterType((*MarkdownData)(nil), "adarender.MarkdownData")
	proto.RegisterMapType((map[string][]byte)(nil), "adarender.MarkdownData.BinaryDataEntry")
	proto.RegisterType((*MarkdownStream)(nil), "adarender.MarkdownStream")
	proto.RegisterType((*HTMLData)(nil), "adarender.HTMLData")
	proto.RegisterMapType((map[string][]byte)(nil), "adarender.HTMLData.BinaryDataEntry")
	proto.RegisterType((*HTMLStream)(nil), "adarender.HTMLStream")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdaRenderServiceClient is the client API for AdaRenderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdaRenderServiceClient interface {
	// render - render markdown
	Render(ctx context.Context, opts ...grpc.CallOption) (AdaRenderService_RenderClient, error)
}

type adaRenderServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdaRenderServiceClient(cc *grpc.ClientConn) AdaRenderServiceClient {
	return &adaRenderServiceClient{cc}
}

func (c *adaRenderServiceClient) Render(ctx context.Context, opts ...grpc.CallOption) (AdaRenderService_RenderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AdaRenderService_serviceDesc.Streams[0], "/adarender.AdaRenderService/render", opts...)
	if err != nil {
		return nil, err
	}
	x := &adaRenderServiceRenderClient{stream}
	return x, nil
}

type AdaRenderService_RenderClient interface {
	Send(*MarkdownStream) error
	Recv() (*HTMLStream, error)
	grpc.ClientStream
}

type adaRenderServiceRenderClient struct {
	grpc.ClientStream
}

func (x *adaRenderServiceRenderClient) Send(m *MarkdownStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *adaRenderServiceRenderClient) Recv() (*HTMLStream, error) {
	m := new(HTMLStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdaRenderServiceServer is the server API for AdaRenderService service.
type AdaRenderServiceServer interface {
	// render - render markdown
	Render(AdaRenderService_RenderServer) error
}

func RegisterAdaRenderServiceServer(s *grpc.Server, srv AdaRenderServiceServer) {
	s.RegisterService(&_AdaRenderService_serviceDesc, srv)
}

func _AdaRenderService_Render_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AdaRenderServiceServer).Render(&adaRenderServiceRenderServer{stream})
}

type AdaRenderService_RenderServer interface {
	Send(*HTMLStream) error
	Recv() (*MarkdownStream, error)
	grpc.ServerStream
}

type adaRenderServiceRenderServer struct {
	grpc.ServerStream
}

func (x *adaRenderServiceRenderServer) Send(m *HTMLStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *adaRenderServiceRenderServer) Recv() (*MarkdownStream, error) {
	m := new(MarkdownStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AdaRenderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adarender.AdaRenderService",
	HandlerType: (*AdaRenderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "render",
			Handler:       _AdaRenderService_Render_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "adarender.proto",
}

func init() { proto.RegisterFile("adarender.proto", fileDescriptor_adarender_909440fee60fb195) }

var fileDescriptor_adarender_909440fee60fb195 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xcd, 0xae, 0xd2, 0x50,
	0x10, 0xf6, 0x94, 0x5b, 0x84, 0x69, 0xf5, 0xde, 0x8c, 0xde, 0x58, 0x89, 0x8b, 0xa6, 0x9a, 0xd8,
	0x15, 0x21, 0xb8, 0x31, 0x46, 0x13, 0x7f, 0x23, 0x0b, 0x70, 0x71, 0xe0, 0x05, 0x0e, 0xf4, 0xc4,
	0x12, 0xfa, 0x43, 0x0e, 0x07, 0x0c, 0x2f, 0xe0, 0xb3, 0xb8, 0xd0, 0xf7, 0xf0, 0x91, 0x5c, 0x9a,
	0x33, 0xa5, 0xa5, 0x25, 0xa8, 0x0b, 0x77, 0x77, 0x77, 0xe6, 0x9b, 0x6f, 0xbe, 0x6f, 0xa6, 0x33,
	0x29, 0x5c, 0x8a, 0x48, 0x28, 0x99, 0x45, 0x52, 0xf5, 0xd7, 0x2a, 0xd7, 0x39, 0x76, 0x2b, 0x20,
	0xf8, 0xc5, 0xc0, 0x9d, 0x08, 0xb5, 0x8a, 0xf2, 0x2f, 0xd9, 0x7b, 0xa1, 0x05, 0x7a, 0x70, 0x7b,
	0xa3, 0x95, 0x79, 0x7a, 0xcc, 0x67, 0x61, 0x97, 0x97, 0x21, 0x7e, 0x04, 0x98, 0x2f, 0x33, 0xa1,
	0xf6, 0x94, 0xb4, 0xfc, 0x56, 0xe8, 0x0c, 0x9f, 0xf6, 0x8f, 0xda, 0x75, 0x99, 0xfe, 0xdb, 0x8a,
	0xf9, 0x21, 0xd3, 0x6a, 0xcf, 0x6b, 0xa5, 0x18, 0x80, 0xab, 0x65, 0xba, 0x4e, 0x84, 0x96, 0x9f,
	0x44, 0x2a, 0x3d, 0x20, 0x9f, 0x06, 0x56, 0xe7, 0x90, 0x9d, 0xd3, 0xe4, 0x18, 0xac, 0xf7, 0x0a,
	0x2e, 0x4f, 0x6c, 0xf0, 0x0a, 0x5a, 0x2b, 0xb9, 0x3f, 0x74, 0x6e, 0x9e, 0x78, 0x1f, 0xec, 0x9d,
	0x48, 0xb6, 0xd2, 0xb3, 0x7c, 0x16, 0xba, 0xbc, 0x08, 0x5e, 0x58, 0xcf, 0x59, 0xf0, 0xc3, 0x82,
	0xbb, 0x65, 0xcf, 0x53, 0xad, 0xa4, 0x48, 0xd1, 0x07, 0x47, 0xe7, 0x5a, 0x24, 0x63, 0x99, 0x7d,
	0xd6, 0x31, 0xc9, 0xd8, 0xbc, 0x0e, 0x61, 0x0f, 0x3a, 0x8b, 0xad, 0x9a, 0x6a, 0xa1, 0x34, 0x29,
	0xda, 0xbc, 0x8a, 0xf1, 0x11, 0x74, 0x17, 0x5b, 0x75, 0xa8, 0x6d, 0x51, 0xf2, 0x08, 0x98, 0xca,
	0x58, 0x6c, 0x62, 0x9a, 0xe6, 0x82, 0xfa, 0xab, 0x62, 0x7c, 0x02, 0x77, 0xc8, 0x64, 0x54, 0x12,
	0x6c, 0x22, 0x34, 0x41, 0x44, 0xb8, 0x88, 0x4c, 0xb2, 0x4d, 0x93, 0xd0, 0xdb, 0x8c, 0x27, 0x95,
	0xca, 0x95, 0x17, 0x51, 0x45, 0x11, 0xe0, 0x4b, 0x70, 0xd3, 0xda, 0x36, 0xbc, 0x9f, 0x66, 0x12,
	0x67, 0xf8, 0xe0, 0x0f, 0xdb, 0xe2, 0x0d, 0x36, 0x5e, 0x83, 0xad, 0xf3, 0x95, 0xcc, 0xbc, 0xef,
	0x56, 0x21, 0x4a, 0x51, 0xf0, 0x8d, 0x41, 0x67, 0x34, 0x9b, 0x8c, 0xff, 0x71, 0x26, 0xef, 0xce,
	0x9c, 0xc9, 0xe3, 0x9a, 0x71, 0x29, 0xf1, 0xb7, 0x13, 0xf9, 0xdf, 0xd5, 0x7e, 0xb5, 0x00, 0x8c,
	0xcf, 0x0d, 0x5b, 0xeb, 0x00, 0x3a, 0xb1, 0x4e, 0x93, 0xfa, 0x4a, 0xef, 0x9d, 0xf9, 0xb2, 0xbc,
	0x62, 0x0d, 0x67, 0x70, 0xf5, 0x26, 0x12, 0x9c, 0xf2, 0x53, 0xa9, 0x76, 0xcb, 0x85, 0xc4, 0xd7,
	0xd0, 0x2e, 0x0a, 0xf0, 0xe1, 0x99, 0x7b, 0x28, 0x3e, 0x59, 0xef, 0xfa, 0x44, 0xb7, 0x80, 0x83,
	0x5b, 0x21, 0x1b, 0xb0, 0x79, 0x9b, 0x7e, 0x23, 0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x05,
	0x20, 0x96, 0xec, 0x59, 0x04, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package grpc_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type IdReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReq) Reset()         { *m = IdReq{} }
func (m *IdReq) String() string { return proto.CompactTextString(m) }
func (*IdReq) ProtoMessage()    {}
func (*IdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *IdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdReq.Unmarshal(m, b)
}
func (m *IdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdReq.Marshal(b, m, deterministic)
}
func (m *IdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReq.Merge(m, src)
}
func (m *IdReq) XXX_Size() int {
	return xxx_messageInfo_IdReq.Size(m)
}
func (m *IdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReq.DiscardUnknown(m)
}

var xxx_messageInfo_IdReq proto.InternalMessageInfo

func (m *IdReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CustomerCreateResp struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sfid                 string   `protobuf:"bytes,2,opt,name=sfid,proto3" json:"sfid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerCreateResp) Reset()         { *m = CustomerCreateResp{} }
func (m *CustomerCreateResp) String() string { return proto.CompactTextString(m) }
func (*CustomerCreateResp) ProtoMessage()    {}
func (*CustomerCreateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *CustomerCreateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerCreateResp.Unmarshal(m, b)
}
func (m *CustomerCreateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerCreateResp.Marshal(b, m, deterministic)
}
func (m *CustomerCreateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerCreateResp.Merge(m, src)
}
func (m *CustomerCreateResp) XXX_Size() int {
	return xxx_messageInfo_CustomerCreateResp.Size(m)
}
func (m *CustomerCreateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerCreateResp.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerCreateResp proto.InternalMessageInfo

func (m *CustomerCreateResp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerCreateResp) GetSfid() string {
	if m != nil {
		return m.Sfid
	}
	return ""
}

type CustomerCreateReq struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneAreaCode        string   `protobuf:"bytes,3,opt,name=phone_area_code,json=phoneAreaCode,proto3" json:"phone_area_code,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	FirstName            string   `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender               int32    `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday             string   `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerCreateReq) Reset()         { *m = CustomerCreateReq{} }
func (m *CustomerCreateReq) String() string { return proto.CompactTextString(m) }
func (*CustomerCreateReq) ProtoMessage()    {}
func (*CustomerCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *CustomerCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerCreateReq.Unmarshal(m, b)
}
func (m *CustomerCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerCreateReq.Marshal(b, m, deterministic)
}
func (m *CustomerCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerCreateReq.Merge(m, src)
}
func (m *CustomerCreateReq) XXX_Size() int {
	return xxx_messageInfo_CustomerCreateReq.Size(m)
}
func (m *CustomerCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerCreateReq proto.InternalMessageInfo

func (m *CustomerCreateReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *CustomerCreateReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerCreateReq) GetPhoneAreaCode() string {
	if m != nil {
		return m.PhoneAreaCode
	}
	return ""
}

func (m *CustomerCreateReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CustomerCreateReq) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CustomerCreateReq) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CustomerCreateReq) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *CustomerCreateReq) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

type CustomerListReq struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageIndex            int32    `protobuf:"varint,2,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	Keyword              string   `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerListReq) Reset()         { *m = CustomerListReq{} }
func (m *CustomerListReq) String() string { return proto.CompactTextString(m) }
func (*CustomerListReq) ProtoMessage()    {}
func (*CustomerListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{3}
}

func (m *CustomerListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerListReq.Unmarshal(m, b)
}
func (m *CustomerListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerListReq.Marshal(b, m, deterministic)
}
func (m *CustomerListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerListReq.Merge(m, src)
}
func (m *CustomerListReq) XXX_Size() int {
	return xxx_messageInfo_CustomerListReq.Size(m)
}
func (m *CustomerListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerListReq.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerListReq proto.InternalMessageInfo

func (m *CustomerListReq) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *CustomerListReq) GetPageIndex() int32 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *CustomerListReq) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type CustomerListResp struct {
	Data                 []*CustomerListResp_One `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total                int64                   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CustomerListResp) Reset()         { *m = CustomerListResp{} }
func (m *CustomerListResp) String() string { return proto.CompactTextString(m) }
func (*CustomerListResp) ProtoMessage()    {}
func (*CustomerListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{4}
}

func (m *CustomerListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerListResp.Unmarshal(m, b)
}
func (m *CustomerListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerListResp.Marshal(b, m, deterministic)
}
func (m *CustomerListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerListResp.Merge(m, src)
}
func (m *CustomerListResp) XXX_Size() int {
	return xxx_messageInfo_CustomerListResp.Size(m)
}
func (m *CustomerListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerListResp.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerListResp proto.InternalMessageInfo

func (m *CustomerListResp) GetData() []*CustomerListResp_One {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CustomerListResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type CustomerListResp_One struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sfid                 string   `protobuf:"bytes,2,opt,name=sfid,proto3" json:"sfid,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneAreaCode        string   `protobuf:"bytes,5,opt,name=phone_area_code,json=phoneAreaCode,proto3" json:"phone_area_code,omitempty"`
	FirstName            string   `protobuf:"bytes,6,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	CurrentUniversity    string   `protobuf:"bytes,8,opt,name=current_university,json=currentUniversity,proto3" json:"current_university,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerListResp_One) Reset()         { *m = CustomerListResp_One{} }
func (m *CustomerListResp_One) String() string { return proto.CompactTextString(m) }
func (*CustomerListResp_One) ProtoMessage()    {}
func (*CustomerListResp_One) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{4, 0}
}

func (m *CustomerListResp_One) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerListResp_One.Unmarshal(m, b)
}
func (m *CustomerListResp_One) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerListResp_One.Marshal(b, m, deterministic)
}
func (m *CustomerListResp_One) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerListResp_One.Merge(m, src)
}
func (m *CustomerListResp_One) XXX_Size() int {
	return xxx_messageInfo_CustomerListResp_One.Size(m)
}
func (m *CustomerListResp_One) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerListResp_One.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerListResp_One proto.InternalMessageInfo

func (m *CustomerListResp_One) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerListResp_One) GetSfid() string {
	if m != nil {
		return m.Sfid
	}
	return ""
}

func (m *CustomerListResp_One) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *CustomerListResp_One) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerListResp_One) GetPhoneAreaCode() string {
	if m != nil {
		return m.PhoneAreaCode
	}
	return ""
}

func (m *CustomerListResp_One) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CustomerListResp_One) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CustomerListResp_One) GetCurrentUniversity() string {
	if m != nil {
		return m.CurrentUniversity
	}
	return ""
}

func init() {
	proto.RegisterType((*IdReq)(nil), "grpc_pb.IdReq")
	proto.RegisterType((*CustomerCreateResp)(nil), "grpc_pb.CustomerCreateResp")
	proto.RegisterType((*CustomerCreateReq)(nil), "grpc_pb.CustomerCreateReq")
	proto.RegisterType((*CustomerListReq)(nil), "grpc_pb.CustomerListReq")
	proto.RegisterType((*CustomerListResp)(nil), "grpc_pb.CustomerListResp")
	proto.RegisterType((*CustomerListResp_One)(nil), "grpc_pb.CustomerListResp.One")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0xf8, 0x4f, 0xe2, 0x29, 0xa4, 0x74, 0x84, 0x60, 0x71, 0x54, 0x29, 0xf8, 0x80,
	0x72, 0x21, 0x12, 0xed, 0x05, 0x89, 0x53, 0x15, 0x2e, 0x95, 0x50, 0x2b, 0xb9, 0xe2, 0x6c, 0x6d,
	0xbc, 0xd3, 0x74, 0x45, 0xfc, 0x27, 0xbb, 0x9b, 0x42, 0x7b, 0xe2, 0xe9, 0x78, 0x01, 0x9e, 0x85,
	0x3b, 0xf2, 0xda, 0x0e, 0x6a, 0x4a, 0x02, 0xb7, 0xcc, 0xf7, 0xcd, 0x78, 0x33, 0xdf, 0x6f, 0x17,
	0x86, 0xd9, 0x5a, 0x9b, 0x32, 0x27, 0x35, 0xad, 0x54, 0x69, 0x4a, 0xec, 0x2f, 0x54, 0x95, 0xa5,
	0xd5, 0x3c, 0x7e, 0x09, 0xfe, 0xb9, 0x48, 0x68, 0x85, 0x43, 0xe8, 0x49, 0xc1, 0x9c, 0xb1, 0x33,
	0xf1, 0x93, 0x9e, 0x14, 0xf1, 0x7b, 0xc0, 0x59, 0x3b, 0x33, 0x53, 0xc4, 0x0d, 0x25, 0xa4, 0xab,
	0xed, 0x2e, 0x44, 0xf0, 0xf4, 0xb5, 0x14, 0xac, 0x37, 0x76, 0x26, 0x61, 0x62, 0x7f, 0xc7, 0xdf,
	0x7b, 0x70, 0xb4, 0x3d, 0xba, 0xc2, 0xd7, 0xf0, 0xa4, 0xba, 0x29, 0x0b, 0x4a, 0x8b, 0x75, 0x3e,
	0x27, 0x65, 0xbf, 0x11, 0x26, 0x07, 0x56, 0xbb, 0xb0, 0x12, 0x3e, 0x07, 0x9f, 0x72, 0x2e, 0x97,
	0xed, 0xd7, 0x9a, 0x02, 0xdf, 0xc0, 0x61, 0x33, 0xc8, 0x15, 0xf1, 0x34, 0x2b, 0x05, 0x31, 0xd7,
	0xfa, 0x4f, 0xad, 0x7c, 0xa6, 0x88, 0xcf, 0x4a, 0x41, 0x18, 0xc1, 0xa0, 0xe2, 0x5a, 0x7f, 0x2d,
	0x95, 0x60, 0x9e, 0x6d, 0xd8, 0xd4, 0x78, 0x0c, 0x70, 0x2d, 0x95, 0x36, 0x69, 0xc1, 0x73, 0x62,
	0xbe, 0x75, 0x43, 0xab, 0x5c, 0xf0, 0x9c, 0x70, 0x04, 0xe1, 0x92, 0x77, 0x6e, 0xd0, 0xcc, 0xd6,
	0x82, 0x35, 0x5f, 0x40, 0xb0, 0xa0, 0x42, 0x90, 0x62, 0x7d, 0xbb, 0x76, 0x5b, 0xd5, 0xe7, 0xcd,
	0xa5, 0x32, 0x37, 0x82, 0xdf, 0xb1, 0x41, 0x33, 0xd3, 0xd5, 0xf1, 0x02, 0x0e, 0xbb, 0x04, 0x3e,
	0x49, 0x6d, 0xea, 0xfd, 0x47, 0x10, 0x56, 0x7c, 0x41, 0xa9, 0x96, 0xf7, 0xd4, 0x06, 0x38, 0xa8,
	0x85, 0x2b, 0x79, 0x4f, 0xf5, 0xff, 0xb3, 0xa6, 0x2c, 0x04, 0x7d, 0xb3, 0xeb, 0xfb, 0x89, 0x6d,
	0x3f, 0xaf, 0x05, 0x64, 0xd0, 0xff, 0x42, 0x77, 0x76, 0xb3, 0x66, 0xf5, 0xae, 0x8c, 0x7f, 0xf6,
	0xe0, 0xd9, 0xc3, 0x93, 0x74, 0x85, 0xef, 0xc0, 0x13, 0xdc, 0x70, 0xe6, 0x8c, 0xdd, 0xc9, 0xc1,
	0xc9, 0xf1, 0xb4, 0x65, 0x3d, 0xdd, 0x6e, 0x9c, 0x5e, 0x16, 0x94, 0xd8, 0xd6, 0x3a, 0x7a, 0x53,
	0x1a, 0xde, 0x44, 0xef, 0x26, 0x4d, 0x11, 0xfd, 0x72, 0xc0, 0xbd, 0x2c, 0xe8, 0x7f, 0xa8, 0x3f,
	0xe2, 0xeb, 0xee, 0xe1, 0xeb, 0xfd, 0x83, 0xaf, 0xff, 0x37, 0xbe, 0x0f, 0x19, 0x06, 0x7b, 0x19,
	0xf6, 0xb7, 0x18, 0xbe, 0x05, 0xcc, 0xd6, 0x4a, 0x51, 0x61, 0xd2, 0x75, 0x21, 0x6f, 0x49, 0x69,
	0x69, 0x3a, 0x6a, 0x47, 0xad, 0xf3, 0x79, 0x63, 0x9c, 0xfc, 0x70, 0xfe, 0xf0, 0xbb, 0x22, 0x75,
	0x2b, 0x33, 0xc2, 0x33, 0x08, 0x9a, 0xcb, 0x8c, 0xd1, 0xa3, 0x40, 0x37, 0xb7, 0x3c, 0x1a, 0xed,
	0xf4, 0x74, 0x85, 0x1f, 0xc0, 0xab, 0xa3, 0x47, 0xb6, 0x83, 0xc8, 0x2a, 0x7a, 0xb5, 0x93, 0x15,
	0x9e, 0x42, 0xf0, 0x91, 0x96, 0x64, 0x08, 0x87, 0x9b, 0x26, 0xfb, 0x72, 0xf7, 0x0c, 0xcd, 0x03,
	0xfb, 0xda, 0x4f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xec, 0x38, 0xab, 0x41, 0xff, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	Create(ctx context.Context, in *CustomerCreateReq, opts ...grpc.CallOption) (*CustomerCreateResp, error)
	List(ctx context.Context, in *CustomerListReq, opts ...grpc.CallOption) (*CustomerListResp, error)
	Delete(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*CustomerListResp, error)
}

type customerServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerServiceClient(cc *grpc.ClientConn) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) Create(ctx context.Context, in *CustomerCreateReq, opts ...grpc.CallOption) (*CustomerCreateResp, error) {
	out := new(CustomerCreateResp)
	err := c.cc.Invoke(ctx, "/grpc_pb.CustomerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) List(ctx context.Context, in *CustomerListReq, opts ...grpc.CallOption) (*CustomerListResp, error) {
	out := new(CustomerListResp)
	err := c.cc.Invoke(ctx, "/grpc_pb.CustomerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Delete(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*CustomerListResp, error) {
	out := new(CustomerListResp)
	err := c.cc.Invoke(ctx, "/grpc_pb.CustomerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	Create(context.Context, *CustomerCreateReq) (*CustomerCreateResp, error)
	List(context.Context, *CustomerListReq) (*CustomerListResp, error)
	Delete(context.Context, *IdReq) (*CustomerListResp, error)
}

// UnimplementedCustomerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (*UnimplementedCustomerServiceServer) Create(ctx context.Context, req *CustomerCreateReq) (*CustomerCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCustomerServiceServer) List(ctx context.Context, req *CustomerListReq) (*CustomerListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCustomerServiceServer) Delete(ctx context.Context, req *IdReq) (*CustomerListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_pb.CustomerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Create(ctx, req.(*CustomerCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_pb.CustomerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).List(ctx, req.(*CustomerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_pb.CustomerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Delete(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_pb.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CustomerService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CustomerService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CustomerService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
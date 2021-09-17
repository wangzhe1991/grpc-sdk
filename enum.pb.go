// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enum.proto

package grpc_pb

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

// 性别
type Gender int32

const (
	Gender_GENDER_UNKNOW Gender = 0
	Gender_GENDER_MALE   Gender = 1
	Gender_GENDER_FEMALE Gender = 2
)

var Gender_name = map[int32]string{
	0: "GENDER_UNKNOW",
	1: "GENDER_MALE",
	2: "GENDER_FEMALE",
}

var Gender_value = map[string]int32{
	"GENDER_UNKNOW": 0,
	"GENDER_MALE":   1,
	"GENDER_FEMALE": 2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{0}
}

// 会员等级
type CustomerLevel int32

const (
	CustomerLevel_CUSTOMER_LEVEL_ORDINARY CustomerLevel = 0
	CustomerLevel_CUSTOMER_LEVEL_BRONZE   CustomerLevel = 10
	CustomerLevel_CUSTOMER_LEVEL_SILVER   CustomerLevel = 20
	CustomerLevel_CUSTOMER_LEVEL_GOLD     CustomerLevel = 30
)

var CustomerLevel_name = map[int32]string{
	0:  "CUSTOMER_LEVEL_ORDINARY",
	10: "CUSTOMER_LEVEL_BRONZE",
	20: "CUSTOMER_LEVEL_SILVER",
	30: "CUSTOMER_LEVEL_GOLD",
}

var CustomerLevel_value = map[string]int32{
	"CUSTOMER_LEVEL_ORDINARY": 0,
	"CUSTOMER_LEVEL_BRONZE":   10,
	"CUSTOMER_LEVEL_SILVER":   20,
	"CUSTOMER_LEVEL_GOLD":     30,
}

func (x CustomerLevel) String() string {
	return proto.EnumName(CustomerLevel_name, int32(x))
}

func (CustomerLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{1}
}

func init() {
	proto.RegisterEnum("grpc_pb.Gender", Gender_name, Gender_value)
	proto.RegisterEnum("grpc_pb.CustomerLevel", CustomerLevel_name, CustomerLevel_value)
}

func init() { proto.RegisterFile("enum.proto", fileDescriptor_13a9f1b5947140c8) }

var fileDescriptor_13a9f1b5947140c8 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcd, 0x2b, 0xcd,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x48, 0xd2,
	0xb2, 0xe7, 0x62, 0x73, 0x4f, 0xcd, 0x4b, 0x49, 0x2d, 0x12, 0x12, 0xe4, 0xe2, 0x75, 0x77, 0xf5,
	0x73, 0x71, 0x0d, 0x8a, 0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0x17, 0x60, 0x10, 0xe2, 0xe7, 0xe2,
	0x86, 0x0a, 0xf9, 0x3a, 0xfa, 0xb8, 0x0a, 0x30, 0x22, 0xa9, 0x71, 0x73, 0x05, 0x0b, 0x31, 0x69,
	0x55, 0x73, 0xf1, 0x3a, 0x97, 0x16, 0x97, 0xe4, 0xe7, 0xa6, 0x16, 0xf9, 0xa4, 0x96, 0xa5, 0xe6,
	0x08, 0x49, 0x73, 0x89, 0x3b, 0x87, 0x06, 0x87, 0xf8, 0xfb, 0xba, 0x06, 0xc5, 0xfb, 0xb8, 0x86,
	0xb9, 0xfa, 0xc4, 0xfb, 0x07, 0xb9, 0x78, 0xfa, 0x39, 0x06, 0x45, 0x0a, 0x30, 0x08, 0x49, 0x72,
	0x89, 0xa2, 0x49, 0x3a, 0x05, 0xf9, 0xfb, 0x45, 0xb9, 0x0a, 0x70, 0x61, 0x91, 0x0a, 0xf6, 0xf4,
	0x09, 0x73, 0x0d, 0x12, 0x10, 0x11, 0x12, 0xe7, 0x12, 0x46, 0x93, 0x72, 0xf7, 0xf7, 0x71, 0x11,
	0x90, 0x4b, 0x62, 0x03, 0xfb, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xac, 0x9c, 0x34, 0xed,
	0xdb, 0x00, 0x00, 0x00,
}
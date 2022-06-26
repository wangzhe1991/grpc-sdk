// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.0
// source: region.proto

package region

import (
	common "github.com/wangzhe1991/grpc-sdk/pb/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListRegionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*One `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListRegionResp) Reset() {
	*x = ListRegionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRegionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegionResp) ProtoMessage() {}

func (x *ListRegionResp) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegionResp.ProtoReflect.Descriptor instead.
func (*ListRegionResp) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{0}
}

func (x *ListRegionResp) GetData() []*One {
	if x != nil {
		return x.Data
	}
	return nil
}

type One struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pid      int32  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Level    int32  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	CityCode string `protobuf:"bytes,5,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	MerId    string `protobuf:"bytes,6,opt,name=mer_id,json=merId,proto3" json:"mer_id,omitempty"`
	MerName  string `protobuf:"bytes,7,opt,name=mer_name,json=merName,proto3" json:"mer_name,omitempty"`
	Center   string `protobuf:"bytes,8,opt,name=center,proto3" json:"center,omitempty"`
	Data     []*One `protobuf:"bytes,9,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *One) Reset() {
	*x = One{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *One) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*One) ProtoMessage() {}

func (x *One) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use One.ProtoReflect.Descriptor instead.
func (*One) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{1}
}

func (x *One) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *One) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *One) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *One) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *One) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

func (x *One) GetMerId() string {
	if x != nil {
		return x.MerId
	}
	return ""
}

func (x *One) GetMerName() string {
	if x != nil {
		return x.MerName
	}
	return ""
}

func (x *One) GetCenter() string {
	if x != nil {
		return x.Center
	}
	return ""
}

func (x *One) GetData() []*One {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_region_proto protoreflect.FileDescriptor

var file_region_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x1a,
	0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x6e, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe2, 0x01, 0x0a, 0x03, 0x4f, 0x6e,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x6e, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x9e,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x4d, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61,
	0x6e, 0x67, 0x7a, 0x68, 0x65, 0x31, 0x39, 0x39, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_region_proto_rawDescOnce sync.Once
	file_region_proto_rawDescData = file_region_proto_rawDesc
)

func file_region_proto_rawDescGZIP() []byte {
	file_region_proto_rawDescOnce.Do(func() {
		file_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_region_proto_rawDescData)
	})
	return file_region_proto_rawDescData
}

var file_region_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_region_proto_goTypes = []interface{}{
	(*ListRegionResp)(nil),       // 0: grpc_sdk.region.ListRegionResp
	(*One)(nil),                  // 1: grpc_sdk.region.One
	(*common.ListRequest)(nil),   // 2: grpc_sdk.common.ListRequest
	(*common.EmptyRequest)(nil),  // 3: grpc_sdk.common.EmptyRequest
	(*common.EmptyResponse)(nil), // 4: grpc_sdk.common.EmptyResponse
}
var file_region_proto_depIdxs = []int32{
	1, // 0: grpc_sdk.region.ListRegionResp.data:type_name -> grpc_sdk.region.One
	1, // 1: grpc_sdk.region.One.data:type_name -> grpc_sdk.region.One
	2, // 2: grpc_sdk.region.Region.List:input_type -> grpc_sdk.common.ListRequest
	3, // 3: grpc_sdk.region.Region.ManualUpdate:input_type -> grpc_sdk.common.EmptyRequest
	0, // 4: grpc_sdk.region.Region.List:output_type -> grpc_sdk.region.ListRegionResp
	4, // 5: grpc_sdk.region.Region.ManualUpdate:output_type -> grpc_sdk.common.EmptyResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_region_proto_init() }
func file_region_proto_init() {
	if File_region_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_region_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRegionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_region_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*One); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_region_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_region_proto_goTypes,
		DependencyIndexes: file_region_proto_depIdxs,
		MessageInfos:      file_region_proto_msgTypes,
	}.Build()
	File_region_proto = out.File
	file_region_proto_rawDesc = nil
	file_region_proto_goTypes = nil
	file_region_proto_depIdxs = nil
}

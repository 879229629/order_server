// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: order_server/common/v1/enum.proto

package commonv1

import (
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

type OrderStatus int32

const (
	OrderStatus_UNKNOWN    OrderStatus = 0
	OrderStatus_Waiting    OrderStatus = 1
	OrderStatus_Success    OrderStatus = 2
	OrderStatus_Processing OrderStatus = 3
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "Waiting",
		2: "Success",
		3: "Processing",
	}
	OrderStatus_value = map[string]int32{
		"UNKNOWN":    0,
		"Waiting":    1,
		"Success":    2,
		"Processing": 3,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_order_server_common_v1_enum_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_order_server_common_v1_enum_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_order_server_common_v1_enum_proto_rawDescGZIP(), []int{0}
}

var File_order_server_common_v1_enum_proto protoreflect.FileDescriptor

var file_order_server_common_v1_enum_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x44, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x10,
	0x03, 0x42, 0xe9, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x09, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x38, 0x37, 0x39, 0x32, 0x32, 0x39,
	0x36, 0x32, 0x39, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x43, 0x58, 0xaa,
	0x02, 0x15, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x21, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_server_common_v1_enum_proto_rawDescOnce sync.Once
	file_order_server_common_v1_enum_proto_rawDescData = file_order_server_common_v1_enum_proto_rawDesc
)

func file_order_server_common_v1_enum_proto_rawDescGZIP() []byte {
	file_order_server_common_v1_enum_proto_rawDescOnce.Do(func() {
		file_order_server_common_v1_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_server_common_v1_enum_proto_rawDescData)
	})
	return file_order_server_common_v1_enum_proto_rawDescData
}

var file_order_server_common_v1_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_order_server_common_v1_enum_proto_goTypes = []interface{}{
	(OrderStatus)(0), // 0: order_server.common.v1.OrderStatus
}
var file_order_server_common_v1_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_server_common_v1_enum_proto_init() }
func file_order_server_common_v1_enum_proto_init() {
	if File_order_server_common_v1_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_server_common_v1_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_server_common_v1_enum_proto_goTypes,
		DependencyIndexes: file_order_server_common_v1_enum_proto_depIdxs,
		EnumInfos:         file_order_server_common_v1_enum_proto_enumTypes,
	}.Build()
	File_order_server_common_v1_enum_proto = out.File
	file_order_server_common_v1_enum_proto_rawDesc = nil
	file_order_server_common_v1_enum_proto_goTypes = nil
	file_order_server_common_v1_enum_proto_depIdxs = nil
}

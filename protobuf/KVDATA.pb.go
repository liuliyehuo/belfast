// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: KVDATA.proto

package protobuf

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

type KVDATA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *uint32 `protobuf:"varint,1,req,name=key" json:"key,omitempty"`
	Value *uint32 `protobuf:"varint,2,req,name=value" json:"value,omitempty"`
}

func (x *KVDATA) Reset() {
	*x = KVDATA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KVDATA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVDATA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVDATA) ProtoMessage() {}

func (x *KVDATA) ProtoReflect() protoreflect.Message {
	mi := &file_KVDATA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVDATA.ProtoReflect.Descriptor instead.
func (*KVDATA) Descriptor() ([]byte, []int) {
	return file_KVDATA_proto_rawDescGZIP(), []int{0}
}

func (x *KVDATA) GetKey() uint32 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *KVDATA) GetValue() uint32 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

var File_KVDATA_proto protoreflect.FileDescriptor

var file_KVDATA_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x4b, 0x56, 0x44, 0x41, 0x54, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x06, 0x4b, 0x56, 0x44, 0x41, 0x54,
	0x41, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_KVDATA_proto_rawDescOnce sync.Once
	file_KVDATA_proto_rawDescData = file_KVDATA_proto_rawDesc
)

func file_KVDATA_proto_rawDescGZIP() []byte {
	file_KVDATA_proto_rawDescOnce.Do(func() {
		file_KVDATA_proto_rawDescData = protoimpl.X.CompressGZIP(file_KVDATA_proto_rawDescData)
	})
	return file_KVDATA_proto_rawDescData
}

var file_KVDATA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KVDATA_proto_goTypes = []any{
	(*KVDATA)(nil), // 0: belfast.KVDATA
}
var file_KVDATA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_KVDATA_proto_init() }
func file_KVDATA_proto_init() {
	if File_KVDATA_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_KVDATA_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KVDATA); i {
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
			RawDescriptor: file_KVDATA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KVDATA_proto_goTypes,
		DependencyIndexes: file_KVDATA_proto_depIdxs,
		MessageInfos:      file_KVDATA_proto_msgTypes,
	}.Build()
	File_KVDATA_proto = out.File
	file_KVDATA_proto_rawDesc = nil
	file_KVDATA_proto_goTypes = nil
	file_KVDATA_proto_depIdxs = nil
}

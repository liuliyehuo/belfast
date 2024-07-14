// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: ATTRINFO.proto

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

type ATTRINFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level *uint32 `protobuf:"varint,1,req,name=level" json:"level,omitempty"`
	Exp   *uint32 `protobuf:"varint,2,req,name=exp" json:"exp,omitempty"`
}

func (x *ATTRINFO) Reset() {
	*x = ATTRINFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ATTRINFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ATTRINFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ATTRINFO) ProtoMessage() {}

func (x *ATTRINFO) ProtoReflect() protoreflect.Message {
	mi := &file_ATTRINFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ATTRINFO.ProtoReflect.Descriptor instead.
func (*ATTRINFO) Descriptor() ([]byte, []int) {
	return file_ATTRINFO_proto_rawDescGZIP(), []int{0}
}

func (x *ATTRINFO) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *ATTRINFO) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

var File_ATTRINFO_proto protoreflect.FileDescriptor

var file_ATTRINFO_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x41, 0x54, 0x54, 0x52, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x08, 0x41, 0x54, 0x54,
	0x52, 0x49, 0x4e, 0x46, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x78, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_ATTRINFO_proto_rawDescOnce sync.Once
	file_ATTRINFO_proto_rawDescData = file_ATTRINFO_proto_rawDesc
)

func file_ATTRINFO_proto_rawDescGZIP() []byte {
	file_ATTRINFO_proto_rawDescOnce.Do(func() {
		file_ATTRINFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_ATTRINFO_proto_rawDescData)
	})
	return file_ATTRINFO_proto_rawDescData
}

var file_ATTRINFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ATTRINFO_proto_goTypes = []any{
	(*ATTRINFO)(nil), // 0: belfast.ATTRINFO
}
var file_ATTRINFO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ATTRINFO_proto_init() }
func file_ATTRINFO_proto_init() {
	if File_ATTRINFO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ATTRINFO_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ATTRINFO); i {
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
			RawDescriptor: file_ATTRINFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ATTRINFO_proto_goTypes,
		DependencyIndexes: file_ATTRINFO_proto_depIdxs,
		MessageInfos:      file_ATTRINFO_proto_msgTypes,
	}.Build()
	File_ATTRINFO_proto = out.File
	file_ATTRINFO_proto_rawDesc = nil
	file_ATTRINFO_proto_goTypes = nil
	file_ATTRINFO_proto_depIdxs = nil
}

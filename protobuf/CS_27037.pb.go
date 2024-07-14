// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CS_27037.proto

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

type CS_27037 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type_1     *uint32           `protobuf:"varint,1,req,name=type_1,json=type1" json:"type_1,omitempty"`
	Progresses []*CHILD_PROGRESS `protobuf:"bytes,2,rep,name=progresses" json:"progresses,omitempty"`
}

func (x *CS_27037) Reset() {
	*x = CS_27037{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CS_27037_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_27037) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_27037) ProtoMessage() {}

func (x *CS_27037) ProtoReflect() protoreflect.Message {
	mi := &file_CS_27037_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_27037.ProtoReflect.Descriptor instead.
func (*CS_27037) Descriptor() ([]byte, []int) {
	return file_CS_27037_proto_rawDescGZIP(), []int{0}
}

func (x *CS_27037) GetType_1() uint32 {
	if x != nil && x.Type_1 != nil {
		return *x.Type_1
	}
	return 0
}

func (x *CS_27037) GetProgresses() []*CHILD_PROGRESS {
	if x != nil {
		return x.Progresses
	}
	return nil
}

var File_CS_27037_proto protoreflect.FileDescriptor

var file_CS_27037_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x32, 0x37, 0x30, 0x33, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x14, 0x43, 0x48, 0x49, 0x4c, 0x44,
	0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x08, 0x43, 0x53, 0x5f, 0x32, 0x37, 0x30, 0x33, 0x37, 0x12, 0x15, 0x0a, 0x06, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x31, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CS_27037_proto_rawDescOnce sync.Once
	file_CS_27037_proto_rawDescData = file_CS_27037_proto_rawDesc
)

func file_CS_27037_proto_rawDescGZIP() []byte {
	file_CS_27037_proto_rawDescOnce.Do(func() {
		file_CS_27037_proto_rawDescData = protoimpl.X.CompressGZIP(file_CS_27037_proto_rawDescData)
	})
	return file_CS_27037_proto_rawDescData
}

var file_CS_27037_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CS_27037_proto_goTypes = []any{
	(*CS_27037)(nil),       // 0: belfast.CS_27037
	(*CHILD_PROGRESS)(nil), // 1: belfast.CHILD_PROGRESS
}
var file_CS_27037_proto_depIdxs = []int32{
	1, // 0: belfast.CS_27037.progresses:type_name -> belfast.CHILD_PROGRESS
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CS_27037_proto_init() }
func file_CS_27037_proto_init() {
	if File_CS_27037_proto != nil {
		return
	}
	file_CHILD_PROGRESS_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CS_27037_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CS_27037); i {
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
			RawDescriptor: file_CS_27037_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CS_27037_proto_goTypes,
		DependencyIndexes: file_CS_27037_proto_depIdxs,
		MessageInfos:      file_CS_27037_proto_msgTypes,
	}.Build()
	File_CS_27037_proto = out.File
	file_CS_27037_proto_rawDesc = nil
	file_CS_27037_proto_goTypes = nil
	file_CS_27037_proto_depIdxs = nil
}

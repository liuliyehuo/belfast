// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: ACT_TASK_LIST.proto

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

type ACT_TASK_LIST struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActId *uint32     `protobuf:"varint,1,req,name=act_id,json=actId" json:"act_id,omitempty"`
	Tasks []*ACT_TASK `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty"`
}

func (x *ACT_TASK_LIST) Reset() {
	*x = ACT_TASK_LIST{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ACT_TASK_LIST_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACT_TASK_LIST) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACT_TASK_LIST) ProtoMessage() {}

func (x *ACT_TASK_LIST) ProtoReflect() protoreflect.Message {
	mi := &file_ACT_TASK_LIST_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACT_TASK_LIST.ProtoReflect.Descriptor instead.
func (*ACT_TASK_LIST) Descriptor() ([]byte, []int) {
	return file_ACT_TASK_LIST_proto_rawDescGZIP(), []int{0}
}

func (x *ACT_TASK_LIST) GetActId() uint32 {
	if x != nil && x.ActId != nil {
		return *x.ActId
	}
	return 0
}

func (x *ACT_TASK_LIST) GetTasks() []*ACT_TASK {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_ACT_TASK_LIST_proto protoreflect.FileDescriptor

var file_ACT_TASK_LIST_proto_rawDesc = []byte{
	0x0a, 0x13, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0e,
	0x41, 0x43, 0x54, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f,
	0x0a, 0x0d, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x12,
	0x15, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x05, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e,
	0x41, 0x43, 0x54, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_ACT_TASK_LIST_proto_rawDescOnce sync.Once
	file_ACT_TASK_LIST_proto_rawDescData = file_ACT_TASK_LIST_proto_rawDesc
)

func file_ACT_TASK_LIST_proto_rawDescGZIP() []byte {
	file_ACT_TASK_LIST_proto_rawDescOnce.Do(func() {
		file_ACT_TASK_LIST_proto_rawDescData = protoimpl.X.CompressGZIP(file_ACT_TASK_LIST_proto_rawDescData)
	})
	return file_ACT_TASK_LIST_proto_rawDescData
}

var file_ACT_TASK_LIST_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ACT_TASK_LIST_proto_goTypes = []any{
	(*ACT_TASK_LIST)(nil), // 0: belfast.ACT_TASK_LIST
	(*ACT_TASK)(nil),      // 1: belfast.ACT_TASK
}
var file_ACT_TASK_LIST_proto_depIdxs = []int32{
	1, // 0: belfast.ACT_TASK_LIST.tasks:type_name -> belfast.ACT_TASK
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ACT_TASK_LIST_proto_init() }
func file_ACT_TASK_LIST_proto_init() {
	if File_ACT_TASK_LIST_proto != nil {
		return
	}
	file_ACT_TASK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ACT_TASK_LIST_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ACT_TASK_LIST); i {
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
			RawDescriptor: file_ACT_TASK_LIST_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ACT_TASK_LIST_proto_goTypes,
		DependencyIndexes: file_ACT_TASK_LIST_proto_depIdxs,
		MessageInfos:      file_ACT_TASK_LIST_proto_msgTypes,
	}.Build()
	File_ACT_TASK_LIST_proto = out.File
	file_ACT_TASK_LIST_proto_rawDesc = nil
	file_ACT_TASK_LIST_proto_goTypes = nil
	file_ACT_TASK_LIST_proto_depIdxs = nil
}

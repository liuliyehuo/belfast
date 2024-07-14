// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: SHIPCOREINFO.proto

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

type SHIPCOREINFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *uint32          `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Exp       *uint32          `protobuf:"varint,2,req,name=exp" json:"exp,omitempty"`
	ModelList []*SHIPMODELINFO `protobuf:"bytes,3,rep,name=model_list,json=modelList" json:"model_list,omitempty"`
}

func (x *SHIPCOREINFO) Reset() {
	*x = SHIPCOREINFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SHIPCOREINFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SHIPCOREINFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SHIPCOREINFO) ProtoMessage() {}

func (x *SHIPCOREINFO) ProtoReflect() protoreflect.Message {
	mi := &file_SHIPCOREINFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SHIPCOREINFO.ProtoReflect.Descriptor instead.
func (*SHIPCOREINFO) Descriptor() ([]byte, []int) {
	return file_SHIPCOREINFO_proto_rawDescGZIP(), []int{0}
}

func (x *SHIPCOREINFO) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SHIPCOREINFO) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

func (x *SHIPCOREINFO) GetModelList() []*SHIPMODELINFO {
	if x != nil {
		return x.ModelList
	}
	return nil
}

var File_SHIPCOREINFO_proto protoreflect.FileDescriptor

var file_SHIPCOREINFO_proto_rawDesc = []byte{
	0x0a, 0x12, 0x53, 0x48, 0x49, 0x50, 0x43, 0x4f, 0x52, 0x45, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x13, 0x53,
	0x48, 0x49, 0x50, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0c, 0x53, 0x48, 0x49, 0x50, 0x43, 0x4f, 0x52, 0x45, 0x49, 0x4e,
	0x46, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x03, 0x65, 0x78, 0x70, 0x12, 0x35, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61,
	0x73, 0x74, 0x2e, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x49, 0x4e, 0x46, 0x4f,
	0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SHIPCOREINFO_proto_rawDescOnce sync.Once
	file_SHIPCOREINFO_proto_rawDescData = file_SHIPCOREINFO_proto_rawDesc
)

func file_SHIPCOREINFO_proto_rawDescGZIP() []byte {
	file_SHIPCOREINFO_proto_rawDescOnce.Do(func() {
		file_SHIPCOREINFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_SHIPCOREINFO_proto_rawDescData)
	})
	return file_SHIPCOREINFO_proto_rawDescData
}

var file_SHIPCOREINFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SHIPCOREINFO_proto_goTypes = []any{
	(*SHIPCOREINFO)(nil),  // 0: belfast.SHIPCOREINFO
	(*SHIPMODELINFO)(nil), // 1: belfast.SHIPMODELINFO
}
var file_SHIPCOREINFO_proto_depIdxs = []int32{
	1, // 0: belfast.SHIPCOREINFO.model_list:type_name -> belfast.SHIPMODELINFO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SHIPCOREINFO_proto_init() }
func file_SHIPCOREINFO_proto_init() {
	if File_SHIPCOREINFO_proto != nil {
		return
	}
	file_SHIPMODELINFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SHIPCOREINFO_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SHIPCOREINFO); i {
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
			RawDescriptor: file_SHIPCOREINFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SHIPCOREINFO_proto_goTypes,
		DependencyIndexes: file_SHIPCOREINFO_proto_depIdxs,
		MessageInfos:      file_SHIPCOREINFO_proto_msgTypes,
	}.Build()
	File_SHIPCOREINFO_proto = out.File
	file_SHIPCOREINFO_proto_rawDesc = nil
	file_SHIPCOREINFO_proto_goTypes = nil
	file_SHIPCOREINFO_proto_depIdxs = nil
}

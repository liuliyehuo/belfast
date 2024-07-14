// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: SC_50103.proto

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

type SC_50103 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId    *uint32   `protobuf:"varint,1,req,name=ad_id,json=adId" json:"ad_id,omitempty"`
	ArgList []*AD_ARG `protobuf:"bytes,2,rep,name=arg_list,json=argList" json:"arg_list,omitempty"`
}

func (x *SC_50103) Reset() {
	*x = SC_50103{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_50103_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_50103) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_50103) ProtoMessage() {}

func (x *SC_50103) ProtoReflect() protoreflect.Message {
	mi := &file_SC_50103_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_50103.ProtoReflect.Descriptor instead.
func (*SC_50103) Descriptor() ([]byte, []int) {
	return file_SC_50103_proto_rawDescGZIP(), []int{0}
}

func (x *SC_50103) GetAdId() uint32 {
	if x != nil && x.AdId != nil {
		return *x.AdId
	}
	return 0
}

func (x *SC_50103) GetArgList() []*AD_ARG {
	if x != nil {
		return x.ArgList
	}
	return nil
}

var File_SC_50103_proto protoreflect.FileDescriptor

var file_SC_50103_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x35, 0x30, 0x31, 0x30, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0c, 0x41, 0x44, 0x5f, 0x41, 0x52,
	0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x08, 0x53, 0x43, 0x5f, 0x35, 0x30,
	0x31, 0x30, 0x33, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x65, 0x6c,
	0x66, 0x61, 0x73, 0x74, 0x2e, 0x41, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x52, 0x07, 0x61, 0x72, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66,
}

var (
	file_SC_50103_proto_rawDescOnce sync.Once
	file_SC_50103_proto_rawDescData = file_SC_50103_proto_rawDesc
)

func file_SC_50103_proto_rawDescGZIP() []byte {
	file_SC_50103_proto_rawDescOnce.Do(func() {
		file_SC_50103_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_50103_proto_rawDescData)
	})
	return file_SC_50103_proto_rawDescData
}

var file_SC_50103_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_50103_proto_goTypes = []any{
	(*SC_50103)(nil), // 0: belfast.SC_50103
	(*AD_ARG)(nil),   // 1: belfast.AD_ARG
}
var file_SC_50103_proto_depIdxs = []int32{
	1, // 0: belfast.SC_50103.arg_list:type_name -> belfast.AD_ARG
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SC_50103_proto_init() }
func file_SC_50103_proto_init() {
	if File_SC_50103_proto != nil {
		return
	}
	file_AD_ARG_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_50103_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SC_50103); i {
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
			RawDescriptor: file_SC_50103_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_50103_proto_goTypes,
		DependencyIndexes: file_SC_50103_proto_depIdxs,
		MessageInfos:      file_SC_50103_proto_msgTypes,
	}.Build()
	File_SC_50103_proto = out.File
	file_SC_50103_proto_rawDesc = nil
	file_SC_50103_proto_goTypes = nil
	file_SC_50103_proto_depIdxs = nil
}

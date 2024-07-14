// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: COMMANDERHOMESLOT.proto

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

type COMMANDERHOMESLOT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	OpFlag         *uint32 `protobuf:"varint,2,req,name=op_flag,json=opFlag" json:"op_flag,omitempty"`
	ExpTime        *uint32 `protobuf:"varint,3,req,name=exp_time,json=expTime" json:"exp_time,omitempty"`
	CommanderId    *uint32 `protobuf:"varint,4,req,name=commander_id,json=commanderId" json:"commander_id,omitempty"`
	Style          *uint32 `protobuf:"varint,5,req,name=style" json:"style,omitempty"`
	CommanderLevel *uint32 `protobuf:"varint,6,opt,name=commander_level,json=commanderLevel" json:"commander_level,omitempty"`
	CommanderExp   *uint32 `protobuf:"varint,7,opt,name=commander_exp,json=commanderExp" json:"commander_exp,omitempty"`
	CacheExp       *uint32 `protobuf:"varint,8,req,name=cache_exp,json=cacheExp" json:"cache_exp,omitempty"`
}

func (x *COMMANDERHOMESLOT) Reset() {
	*x = COMMANDERHOMESLOT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_COMMANDERHOMESLOT_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COMMANDERHOMESLOT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COMMANDERHOMESLOT) ProtoMessage() {}

func (x *COMMANDERHOMESLOT) ProtoReflect() protoreflect.Message {
	mi := &file_COMMANDERHOMESLOT_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COMMANDERHOMESLOT.ProtoReflect.Descriptor instead.
func (*COMMANDERHOMESLOT) Descriptor() ([]byte, []int) {
	return file_COMMANDERHOMESLOT_proto_rawDescGZIP(), []int{0}
}

func (x *COMMANDERHOMESLOT) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *COMMANDERHOMESLOT) GetOpFlag() uint32 {
	if x != nil && x.OpFlag != nil {
		return *x.OpFlag
	}
	return 0
}

func (x *COMMANDERHOMESLOT) GetExpTime() uint32 {
	if x != nil && x.ExpTime != nil {
		return *x.ExpTime
	}
	return 0
}

func (x *COMMANDERHOMESLOT) GetCommanderId() uint32 {
	if x != nil && x.CommanderId != nil {
		return *x.CommanderId
	}
	return 0
}

func (x *COMMANDERHOMESLOT) GetStyle() uint32 {
	if x != nil && x.Style != nil {
		return *x.Style
	}
	return 0
}

func (x *COMMANDERHOMESLOT) GetCommanderLevel() uint32 {
	if x != nil && x.CommanderLevel != nil {
		return *x.CommanderLevel
	}
	return 0
}

func (x *COMMANDERHOMESLOT) GetCommanderExp() uint32 {
	if x != nil && x.CommanderExp != nil {
		return *x.CommanderExp
	}
	return 0
}

func (x *COMMANDERHOMESLOT) GetCacheExp() uint32 {
	if x != nil && x.CacheExp != nil {
		return *x.CacheExp
	}
	return 0
}

var File_COMMANDERHOMESLOT_proto protoreflect.FileDescriptor

var file_COMMANDERHOMESLOT_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x45, 0x52, 0x48, 0x4f, 0x4d, 0x45, 0x53,
	0x4c, 0x4f, 0x54, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61,
	0x73, 0x74, 0x22, 0xfb, 0x01, 0x0a, 0x11, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x45, 0x52,
	0x48, 0x4f, 0x4d, 0x45, 0x53, 0x4c, 0x4f, 0x54, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x66,
	0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x70, 0x46, 0x6c, 0x61,
	0x67, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x78, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x45, 0x78, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x65, 0x78, 0x70,
	0x18, 0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x45, 0x78, 0x70,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_COMMANDERHOMESLOT_proto_rawDescOnce sync.Once
	file_COMMANDERHOMESLOT_proto_rawDescData = file_COMMANDERHOMESLOT_proto_rawDesc
)

func file_COMMANDERHOMESLOT_proto_rawDescGZIP() []byte {
	file_COMMANDERHOMESLOT_proto_rawDescOnce.Do(func() {
		file_COMMANDERHOMESLOT_proto_rawDescData = protoimpl.X.CompressGZIP(file_COMMANDERHOMESLOT_proto_rawDescData)
	})
	return file_COMMANDERHOMESLOT_proto_rawDescData
}

var file_COMMANDERHOMESLOT_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_COMMANDERHOMESLOT_proto_goTypes = []any{
	(*COMMANDERHOMESLOT)(nil), // 0: belfast.COMMANDERHOMESLOT
}
var file_COMMANDERHOMESLOT_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_COMMANDERHOMESLOT_proto_init() }
func file_COMMANDERHOMESLOT_proto_init() {
	if File_COMMANDERHOMESLOT_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_COMMANDERHOMESLOT_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*COMMANDERHOMESLOT); i {
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
			RawDescriptor: file_COMMANDERHOMESLOT_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_COMMANDERHOMESLOT_proto_goTypes,
		DependencyIndexes: file_COMMANDERHOMESLOT_proto_depIdxs,
		MessageInfos:      file_COMMANDERHOMESLOT_proto_msgTypes,
	}.Build()
	File_COMMANDERHOMESLOT_proto = out.File
	file_COMMANDERHOMESLOT_proto_rawDesc = nil
	file_COMMANDERHOMESLOT_proto_goTypes = nil
	file_COMMANDERHOMESLOT_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CS_28009.proto

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

type CS_28009 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipGroup *uint32                `protobuf:"varint,1,req,name=ship_group,json=shipGroup" json:"ship_group,omitempty"`
	Gifts     []*APARTMENT_GIVE_GIFT `protobuf:"bytes,2,rep,name=gifts" json:"gifts,omitempty"`
}

func (x *CS_28009) Reset() {
	*x = CS_28009{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CS_28009_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_28009) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_28009) ProtoMessage() {}

func (x *CS_28009) ProtoReflect() protoreflect.Message {
	mi := &file_CS_28009_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_28009.ProtoReflect.Descriptor instead.
func (*CS_28009) Descriptor() ([]byte, []int) {
	return file_CS_28009_proto_rawDescGZIP(), []int{0}
}

func (x *CS_28009) GetShipGroup() uint32 {
	if x != nil && x.ShipGroup != nil {
		return *x.ShipGroup
	}
	return 0
}

func (x *CS_28009) GetGifts() []*APARTMENT_GIVE_GIFT {
	if x != nil {
		return x.Gifts
	}
	return nil
}

var File_CS_28009_proto protoreflect.FileDescriptor

var file_CS_28009_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x32, 0x38, 0x30, 0x30, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x19, 0x41, 0x50, 0x41, 0x52, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x49, 0x56, 0x45, 0x5f, 0x47, 0x49, 0x46, 0x54, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x08, 0x43, 0x53, 0x5f, 0x32, 0x38, 0x30, 0x30, 0x39,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x32, 0x0a, 0x05, 0x67, 0x69, 0x66, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x41, 0x50, 0x41, 0x52, 0x54, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x47, 0x49, 0x56, 0x45, 0x5f, 0x47, 0x49, 0x46, 0x54, 0x52, 0x05, 0x67, 0x69,
	0x66, 0x74, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66,
}

var (
	file_CS_28009_proto_rawDescOnce sync.Once
	file_CS_28009_proto_rawDescData = file_CS_28009_proto_rawDesc
)

func file_CS_28009_proto_rawDescGZIP() []byte {
	file_CS_28009_proto_rawDescOnce.Do(func() {
		file_CS_28009_proto_rawDescData = protoimpl.X.CompressGZIP(file_CS_28009_proto_rawDescData)
	})
	return file_CS_28009_proto_rawDescData
}

var file_CS_28009_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CS_28009_proto_goTypes = []any{
	(*CS_28009)(nil),            // 0: belfast.CS_28009
	(*APARTMENT_GIVE_GIFT)(nil), // 1: belfast.APARTMENT_GIVE_GIFT
}
var file_CS_28009_proto_depIdxs = []int32{
	1, // 0: belfast.CS_28009.gifts:type_name -> belfast.APARTMENT_GIVE_GIFT
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CS_28009_proto_init() }
func file_CS_28009_proto_init() {
	if File_CS_28009_proto != nil {
		return
	}
	file_APARTMENT_GIVE_GIFT_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CS_28009_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CS_28009); i {
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
			RawDescriptor: file_CS_28009_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CS_28009_proto_goTypes,
		DependencyIndexes: file_CS_28009_proto_depIdxs,
		MessageInfos:      file_CS_28009_proto_msgTypes,
	}.Build()
	File_CS_28009_proto = out.File
	file_CS_28009_proto_rawDesc = nil
	file_CS_28009_proto_goTypes = nil
	file_CS_28009_proto_depIdxs = nil
}

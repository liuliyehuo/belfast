// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: SC_63004.proto

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

type SC_63004 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result         *uint32              `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	CommonList     []*DROPINFO          `protobuf:"bytes,2,rep,name=common_list,json=commonList" json:"common_list,omitempty"`
	RefreshList    []*TECHNOLOGYREFRESH `protobuf:"bytes,3,rep,name=refresh_list,json=refreshList" json:"refresh_list,omitempty"`
	DropList       []*DROPINFO          `protobuf:"bytes,4,rep,name=drop_list,json=dropList" json:"drop_list,omitempty"`
	CatchupList    []*DROPINFO          `protobuf:"bytes,5,rep,name=catchup_list,json=catchupList" json:"catchup_list,omitempty"`
	CatchupactList []*DROPINFO          `protobuf:"bytes,6,rep,name=catchupact_list,json=catchupactList" json:"catchupact_list,omitempty"`
}

func (x *SC_63004) Reset() {
	*x = SC_63004{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_63004_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_63004) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_63004) ProtoMessage() {}

func (x *SC_63004) ProtoReflect() protoreflect.Message {
	mi := &file_SC_63004_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_63004.ProtoReflect.Descriptor instead.
func (*SC_63004) Descriptor() ([]byte, []int) {
	return file_SC_63004_proto_rawDescGZIP(), []int{0}
}

func (x *SC_63004) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *SC_63004) GetCommonList() []*DROPINFO {
	if x != nil {
		return x.CommonList
	}
	return nil
}

func (x *SC_63004) GetRefreshList() []*TECHNOLOGYREFRESH {
	if x != nil {
		return x.RefreshList
	}
	return nil
}

func (x *SC_63004) GetDropList() []*DROPINFO {
	if x != nil {
		return x.DropList
	}
	return nil
}

func (x *SC_63004) GetCatchupList() []*DROPINFO {
	if x != nil {
		return x.CatchupList
	}
	return nil
}

func (x *SC_63004) GetCatchupactList() []*DROPINFO {
	if x != nil {
		return x.CatchupactList
	}
	return nil
}

var File_SC_63004_proto protoreflect.FileDescriptor

var file_SC_63004_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x36, 0x33, 0x30, 0x30, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0e, 0x44, 0x52, 0x4f, 0x50, 0x49,
	0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x54, 0x45, 0x43, 0x48, 0x4e,
	0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x08, 0x53, 0x43, 0x5f, 0x36, 0x33, 0x30, 0x30, 0x34, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62,
	0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x44, 0x52, 0x4f, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x54, 0x45, 0x43, 0x48,
	0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x52, 0x0b, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x64, 0x72,
	0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x44, 0x52, 0x4f, 0x50, 0x49, 0x4e, 0x46, 0x4f,
	0x52, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x63, 0x61,
	0x74, 0x63, 0x68, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x44, 0x52, 0x4f, 0x50, 0x49,
	0x4e, 0x46, 0x4f, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x0f, 0x63, 0x61, 0x74, 0x63, 0x68, 0x75, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66,
	0x61, 0x73, 0x74, 0x2e, 0x44, 0x52, 0x4f, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x0e, 0x63, 0x61,
	0x74, 0x63, 0x68, 0x75, 0x70, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_63004_proto_rawDescOnce sync.Once
	file_SC_63004_proto_rawDescData = file_SC_63004_proto_rawDesc
)

func file_SC_63004_proto_rawDescGZIP() []byte {
	file_SC_63004_proto_rawDescOnce.Do(func() {
		file_SC_63004_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_63004_proto_rawDescData)
	})
	return file_SC_63004_proto_rawDescData
}

var file_SC_63004_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_63004_proto_goTypes = []any{
	(*SC_63004)(nil),          // 0: belfast.SC_63004
	(*DROPINFO)(nil),          // 1: belfast.DROPINFO
	(*TECHNOLOGYREFRESH)(nil), // 2: belfast.TECHNOLOGYREFRESH
}
var file_SC_63004_proto_depIdxs = []int32{
	1, // 0: belfast.SC_63004.common_list:type_name -> belfast.DROPINFO
	2, // 1: belfast.SC_63004.refresh_list:type_name -> belfast.TECHNOLOGYREFRESH
	1, // 2: belfast.SC_63004.drop_list:type_name -> belfast.DROPINFO
	1, // 3: belfast.SC_63004.catchup_list:type_name -> belfast.DROPINFO
	1, // 4: belfast.SC_63004.catchupact_list:type_name -> belfast.DROPINFO
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_SC_63004_proto_init() }
func file_SC_63004_proto_init() {
	if File_SC_63004_proto != nil {
		return
	}
	file_DROPINFO_proto_init()
	file_TECHNOLOGYREFRESH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_63004_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SC_63004); i {
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
			RawDescriptor: file_SC_63004_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_63004_proto_goTypes,
		DependencyIndexes: file_SC_63004_proto_depIdxs,
		MessageInfos:      file_SC_63004_proto_msgTypes,
	}.Build()
	File_SC_63004_proto = out.File
	file_SC_63004_proto_rawDesc = nil
	file_SC_63004_proto_goTypes = nil
	file_SC_63004_proto_depIdxs = nil
}

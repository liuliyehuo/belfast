// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: USERCHALLENGEINFO.proto

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

type USERCHALLENGEINFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentScore  *uint32                 `protobuf:"varint,1,req,name=current_score,json=currentScore" json:"current_score,omitempty"`
	Level         *uint32                 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	GroupincList  []*GROUPINFOINCHALLENGE `protobuf:"bytes,3,rep,name=groupinc_list,json=groupincList" json:"groupinc_list,omitempty"`
	Mode          *uint32                 `protobuf:"varint,4,req,name=mode" json:"mode,omitempty"`
	Issl          *uint32                 `protobuf:"varint,5,req,name=issl" json:"issl,omitempty"`
	SeasonId      *uint32                 `protobuf:"varint,6,req,name=season_id,json=seasonId" json:"season_id,omitempty"`
	DungeonIdList []uint32                `protobuf:"varint,7,rep,name=dungeon_id_list,json=dungeonIdList" json:"dungeon_id_list,omitempty"`
	BuffList      []uint32                `protobuf:"varint,8,rep,name=buff_list,json=buffList" json:"buff_list,omitempty"`
}

func (x *USERCHALLENGEINFO) Reset() {
	*x = USERCHALLENGEINFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_USERCHALLENGEINFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *USERCHALLENGEINFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*USERCHALLENGEINFO) ProtoMessage() {}

func (x *USERCHALLENGEINFO) ProtoReflect() protoreflect.Message {
	mi := &file_USERCHALLENGEINFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use USERCHALLENGEINFO.ProtoReflect.Descriptor instead.
func (*USERCHALLENGEINFO) Descriptor() ([]byte, []int) {
	return file_USERCHALLENGEINFO_proto_rawDescGZIP(), []int{0}
}

func (x *USERCHALLENGEINFO) GetCurrentScore() uint32 {
	if x != nil && x.CurrentScore != nil {
		return *x.CurrentScore
	}
	return 0
}

func (x *USERCHALLENGEINFO) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *USERCHALLENGEINFO) GetGroupincList() []*GROUPINFOINCHALLENGE {
	if x != nil {
		return x.GroupincList
	}
	return nil
}

func (x *USERCHALLENGEINFO) GetMode() uint32 {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return 0
}

func (x *USERCHALLENGEINFO) GetIssl() uint32 {
	if x != nil && x.Issl != nil {
		return *x.Issl
	}
	return 0
}

func (x *USERCHALLENGEINFO) GetSeasonId() uint32 {
	if x != nil && x.SeasonId != nil {
		return *x.SeasonId
	}
	return 0
}

func (x *USERCHALLENGEINFO) GetDungeonIdList() []uint32 {
	if x != nil {
		return x.DungeonIdList
	}
	return nil
}

func (x *USERCHALLENGEINFO) GetBuffList() []uint32 {
	if x != nil {
		return x.BuffList
	}
	return nil
}

var File_USERCHALLENGEINFO_proto protoreflect.FileDescriptor

var file_USERCHALLENGEINFO_proto_rawDesc = []byte{
	0x0a, 0x17, 0x55, 0x53, 0x45, 0x52, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x49,
	0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61,
	0x73, 0x74, 0x1a, 0x1a, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x49, 0x4e, 0x43,
	0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x02, 0x0a, 0x11, 0x55, 0x53, 0x45, 0x52, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45,
	0x49, 0x4e, 0x46, 0x4f, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x42, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x49, 0x4e, 0x43, 0x48, 0x41, 0x4c,
	0x4c, 0x45, 0x4e, 0x47, 0x45, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x63, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x73, 0x6c, 0x18,
	0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x69, 0x73, 0x73, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0d, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_USERCHALLENGEINFO_proto_rawDescOnce sync.Once
	file_USERCHALLENGEINFO_proto_rawDescData = file_USERCHALLENGEINFO_proto_rawDesc
)

func file_USERCHALLENGEINFO_proto_rawDescGZIP() []byte {
	file_USERCHALLENGEINFO_proto_rawDescOnce.Do(func() {
		file_USERCHALLENGEINFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_USERCHALLENGEINFO_proto_rawDescData)
	})
	return file_USERCHALLENGEINFO_proto_rawDescData
}

var file_USERCHALLENGEINFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_USERCHALLENGEINFO_proto_goTypes = []any{
	(*USERCHALLENGEINFO)(nil),    // 0: belfast.USERCHALLENGEINFO
	(*GROUPINFOINCHALLENGE)(nil), // 1: belfast.GROUPINFOINCHALLENGE
}
var file_USERCHALLENGEINFO_proto_depIdxs = []int32{
	1, // 0: belfast.USERCHALLENGEINFO.groupinc_list:type_name -> belfast.GROUPINFOINCHALLENGE
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_USERCHALLENGEINFO_proto_init() }
func file_USERCHALLENGEINFO_proto_init() {
	if File_USERCHALLENGEINFO_proto != nil {
		return
	}
	file_GROUPINFOINCHALLENGE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_USERCHALLENGEINFO_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*USERCHALLENGEINFO); i {
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
			RawDescriptor: file_USERCHALLENGEINFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_USERCHALLENGEINFO_proto_goTypes,
		DependencyIndexes: file_USERCHALLENGEINFO_proto_depIdxs,
		MessageInfos:      file_USERCHALLENGEINFO_proto_msgTypes,
	}.Build()
	File_USERCHALLENGEINFO_proto = out.File
	file_USERCHALLENGEINFO_proto_rawDesc = nil
	file_USERCHALLENGEINFO_proto_goTypes = nil
	file_USERCHALLENGEINFO_proto_depIdxs = nil
}

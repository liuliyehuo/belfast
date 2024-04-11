// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_40004.proto

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

type SC_40004 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result        *uint32          `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	DropInfo      []*DROPINFO      `protobuf:"bytes,2,rep,name=drop_info,json=dropInfo" json:"drop_info,omitempty"`
	ExtraDropInfo []*DROPINFO      `protobuf:"bytes,3,rep,name=extra_drop_info,json=extraDropInfo" json:"extra_drop_info,omitempty"`
	PlayerExp     *uint32          `protobuf:"varint,4,req,name=player_exp,json=playerExp" json:"player_exp,omitempty"`
	ShipExpList   []*SHIP_EXP      `protobuf:"bytes,5,rep,name=ship_exp_list,json=shipExpList" json:"ship_exp_list,omitempty"`
	Mvp           *uint32          `protobuf:"varint,6,req,name=mvp" json:"mvp,omitempty"`
	CommanderExp  []*COMMANDER_EXP `protobuf:"bytes,7,rep,name=commander_exp,json=commanderExp" json:"commander_exp,omitempty"`
	HpDropInfo    []*HPDROPINFO    `protobuf:"bytes,8,rep,name=hp_drop_info,json=hpDropInfo" json:"hp_drop_info,omitempty"`
}

func (x *SC_40004) Reset() {
	*x = SC_40004{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_40004_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_40004) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_40004) ProtoMessage() {}

func (x *SC_40004) ProtoReflect() protoreflect.Message {
	mi := &file_SC_40004_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_40004.ProtoReflect.Descriptor instead.
func (*SC_40004) Descriptor() ([]byte, []int) {
	return file_SC_40004_proto_rawDescGZIP(), []int{0}
}

func (x *SC_40004) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *SC_40004) GetDropInfo() []*DROPINFO {
	if x != nil {
		return x.DropInfo
	}
	return nil
}

func (x *SC_40004) GetExtraDropInfo() []*DROPINFO {
	if x != nil {
		return x.ExtraDropInfo
	}
	return nil
}

func (x *SC_40004) GetPlayerExp() uint32 {
	if x != nil && x.PlayerExp != nil {
		return *x.PlayerExp
	}
	return 0
}

func (x *SC_40004) GetShipExpList() []*SHIP_EXP {
	if x != nil {
		return x.ShipExpList
	}
	return nil
}

func (x *SC_40004) GetMvp() uint32 {
	if x != nil && x.Mvp != nil {
		return *x.Mvp
	}
	return 0
}

func (x *SC_40004) GetCommanderExp() []*COMMANDER_EXP {
	if x != nil {
		return x.CommanderExp
	}
	return nil
}

func (x *SC_40004) GetHpDropInfo() []*HPDROPINFO {
	if x != nil {
		return x.HpDropInfo
	}
	return nil
}

var File_SC_40004_proto protoreflect.FileDescriptor

var file_SC_40004_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x34, 0x30, 0x30, 0x30, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0e, 0x44, 0x52, 0x4f, 0x50, 0x49,
	0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x53, 0x48, 0x49, 0x50, 0x5f,
	0x45, 0x58, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x4e, 0x44, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x48, 0x50, 0x44, 0x52, 0x4f, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe9, 0x02, 0x0a, 0x08, 0x53, 0x43, 0x5f, 0x34, 0x30, 0x30, 0x30, 0x34, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61,
	0x73, 0x74, 0x2e, 0x44, 0x52, 0x4f, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x08, 0x64, 0x72, 0x6f,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64,
	0x72, 0x6f, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x44, 0x52, 0x4f, 0x50, 0x49, 0x4e, 0x46,
	0x4f, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x72, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x04,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x12,
	0x35, 0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x45,
	0x78, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x76, 0x70, 0x18, 0x06, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x76, 0x70, 0x12, 0x3b, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x65, 0x72, 0x45, 0x78, 0x70, 0x12, 0x35, 0x0a, 0x0c, 0x68, 0x70, 0x5f, 0x64, 0x72, 0x6f, 0x70,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65,
	0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x48, 0x50, 0x44, 0x52, 0x4f, 0x50, 0x49, 0x4e, 0x46, 0x4f,
	0x52, 0x0a, 0x68, 0x70, 0x44, 0x72, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_40004_proto_rawDescOnce sync.Once
	file_SC_40004_proto_rawDescData = file_SC_40004_proto_rawDesc
)

func file_SC_40004_proto_rawDescGZIP() []byte {
	file_SC_40004_proto_rawDescOnce.Do(func() {
		file_SC_40004_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_40004_proto_rawDescData)
	})
	return file_SC_40004_proto_rawDescData
}

var file_SC_40004_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_40004_proto_goTypes = []interface{}{
	(*SC_40004)(nil),      // 0: belfast.SC_40004
	(*DROPINFO)(nil),      // 1: belfast.DROPINFO
	(*SHIP_EXP)(nil),      // 2: belfast.SHIP_EXP
	(*COMMANDER_EXP)(nil), // 3: belfast.COMMANDER_EXP
	(*HPDROPINFO)(nil),    // 4: belfast.HPDROPINFO
}
var file_SC_40004_proto_depIdxs = []int32{
	1, // 0: belfast.SC_40004.drop_info:type_name -> belfast.DROPINFO
	1, // 1: belfast.SC_40004.extra_drop_info:type_name -> belfast.DROPINFO
	2, // 2: belfast.SC_40004.ship_exp_list:type_name -> belfast.SHIP_EXP
	3, // 3: belfast.SC_40004.commander_exp:type_name -> belfast.COMMANDER_EXP
	4, // 4: belfast.SC_40004.hp_drop_info:type_name -> belfast.HPDROPINFO
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_SC_40004_proto_init() }
func file_SC_40004_proto_init() {
	if File_SC_40004_proto != nil {
		return
	}
	file_DROPINFO_proto_init()
	file_SHIP_EXP_proto_init()
	file_COMMANDER_EXP_proto_init()
	file_HPDROPINFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_40004_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_40004); i {
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
			RawDescriptor: file_SC_40004_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_40004_proto_goTypes,
		DependencyIndexes: file_SC_40004_proto_depIdxs,
		MessageInfos:      file_SC_40004_proto_msgTypes,
	}.Build()
	File_SC_40004_proto = out.File
	file_SC_40004_proto_rawDesc = nil
	file_SC_40004_proto_goTypes = nil
	file_SC_40004_proto_depIdxs = nil
}
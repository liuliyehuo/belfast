// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: FRIENDSCORE.proto

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

type FRIENDSCORE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *uint32      `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name     *string      `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Score    *uint32      `protobuf:"varint,3,req,name=score" json:"score,omitempty"`
	Display  *DISPLAYINFO `protobuf:"bytes,4,req,name=display" json:"display,omitempty"`
	TimeData *uint32      `protobuf:"varint,5,req,name=time_data,json=timeData" json:"time_data,omitempty"`
}

func (x *FRIENDSCORE) Reset() {
	*x = FRIENDSCORE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FRIENDSCORE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FRIENDSCORE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FRIENDSCORE) ProtoMessage() {}

func (x *FRIENDSCORE) ProtoReflect() protoreflect.Message {
	mi := &file_FRIENDSCORE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FRIENDSCORE.ProtoReflect.Descriptor instead.
func (*FRIENDSCORE) Descriptor() ([]byte, []int) {
	return file_FRIENDSCORE_proto_rawDescGZIP(), []int{0}
}

func (x *FRIENDSCORE) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *FRIENDSCORE) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *FRIENDSCORE) GetScore() uint32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *FRIENDSCORE) GetDisplay() *DISPLAYINFO {
	if x != nil {
		return x.Display
	}
	return nil
}

func (x *FRIENDSCORE) GetTimeData() uint32 {
	if x != nil && x.TimeData != nil {
		return *x.TimeData
	}
	return 0
}

var File_FRIENDSCORE_proto protoreflect.FileDescriptor

var file_FRIENDSCORE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x11, 0x44, 0x49,
	0x53, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x01, 0x0a, 0x0b, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x65, 0x6c,
	0x66, 0x61, 0x73, 0x74, 0x2e, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x46, 0x4f,
	0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66,
}

var (
	file_FRIENDSCORE_proto_rawDescOnce sync.Once
	file_FRIENDSCORE_proto_rawDescData = file_FRIENDSCORE_proto_rawDesc
)

func file_FRIENDSCORE_proto_rawDescGZIP() []byte {
	file_FRIENDSCORE_proto_rawDescOnce.Do(func() {
		file_FRIENDSCORE_proto_rawDescData = protoimpl.X.CompressGZIP(file_FRIENDSCORE_proto_rawDescData)
	})
	return file_FRIENDSCORE_proto_rawDescData
}

var file_FRIENDSCORE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FRIENDSCORE_proto_goTypes = []any{
	(*FRIENDSCORE)(nil), // 0: belfast.FRIENDSCORE
	(*DISPLAYINFO)(nil), // 1: belfast.DISPLAYINFO
}
var file_FRIENDSCORE_proto_depIdxs = []int32{
	1, // 0: belfast.FRIENDSCORE.display:type_name -> belfast.DISPLAYINFO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FRIENDSCORE_proto_init() }
func file_FRIENDSCORE_proto_init() {
	if File_FRIENDSCORE_proto != nil {
		return
	}
	file_DISPLAYINFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FRIENDSCORE_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FRIENDSCORE); i {
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
			RawDescriptor: file_FRIENDSCORE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FRIENDSCORE_proto_goTypes,
		DependencyIndexes: file_FRIENDSCORE_proto_depIdxs,
		MessageInfos:      file_FRIENDSCORE_proto_msgTypes,
	}.Build()
	File_FRIENDSCORE_proto = out.File
	file_FRIENDSCORE_proto_rawDesc = nil
	file_FRIENDSCORE_proto_goTypes = nil
	file_FRIENDSCORE_proto_depIdxs = nil
}

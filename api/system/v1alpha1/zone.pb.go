// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: system/v1alpha1/zone.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Zone defines the Zone configuration used at the Global Control Plane
// within a distributed deployment
type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enable allows to turn the zone on/off and exclude the whole zone from
	// balancing traffic on it
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Zone) Reset() {
	*x = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1alpha1_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zone) ProtoMessage() {}

func (x *Zone) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1alpha1_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zone.ProtoReflect.Descriptor instead.
func (*Zone) Descriptor() ([]byte, []int) {
	return file_system_v1alpha1_zone_proto_rawDescGZIP(), []int{0}
}

func (x *Zone) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

var File_system_v1alpha1_zone_proto protoreflect.FileDescriptor

var file_system_v1alpha1_zone_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6b, 0x75,
	0x6d, 0x61, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_v1alpha1_zone_proto_rawDescOnce sync.Once
	file_system_v1alpha1_zone_proto_rawDescData = file_system_v1alpha1_zone_proto_rawDesc
)

func file_system_v1alpha1_zone_proto_rawDescGZIP() []byte {
	file_system_v1alpha1_zone_proto_rawDescOnce.Do(func() {
		file_system_v1alpha1_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_v1alpha1_zone_proto_rawDescData)
	})
	return file_system_v1alpha1_zone_proto_rawDescData
}

var file_system_v1alpha1_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_system_v1alpha1_zone_proto_goTypes = []interface{}{
	(*Zone)(nil),               // 0: kuma.system.v1alpha1.Zone
	(*wrappers.BoolValue)(nil), // 1: google.protobuf.BoolValue
}
var file_system_v1alpha1_zone_proto_depIdxs = []int32{
	1, // 0: kuma.system.v1alpha1.Zone.enabled:type_name -> google.protobuf.BoolValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_system_v1alpha1_zone_proto_init() }
func file_system_v1alpha1_zone_proto_init() {
	if File_system_v1alpha1_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_v1alpha1_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone); i {
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
			RawDescriptor: file_system_v1alpha1_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_system_v1alpha1_zone_proto_goTypes,
		DependencyIndexes: file_system_v1alpha1_zone_proto_depIdxs,
		MessageInfos:      file_system_v1alpha1_zone_proto_msgTypes,
	}.Build()
	File_system_v1alpha1_zone_proto = out.File
	file_system_v1alpha1_zone_proto_rawDesc = nil
	file_system_v1alpha1_zone_proto_goTypes = nil
	file_system_v1alpha1_zone_proto_depIdxs = nil
}

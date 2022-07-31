// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: examplepb/example.proto

package examplepb

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

type Msg1_Enum1 int32

const (
	Msg1_E1 Msg1_Enum1 = 0
	Msg1_E2 Msg1_Enum1 = 1
	Msg1_E3 Msg1_Enum1 = 2
)

// Enum value maps for Msg1_Enum1.
var (
	Msg1_Enum1_name = map[int32]string{
		0: "E1",
		1: "E2",
		2: "E3",
	}
	Msg1_Enum1_value = map[string]int32{
		"E1": 0,
		"E2": 1,
		"E3": 2,
	}
)

func (x Msg1_Enum1) Enum() *Msg1_Enum1 {
	p := new(Msg1_Enum1)
	*p = x
	return p
}

func (x Msg1_Enum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Msg1_Enum1) Descriptor() protoreflect.EnumDescriptor {
	return file_examplepb_example_proto_enumTypes[0].Descriptor()
}

func (Msg1_Enum1) Type() protoreflect.EnumType {
	return &file_examplepb_example_proto_enumTypes[0]
}

func (x Msg1_Enum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Msg1_Enum1) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Msg1_Enum1(num)
	return nil
}

// Deprecated: Use Msg1_Enum1.Descriptor instead.
func (Msg1_Enum1) EnumDescriptor() ([]byte, []int) {
	return file_examplepb_example_proto_rawDescGZIP(), []int{0, 0}
}

type Msg1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F1 []byte      `protobuf:"bytes,1,opt,name=f1" json:"f1,omitempty"`
	F2 *int64      `protobuf:"varint,2,opt,name=f2" json:"f2,omitempty"`
	F3 *Msg1_Enum1 `protobuf:"varint,3,opt,name=f3,enum=examplepb.Msg1_Enum1" json:"f3,omitempty"`
	F4 *bool       `protobuf:"varint,4,opt,name=f4" json:"f4,omitempty"`
	F5 *int64      `protobuf:"varint,5,opt,name=f5" json:"f5,omitempty"`
	F6 *Msg1_Enum1 `protobuf:"varint,6,opt,name=f6,enum=examplepb.Msg1_Enum1" json:"f6,omitempty"`
}

func (x *Msg1) Reset() {
	*x = Msg1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examplepb_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg1) ProtoMessage() {}

func (x *Msg1) ProtoReflect() protoreflect.Message {
	mi := &file_examplepb_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg1.ProtoReflect.Descriptor instead.
func (*Msg1) Descriptor() ([]byte, []int) {
	return file_examplepb_example_proto_rawDescGZIP(), []int{0}
}

func (x *Msg1) GetF1() []byte {
	if x != nil {
		return x.F1
	}
	return nil
}

func (x *Msg1) GetF2() int64 {
	if x != nil && x.F2 != nil {
		return *x.F2
	}
	return 0
}

func (x *Msg1) GetF3() Msg1_Enum1 {
	if x != nil && x.F3 != nil {
		return *x.F3
	}
	return Msg1_E1
}

func (x *Msg1) GetF4() bool {
	if x != nil && x.F4 != nil {
		return *x.F4
	}
	return false
}

func (x *Msg1) GetF5() int64 {
	if x != nil && x.F5 != nil {
		return *x.F5
	}
	return 0
}

func (x *Msg1) GetF6() Msg1_Enum1 {
	if x != nil && x.F6 != nil {
		return *x.F6
	}
	return Msg1_E1
}

var File_examplepb_example_proto protoreflect.FileDescriptor

var file_examplepb_example_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x1a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x61, 0x6e, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x04, 0x4d, 0x73, 0x67, 0x31, 0x12, 0x0e, 0x0a, 0x02,
	0x66, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x66, 0x31, 0x12, 0x0e, 0x0a, 0x02,
	0x66, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x66, 0x32, 0x12, 0x25, 0x0a, 0x02,
	0x66, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4d, 0x73, 0x67, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x52,
	0x02, 0x66, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x66, 0x34, 0x12, 0x17, 0x0a, 0x02, 0x66, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0x8a, 0xb5, 0x18, 0x03, 0x33, 0x2e, 0x31, 0x52, 0x02, 0x66, 0x35, 0x12, 0x2e, 0x0a, 0x02,
	0x66, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4d, 0x73, 0x67, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x42,
	0x07, 0x8a, 0xb5, 0x18, 0x03, 0x33, 0x2e, 0x32, 0x52, 0x02, 0x66, 0x36, 0x22, 0x31, 0x0a, 0x05,
	0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x31, 0x10, 0x00, 0x12, 0x06, 0x0a,
	0x02, 0x45, 0x32, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x02, 0x45, 0x33, 0x10, 0x02, 0x1a, 0x07, 0x9a,
	0xb5, 0x18, 0x03, 0x33, 0x2e, 0x33, 0x1a, 0x07, 0x92, 0xb5, 0x18, 0x03, 0x33, 0x2e, 0x30, 0x3a,
	0x07, 0x82, 0xb5, 0x18, 0x03, 0x33, 0x2e, 0x30, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62,
}

var (
	file_examplepb_example_proto_rawDescOnce sync.Once
	file_examplepb_example_proto_rawDescData = file_examplepb_example_proto_rawDesc
)

func file_examplepb_example_proto_rawDescGZIP() []byte {
	file_examplepb_example_proto_rawDescOnce.Do(func() {
		file_examplepb_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_examplepb_example_proto_rawDescData)
	})
	return file_examplepb_example_proto_rawDescData
}

var file_examplepb_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_examplepb_example_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_examplepb_example_proto_goTypes = []interface{}{
	(Msg1_Enum1)(0), // 0: examplepb.Msg1.Enum1
	(*Msg1)(nil),    // 1: examplepb.Msg1
}
var file_examplepb_example_proto_depIdxs = []int32{
	0, // 0: examplepb.Msg1.f3:type_name -> examplepb.Msg1.Enum1
	0, // 1: examplepb.Msg1.f6:type_name -> examplepb.Msg1.Enum1
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_examplepb_example_proto_init() }
func file_examplepb_example_proto_init() {
	if File_examplepb_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examplepb_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg1); i {
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
			RawDescriptor: file_examplepb_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examplepb_example_proto_goTypes,
		DependencyIndexes: file_examplepb_example_proto_depIdxs,
		EnumInfos:         file_examplepb_example_proto_enumTypes,
		MessageInfos:      file_examplepb_example_proto_msgTypes,
	}.Build()
	File_examplepb_example_proto = out.File
	file_examplepb_example_proto_rawDesc = nil
	file_examplepb_example_proto_goTypes = nil
	file_examplepb_example_proto_depIdxs = nil
}

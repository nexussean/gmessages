// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: pblite.proto

package gmproto

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_pblite_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50000,
		Name:          "pblite.pblite_binary",
		Tag:           "varint,50000,opt,name=pblite_binary",
		Filename:      "pblite.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool pblite_binary = 50000;
	E_PbliteBinary = &file_pblite_proto_extTypes[0]
)

var File_pblite_proto protoreflect.FileDescriptor

//go:embed pblite.pb.raw
var file_pblite_proto_rawDesc []byte

var file_pblite_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_pblite_proto_depIdxs = []int32{
	0, // 0: pblite.pblite_binary:extendee -> google.protobuf.FieldOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pblite_proto_init() }
func file_pblite_proto_init() {
	if File_pblite_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pblite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_pblite_proto_goTypes,
		DependencyIndexes: file_pblite_proto_depIdxs,
		ExtensionInfos:    file_pblite_proto_extTypes,
	}.Build()
	File_pblite_proto = out.File
	file_pblite_proto_rawDesc = nil
	file_pblite_proto_goTypes = nil
	file_pblite_proto_depIdxs = nil
}

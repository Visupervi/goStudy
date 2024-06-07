// Protocol Buffers - Google's data interchange format
// Copyright 2023 Google Inc.  All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/go_features.proto

package gofeaturespb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

type GoFeatures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not to generate the deprecated UnmarshalJSON method for enums.
	LegacyUnmarshalJsonEnum *bool `protobuf:"varint,1,opt,name=legacy_unmarshal_json_enum,json=legacyUnmarshalJsonEnum" json:"legacy_unmarshal_json_enum,omitempty"`
}

func (x *GoFeatures) Reset() {
	*x = GoFeatures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_go_features_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoFeatures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoFeatures) ProtoMessage() {}

func (x *GoFeatures) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_go_features_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoFeatures.ProtoReflect.Descriptor instead.
func (*GoFeatures) Descriptor() ([]byte, []int) {
	return file_google_protobuf_go_features_proto_rawDescGZIP(), []int{0}
}

func (x *GoFeatures) GetLegacyUnmarshalJsonEnum() bool {
	if x != nil && x.LegacyUnmarshalJsonEnum != nil {
		return *x.LegacyUnmarshalJsonEnum
	}
	return false
}

var file_google_protobuf_go_features_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FeatureSet)(nil),
		ExtensionType: (*GoFeatures)(nil),
		Field:         1002,
		Name:          "pb.go",
		Tag:           "bytes,1002,opt,name=go",
		Filename:      "google/protobuf/go_features.proto",
	},
}

// Extension fields to descriptorpb.FeatureSet.
var (
	// optional pb.GoFeatures go = 1002;
	E_Go = &file_google_protobuf_go_features_proto_extTypes[0]
)

var File_google_protobuf_go_features_proto protoreflect.FileDescriptor

var file_google_protobuf_go_features_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x0a, 0x47, 0x6f,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0xba, 0x01, 0x0a, 0x1a, 0x6c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x5f, 0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x5f, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x7d, 0x88,
	0x01, 0x01, 0x98, 0x01, 0x06, 0xa2, 0x01, 0x09, 0x12, 0x04, 0x74, 0x72, 0x75, 0x65, 0x18, 0x84,
	0x07, 0xa2, 0x01, 0x0a, 0x12, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x18, 0xe7, 0x07, 0xb2, 0x01,
	0x5b, 0x08, 0xe8, 0x07, 0x10, 0xe8, 0x07, 0x1a, 0x53, 0x54, 0x68, 0x65, 0x20, 0x6c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x20, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4a, 0x53, 0x4f,
	0x4e, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x66, 0x75, 0x74,
	0x75, 0x72, 0x65, 0x20, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x17, 0x6c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4a, 0x73, 0x6f,
	0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x3a, 0x3c, 0x0a, 0x02, 0x67, 0x6f, 0x12, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x02, 0x67, 0x6f, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x70, 0x62,
}

var (
	file_google_protobuf_go_features_proto_rawDescOnce sync.Once
	file_google_protobuf_go_features_proto_rawDescData = file_google_protobuf_go_features_proto_rawDesc
)

func file_google_protobuf_go_features_proto_rawDescGZIP() []byte {
	file_google_protobuf_go_features_proto_rawDescOnce.Do(func() {
		file_google_protobuf_go_features_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_go_features_proto_rawDescData)
	})
	return file_google_protobuf_go_features_proto_rawDescData
}

var file_google_protobuf_go_features_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_go_features_proto_goTypes = []any{
	(*GoFeatures)(nil),              // 0: pb.GoFeatures
	(*descriptorpb.FeatureSet)(nil), // 1: google.protobuf.FeatureSet
}
var file_google_protobuf_go_features_proto_depIdxs = []int32{
	1, // 0: pb.go:extendee -> google.protobuf.FeatureSet
	0, // 1: pb.go:type_name -> pb.GoFeatures
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_protobuf_go_features_proto_init() }
func file_google_protobuf_go_features_proto_init() {
	if File_google_protobuf_go_features_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_protobuf_go_features_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GoFeatures); i {
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
			RawDescriptor: file_google_protobuf_go_features_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_go_features_proto_goTypes,
		DependencyIndexes: file_google_protobuf_go_features_proto_depIdxs,
		MessageInfos:      file_google_protobuf_go_features_proto_msgTypes,
		ExtensionInfos:    file_google_protobuf_go_features_proto_extTypes,
	}.Build()
	File_google_protobuf_go_features_proto = out.File
	file_google_protobuf_go_features_proto_rawDesc = nil
	file_google_protobuf_go_features_proto_goTypes = nil
	file_google_protobuf_go_features_proto_depIdxs = nil
}

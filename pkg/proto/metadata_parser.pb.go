// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: pkg/proto/metadata_parser.proto

package proto

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

type ExtractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath  string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileType  string `protobuf:"bytes,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	FileChunk []byte `protobuf:"bytes,4,opt,name=file_chunk,json=fileChunk,proto3" json:"file_chunk,omitempty"`
}

func (x *ExtractRequest) Reset() {
	*x = ExtractRequest{}
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractRequest) ProtoMessage() {}

func (x *ExtractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractRequest.ProtoReflect.Descriptor instead.
func (*ExtractRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_metadata_parser_proto_rawDescGZIP(), []int{0}
}

func (x *ExtractRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ExtractRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *ExtractRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *ExtractRequest) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

type ExtractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata []byte `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ExtractResponse) Reset() {
	*x = ExtractResponse{}
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractResponse) ProtoMessage() {}

func (x *ExtractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractResponse.ProtoReflect.Descriptor instead.
func (*ExtractResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_metadata_parser_proto_rawDescGZIP(), []int{1}
}

func (x *ExtractResponse) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type GetRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRulesRequest) Reset() {
	*x = GetRulesRequest{}
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRulesRequest) ProtoMessage() {}

func (x *GetRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRulesRequest.ProtoReflect.Descriptor instead.
func (*GetRulesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_metadata_parser_proto_rawDescGZIP(), []int{2}
}

type GetRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules []*GetRulesResponseItem `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *GetRulesResponse) Reset() {
	*x = GetRulesResponse{}
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRulesResponse) ProtoMessage() {}

func (x *GetRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRulesResponse.ProtoReflect.Descriptor instead.
func (*GetRulesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_metadata_parser_proto_rawDescGZIP(), []int{3}
}

func (x *GetRulesResponse) GetRules() []*GetRulesResponseItem {
	if x != nil {
		return x.Rules
	}
	return nil
}

type GetRulesResponseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileNameRegexp string `protobuf:"bytes,1,opt,name=file_name_regexp,json=fileNameRegexp,proto3" json:"file_name_regexp,omitempty"`
	FilePathRegexp string `protobuf:"bytes,2,opt,name=file_path_regexp,json=filePathRegexp,proto3" json:"file_path_regexp,omitempty"`
	FileTypeRegexp string `protobuf:"bytes,3,opt,name=file_type_regexp,json=fileTypeRegexp,proto3" json:"file_type_regexp,omitempty"`
	FileChunkSize  int64  `protobuf:"varint,4,opt,name=file_chunk_size,json=fileChunkSize,proto3" json:"file_chunk_size,omitempty"`
}

func (x *GetRulesResponseItem) Reset() {
	*x = GetRulesResponseItem{}
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRulesResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRulesResponseItem) ProtoMessage() {}

func (x *GetRulesResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_metadata_parser_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRulesResponseItem.ProtoReflect.Descriptor instead.
func (*GetRulesResponseItem) Descriptor() ([]byte, []int) {
	return file_pkg_proto_metadata_parser_proto_rawDescGZIP(), []int{4}
}

func (x *GetRulesResponseItem) GetFileNameRegexp() string {
	if x != nil {
		return x.FileNameRegexp
	}
	return ""
}

func (x *GetRulesResponseItem) GetFilePathRegexp() string {
	if x != nil {
		return x.FilePathRegexp
	}
	return ""
}

func (x *GetRulesResponseItem) GetFileTypeRegexp() string {
	if x != nil {
		return x.FileTypeRegexp
	}
	return ""
}

func (x *GetRulesResponseItem) GetFileChunkSize() int64 {
	if x != nil {
		return x.FileChunkSize
	}
	return 0
}

var File_pkg_proto_metadata_parser_proto protoreflect.FileDescriptor

var file_pkg_proto_metadata_parser_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x86, 0x01, 0x0a, 0x0e,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x22, 0x2d, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x22, 0xbc, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67,
	0x65, 0x78, 0x70, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x28, 0x0a,
	0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x32,
	0x93, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x07, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x19,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_metadata_parser_proto_rawDescOnce sync.Once
	file_pkg_proto_metadata_parser_proto_rawDescData = file_pkg_proto_metadata_parser_proto_rawDesc
)

func file_pkg_proto_metadata_parser_proto_rawDescGZIP() []byte {
	file_pkg_proto_metadata_parser_proto_rawDescOnce.Do(func() {
		file_pkg_proto_metadata_parser_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_metadata_parser_proto_rawDescData)
	})
	return file_pkg_proto_metadata_parser_proto_rawDescData
}

var file_pkg_proto_metadata_parser_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_metadata_parser_proto_goTypes = []any{
	(*ExtractRequest)(nil),       // 0: metadata.ExtractRequest
	(*ExtractResponse)(nil),      // 1: metadata.ExtractResponse
	(*GetRulesRequest)(nil),      // 2: metadata.GetRulesRequest
	(*GetRulesResponse)(nil),     // 3: metadata.GetRulesResponse
	(*GetRulesResponseItem)(nil), // 4: metadata.GetRulesResponseItem
}
var file_pkg_proto_metadata_parser_proto_depIdxs = []int32{
	4, // 0: metadata.GetRulesResponse.rules:type_name -> metadata.GetRulesResponseItem
	0, // 1: metadata.MetadataParser.Extract:input_type -> metadata.ExtractRequest
	2, // 2: metadata.MetadataParser.GetRules:input_type -> metadata.GetRulesRequest
	1, // 3: metadata.MetadataParser.Extract:output_type -> metadata.ExtractResponse
	3, // 4: metadata.MetadataParser.GetRules:output_type -> metadata.GetRulesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_metadata_parser_proto_init() }
func file_pkg_proto_metadata_parser_proto_init() {
	if File_pkg_proto_metadata_parser_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_metadata_parser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_metadata_parser_proto_goTypes,
		DependencyIndexes: file_pkg_proto_metadata_parser_proto_depIdxs,
		MessageInfos:      file_pkg_proto_metadata_parser_proto_msgTypes,
	}.Build()
	File_pkg_proto_metadata_parser_proto = out.File
	file_pkg_proto_metadata_parser_proto_rawDesc = nil
	file_pkg_proto_metadata_parser_proto_goTypes = nil
	file_pkg_proto_metadata_parser_proto_depIdxs = nil
}
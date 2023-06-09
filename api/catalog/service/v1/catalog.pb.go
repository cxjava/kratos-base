// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: v1/catalog.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateBeerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBeerReq) Reset() {
	*x = CreateBeerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBeerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBeerReq) ProtoMessage() {}

func (x *CreateBeerReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBeerReq.ProtoReflect.Descriptor instead.
func (*CreateBeerReq) Descriptor() ([]byte, []int) {
	return file_v1_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBeerReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateBeerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBeerReply) Reset() {
	*x = CreateBeerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBeerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBeerReply) ProtoMessage() {}

func (x *CreateBeerReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBeerReply.ProtoReflect.Descriptor instead.
func (*CreateBeerReply) Descriptor() ([]byte, []int) {
	return file_v1_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBeerReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateBeerReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetBeerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBeerReq) Reset() {
	*x = GetBeerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeerReq) ProtoMessage() {}

func (x *GetBeerReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeerReq.ProtoReflect.Descriptor instead.
func (*GetBeerReq) Descriptor() ([]byte, []int) {
	return file_v1_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *GetBeerReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBeerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetBeerReply) Reset() {
	*x = GetBeerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeerReply) ProtoMessage() {}

func (x *GetBeerReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeerReply.ProtoReflect.Descriptor instead.
func (*GetBeerReply) Descriptor() ([]byte, []int) {
	return file_v1_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *GetBeerReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBeerReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_v1_catalog_proto protoreflect.FileDescriptor

var file_v1_catalog_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0xb0, 0x01, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x56,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x23, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_catalog_proto_rawDescOnce sync.Once
	file_v1_catalog_proto_rawDescData = file_v1_catalog_proto_rawDesc
)

func file_v1_catalog_proto_rawDescGZIP() []byte {
	file_v1_catalog_proto_rawDescOnce.Do(func() {
		file_v1_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_catalog_proto_rawDescData)
	})
	return file_v1_catalog_proto_rawDescData
}

var file_v1_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_catalog_proto_goTypes = []interface{}{
	(*CreateBeerReq)(nil),   // 0: catalog.service.v1.CreateBeerReq
	(*CreateBeerReply)(nil), // 1: catalog.service.v1.CreateBeerReply
	(*GetBeerReq)(nil),      // 2: catalog.service.v1.GetBeerReq
	(*GetBeerReply)(nil),    // 3: catalog.service.v1.GetBeerReply
}
var file_v1_catalog_proto_depIdxs = []int32{
	0, // 0: catalog.service.v1.Catalog.CreateBeer:input_type -> catalog.service.v1.CreateBeerReq
	2, // 1: catalog.service.v1.Catalog.GetBeer:input_type -> catalog.service.v1.GetBeerReq
	1, // 2: catalog.service.v1.Catalog.CreateBeer:output_type -> catalog.service.v1.CreateBeerReply
	3, // 3: catalog.service.v1.Catalog.GetBeer:output_type -> catalog.service.v1.GetBeerReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_catalog_proto_init() }
func file_v1_catalog_proto_init() {
	if File_v1_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBeerReq); i {
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
		file_v1_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBeerReply); i {
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
		file_v1_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBeerReq); i {
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
		file_v1_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBeerReply); i {
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
			RawDescriptor: file_v1_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_catalog_proto_goTypes,
		DependencyIndexes: file_v1_catalog_proto_depIdxs,
		MessageInfos:      file_v1_catalog_proto_msgTypes,
	}.Build()
	File_v1_catalog_proto = out.File
	file_v1_catalog_proto_rawDesc = nil
	file_v1_catalog_proto_goTypes = nil
	file_v1_catalog_proto_depIdxs = nil
}

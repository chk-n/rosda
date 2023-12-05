// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: internal/proto/api_slave.proto

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

// RecordServiceRequest definition in protobuf
type RecordServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=ResourceId,proto3" json:"ResourceId,omitempty"` // Resource ID
	// where image can be fetched from
	ImageUrl string `protobuf:"bytes,2,opt,name=ImageUrl,proto3" json:"ImageUrl,omitempty"`
	Cpu      uint32 `protobuf:"varint,3,opt,name=Cpu,proto3" json:"Cpu,omitempty"`
	Ram      uint32 `protobuf:"varint,4,opt,name=Ram,proto3" json:"Ram,omitempty"`
	// absolute number of instances that should be running
	InstanceCount uint32 `protobuf:"varint,5,opt,name=InstanceCount,proto3" json:"InstanceCount,omitempty"`
	EnvVariables  string `protobuf:"bytes,6,opt,name=EnvVariables,proto3" json:"EnvVariables,omitempty"`
}

func (x *RecordServiceRequest) Reset() {
	*x = RecordServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_slave_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordServiceRequest) ProtoMessage() {}

func (x *RecordServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_slave_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordServiceRequest.ProtoReflect.Descriptor instead.
func (*RecordServiceRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_slave_proto_rawDescGZIP(), []int{0}
}

func (x *RecordServiceRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RecordServiceRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *RecordServiceRequest) GetCpu() uint32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *RecordServiceRequest) GetRam() uint32 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *RecordServiceRequest) GetInstanceCount() uint32 {
	if x != nil {
		return x.InstanceCount
	}
	return 0
}

func (x *RecordServiceRequest) GetEnvVariables() string {
	if x != nil {
		return x.EnvVariables
	}
	return ""
}

// RecordServiceResponse definition in protobuf
type RecordServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// which port on host is open for connections to service
	Port uint32 `protobuf:"varint,1,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *RecordServiceResponse) Reset() {
	*x = RecordServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_slave_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordServiceResponse) ProtoMessage() {}

func (x *RecordServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_slave_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordServiceResponse.ProtoReflect.Descriptor instead.
func (*RecordServiceResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_slave_proto_rawDescGZIP(), []int{1}
}

func (x *RecordServiceResponse) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// RemoveServiceRequest definition in protobuf
type RemoveServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=ResourceId,proto3" json:"ResourceId,omitempty"`
	// how many seconds slave should wait before removing service
	TeardownDelay string `protobuf:"bytes,2,opt,name=TeardownDelay,proto3" json:"TeardownDelay,omitempty"`
}

func (x *RemoveServiceRequest) Reset() {
	*x = RemoveServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_slave_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveServiceRequest) ProtoMessage() {}

func (x *RemoveServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_slave_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveServiceRequest.ProtoReflect.Descriptor instead.
func (*RemoveServiceRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_slave_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveServiceRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RemoveServiceRequest) GetTeardownDelay() string {
	if x != nil {
		return x.TeardownDelay
	}
	return ""
}

// RemoveServiceResponse definition in protobuf
type RemoveServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=ResourceId,proto3" json:"ResourceId,omitempty"`
}

func (x *RemoveServiceResponse) Reset() {
	*x = RemoveServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_api_slave_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveServiceResponse) ProtoMessage() {}

func (x *RemoveServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_api_slave_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveServiceResponse.ProtoReflect.Descriptor instead.
func (*RemoveServiceResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_api_slave_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveServiceResponse) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

var File_internal_proto_api_slave_proto protoreflect.FileDescriptor

var file_internal_proto_api_slave_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc0, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x43, 0x70, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x61, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x52, 0x61, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x5c, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x65, 0x61, 0x72,
	0x64, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x54, 0x65, 0x61, 0x72, 0x64, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x37,
	0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x32, 0x8e, 0x01, 0x0a, 0x08, 0x41, 0x70, 0x69, 0x53,
	0x6c, 0x61, 0x76, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_proto_api_slave_proto_rawDescOnce sync.Once
	file_internal_proto_api_slave_proto_rawDescData = file_internal_proto_api_slave_proto_rawDesc
)

func file_internal_proto_api_slave_proto_rawDescGZIP() []byte {
	file_internal_proto_api_slave_proto_rawDescOnce.Do(func() {
		file_internal_proto_api_slave_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_api_slave_proto_rawDescData)
	})
	return file_internal_proto_api_slave_proto_rawDescData
}

var file_internal_proto_api_slave_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_proto_api_slave_proto_goTypes = []interface{}{
	(*RecordServiceRequest)(nil),  // 0: RecordServiceRequest
	(*RecordServiceResponse)(nil), // 1: RecordServiceResponse
	(*RemoveServiceRequest)(nil),  // 2: RemoveServiceRequest
	(*RemoveServiceResponse)(nil), // 3: RemoveServiceResponse
}
var file_internal_proto_api_slave_proto_depIdxs = []int32{
	0, // 0: ApiSlave.RecordService:input_type -> RecordServiceRequest
	2, // 1: ApiSlave.RemoveService:input_type -> RemoveServiceRequest
	1, // 2: ApiSlave.RecordService:output_type -> RecordServiceResponse
	3, // 3: ApiSlave.RemoveService:output_type -> RemoveServiceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_api_slave_proto_init() }
func file_internal_proto_api_slave_proto_init() {
	if File_internal_proto_api_slave_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_api_slave_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordServiceRequest); i {
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
		file_internal_proto_api_slave_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordServiceResponse); i {
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
		file_internal_proto_api_slave_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveServiceRequest); i {
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
		file_internal_proto_api_slave_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveServiceResponse); i {
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
			RawDescriptor: file_internal_proto_api_slave_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_api_slave_proto_goTypes,
		DependencyIndexes: file_internal_proto_api_slave_proto_depIdxs,
		MessageInfos:      file_internal_proto_api_slave_proto_msgTypes,
	}.Build()
	File_internal_proto_api_slave_proto = out.File
	file_internal_proto_api_slave_proto_rawDesc = nil
	file_internal_proto_api_slave_proto_goTypes = nil
	file_internal_proto_api_slave_proto_depIdxs = nil
}

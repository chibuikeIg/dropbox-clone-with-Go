// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: upload_data.proto

package uploaddataservice

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

// save upload data request and response
type SaveUploadDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UploadId string `protobuf:"bytes,2,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	PartNum  int32  `protobuf:"varint,3,opt,name=partNum,proto3" json:"partNum,omitempty"`
	Etag     string `protobuf:"bytes,4,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *SaveUploadDataRequest) Reset() {
	*x = SaveUploadDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUploadDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUploadDataRequest) ProtoMessage() {}

func (x *SaveUploadDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUploadDataRequest.ProtoReflect.Descriptor instead.
func (*SaveUploadDataRequest) Descriptor() ([]byte, []int) {
	return file_upload_data_proto_rawDescGZIP(), []int{0}
}

func (x *SaveUploadDataRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SaveUploadDataRequest) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *SaveUploadDataRequest) GetPartNum() int32 {
	if x != nil {
		return x.PartNum
	}
	return 0
}

func (x *SaveUploadDataRequest) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type SaveUploadDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	StatusCode int32  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *SaveUploadDataResponse) Reset() {
	*x = SaveUploadDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUploadDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUploadDataResponse) ProtoMessage() {}

func (x *SaveUploadDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUploadDataResponse.ProtoReflect.Descriptor instead.
func (*SaveUploadDataResponse) Descriptor() ([]byte, []int) {
	return file_upload_data_proto_rawDescGZIP(), []int{1}
}

func (x *SaveUploadDataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SaveUploadDataResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

// fetch upload data request and response
type UploadDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadId string `protobuf:"bytes,1,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
}

func (x *UploadDataRequest) Reset() {
	*x = UploadDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDataRequest) ProtoMessage() {}

func (x *UploadDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDataRequest.ProtoReflect.Descriptor instead.
func (*UploadDataRequest) Descriptor() ([]byte, []int) {
	return file_upload_data_proto_rawDescGZIP(), []int{2}
}

func (x *UploadDataRequest) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

type UploadDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	StatusCode int32  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *UploadDataResponse) Reset() {
	*x = UploadDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDataResponse) ProtoMessage() {}

func (x *UploadDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDataResponse.ProtoReflect.Descriptor instead.
func (*UploadDataResponse) Descriptor() ([]byte, []int) {
	return file_upload_data_proto_rawDescGZIP(), []int{3}
}

func (x *UploadDataResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *UploadDataResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

// delete upload data request and response
type DeleteUploadDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UploadId string `protobuf:"bytes,2,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
}

func (x *DeleteUploadDataRequest) Reset() {
	*x = DeleteUploadDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUploadDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUploadDataRequest) ProtoMessage() {}

func (x *DeleteUploadDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUploadDataRequest.ProtoReflect.Descriptor instead.
func (*DeleteUploadDataRequest) Descriptor() ([]byte, []int) {
	return file_upload_data_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUploadDataRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteUploadDataRequest) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

type DeleteUploadDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	StatusCode int32  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *DeleteUploadDataResponse) Reset() {
	*x = DeleteUploadDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUploadDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUploadDataResponse) ProtoMessage() {}

func (x *DeleteUploadDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUploadDataResponse.ProtoReflect.Descriptor instead.
func (*DeleteUploadDataResponse) Descriptor() ([]byte, []int) {
	return file_upload_data_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUploadDataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteUploadDataResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_upload_data_proto protoreflect.FileDescriptor

var file_upload_data_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74,
	0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x52,
	0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x4d, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x18,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0xda, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x43, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upload_data_proto_rawDescOnce sync.Once
	file_upload_data_proto_rawDescData = file_upload_data_proto_rawDesc
)

func file_upload_data_proto_rawDescGZIP() []byte {
	file_upload_data_proto_rawDescOnce.Do(func() {
		file_upload_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_data_proto_rawDescData)
	})
	return file_upload_data_proto_rawDescData
}

var file_upload_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_upload_data_proto_goTypes = []interface{}{
	(*SaveUploadDataRequest)(nil),    // 0: SaveUploadDataRequest
	(*SaveUploadDataResponse)(nil),   // 1: SaveUploadDataResponse
	(*UploadDataRequest)(nil),        // 2: UploadDataRequest
	(*UploadDataResponse)(nil),       // 3: UploadDataResponse
	(*DeleteUploadDataRequest)(nil),  // 4: DeleteUploadDataRequest
	(*DeleteUploadDataResponse)(nil), // 5: DeleteUploadDataResponse
}
var file_upload_data_proto_depIdxs = []int32{
	0, // 0: UploadData.SaveUploadData:input_type -> SaveUploadDataRequest
	2, // 1: UploadData.FetchUploadData:input_type -> UploadDataRequest
	4, // 2: UploadData.DeleteUploadData:input_type -> DeleteUploadDataRequest
	1, // 3: UploadData.SaveUploadData:output_type -> SaveUploadDataResponse
	3, // 4: UploadData.FetchUploadData:output_type -> UploadDataResponse
	5, // 5: UploadData.DeleteUploadData:output_type -> DeleteUploadDataResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upload_data_proto_init() }
func file_upload_data_proto_init() {
	if File_upload_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUploadDataRequest); i {
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
		file_upload_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUploadDataResponse); i {
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
		file_upload_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDataRequest); i {
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
		file_upload_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDataResponse); i {
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
		file_upload_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUploadDataRequest); i {
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
		file_upload_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUploadDataResponse); i {
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
			RawDescriptor: file_upload_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_data_proto_goTypes,
		DependencyIndexes: file_upload_data_proto_depIdxs,
		MessageInfos:      file_upload_data_proto_msgTypes,
	}.Build()
	File_upload_data_proto = out.File
	file_upload_data_proto_rawDesc = nil
	file_upload_data_proto_goTypes = nil
	file_upload_data_proto_depIdxs = nil
}

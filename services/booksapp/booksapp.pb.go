// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: proto/booksapp.proto

package booksapp

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

type BookInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Genre  string `protobuf:"bytes,3,opt,name=genre,proto3" json:"genre,omitempty"`
	Author string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *BookInfo) Reset() {
	*x = BookInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookInfo) ProtoMessage() {}

func (x *BookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookInfo.ProtoReflect.Descriptor instead.
func (*BookInfo) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{0}
}

func (x *BookInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookInfo) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *BookInfo) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{1}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type GetBooksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested page
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// Requested number of items per page
	PerPage int64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *GetBooksReq) Reset() {
	*x = GetBooksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksReq) ProtoMessage() {}

func (x *GetBooksReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksReq.ProtoReflect.Descriptor instead.
func (*GetBooksReq) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{3}
}

func (x *GetBooksReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetBooksReq) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type GetBooksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection of books
	Books []*BookInfo `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBooksResp) Reset() {
	*x = GetBooksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksResp) ProtoMessage() {}

func (x *GetBooksResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksResp.ProtoReflect.Descriptor instead.
func (*GetBooksResp) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{4}
}

func (x *GetBooksResp) GetBooks() []*BookInfo {
	if x != nil {
		return x.Books
	}
	return nil
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime  string `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{5}
}

func (x *UploadRequest) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *UploadRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{6}
}

func (x *UploadResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Mime  string `protobuf:"bytes,2,opt,name=mime,proto3" json:"mime,omitempty"`
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booksapp_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booksapp_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_proto_booksapp_proto_rawDescGZIP(), []int{8}
}

func (x *DownloadResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *DownloadResponse) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

var File_proto_booksapp_proto protoreflect.FileDescriptor

var file_proto_booksapp_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e,
	0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x14,
	0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x39, 0x0a, 0x0d,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x24, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a,
	0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69,
	0x6d, 0x65, 0x32, 0x9a, 0x04, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x4b, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61,
	0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49,
	0x64, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x46, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61,
	0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x22, 0x0b, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x01,
	0x2a, 0x12, 0x44, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x1a, 0x05, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x40, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x55, 0x0a, 0x06, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b,
	0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a, 0x28, 0x01,
	0x12, 0x5d, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61,
	0x70, 0x70, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a, 0x30, 0x01, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_booksapp_proto_rawDescOnce sync.Once
	file_proto_booksapp_proto_rawDescData = file_proto_booksapp_proto_rawDesc
)

func file_proto_booksapp_proto_rawDescGZIP() []byte {
	file_proto_booksapp_proto_rawDescOnce.Do(func() {
		file_proto_booksapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_booksapp_proto_rawDescData)
	})
	return file_proto_booksapp_proto_rawDescData
}

var file_proto_booksapp_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_booksapp_proto_goTypes = []interface{}{
	(*BookInfo)(nil),         // 0: booksapp.BookInfo
	(*Id)(nil),               // 1: booksapp.Id
	(*Status)(nil),           // 2: booksapp.Status
	(*GetBooksReq)(nil),      // 3: booksapp.GetBooksReq
	(*GetBooksResp)(nil),     // 4: booksapp.GetBooksResp
	(*UploadRequest)(nil),    // 5: booksapp.UploadRequest
	(*UploadResponse)(nil),   // 6: booksapp.UploadResponse
	(*DownloadRequest)(nil),  // 7: booksapp.DownloadRequest
	(*DownloadResponse)(nil), // 8: booksapp.DownloadResponse
}
var file_proto_booksapp_proto_depIdxs = []int32{
	0, // 0: booksapp.GetBooksResp.books:type_name -> booksapp.BookInfo
	3, // 1: booksapp.Book.GetBooks:input_type -> booksapp.GetBooksReq
	1, // 2: booksapp.Book.GetBook:input_type -> booksapp.Id
	0, // 3: booksapp.Book.CreateBook:input_type -> booksapp.BookInfo
	0, // 4: booksapp.Book.UpdateBook:input_type -> booksapp.BookInfo
	1, // 5: booksapp.Book.DeleteBook:input_type -> booksapp.Id
	5, // 6: booksapp.Book.Upload:input_type -> booksapp.UploadRequest
	7, // 7: booksapp.Book.Download:input_type -> booksapp.DownloadRequest
	4, // 8: booksapp.Book.GetBooks:output_type -> booksapp.GetBooksResp
	0, // 9: booksapp.Book.GetBook:output_type -> booksapp.BookInfo
	1, // 10: booksapp.Book.CreateBook:output_type -> booksapp.Id
	2, // 11: booksapp.Book.UpdateBook:output_type -> booksapp.Status
	2, // 12: booksapp.Book.DeleteBook:output_type -> booksapp.Status
	6, // 13: booksapp.Book.Upload:output_type -> booksapp.UploadResponse
	8, // 14: booksapp.Book.Download:output_type -> booksapp.DownloadResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_booksapp_proto_init() }
func file_proto_booksapp_proto_init() {
	if File_proto_booksapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_booksapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookInfo); i {
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
		file_proto_booksapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_proto_booksapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_proto_booksapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksReq); i {
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
		file_proto_booksapp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksResp); i {
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
		file_proto_booksapp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_proto_booksapp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_proto_booksapp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_proto_booksapp_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadResponse); i {
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
			RawDescriptor: file_proto_booksapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_booksapp_proto_goTypes,
		DependencyIndexes: file_proto_booksapp_proto_depIdxs,
		MessageInfos:      file_proto_booksapp_proto_msgTypes,
	}.Build()
	File_proto_booksapp_proto = out.File
	file_proto_booksapp_proto_rawDesc = nil
	file_proto_booksapp_proto_goTypes = nil
	file_proto_booksapp_proto_depIdxs = nil
}

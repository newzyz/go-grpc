// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: proto/customersapp.proto

package customersapp

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

type CustomerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Age       int64  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *CustomerInfo) Reset() {
	*x = CustomerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customersapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerInfo) ProtoMessage() {}

func (x *CustomerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customersapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerInfo.ProtoReflect.Descriptor instead.
func (*CustomerInfo) Descriptor() ([]byte, []int) {
	return file_proto_customersapp_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerInfo) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CustomerInfo) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CustomerInfo) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
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
		mi := &file_proto_customersapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customersapp_proto_msgTypes[1]
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
	return file_proto_customersapp_proto_rawDescGZIP(), []int{1}
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
		mi := &file_proto_customersapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customersapp_proto_msgTypes[2]
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
	return file_proto_customersapp_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type GetCustomersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested page
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// Requested number of items per page
	PerPage int64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *GetCustomersReq) Reset() {
	*x = GetCustomersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customersapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomersReq) ProtoMessage() {}

func (x *GetCustomersReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customersapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomersReq.ProtoReflect.Descriptor instead.
func (*GetCustomersReq) Descriptor() ([]byte, []int) {
	return file_proto_customersapp_proto_rawDescGZIP(), []int{3}
}

func (x *GetCustomersReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetCustomersReq) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type GetCustomersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection of books
	Customers []*CustomerInfo `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
}

func (x *GetCustomersResp) Reset() {
	*x = GetCustomersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customersapp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomersResp) ProtoMessage() {}

func (x *GetCustomersResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customersapp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomersResp.ProtoReflect.Descriptor instead.
func (*GetCustomersResp) Descriptor() ([]byte, []int) {
	return file_proto_customersapp_proto_rawDescGZIP(), []int{4}
}

func (x *GetCustomersResp) GetCustomers() []*CustomerInfo {
	if x != nil {
		return x.Customers
	}
	return nil
}

var File_proto_customersapp_proto protoreflect.FileDescriptor

var file_proto_customersapp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x32, 0xcc, 0x03, 0x0a, 0x08, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x63, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12,
	0x0e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x5a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x1a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x10, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x14, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x1a, 0x09, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x50, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x26, 0x5a, 0x24, 0x2e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61,
	0x70, 0x70, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_customersapp_proto_rawDescOnce sync.Once
	file_proto_customersapp_proto_rawDescData = file_proto_customersapp_proto_rawDesc
)

func file_proto_customersapp_proto_rawDescGZIP() []byte {
	file_proto_customersapp_proto_rawDescOnce.Do(func() {
		file_proto_customersapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_customersapp_proto_rawDescData)
	})
	return file_proto_customersapp_proto_rawDescData
}

var file_proto_customersapp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_customersapp_proto_goTypes = []interface{}{
	(*CustomerInfo)(nil),     // 0: customersapp.CustomerInfo
	(*Id)(nil),               // 1: customersapp.Id
	(*Status)(nil),           // 2: customersapp.Status
	(*GetCustomersReq)(nil),  // 3: customersapp.GetCustomersReq
	(*GetCustomersResp)(nil), // 4: customersapp.GetCustomersResp
}
var file_proto_customersapp_proto_depIdxs = []int32{
	0, // 0: customersapp.GetCustomersResp.customers:type_name -> customersapp.CustomerInfo
	3, // 1: customersapp.Customer.GetCustomers:input_type -> customersapp.GetCustomersReq
	1, // 2: customersapp.Customer.GetCustomer:input_type -> customersapp.Id
	0, // 3: customersapp.Customer.CreateCustomer:input_type -> customersapp.CustomerInfo
	0, // 4: customersapp.Customer.UpdateCustomer:input_type -> customersapp.CustomerInfo
	1, // 5: customersapp.Customer.DeleteCustomer:input_type -> customersapp.Id
	4, // 6: customersapp.Customer.GetCustomers:output_type -> customersapp.GetCustomersResp
	0, // 7: customersapp.Customer.GetCustomer:output_type -> customersapp.CustomerInfo
	1, // 8: customersapp.Customer.CreateCustomer:output_type -> customersapp.Id
	2, // 9: customersapp.Customer.UpdateCustomer:output_type -> customersapp.Status
	2, // 10: customersapp.Customer.DeleteCustomer:output_type -> customersapp.Status
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_customersapp_proto_init() }
func file_proto_customersapp_proto_init() {
	if File_proto_customersapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_customersapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerInfo); i {
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
		file_proto_customersapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_customersapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_customersapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomersReq); i {
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
		file_proto_customersapp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomersResp); i {
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
			RawDescriptor: file_proto_customersapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_customersapp_proto_goTypes,
		DependencyIndexes: file_proto_customersapp_proto_depIdxs,
		MessageInfos:      file_proto_customersapp_proto_msgTypes,
	}.Build()
	File_proto_customersapp_proto = out.File
	file_proto_customersapp_proto_rawDesc = nil
	file_proto_customersapp_proto_goTypes = nil
	file_proto_customersapp_proto_depIdxs = nil
}

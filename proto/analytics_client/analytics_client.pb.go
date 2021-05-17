// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: analytics_client.proto

package analytics_client_pb

import (
	proto "github.com/golang/protobuf/proto"
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

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_analytics_client_proto_rawDescGZIP(), []int{0}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_analytics_client_proto_rawDescGZIP(), []int{1}
}

type Dish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Dish) Reset() {
	*x = Dish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dish) ProtoMessage() {}

func (x *Dish) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dish.ProtoReflect.Descriptor instead.
func (*Dish) Descriptor() ([]byte, []int) {
	return file_analytics_client_proto_rawDescGZIP(), []int{2}
}

func (x *Dish) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dish) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type AddBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelName   string  `protobuf:"bytes,1,opt,name=channelName,proto3" json:"channelName,omitempty"`
	StoreId       string  `protobuf:"bytes,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
	CustomerName  string  `protobuf:"bytes,3,opt,name=customerName,proto3" json:"customerName,omitempty"`
	Timestamp     int64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Dishes        []*Dish `protobuf:"bytes,5,rep,name=dishes,proto3" json:"dishes,omitempty"`
	TotalQuantity int32   `protobuf:"varint,6,opt,name=totalQuantity,proto3" json:"totalQuantity,omitempty"`
	Total         int64   `protobuf:"varint,7,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *AddBillRequest) Reset() {
	*x = AddBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBillRequest) ProtoMessage() {}

func (x *AddBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBillRequest.ProtoReflect.Descriptor instead.
func (*AddBillRequest) Descriptor() ([]byte, []int) {
	return file_analytics_client_proto_rawDescGZIP(), []int{3}
}

func (x *AddBillRequest) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *AddBillRequest) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *AddBillRequest) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *AddBillRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *AddBillRequest) GetDishes() []*Dish {
	if x != nil {
		return x.Dishes
	}
	return nil
}

func (x *AddBillRequest) GetTotalQuantity() int32 {
	if x != nil {
		return x.TotalQuantity
	}
	return 0
}

func (x *AddBillRequest) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AddBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddBillResponse) Reset() {
	*x = AddBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBillResponse) ProtoMessage() {}

func (x *AddBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBillResponse.ProtoReflect.Descriptor instead.
func (*AddBillResponse) Descriptor() ([]byte, []int) {
	return file_analytics_client_proto_rawDescGZIP(), []int{4}
}

func (x *AddBillResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_analytics_client_proto protoreflect.FileDescriptor

var file_analytics_client_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x15, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x06, 0x64, 0x69,
	0x73, 0x68, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x69,
	0x73, 0x68, 0x52, 0x06, 0x64, 0x69, 0x73, 0x68, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x69, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xc1, 0x01, 0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x69, 0x6c, 0x6c,
	0x12, 0x20, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x3b, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analytics_client_proto_rawDescOnce sync.Once
	file_analytics_client_proto_rawDescData = file_analytics_client_proto_rawDesc
)

func file_analytics_client_proto_rawDescGZIP() []byte {
	file_analytics_client_proto_rawDescOnce.Do(func() {
		file_analytics_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_client_proto_rawDescData)
	})
	return file_analytics_client_proto_rawDescData
}

var file_analytics_client_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_analytics_client_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),  // 0: analytics_client.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 1: analytics_client.HealthCheckResponse
	(*Dish)(nil),                // 2: analytics_client.Dish
	(*AddBillRequest)(nil),      // 3: analytics_client.AddBillRequest
	(*AddBillResponse)(nil),     // 4: analytics_client.AddBillResponse
}
var file_analytics_client_proto_depIdxs = []int32{
	2, // 0: analytics_client.AddBillRequest.dishes:type_name -> analytics_client.Dish
	0, // 1: analytics_client.AnalyticsClient.HealthCheck:input_type -> analytics_client.HealthCheckRequest
	3, // 2: analytics_client.AnalyticsClient.AddBill:input_type -> analytics_client.AddBillRequest
	1, // 3: analytics_client.AnalyticsClient.HealthCheck:output_type -> analytics_client.HealthCheckResponse
	4, // 4: analytics_client.AnalyticsClient.AddBill:output_type -> analytics_client.AddBillResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_analytics_client_proto_init() }
func file_analytics_client_proto_init() {
	if File_analytics_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analytics_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_analytics_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_analytics_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dish); i {
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
		file_analytics_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBillRequest); i {
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
		file_analytics_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBillResponse); i {
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
			RawDescriptor: file_analytics_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_client_proto_goTypes,
		DependencyIndexes: file_analytics_client_proto_depIdxs,
		MessageInfos:      file_analytics_client_proto_msgTypes,
	}.Build()
	File_analytics_client_proto = out.File
	file_analytics_client_proto_rawDesc = nil
	file_analytics_client_proto_goTypes = nil
	file_analytics_client_proto_depIdxs = nil
}

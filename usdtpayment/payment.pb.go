// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: payment.proto

package usdtpayment

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

type Chain int32

const (
	Chain_UnknownChain Chain = 0
	Chain_Trc20        Chain = 1
)

// Enum value maps for Chain.
var (
	Chain_name = map[int32]string{
		0: "UnknownChain",
		1: "Trc20",
	}
	Chain_value = map[string]int32{
		"UnknownChain": 0,
		"Trc20":        1,
	}
)

func (x Chain) Enum() *Chain {
	p := new(Chain)
	*p = x
	return p
}

func (x Chain) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Chain) Descriptor() protoreflect.EnumDescriptor {
	return file_payment_proto_enumTypes[0].Descriptor()
}

func (Chain) Type() protoreflect.EnumType {
	return &file_payment_proto_enumTypes[0]
}

func (x Chain) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Chain.Descriptor instead.
func (Chain) EnumDescriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price          float32 `protobuf:"fixed32,1,opt,name=price,proto3" json:"price,omitempty"`
	Desc           string  `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	OrderId        string  `protobuf:"bytes,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	AsyncNotifyUrl string  `protobuf:"bytes,4,opt,name=async_notify_url,json=asyncNotifyUrl,proto3" json:"async_notify_url,omitempty"`
	ReturnUrl      string  `protobuf:"bytes,5,opt,name=return_url,json=returnUrl,proto3" json:"return_url,omitempty"`
	Timeout        int64   `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Chain          Chain   `protobuf:"varint,7,opt,name=chain,proto3,enum=proto.Chain" json:"chain,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateOrderRequest) GetAsyncNotifyUrl() string {
	if x != nil {
		return x.AsyncNotifyUrl
	}
	return ""
}

func (x *CreateOrderRequest) GetReturnUrl() string {
	if x != nil {
		return x.ReturnUrl
	}
	return ""
}

func (x *CreateOrderRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *CreateOrderRequest) GetChain() Chain {
	if x != nil {
		return x.Chain
	}
	return Chain_UnknownChain
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemOrderId string `protobuf:"bytes,1,opt,name=system_order_id,json=systemOrderId,proto3" json:"system_order_id,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetSystemOrderId() string {
	if x != nil {
		return x.SystemOrderId
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x73, 0x79,
	0x6e, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x24, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x10, 0x0a, 0x0c, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x72, 0x63, 0x32, 0x30, 0x10, 0x01, 0x32, 0x56, 0x0a,
	0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x75, 0x73, 0x64, 0x74, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payment_proto_goTypes = []interface{}{
	(Chain)(0),                  // 0: proto.Chain
	(*CreateOrderRequest)(nil),  // 1: proto.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 2: proto.CreateOrderResponse
}
var file_payment_proto_depIdxs = []int32{
	0, // 0: proto.CreateOrderRequest.chain:type_name -> proto.Chain
	1, // 1: proto.PaymentService.CreateOrder:input_type -> proto.CreateOrderRequest
	2, // 2: proto.PaymentService.CreateOrder:output_type -> proto.CreateOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		EnumInfos:         file_payment_proto_enumTypes,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}

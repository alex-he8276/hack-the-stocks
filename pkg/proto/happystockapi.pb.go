// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: happystock/happystockapi.proto

package happystock

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StockSentiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Date      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Sentiment int32                  `protobuf:"varint,3,opt,name=sentiment,proto3" json:"sentiment,omitempty"`
}

func (x *StockSentiment) Reset() {
	*x = StockSentiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_happystock_happystockapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockSentiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockSentiment) ProtoMessage() {}

func (x *StockSentiment) ProtoReflect() protoreflect.Message {
	mi := &file_happystock_happystockapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockSentiment.ProtoReflect.Descriptor instead.
func (*StockSentiment) Descriptor() ([]byte, []int) {
	return file_happystock_happystockapi_proto_rawDescGZIP(), []int{0}
}

func (x *StockSentiment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StockSentiment) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *StockSentiment) GetSentiment() int32 {
	if x != nil {
		return x.Sentiment
	}
	return 0
}

type StockPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Date  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Price float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *StockPrice) Reset() {
	*x = StockPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_happystock_happystockapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockPrice) ProtoMessage() {}

func (x *StockPrice) ProtoReflect() protoreflect.Message {
	mi := &file_happystock_happystockapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockPrice.ProtoReflect.Descriptor instead.
func (*StockPrice) Descriptor() ([]byte, []int) {
	return file_happystock_happystockapi_proto_rawDescGZIP(), []int{1}
}

func (x *StockPrice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StockPrice) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *StockPrice) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ListStockSentiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SentimentList []*StockSentiment `protobuf:"bytes,1,rep,name=sentimentList,proto3" json:"sentimentList,omitempty"`
}

func (x *ListStockSentiment) Reset() {
	*x = ListStockSentiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_happystock_happystockapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStockSentiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStockSentiment) ProtoMessage() {}

func (x *ListStockSentiment) ProtoReflect() protoreflect.Message {
	mi := &file_happystock_happystockapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStockSentiment.ProtoReflect.Descriptor instead.
func (*ListStockSentiment) Descriptor() ([]byte, []int) {
	return file_happystock_happystockapi_proto_rawDescGZIP(), []int{2}
}

func (x *ListStockSentiment) GetSentimentList() []*StockSentiment {
	if x != nil {
		return x.SentimentList
	}
	return nil
}

type ListStockPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PriceList []*StockPrice `protobuf:"bytes,1,rep,name=priceList,proto3" json:"priceList,omitempty"`
}

func (x *ListStockPrice) Reset() {
	*x = ListStockPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_happystock_happystockapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStockPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStockPrice) ProtoMessage() {}

func (x *ListStockPrice) ProtoReflect() protoreflect.Message {
	mi := &file_happystock_happystockapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStockPrice.ProtoReflect.Descriptor instead.
func (*ListStockPrice) Descriptor() ([]byte, []int) {
	return file_happystock_happystockapi_proto_rawDescGZIP(), []int{3}
}

func (x *ListStockPrice) GetPriceList() []*StockPrice {
	if x != nil {
		return x.PriceList
	}
	return nil
}

var File_happystock_happystockapi_proto protoreflect.FileDescriptor

var file_happystock_happystockapi_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x68, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x68, 0x61, 0x70,
	0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x68, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x68, 0x61, 0x70,
	0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x0e, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x66, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x64, 0x0a, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a,
	0x0d, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x68, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x68, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d,
	0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x54, 0x0a,
	0x0e, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x42, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x68, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0xde, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x70, 0x70,
	0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x68, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x61, 0x70, 0x69, 0x42, 0x12, 0x48, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x61, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x62, 0x75, 0x66, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x6d, 0x65, 0x77, 0x63, 0x69, 0x66, 0x65, 0x72, 0x2f, 0x68,
	0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x68, 0x61, 0x70, 0x70, 0x79,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0xa2, 0x02, 0x03, 0x48, 0x48, 0x58, 0xaa, 0x02, 0x18, 0x48, 0x61,
	0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x48, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x61, 0x70, 0x69, 0xca, 0x02, 0x18, 0x48, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x5c, 0x48, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x61, 0x70,
	0x69, 0xe2, 0x02, 0x24, 0x48, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5c, 0x48,
	0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x61, 0x70, 0x69, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x48, 0x61, 0x70, 0x70, 0x79,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x3a, 0x3a, 0x48, 0x61, 0x70, 0x70, 0x79, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_happystock_happystockapi_proto_rawDescOnce sync.Once
	file_happystock_happystockapi_proto_rawDescData = file_happystock_happystockapi_proto_rawDesc
)

func file_happystock_happystockapi_proto_rawDescGZIP() []byte {
	file_happystock_happystockapi_proto_rawDescOnce.Do(func() {
		file_happystock_happystockapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_happystock_happystockapi_proto_rawDescData)
	})
	return file_happystock_happystockapi_proto_rawDescData
}

var file_happystock_happystockapi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_happystock_happystockapi_proto_goTypes = []interface{}{
	(*StockSentiment)(nil),        // 0: happystock.happystockapi.stockSentiment
	(*StockPrice)(nil),            // 1: happystock.happystockapi.stockPrice
	(*ListStockSentiment)(nil),    // 2: happystock.happystockapi.listStockSentiment
	(*ListStockPrice)(nil),        // 3: happystock.happystockapi.listStockPrice
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_happystock_happystockapi_proto_depIdxs = []int32{
	4, // 0: happystock.happystockapi.stockSentiment.date:type_name -> google.protobuf.Timestamp
	4, // 1: happystock.happystockapi.stockPrice.date:type_name -> google.protobuf.Timestamp
	0, // 2: happystock.happystockapi.listStockSentiment.sentimentList:type_name -> happystock.happystockapi.stockSentiment
	1, // 3: happystock.happystockapi.listStockPrice.priceList:type_name -> happystock.happystockapi.stockPrice
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_happystock_happystockapi_proto_init() }
func file_happystock_happystockapi_proto_init() {
	if File_happystock_happystockapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_happystock_happystockapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockSentiment); i {
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
		file_happystock_happystockapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockPrice); i {
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
		file_happystock_happystockapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStockSentiment); i {
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
		file_happystock_happystockapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStockPrice); i {
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
			RawDescriptor: file_happystock_happystockapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_happystock_happystockapi_proto_goTypes,
		DependencyIndexes: file_happystock_happystockapi_proto_depIdxs,
		MessageInfos:      file_happystock_happystockapi_proto_msgTypes,
	}.Build()
	File_happystock_happystockapi_proto = out.File
	file_happystock_happystockapi_proto_rawDesc = nil
	file_happystock_happystockapi_proto_goTypes = nil
	file_happystock_happystockapi_proto_depIdxs = nil
}

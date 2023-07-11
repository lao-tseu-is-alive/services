// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/stock.proto

package stock

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

// Get the last price for a given stock ticker
type PriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stock symbol e.g AAPL
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *PriceRequest) Reset() {
	*x = PriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceRequest) ProtoMessage() {}

func (x *PriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceRequest.ProtoReflect.Descriptor instead.
func (*PriceRequest) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{0}
}

func (x *PriceRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type PriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the stock symbol e.g AAPL
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// the last price
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PriceResponse) Reset() {
	*x = PriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceResponse) ProtoMessage() {}

func (x *PriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceResponse.ProtoReflect.Descriptor instead.
func (*PriceResponse) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{1}
}

func (x *PriceResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *PriceResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// Get the last quote for the stock
type QuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the stock symbol e.g AAPL
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *QuoteRequest) Reset() {
	*x = QuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteRequest) ProtoMessage() {}

func (x *QuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteRequest.ProtoReflect.Descriptor instead.
func (*QuoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{2}
}

func (x *QuoteRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type QuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the stock symbol
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// the asking price
	AskPrice float64 `protobuf:"fixed64,2,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	// the bidding price
	BidPrice float64 `protobuf:"fixed64,3,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	// the ask size
	AskSize int32 `protobuf:"varint,4,opt,name=ask_size,json=askSize,proto3" json:"ask_size,omitempty"`
	// the bid size
	BidSize int32 `protobuf:"varint,5,opt,name=bid_size,json=bidSize,proto3" json:"bid_size,omitempty"`
	// the UTC timestamp of the quote
	Timestamp string `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *QuoteResponse) Reset() {
	*x = QuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteResponse) ProtoMessage() {}

func (x *QuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteResponse.ProtoReflect.Descriptor instead.
func (*QuoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{3}
}

func (x *QuoteResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *QuoteResponse) GetAskPrice() float64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *QuoteResponse) GetBidPrice() float64 {
	if x != nil {
		return x.BidPrice
	}
	return 0
}

func (x *QuoteResponse) GetAskSize() int32 {
	if x != nil {
		return x.AskSize
	}
	return 0
}

func (x *QuoteResponse) GetBidSize() int32 {
	if x != nil {
		return x.BidSize
	}
	return 0
}

func (x *QuoteResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

// Get the historic open-close for a given day
type HistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the stock symbol e.g AAPL
	Stock string `protobuf:"bytes,1,opt,name=stock,proto3" json:"stock,omitempty"`
	// date to retrieve as YYYY-MM-DD
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *HistoryRequest) Reset() {
	*x = HistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryRequest) ProtoMessage() {}

func (x *HistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryRequest.ProtoReflect.Descriptor instead.
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{4}
}

func (x *HistoryRequest) GetStock() string {
	if x != nil {
		return x.Stock
	}
	return ""
}

func (x *HistoryRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type HistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the stock symbol
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// the open price
	Open float64 `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	// the close price
	Close float64 `protobuf:"fixed64,3,opt,name=close,proto3" json:"close,omitempty"`
	// the peak price
	High float64 `protobuf:"fixed64,4,opt,name=high,proto3" json:"high,omitempty"`
	// the low price
	Low float64 `protobuf:"fixed64,5,opt,name=low,proto3" json:"low,omitempty"`
	// the volume
	Volume int32 `protobuf:"varint,6,opt,name=volume,proto3" json:"volume,omitempty"`
	// the date
	Date string `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *HistoryResponse) Reset() {
	*x = HistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryResponse) ProtoMessage() {}

func (x *HistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryResponse.ProtoReflect.Descriptor instead.
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{5}
}

func (x *HistoryResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *HistoryResponse) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *HistoryResponse) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *HistoryResponse) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *HistoryResponse) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *HistoryResponse) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *HistoryResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_proto_stock_proto protoreflect.FileDescriptor

var file_proto_stock_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x22, 0x3d, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x26, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0xb5, 0x01, 0x0a, 0x0d, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x69, 0x64, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x3a, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xa5, 0x01,
	0x0a, 0x0f, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x32, 0xaf, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x34, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x13,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_stock_proto_rawDescOnce sync.Once
	file_proto_stock_proto_rawDescData = file_proto_stock_proto_rawDesc
)

func file_proto_stock_proto_rawDescGZIP() []byte {
	file_proto_stock_proto_rawDescOnce.Do(func() {
		file_proto_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stock_proto_rawDescData)
	})
	return file_proto_stock_proto_rawDescData
}

var file_proto_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_stock_proto_goTypes = []interface{}{
	(*PriceRequest)(nil),    // 0: stock.PriceRequest
	(*PriceResponse)(nil),   // 1: stock.PriceResponse
	(*QuoteRequest)(nil),    // 2: stock.QuoteRequest
	(*QuoteResponse)(nil),   // 3: stock.QuoteResponse
	(*HistoryRequest)(nil),  // 4: stock.HistoryRequest
	(*HistoryResponse)(nil), // 5: stock.HistoryResponse
}
var file_proto_stock_proto_depIdxs = []int32{
	2, // 0: stock.Stock.Quote:input_type -> stock.QuoteRequest
	0, // 1: stock.Stock.Price:input_type -> stock.PriceRequest
	4, // 2: stock.Stock.History:input_type -> stock.HistoryRequest
	3, // 3: stock.Stock.Quote:output_type -> stock.QuoteResponse
	1, // 4: stock.Stock.Price:output_type -> stock.PriceResponse
	5, // 5: stock.Stock.History:output_type -> stock.HistoryResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_stock_proto_init() }
func file_proto_stock_proto_init() {
	if File_proto_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceRequest); i {
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
		file_proto_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceResponse); i {
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
		file_proto_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteRequest); i {
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
		file_proto_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteResponse); i {
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
		file_proto_stock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryRequest); i {
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
		file_proto_stock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryResponse); i {
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
			RawDescriptor: file_proto_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_stock_proto_goTypes,
		DependencyIndexes: file_proto_stock_proto_depIdxs,
		MessageInfos:      file_proto_stock_proto_msgTypes,
	}.Build()
	File_proto_stock_proto = out.File
	file_proto_stock_proto_rawDesc = nil
	file_proto_stock_proto_goTypes = nil
	file_proto_stock_proto_depIdxs = nil
}

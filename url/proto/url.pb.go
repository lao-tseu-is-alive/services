// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/url.proto

package url

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

// Create a URL
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a unique id e.g uuid or my-url
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// destination url
	DestinationURL string `protobuf:"bytes,2,opt,name=destinationURL,proto3" json:"destinationURL,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateRequest) GetDestinationURL() string {
	if x != nil {
		return x.DestinationURL
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url *URLPair `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetUrl() *URLPair {
	if x != nil {
		return x.Url
	}
	return nil
}

// Delete a URL
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// specify id or shortURL
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// optional shortURL
	ShortURL string `protobuf:"bytes,2,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{3}
}

// Shorten a URL
type ShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the url to shorten
	DestinationURL string `protobuf:"bytes,1,opt,name=destinationURL,proto3" json:"destinationURL,omitempty"`
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{4}
}

func (x *ShortenRequest) GetDestinationURL() string {
	if x != nil {
		return x.DestinationURL
	}
	return ""
}

type ShortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the shortened url
	ShortURL string `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResponse.ProtoReflect.Descriptor instead.
func (*ShortenResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{5}
}

func (x *ShortenResponse) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type URLPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// url id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// destination url
	DestinationURL string `protobuf:"bytes,2,opt,name=destinationURL,proto3" json:"destinationURL,omitempty"`
	// shortened url
	ShortURL string `protobuf:"bytes,3,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
	// time of creation
	Created string `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	// The number of times the short URL has been resolved
	HitCount int64 `protobuf:"varint,5,opt,name=hitCount,proto3" json:"hitCount,omitempty"`
}

func (x *URLPair) Reset() {
	*x = URLPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLPair) ProtoMessage() {}

func (x *URLPair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLPair.ProtoReflect.Descriptor instead.
func (*URLPair) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{6}
}

func (x *URLPair) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *URLPair) GetDestinationURL() string {
	if x != nil {
		return x.DestinationURL
	}
	return ""
}

func (x *URLPair) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

func (x *URLPair) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *URLPair) GetHitCount() int64 {
	if x != nil {
		return x.HitCount
	}
	return 0
}

// List all the shortened URLs
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// filter by short URL, optional
	ShortURL string `protobuf:"bytes,2,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{7}
}

func (x *ListRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlPairs []*URLPair `protobuf:"bytes,1,rep,name=urlPairs,proto3" json:"urlPairs,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{8}
}

func (x *ListResponse) GetUrlPairs() []*URLPair {
	if x != nil {
		return x.UrlPairs
	}
	return nil
}

// Update the destination for a short URL
type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// update by id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// update by short url
	ShortURL string `protobuf:"bytes,2,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
	// the destination to update to
	DestinationURL string `protobuf:"bytes,3,opt,name=destinationURL,proto3" json:"destinationURL,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

func (x *UpdateRequest) GetDestinationURL() string {
	if x != nil {
		return x.DestinationURL
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{10}
}

// Resolve returns the destination URL of a short URL.
type ResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// short url to resolve
	ShortURL string `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *ResolveRequest) Reset() {
	*x = ResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequest) ProtoMessage() {}

func (x *ResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequest.ProtoReflect.Descriptor instead.
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{11}
}

func (x *ResolveRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type ResolveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationURL string `protobuf:"bytes,1,opt,name=destinationURL,proto3" json:"destinationURL,omitempty"`
}

func (x *ResolveResponse) Reset() {
	*x = ResolveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveResponse) ProtoMessage() {}

func (x *ResolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveResponse.ProtoReflect.Descriptor instead.
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{12}
}

func (x *ResolveResponse) GetDestinationURL() string {
	if x != nil {
		return x.DestinationURL
	}
	return ""
}

var File_proto_url_proto protoreflect.FileDescriptor

var file_proto_url_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x47, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x22,
	0x30, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x50, 0x61, 0x69, 0x72, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x3b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x10,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x38, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x55, 0x52,
	0x4c, 0x50, 0x61, 0x69, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x68, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x29, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x38, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x75, 0x72,
	0x6c, 0x50, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75,
	0x72, 0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x50, 0x61, 0x69, 0x72, 0x52, 0x08, 0x75, 0x72, 0x6c, 0x50,
	0x61, 0x69, 0x72, 0x73, 0x22, 0x63, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x39, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x52, 0x4c, 0x32, 0xc3, 0x02, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x33, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x72, 0x6c,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x75,
	0x72, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x10, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x75, 0x72, 0x6c,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x75, 0x72, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_url_proto_rawDescOnce sync.Once
	file_proto_url_proto_rawDescData = file_proto_url_proto_rawDesc
)

func file_proto_url_proto_rawDescGZIP() []byte {
	file_proto_url_proto_rawDescOnce.Do(func() {
		file_proto_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_url_proto_rawDescData)
	})
	return file_proto_url_proto_rawDescData
}

var file_proto_url_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_url_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),   // 0: url.CreateRequest
	(*CreateResponse)(nil),  // 1: url.CreateResponse
	(*DeleteRequest)(nil),   // 2: url.DeleteRequest
	(*DeleteResponse)(nil),  // 3: url.DeleteResponse
	(*ShortenRequest)(nil),  // 4: url.ShortenRequest
	(*ShortenResponse)(nil), // 5: url.ShortenResponse
	(*URLPair)(nil),         // 6: url.URLPair
	(*ListRequest)(nil),     // 7: url.ListRequest
	(*ListResponse)(nil),    // 8: url.ListResponse
	(*UpdateRequest)(nil),   // 9: url.UpdateRequest
	(*UpdateResponse)(nil),  // 10: url.UpdateResponse
	(*ResolveRequest)(nil),  // 11: url.ResolveRequest
	(*ResolveResponse)(nil), // 12: url.ResolveResponse
}
var file_proto_url_proto_depIdxs = []int32{
	6,  // 0: url.CreateResponse.url:type_name -> url.URLPair
	6,  // 1: url.ListResponse.urlPairs:type_name -> url.URLPair
	0,  // 2: url.Url.Create:input_type -> url.CreateRequest
	4,  // 3: url.Url.Shorten:input_type -> url.ShortenRequest
	7,  // 4: url.Url.List:input_type -> url.ListRequest
	11, // 5: url.Url.Resolve:input_type -> url.ResolveRequest
	9,  // 6: url.Url.Update:input_type -> url.UpdateRequest
	2,  // 7: url.Url.Delete:input_type -> url.DeleteRequest
	1,  // 8: url.Url.Create:output_type -> url.CreateResponse
	5,  // 9: url.Url.Shorten:output_type -> url.ShortenResponse
	8,  // 10: url.Url.List:output_type -> url.ListResponse
	12, // 11: url.Url.Resolve:output_type -> url.ResolveResponse
	10, // 12: url.Url.Update:output_type -> url.UpdateResponse
	3,  // 13: url.Url.Delete:output_type -> url.DeleteResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_url_proto_init() }
func file_proto_url_proto_init() {
	if File_proto_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_proto_url_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_url_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_proto_url_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenRequest); i {
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
		file_proto_url_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenResponse); i {
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
		file_proto_url_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLPair); i {
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
		file_proto_url_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_url_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_proto_url_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_url_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_proto_url_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveRequest); i {
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
		file_proto_url_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveResponse); i {
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
			RawDescriptor: file_proto_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_url_proto_goTypes,
		DependencyIndexes: file_proto_url_proto_depIdxs,
		MessageInfos:      file_proto_url_proto_msgTypes,
	}.Build()
	File_proto_url_proto = out.File
	file_proto_url_proto_rawDesc = nil
	file_proto_url_proto_goTypes = nil
	file_proto_url_proto_depIdxs = nil
}

// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/lists.proto

package lists

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "micro.dev/v4/service/client"
	server "micro.dev/v4/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Lists service

type ListsService interface {
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Events(ctx context.Context, in *EventsRequest, opts ...client.CallOption) (Lists_EventsService, error)
}

type listsService struct {
	c    client.Client
	name string
}

func NewListsService(name string, c client.Client) ListsService {
	return &listsService{
		c:    c,
		name: name,
	}
}

func (c *listsService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Lists.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Lists.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Lists.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Lists.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Lists.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsService) Events(ctx context.Context, in *EventsRequest, opts ...client.CallOption) (Lists_EventsService, error) {
	req := c.c.NewRequest(c.name, "Lists.Events", &EventsRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &listsServiceEvents{stream}, nil
}

type Lists_EventsService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*EventsResponse, error)
}

type listsServiceEvents struct {
	stream client.Stream
}

func (x *listsServiceEvents) Close() error {
	return x.stream.Close()
}

func (x *listsServiceEvents) Context() context.Context {
	return x.stream.Context()
}

func (x *listsServiceEvents) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *listsServiceEvents) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *listsServiceEvents) Recv() (*EventsResponse, error) {
	m := new(EventsResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Lists service

type ListsHandler interface {
	List(context.Context, *ListRequest, *ListResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Events(context.Context, *EventsRequest, Lists_EventsStream) error
}

func RegisterListsHandler(s server.Server, hdlr ListsHandler, opts ...server.HandlerOption) error {
	type lists interface {
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Events(ctx context.Context, stream server.Stream) error
	}
	type Lists struct {
		lists
	}
	h := &listsHandler{hdlr}
	return s.Handle(s.NewHandler(&Lists{h}, opts...))
}

type listsHandler struct {
	ListsHandler
}

func (h *listsHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.ListsHandler.List(ctx, in, out)
}

func (h *listsHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ListsHandler.Create(ctx, in, out)
}

func (h *listsHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.ListsHandler.Read(ctx, in, out)
}

func (h *listsHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ListsHandler.Delete(ctx, in, out)
}

func (h *listsHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ListsHandler.Update(ctx, in, out)
}

func (h *listsHandler) Events(ctx context.Context, stream server.Stream) error {
	m := new(EventsRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ListsHandler.Events(ctx, m, &listsEventsStream{stream})
}

type Lists_EventsStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*EventsResponse) error
}

type listsEventsStream struct {
	stream server.Stream
}

func (x *listsEventsStream) Close() error {
	return x.stream.Close()
}

func (x *listsEventsStream) Context() context.Context {
	return x.stream.Context()
}

func (x *listsEventsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *listsEventsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *listsEventsStream) Send(m *EventsResponse) error {
	return x.stream.Send(m)
}

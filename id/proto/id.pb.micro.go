// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/id.proto

package id

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

// Client API for Id service

type IdService interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error)
	Types(ctx context.Context, in *TypesRequest, opts ...client.CallOption) (*TypesResponse, error)
}

type idService struct {
	c    client.Client
	name string
}

func NewIdService(name string, c client.Client) IdService {
	return &idService{
		c:    c,
		name: name,
	}
}

func (c *idService) Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error) {
	req := c.c.NewRequest(c.name, "Id.Generate", in)
	out := new(GenerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idService) Types(ctx context.Context, in *TypesRequest, opts ...client.CallOption) (*TypesResponse, error) {
	req := c.c.NewRequest(c.name, "Id.Types", in)
	out := new(TypesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Id service

type IdHandler interface {
	Generate(context.Context, *GenerateRequest, *GenerateResponse) error
	Types(context.Context, *TypesRequest, *TypesResponse) error
}

func RegisterIdHandler(s server.Server, hdlr IdHandler, opts ...server.HandlerOption) error {
	type id interface {
		Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error
		Types(ctx context.Context, in *TypesRequest, out *TypesResponse) error
	}
	type Id struct {
		id
	}
	h := &idHandler{hdlr}
	return s.Handle(s.NewHandler(&Id{h}, opts...))
}

type idHandler struct {
	IdHandler
}

func (h *idHandler) Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error {
	return h.IdHandler.Generate(ctx, in, out)
}

func (h *idHandler) Types(ctx context.Context, in *TypesRequest, out *TypesResponse) error {
	return h.IdHandler.Types(ctx, in, out)
}

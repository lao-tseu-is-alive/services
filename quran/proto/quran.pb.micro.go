// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/quran.proto

package quran

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

// Client API for Quran service

type QuranService interface {
	Chapters(ctx context.Context, in *ChaptersRequest, opts ...client.CallOption) (*ChaptersResponse, error)
	Summary(ctx context.Context, in *SummaryRequest, opts ...client.CallOption) (*SummaryResponse, error)
	Verses(ctx context.Context, in *VersesRequest, opts ...client.CallOption) (*VersesResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type quranService struct {
	c    client.Client
	name string
}

func NewQuranService(name string, c client.Client) QuranService {
	return &quranService{
		c:    c,
		name: name,
	}
}

func (c *quranService) Chapters(ctx context.Context, in *ChaptersRequest, opts ...client.CallOption) (*ChaptersResponse, error) {
	req := c.c.NewRequest(c.name, "Quran.Chapters", in)
	out := new(ChaptersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quranService) Summary(ctx context.Context, in *SummaryRequest, opts ...client.CallOption) (*SummaryResponse, error) {
	req := c.c.NewRequest(c.name, "Quran.Summary", in)
	out := new(SummaryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quranService) Verses(ctx context.Context, in *VersesRequest, opts ...client.CallOption) (*VersesResponse, error) {
	req := c.c.NewRequest(c.name, "Quran.Verses", in)
	out := new(VersesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quranService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Quran.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Quran service

type QuranHandler interface {
	Chapters(context.Context, *ChaptersRequest, *ChaptersResponse) error
	Summary(context.Context, *SummaryRequest, *SummaryResponse) error
	Verses(context.Context, *VersesRequest, *VersesResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterQuranHandler(s server.Server, hdlr QuranHandler, opts ...server.HandlerOption) error {
	type quran interface {
		Chapters(ctx context.Context, in *ChaptersRequest, out *ChaptersResponse) error
		Summary(ctx context.Context, in *SummaryRequest, out *SummaryResponse) error
		Verses(ctx context.Context, in *VersesRequest, out *VersesResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
	}
	type Quran struct {
		quran
	}
	h := &quranHandler{hdlr}
	return s.Handle(s.NewHandler(&Quran{h}, opts...))
}

type quranHandler struct {
	QuranHandler
}

func (h *quranHandler) Chapters(ctx context.Context, in *ChaptersRequest, out *ChaptersResponse) error {
	return h.QuranHandler.Chapters(ctx, in, out)
}

func (h *quranHandler) Summary(ctx context.Context, in *SummaryRequest, out *SummaryResponse) error {
	return h.QuranHandler.Summary(ctx, in, out)
}

func (h *quranHandler) Verses(ctx context.Context, in *VersesRequest, out *VersesResponse) error {
	return h.QuranHandler.Verses(ctx, in, out)
}

func (h *quranHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.QuranHandler.Search(ctx, in, out)
}

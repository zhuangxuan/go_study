// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/getCaptcha/getCaptcha.proto

package go_micro_srv_getCaptcha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for GetCaptcha service

type GetCaptchaService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type getCaptchaService struct {
	c    client.Client
	name string
}

func NewGetCaptchaService(name string, c client.Client) GetCaptchaService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.getCaptcha"
	}
	return &getCaptchaService{
		c:    c,
		name: name,
	}
}

func (c *getCaptchaService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetCaptcha.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetCaptcha service

type GetCaptchaHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterGetCaptchaHandler(s server.Server, hdlr GetCaptchaHandler, opts ...server.HandlerOption) error {
	type getCaptcha interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type GetCaptcha struct {
		getCaptcha
	}
	h := &getCaptchaHandler{hdlr}
	return s.Handle(s.NewHandler(&GetCaptcha{h}, opts...))
}

type getCaptchaHandler struct {
	GetCaptchaHandler
}

func (h *getCaptchaHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.GetCaptchaHandler.Call(ctx, in, out)
}

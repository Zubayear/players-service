// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/players.proto

package players

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Players service

func NewPlayersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Players service

type PlayersService interface {
	Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error)
	Get(ctx context.Context, in *PlayerRequest, opts ...client.CallOption) (*PlayerResponse, error)
	GetAll(ctx context.Context, in *PlayersRequest, opts ...client.CallOption) (*PlayersResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Filter(ctx context.Context, in *FilterRequest, opts ...client.CallOption) (*PlayersResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
}

type playersService struct {
	c    client.Client
	name string
}

func NewPlayersService(name string, c client.Client) PlayersService {
	return &playersService{
		c:    c,
		name: name,
	}
}

func (c *playersService) Save(ctx context.Context, in *SaveRequest, opts ...client.CallOption) (*SaveResponse, error) {
	req := c.c.NewRequest(c.name, "Players.Save", in)
	out := new(SaveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playersService) Get(ctx context.Context, in *PlayerRequest, opts ...client.CallOption) (*PlayerResponse, error) {
	req := c.c.NewRequest(c.name, "Players.Get", in)
	out := new(PlayerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playersService) GetAll(ctx context.Context, in *PlayersRequest, opts ...client.CallOption) (*PlayersResponse, error) {
	req := c.c.NewRequest(c.name, "Players.GetAll", in)
	out := new(PlayersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playersService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Players.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playersService) Filter(ctx context.Context, in *FilterRequest, opts ...client.CallOption) (*PlayersResponse, error) {
	req := c.c.NewRequest(c.name, "Players.Filter", in)
	out := new(PlayersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playersService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Players.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Players service

type PlayersHandler interface {
	Save(context.Context, *SaveRequest, *SaveResponse) error
	Get(context.Context, *PlayerRequest, *PlayerResponse) error
	GetAll(context.Context, *PlayersRequest, *PlayersResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Filter(context.Context, *FilterRequest, *PlayersResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
}

func RegisterPlayersHandler(s server.Server, hdlr PlayersHandler, opts ...server.HandlerOption) error {
	type players interface {
		Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error
		Get(ctx context.Context, in *PlayerRequest, out *PlayerResponse) error
		GetAll(ctx context.Context, in *PlayersRequest, out *PlayersResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Filter(ctx context.Context, in *FilterRequest, out *PlayersResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
	}
	type Players struct {
		players
	}
	h := &playersHandler{hdlr}
	return s.Handle(s.NewHandler(&Players{h}, opts...))
}

type playersHandler struct {
	PlayersHandler
}

func (h *playersHandler) Save(ctx context.Context, in *SaveRequest, out *SaveResponse) error {
	return h.PlayersHandler.Save(ctx, in, out)
}

func (h *playersHandler) Get(ctx context.Context, in *PlayerRequest, out *PlayerResponse) error {
	return h.PlayersHandler.Get(ctx, in, out)
}

func (h *playersHandler) GetAll(ctx context.Context, in *PlayersRequest, out *PlayersResponse) error {
	return h.PlayersHandler.GetAll(ctx, in, out)
}

func (h *playersHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.PlayersHandler.Delete(ctx, in, out)
}

func (h *playersHandler) Filter(ctx context.Context, in *FilterRequest, out *PlayersResponse) error {
	return h.PlayersHandler.Filter(ctx, in, out)
}

func (h *playersHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.PlayersHandler.Update(ctx, in, out)
}
// Code generated by protoc-gen-go.
// source: search-only.proto
// DO NOT EDIT!

/*
Package search_only is a generated protocol buffer package.

It is generated from these files:
	search-only.proto

It has these top-level messages:
	Request
	Result
*/
package search_only

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type Result struct {
	Title   string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Snippet string `protobuf:"bytes,3,opt,name=snippet" json:"snippet,omitempty"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Google service

type GoogleClient interface {
	// Search returns a Google search result for the query.
	Search(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type googleClient struct {
	cc *grpc.ClientConn
}

func NewGoogleClient(cc *grpc.ClientConn) GoogleClient {
	return &googleClient{cc}
}

func (c *googleClient) Search(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/.Google/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Google service

type GoogleServer interface {
	// Search returns a Google search result for the query.
	Search(context.Context, *Request) (*Result, error)
}

func RegisterGoogleServer(s *grpc.Server, srv GoogleServer) {
	s.RegisterService(&_Google_serviceDesc, srv)
}

func _Google_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GoogleServer).Search(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Google_serviceDesc = grpc.ServiceDesc{
	ServiceName: ".Google",
	HandlerType: (*GoogleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Google_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

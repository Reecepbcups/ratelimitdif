// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ratelimit/v1/query.proto

package ratelimitv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_AllRateLimits_FullMethodName           = "/ratelimit.v1.Query/AllRateLimits"
	Query_RateLimit_FullMethodName               = "/ratelimit.v1.Query/RateLimit"
	Query_RateLimitsByChainId_FullMethodName     = "/ratelimit.v1.Query/RateLimitsByChainId"
	Query_RateLimitsByChannelId_FullMethodName   = "/ratelimit.v1.Query/RateLimitsByChannelId"
	Query_AllBlacklistedDenoms_FullMethodName    = "/ratelimit.v1.Query/AllBlacklistedDenoms"
	Query_AllWhitelistedAddresses_FullMethodName = "/ratelimit.v1.Query/AllWhitelistedAddresses"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Queries all rate limits
	AllRateLimits(ctx context.Context, in *QueryAllRateLimitsRequest, opts ...grpc.CallOption) (*QueryAllRateLimitsResponse, error)
	// Queries a specific rate limit by channel ID and denom
	// Ex:
	//   - /ratelimit/{channel_id}/by_denom?denom={denom}
	RateLimit(ctx context.Context, in *QueryRateLimitRequest, opts ...grpc.CallOption) (*QueryRateLimitResponse, error)
	// Queries all the rate limits for a given chain
	RateLimitsByChainId(ctx context.Context, in *QueryRateLimitsByChainIdRequest, opts ...grpc.CallOption) (*QueryRateLimitsByChainIdResponse, error)
	// Queries all the rate limits for a given channel ID
	RateLimitsByChannelId(ctx context.Context, in *QueryRateLimitsByChannelIdRequest, opts ...grpc.CallOption) (*QueryRateLimitsByChannelIdResponse, error)
	// Queries all blacklisted denoms
	AllBlacklistedDenoms(ctx context.Context, in *QueryAllBlacklistedDenomsRequest, opts ...grpc.CallOption) (*QueryAllBlacklistedDenomsResponse, error)
	// Queries all whitelisted address pairs
	AllWhitelistedAddresses(ctx context.Context, in *QueryAllWhitelistedAddressesRequest, opts ...grpc.CallOption) (*QueryAllWhitelistedAddressesResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AllRateLimits(ctx context.Context, in *QueryAllRateLimitsRequest, opts ...grpc.CallOption) (*QueryAllRateLimitsResponse, error) {
	out := new(QueryAllRateLimitsResponse)
	err := c.cc.Invoke(ctx, Query_AllRateLimits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RateLimit(ctx context.Context, in *QueryRateLimitRequest, opts ...grpc.CallOption) (*QueryRateLimitResponse, error) {
	out := new(QueryRateLimitResponse)
	err := c.cc.Invoke(ctx, Query_RateLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RateLimitsByChainId(ctx context.Context, in *QueryRateLimitsByChainIdRequest, opts ...grpc.CallOption) (*QueryRateLimitsByChainIdResponse, error) {
	out := new(QueryRateLimitsByChainIdResponse)
	err := c.cc.Invoke(ctx, Query_RateLimitsByChainId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RateLimitsByChannelId(ctx context.Context, in *QueryRateLimitsByChannelIdRequest, opts ...grpc.CallOption) (*QueryRateLimitsByChannelIdResponse, error) {
	out := new(QueryRateLimitsByChannelIdResponse)
	err := c.cc.Invoke(ctx, Query_RateLimitsByChannelId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllBlacklistedDenoms(ctx context.Context, in *QueryAllBlacklistedDenomsRequest, opts ...grpc.CallOption) (*QueryAllBlacklistedDenomsResponse, error) {
	out := new(QueryAllBlacklistedDenomsResponse)
	err := c.cc.Invoke(ctx, Query_AllBlacklistedDenoms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllWhitelistedAddresses(ctx context.Context, in *QueryAllWhitelistedAddressesRequest, opts ...grpc.CallOption) (*QueryAllWhitelistedAddressesResponse, error) {
	out := new(QueryAllWhitelistedAddressesResponse)
	err := c.cc.Invoke(ctx, Query_AllWhitelistedAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Queries all rate limits
	AllRateLimits(context.Context, *QueryAllRateLimitsRequest) (*QueryAllRateLimitsResponse, error)
	// Queries a specific rate limit by channel ID and denom
	// Ex:
	//   - /ratelimit/{channel_id}/by_denom?denom={denom}
	RateLimit(context.Context, *QueryRateLimitRequest) (*QueryRateLimitResponse, error)
	// Queries all the rate limits for a given chain
	RateLimitsByChainId(context.Context, *QueryRateLimitsByChainIdRequest) (*QueryRateLimitsByChainIdResponse, error)
	// Queries all the rate limits for a given channel ID
	RateLimitsByChannelId(context.Context, *QueryRateLimitsByChannelIdRequest) (*QueryRateLimitsByChannelIdResponse, error)
	// Queries all blacklisted denoms
	AllBlacklistedDenoms(context.Context, *QueryAllBlacklistedDenomsRequest) (*QueryAllBlacklistedDenomsResponse, error)
	// Queries all whitelisted address pairs
	AllWhitelistedAddresses(context.Context, *QueryAllWhitelistedAddressesRequest) (*QueryAllWhitelistedAddressesResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) AllRateLimits(context.Context, *QueryAllRateLimitsRequest) (*QueryAllRateLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllRateLimits not implemented")
}
func (UnimplementedQueryServer) RateLimit(context.Context, *QueryRateLimitRequest) (*QueryRateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateLimit not implemented")
}
func (UnimplementedQueryServer) RateLimitsByChainId(context.Context, *QueryRateLimitsByChainIdRequest) (*QueryRateLimitsByChainIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateLimitsByChainId not implemented")
}
func (UnimplementedQueryServer) RateLimitsByChannelId(context.Context, *QueryRateLimitsByChannelIdRequest) (*QueryRateLimitsByChannelIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateLimitsByChannelId not implemented")
}
func (UnimplementedQueryServer) AllBlacklistedDenoms(context.Context, *QueryAllBlacklistedDenomsRequest) (*QueryAllBlacklistedDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllBlacklistedDenoms not implemented")
}
func (UnimplementedQueryServer) AllWhitelistedAddresses(context.Context, *QueryAllWhitelistedAddressesRequest) (*QueryAllWhitelistedAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllWhitelistedAddresses not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_AllRateLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllRateLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllRateLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AllRateLimits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllRateLimits(ctx, req.(*QueryAllRateLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_RateLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RateLimit(ctx, req.(*QueryRateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RateLimitsByChainId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRateLimitsByChainIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RateLimitsByChainId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_RateLimitsByChainId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RateLimitsByChainId(ctx, req.(*QueryRateLimitsByChainIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RateLimitsByChannelId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRateLimitsByChannelIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RateLimitsByChannelId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_RateLimitsByChannelId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RateLimitsByChannelId(ctx, req.(*QueryRateLimitsByChannelIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllBlacklistedDenoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBlacklistedDenomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllBlacklistedDenoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AllBlacklistedDenoms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllBlacklistedDenoms(ctx, req.(*QueryAllBlacklistedDenomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllWhitelistedAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllWhitelistedAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllWhitelistedAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AllWhitelistedAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllWhitelistedAddresses(ctx, req.(*QueryAllWhitelistedAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ratelimit.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllRateLimits",
			Handler:    _Query_AllRateLimits_Handler,
		},
		{
			MethodName: "RateLimit",
			Handler:    _Query_RateLimit_Handler,
		},
		{
			MethodName: "RateLimitsByChainId",
			Handler:    _Query_RateLimitsByChainId_Handler,
		},
		{
			MethodName: "RateLimitsByChannelId",
			Handler:    _Query_RateLimitsByChannelId_Handler,
		},
		{
			MethodName: "AllBlacklistedDenoms",
			Handler:    _Query_AllBlacklistedDenoms_Handler,
		},
		{
			MethodName: "AllWhitelistedAddresses",
			Handler:    _Query_AllWhitelistedAddresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ratelimit/v1/query.proto",
}

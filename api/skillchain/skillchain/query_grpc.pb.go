// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: skillchain/skillchain/query.proto

package skillchain

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
	Query_Params_FullMethodName           = "/skillchain.skillchain.Query/Params"
	Query_TokenInfo_FullMethodName        = "/skillchain.skillchain.Query/TokenInfo"
	Query_VUSDTreasury_FullMethodName     = "/skillchain.skillchain.Query/VUSDTreasury"
	Query_UserVUSDPosition_FullMethodName = "/skillchain.skillchain.Query/UserVUSDPosition"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	TokenInfo(ctx context.Context, in *QueryTokenInfoRequest, opts ...grpc.CallOption) (*QueryTokenInfoResponse, error)
	VUSDTreasury(ctx context.Context, in *QueryVUSDTreasuryRequest, opts ...grpc.CallOption) (*QueryVUSDTreasuryResponse, error)
	UserVUSDPosition(ctx context.Context, in *QueryUserVUSDPositionRequest, opts ...grpc.CallOption) (*QueryUserVUSDPositionResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TokenInfo(ctx context.Context, in *QueryTokenInfoRequest, opts ...grpc.CallOption) (*QueryTokenInfoResponse, error) {
	out := new(QueryTokenInfoResponse)
	err := c.cc.Invoke(ctx, Query_TokenInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VUSDTreasury(ctx context.Context, in *QueryVUSDTreasuryRequest, opts ...grpc.CallOption) (*QueryVUSDTreasuryResponse, error) {
	out := new(QueryVUSDTreasuryResponse)
	err := c.cc.Invoke(ctx, Query_VUSDTreasury_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UserVUSDPosition(ctx context.Context, in *QueryUserVUSDPositionRequest, opts ...grpc.CallOption) (*QueryUserVUSDPositionResponse, error) {
	out := new(QueryUserVUSDPositionResponse)
	err := c.cc.Invoke(ctx, Query_UserVUSDPosition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	TokenInfo(context.Context, *QueryTokenInfoRequest) (*QueryTokenInfoResponse, error)
	VUSDTreasury(context.Context, *QueryVUSDTreasuryRequest) (*QueryVUSDTreasuryResponse, error)
	UserVUSDPosition(context.Context, *QueryUserVUSDPositionRequest) (*QueryUserVUSDPositionResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) TokenInfo(context.Context, *QueryTokenInfoRequest) (*QueryTokenInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenInfo not implemented")
}
func (UnimplementedQueryServer) VUSDTreasury(context.Context, *QueryVUSDTreasuryRequest) (*QueryVUSDTreasuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VUSDTreasury not implemented")
}
func (UnimplementedQueryServer) UserVUSDPosition(context.Context, *QueryUserVUSDPositionRequest) (*QueryUserVUSDPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserVUSDPosition not implemented")
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

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TokenInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenInfo(ctx, req.(*QueryTokenInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VUSDTreasury_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVUSDTreasuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VUSDTreasury(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_VUSDTreasury_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VUSDTreasury(ctx, req.(*QueryVUSDTreasuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UserVUSDPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserVUSDPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UserVUSDPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_UserVUSDPosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UserVUSDPosition(ctx, req.(*QueryUserVUSDPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skillchain.skillchain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "TokenInfo",
			Handler:    _Query_TokenInfo_Handler,
		},
		{
			MethodName: "VUSDTreasury",
			Handler:    _Query_VUSDTreasury_Handler,
		},
		{
			MethodName: "UserVUSDPosition",
			Handler:    _Query_UserVUSDPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skillchain/skillchain/query.proto",
}

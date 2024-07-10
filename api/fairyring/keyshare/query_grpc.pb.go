// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: fairyring/keyshare/query.proto

package keyshare

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
	Query_Params_FullMethodName                = "/fairyring.keyshare.Query/Params"
	Query_Commitments_FullMethodName           = "/fairyring.keyshare.Query/Commitments"
	Query_ValidatorSet_FullMethodName          = "/fairyring.keyshare.Query/ValidatorSet"
	Query_ValidatorSetAll_FullMethodName       = "/fairyring.keyshare.Query/ValidatorSetAll"
	Query_KeyShare_FullMethodName              = "/fairyring.keyshare.Query/KeyShare"
	Query_KeyShareAll_FullMethodName           = "/fairyring.keyshare.Query/KeyShareAll"
	Query_AggregatedKeyShare_FullMethodName    = "/fairyring.keyshare.Query/AggregatedKeyShare"
	Query_AggregatedKeyShareAll_FullMethodName = "/fairyring.keyshare.Query/AggregatedKeyShareAll"
	Query_PubKey_FullMethodName                = "/fairyring.keyshare.Query/PubKey"
	Query_AuthorizedAddress_FullMethodName     = "/fairyring.keyshare.Query/AuthorizedAddress"
	Query_AuthorizedAddressAll_FullMethodName  = "/fairyring.keyshare.Query/AuthorizedAddressAll"
	Query_GeneralKeyShare_FullMethodName       = "/fairyring.keyshare.Query/GeneralKeyShare"
	Query_GeneralKeyShareAll_FullMethodName    = "/fairyring.keyshare.Query/GeneralKeyShareAll"
	Query_VerifiableRandomness_FullMethodName  = "/fairyring.keyshare.Query/VerifiableRandomness"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	Commitments(ctx context.Context, in *QueryCommitmentsRequest, opts ...grpc.CallOption) (*QueryCommitmentsResponse, error)
	// Queries a ValidatorSet by index.
	ValidatorSet(ctx context.Context, in *QueryGetValidatorSetRequest, opts ...grpc.CallOption) (*QueryGetValidatorSetResponse, error)
	// Queries a list of ValidatorSet items.
	ValidatorSetAll(ctx context.Context, in *QueryAllValidatorSetRequest, opts ...grpc.CallOption) (*QueryAllValidatorSetResponse, error)
	// Queries a KeyShare by index.
	KeyShare(ctx context.Context, in *QueryGetKeyShareRequest, opts ...grpc.CallOption) (*QueryGetKeyShareResponse, error)
	// Queries a list of KeyShare items.
	KeyShareAll(ctx context.Context, in *QueryAllKeyShareRequest, opts ...grpc.CallOption) (*QueryAllKeyShareResponse, error)
	// Queries a list of AggregatedKeyShare items.
	AggregatedKeyShare(ctx context.Context, in *QueryGetAggregatedKeyShareRequest, opts ...grpc.CallOption) (*QueryGetAggregatedKeyShareResponse, error)
	AggregatedKeyShareAll(ctx context.Context, in *QueryAllAggregatedKeyShareRequest, opts ...grpc.CallOption) (*QueryAllAggregatedKeyShareResponse, error)
	// Queries the public keys
	PubKey(ctx context.Context, in *QueryPubKeyRequest, opts ...grpc.CallOption) (*QueryPubKeyResponse, error)
	// Queries a list of AuthorizedAddress items.
	AuthorizedAddress(ctx context.Context, in *QueryGetAuthorizedAddressRequest, opts ...grpc.CallOption) (*QueryGetAuthorizedAddressResponse, error)
	AuthorizedAddressAll(ctx context.Context, in *QueryAllAuthorizedAddressRequest, opts ...grpc.CallOption) (*QueryAllAuthorizedAddressResponse, error)
	// Queries a list of GeneralKeyShare items.
	GeneralKeyShare(ctx context.Context, in *QueryGetGeneralKeyShareRequest, opts ...grpc.CallOption) (*QueryGetGeneralKeyShareResponse, error)
	GeneralKeyShareAll(ctx context.Context, in *QueryAllGeneralKeyShareRequest, opts ...grpc.CallOption) (*QueryAllGeneralKeyShareResponse, error)
	VerifiableRandomness(ctx context.Context, in *QueryVerifiableRandomnessQuery, opts ...grpc.CallOption) (*QueryVerifiableRandomnessResponse, error)
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

func (c *queryClient) Commitments(ctx context.Context, in *QueryCommitmentsRequest, opts ...grpc.CallOption) (*QueryCommitmentsResponse, error) {
	out := new(QueryCommitmentsResponse)
	err := c.cc.Invoke(ctx, Query_Commitments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorSet(ctx context.Context, in *QueryGetValidatorSetRequest, opts ...grpc.CallOption) (*QueryGetValidatorSetResponse, error) {
	out := new(QueryGetValidatorSetResponse)
	err := c.cc.Invoke(ctx, Query_ValidatorSet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorSetAll(ctx context.Context, in *QueryAllValidatorSetRequest, opts ...grpc.CallOption) (*QueryAllValidatorSetResponse, error) {
	out := new(QueryAllValidatorSetResponse)
	err := c.cc.Invoke(ctx, Query_ValidatorSetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) KeyShare(ctx context.Context, in *QueryGetKeyShareRequest, opts ...grpc.CallOption) (*QueryGetKeyShareResponse, error) {
	out := new(QueryGetKeyShareResponse)
	err := c.cc.Invoke(ctx, Query_KeyShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) KeyShareAll(ctx context.Context, in *QueryAllKeyShareRequest, opts ...grpc.CallOption) (*QueryAllKeyShareResponse, error) {
	out := new(QueryAllKeyShareResponse)
	err := c.cc.Invoke(ctx, Query_KeyShareAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AggregatedKeyShare(ctx context.Context, in *QueryGetAggregatedKeyShareRequest, opts ...grpc.CallOption) (*QueryGetAggregatedKeyShareResponse, error) {
	out := new(QueryGetAggregatedKeyShareResponse)
	err := c.cc.Invoke(ctx, Query_AggregatedKeyShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AggregatedKeyShareAll(ctx context.Context, in *QueryAllAggregatedKeyShareRequest, opts ...grpc.CallOption) (*QueryAllAggregatedKeyShareResponse, error) {
	out := new(QueryAllAggregatedKeyShareResponse)
	err := c.cc.Invoke(ctx, Query_AggregatedKeyShareAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PubKey(ctx context.Context, in *QueryPubKeyRequest, opts ...grpc.CallOption) (*QueryPubKeyResponse, error) {
	out := new(QueryPubKeyResponse)
	err := c.cc.Invoke(ctx, Query_PubKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AuthorizedAddress(ctx context.Context, in *QueryGetAuthorizedAddressRequest, opts ...grpc.CallOption) (*QueryGetAuthorizedAddressResponse, error) {
	out := new(QueryGetAuthorizedAddressResponse)
	err := c.cc.Invoke(ctx, Query_AuthorizedAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AuthorizedAddressAll(ctx context.Context, in *QueryAllAuthorizedAddressRequest, opts ...grpc.CallOption) (*QueryAllAuthorizedAddressResponse, error) {
	out := new(QueryAllAuthorizedAddressResponse)
	err := c.cc.Invoke(ctx, Query_AuthorizedAddressAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GeneralKeyShare(ctx context.Context, in *QueryGetGeneralKeyShareRequest, opts ...grpc.CallOption) (*QueryGetGeneralKeyShareResponse, error) {
	out := new(QueryGetGeneralKeyShareResponse)
	err := c.cc.Invoke(ctx, Query_GeneralKeyShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GeneralKeyShareAll(ctx context.Context, in *QueryAllGeneralKeyShareRequest, opts ...grpc.CallOption) (*QueryAllGeneralKeyShareResponse, error) {
	out := new(QueryAllGeneralKeyShareResponse)
	err := c.cc.Invoke(ctx, Query_GeneralKeyShareAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VerifiableRandomness(ctx context.Context, in *QueryVerifiableRandomnessQuery, opts ...grpc.CallOption) (*QueryVerifiableRandomnessResponse, error) {
	out := new(QueryVerifiableRandomnessResponse)
	err := c.cc.Invoke(ctx, Query_VerifiableRandomness_FullMethodName, in, out, opts...)
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
	Commitments(context.Context, *QueryCommitmentsRequest) (*QueryCommitmentsResponse, error)
	// Queries a ValidatorSet by index.
	ValidatorSet(context.Context, *QueryGetValidatorSetRequest) (*QueryGetValidatorSetResponse, error)
	// Queries a list of ValidatorSet items.
	ValidatorSetAll(context.Context, *QueryAllValidatorSetRequest) (*QueryAllValidatorSetResponse, error)
	// Queries a KeyShare by index.
	KeyShare(context.Context, *QueryGetKeyShareRequest) (*QueryGetKeyShareResponse, error)
	// Queries a list of KeyShare items.
	KeyShareAll(context.Context, *QueryAllKeyShareRequest) (*QueryAllKeyShareResponse, error)
	// Queries a list of AggregatedKeyShare items.
	AggregatedKeyShare(context.Context, *QueryGetAggregatedKeyShareRequest) (*QueryGetAggregatedKeyShareResponse, error)
	AggregatedKeyShareAll(context.Context, *QueryAllAggregatedKeyShareRequest) (*QueryAllAggregatedKeyShareResponse, error)
	// Queries the public keys
	PubKey(context.Context, *QueryPubKeyRequest) (*QueryPubKeyResponse, error)
	// Queries a list of AuthorizedAddress items.
	AuthorizedAddress(context.Context, *QueryGetAuthorizedAddressRequest) (*QueryGetAuthorizedAddressResponse, error)
	AuthorizedAddressAll(context.Context, *QueryAllAuthorizedAddressRequest) (*QueryAllAuthorizedAddressResponse, error)
	// Queries a list of GeneralKeyShare items.
	GeneralKeyShare(context.Context, *QueryGetGeneralKeyShareRequest) (*QueryGetGeneralKeyShareResponse, error)
	GeneralKeyShareAll(context.Context, *QueryAllGeneralKeyShareRequest) (*QueryAllGeneralKeyShareResponse, error)
	VerifiableRandomness(context.Context, *QueryVerifiableRandomnessQuery) (*QueryVerifiableRandomnessResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Commitments(context.Context, *QueryCommitmentsRequest) (*QueryCommitmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commitments not implemented")
}
func (UnimplementedQueryServer) ValidatorSet(context.Context, *QueryGetValidatorSetRequest) (*QueryGetValidatorSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSet not implemented")
}
func (UnimplementedQueryServer) ValidatorSetAll(context.Context, *QueryAllValidatorSetRequest) (*QueryAllValidatorSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSetAll not implemented")
}
func (UnimplementedQueryServer) KeyShare(context.Context, *QueryGetKeyShareRequest) (*QueryGetKeyShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyShare not implemented")
}
func (UnimplementedQueryServer) KeyShareAll(context.Context, *QueryAllKeyShareRequest) (*QueryAllKeyShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyShareAll not implemented")
}
func (UnimplementedQueryServer) AggregatedKeyShare(context.Context, *QueryGetAggregatedKeyShareRequest) (*QueryGetAggregatedKeyShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AggregatedKeyShare not implemented")
}
func (UnimplementedQueryServer) AggregatedKeyShareAll(context.Context, *QueryAllAggregatedKeyShareRequest) (*QueryAllAggregatedKeyShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AggregatedKeyShareAll not implemented")
}
func (UnimplementedQueryServer) PubKey(context.Context, *QueryPubKeyRequest) (*QueryPubKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PubKey not implemented")
}
func (UnimplementedQueryServer) AuthorizedAddress(context.Context, *QueryGetAuthorizedAddressRequest) (*QueryGetAuthorizedAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizedAddress not implemented")
}
func (UnimplementedQueryServer) AuthorizedAddressAll(context.Context, *QueryAllAuthorizedAddressRequest) (*QueryAllAuthorizedAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizedAddressAll not implemented")
}
func (UnimplementedQueryServer) GeneralKeyShare(context.Context, *QueryGetGeneralKeyShareRequest) (*QueryGetGeneralKeyShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneralKeyShare not implemented")
}
func (UnimplementedQueryServer) GeneralKeyShareAll(context.Context, *QueryAllGeneralKeyShareRequest) (*QueryAllGeneralKeyShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneralKeyShareAll not implemented")
}
func (UnimplementedQueryServer) VerifiableRandomness(context.Context, *QueryVerifiableRandomnessQuery) (*QueryVerifiableRandomnessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifiableRandomness not implemented")
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

func _Query_Commitments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommitmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Commitments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Commitments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Commitments(ctx, req.(*QueryCommitmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetValidatorSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ValidatorSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorSet(ctx, req.(*QueryGetValidatorSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorSetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllValidatorSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorSetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ValidatorSetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorSetAll(ctx, req.(*QueryAllValidatorSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_KeyShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetKeyShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).KeyShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_KeyShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).KeyShare(ctx, req.(*QueryGetKeyShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_KeyShareAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllKeyShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).KeyShareAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_KeyShareAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).KeyShareAll(ctx, req.(*QueryAllKeyShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AggregatedKeyShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAggregatedKeyShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AggregatedKeyShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AggregatedKeyShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AggregatedKeyShare(ctx, req.(*QueryGetAggregatedKeyShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AggregatedKeyShareAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAggregatedKeyShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AggregatedKeyShareAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AggregatedKeyShareAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AggregatedKeyShareAll(ctx, req.(*QueryAllAggregatedKeyShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPubKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PubKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PubKey(ctx, req.(*QueryPubKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AuthorizedAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAuthorizedAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AuthorizedAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AuthorizedAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AuthorizedAddress(ctx, req.(*QueryGetAuthorizedAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AuthorizedAddressAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAuthorizedAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AuthorizedAddressAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AuthorizedAddressAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AuthorizedAddressAll(ctx, req.(*QueryAllAuthorizedAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GeneralKeyShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetGeneralKeyShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GeneralKeyShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GeneralKeyShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GeneralKeyShare(ctx, req.(*QueryGetGeneralKeyShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GeneralKeyShareAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllGeneralKeyShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GeneralKeyShareAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GeneralKeyShareAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GeneralKeyShareAll(ctx, req.(*QueryAllGeneralKeyShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VerifiableRandomness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVerifiableRandomnessQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VerifiableRandomness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_VerifiableRandomness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VerifiableRandomness(ctx, req.(*QueryVerifiableRandomnessQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fairyring.keyshare.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Commitments",
			Handler:    _Query_Commitments_Handler,
		},
		{
			MethodName: "ValidatorSet",
			Handler:    _Query_ValidatorSet_Handler,
		},
		{
			MethodName: "ValidatorSetAll",
			Handler:    _Query_ValidatorSetAll_Handler,
		},
		{
			MethodName: "KeyShare",
			Handler:    _Query_KeyShare_Handler,
		},
		{
			MethodName: "KeyShareAll",
			Handler:    _Query_KeyShareAll_Handler,
		},
		{
			MethodName: "AggregatedKeyShare",
			Handler:    _Query_AggregatedKeyShare_Handler,
		},
		{
			MethodName: "AggregatedKeyShareAll",
			Handler:    _Query_AggregatedKeyShareAll_Handler,
		},
		{
			MethodName: "PubKey",
			Handler:    _Query_PubKey_Handler,
		},
		{
			MethodName: "AuthorizedAddress",
			Handler:    _Query_AuthorizedAddress_Handler,
		},
		{
			MethodName: "AuthorizedAddressAll",
			Handler:    _Query_AuthorizedAddressAll_Handler,
		},
		{
			MethodName: "GeneralKeyShare",
			Handler:    _Query_GeneralKeyShare_Handler,
		},
		{
			MethodName: "GeneralKeyShareAll",
			Handler:    _Query_GeneralKeyShareAll_Handler,
		},
		{
			MethodName: "VerifiableRandomness",
			Handler:    _Query_VerifiableRandomness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fairyring/keyshare/query.proto",
}

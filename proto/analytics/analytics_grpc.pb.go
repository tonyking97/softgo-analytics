// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package analytics_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AnalyticsClient is the client API for Analytics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	AddStore(ctx context.Context, in *AddStoreRequest, opts ...grpc.CallOption) (*AddStoreResponse, error)
	DeleteStore(ctx context.Context, in *DeleteStoreRequest, opts ...grpc.CallOption) (*DeleteStoreResponse, error)
	PeriodicSalesAmount(ctx context.Context, in *PeriodicSalesAmountRequest, opts ...grpc.CallOption) (*PeriodicSalesAmountResponse, error)
	PeriodicStoreSalesAmount(ctx context.Context, in *PeriodicStoreSalesAmountRequest, opts ...grpc.CallOption) (*PeriodicStoreSalesAmountResponse, error)
	TopSoldItems(ctx context.Context, in *TopSoldItemsRequest, opts ...grpc.CallOption) (*TopSoldItemsResponse, error)
	LeastSoldItems(ctx context.Context, in *LeastSoldItemsRequest, opts ...grpc.CallOption) (*LeastSoldItemsResponse, error)
	TopRevenueItems(ctx context.Context, in *TopRevenueItemsRequest, opts ...grpc.CallOption) (*TopRevenueItemsResponse, error)
	LeastRevenueItems(ctx context.Context, in *LeastRevenueItemsRequest, opts ...grpc.CallOption) (*LeastRevenueItemsResponse, error)
	StoreRevenueData(ctx context.Context, in *StoreRevenueDataRequest, opts ...grpc.CallOption) (*StoreRevenueDataResponse, error)
	StoreSalesData(ctx context.Context, in *StoreSalesDataRequest, opts ...grpc.CallOption) (*StoreSalesDataResponse, error)
	AverageSalesAmount(ctx context.Context, in *AverageSalesAmountRequest, opts ...grpc.CallOption) (*AverageSalesAmountResponse, error)
	AverageSoldQuantity(ctx context.Context, in *AverageSoldQuantityRequest, opts ...grpc.CallOption) (*AverageSoldQuantityResponse, error)
	TotalSalesAmount(ctx context.Context, in *TotalSalesAmountRequest, opts ...grpc.CallOption) (*TotalSalesAmountResponse, error)
}

type analyticsClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsClient(cc grpc.ClientConnInterface) AnalyticsClient {
	return &analyticsClient{cc}
}

func (c *analyticsClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) AddStore(ctx context.Context, in *AddStoreRequest, opts ...grpc.CallOption) (*AddStoreResponse, error) {
	out := new(AddStoreResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/AddStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) DeleteStore(ctx context.Context, in *DeleteStoreRequest, opts ...grpc.CallOption) (*DeleteStoreResponse, error) {
	out := new(DeleteStoreResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/DeleteStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) PeriodicSalesAmount(ctx context.Context, in *PeriodicSalesAmountRequest, opts ...grpc.CallOption) (*PeriodicSalesAmountResponse, error) {
	out := new(PeriodicSalesAmountResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/PeriodicSalesAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) PeriodicStoreSalesAmount(ctx context.Context, in *PeriodicStoreSalesAmountRequest, opts ...grpc.CallOption) (*PeriodicStoreSalesAmountResponse, error) {
	out := new(PeriodicStoreSalesAmountResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/PeriodicStoreSalesAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) TopSoldItems(ctx context.Context, in *TopSoldItemsRequest, opts ...grpc.CallOption) (*TopSoldItemsResponse, error) {
	out := new(TopSoldItemsResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/TopSoldItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) LeastSoldItems(ctx context.Context, in *LeastSoldItemsRequest, opts ...grpc.CallOption) (*LeastSoldItemsResponse, error) {
	out := new(LeastSoldItemsResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/LeastSoldItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) TopRevenueItems(ctx context.Context, in *TopRevenueItemsRequest, opts ...grpc.CallOption) (*TopRevenueItemsResponse, error) {
	out := new(TopRevenueItemsResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/TopRevenueItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) LeastRevenueItems(ctx context.Context, in *LeastRevenueItemsRequest, opts ...grpc.CallOption) (*LeastRevenueItemsResponse, error) {
	out := new(LeastRevenueItemsResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/LeastRevenueItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) StoreRevenueData(ctx context.Context, in *StoreRevenueDataRequest, opts ...grpc.CallOption) (*StoreRevenueDataResponse, error) {
	out := new(StoreRevenueDataResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/StoreRevenueData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) StoreSalesData(ctx context.Context, in *StoreSalesDataRequest, opts ...grpc.CallOption) (*StoreSalesDataResponse, error) {
	out := new(StoreSalesDataResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/StoreSalesData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) AverageSalesAmount(ctx context.Context, in *AverageSalesAmountRequest, opts ...grpc.CallOption) (*AverageSalesAmountResponse, error) {
	out := new(AverageSalesAmountResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/AverageSalesAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) AverageSoldQuantity(ctx context.Context, in *AverageSoldQuantityRequest, opts ...grpc.CallOption) (*AverageSoldQuantityResponse, error) {
	out := new(AverageSoldQuantityResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/AverageSoldQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) TotalSalesAmount(ctx context.Context, in *TotalSalesAmountRequest, opts ...grpc.CallOption) (*TotalSalesAmountResponse, error) {
	out := new(TotalSalesAmountResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/TotalSalesAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServer is the server API for Analytics service.
// All implementations must embed UnimplementedAnalyticsServer
// for forward compatibility
type AnalyticsServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	AddStore(context.Context, *AddStoreRequest) (*AddStoreResponse, error)
	DeleteStore(context.Context, *DeleteStoreRequest) (*DeleteStoreResponse, error)
	PeriodicSalesAmount(context.Context, *PeriodicSalesAmountRequest) (*PeriodicSalesAmountResponse, error)
	PeriodicStoreSalesAmount(context.Context, *PeriodicStoreSalesAmountRequest) (*PeriodicStoreSalesAmountResponse, error)
	TopSoldItems(context.Context, *TopSoldItemsRequest) (*TopSoldItemsResponse, error)
	LeastSoldItems(context.Context, *LeastSoldItemsRequest) (*LeastSoldItemsResponse, error)
	TopRevenueItems(context.Context, *TopRevenueItemsRequest) (*TopRevenueItemsResponse, error)
	LeastRevenueItems(context.Context, *LeastRevenueItemsRequest) (*LeastRevenueItemsResponse, error)
	StoreRevenueData(context.Context, *StoreRevenueDataRequest) (*StoreRevenueDataResponse, error)
	StoreSalesData(context.Context, *StoreSalesDataRequest) (*StoreSalesDataResponse, error)
	AverageSalesAmount(context.Context, *AverageSalesAmountRequest) (*AverageSalesAmountResponse, error)
	AverageSoldQuantity(context.Context, *AverageSoldQuantityRequest) (*AverageSoldQuantityResponse, error)
	TotalSalesAmount(context.Context, *TotalSalesAmountRequest) (*TotalSalesAmountResponse, error)
	mustEmbedUnimplementedAnalyticsServer()
}

// UnimplementedAnalyticsServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServer struct {
}

func (UnimplementedAnalyticsServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAnalyticsServer) AddStore(context.Context, *AddStoreRequest) (*AddStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStore not implemented")
}
func (UnimplementedAnalyticsServer) DeleteStore(context.Context, *DeleteStoreRequest) (*DeleteStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStore not implemented")
}
func (UnimplementedAnalyticsServer) PeriodicSalesAmount(context.Context, *PeriodicSalesAmountRequest) (*PeriodicSalesAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeriodicSalesAmount not implemented")
}
func (UnimplementedAnalyticsServer) PeriodicStoreSalesAmount(context.Context, *PeriodicStoreSalesAmountRequest) (*PeriodicStoreSalesAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeriodicStoreSalesAmount not implemented")
}
func (UnimplementedAnalyticsServer) TopSoldItems(context.Context, *TopSoldItemsRequest) (*TopSoldItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopSoldItems not implemented")
}
func (UnimplementedAnalyticsServer) LeastSoldItems(context.Context, *LeastSoldItemsRequest) (*LeastSoldItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeastSoldItems not implemented")
}
func (UnimplementedAnalyticsServer) TopRevenueItems(context.Context, *TopRevenueItemsRequest) (*TopRevenueItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopRevenueItems not implemented")
}
func (UnimplementedAnalyticsServer) LeastRevenueItems(context.Context, *LeastRevenueItemsRequest) (*LeastRevenueItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeastRevenueItems not implemented")
}
func (UnimplementedAnalyticsServer) StoreRevenueData(context.Context, *StoreRevenueDataRequest) (*StoreRevenueDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreRevenueData not implemented")
}
func (UnimplementedAnalyticsServer) StoreSalesData(context.Context, *StoreSalesDataRequest) (*StoreSalesDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreSalesData not implemented")
}
func (UnimplementedAnalyticsServer) AverageSalesAmount(context.Context, *AverageSalesAmountRequest) (*AverageSalesAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AverageSalesAmount not implemented")
}
func (UnimplementedAnalyticsServer) AverageSoldQuantity(context.Context, *AverageSoldQuantityRequest) (*AverageSoldQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AverageSoldQuantity not implemented")
}
func (UnimplementedAnalyticsServer) TotalSalesAmount(context.Context, *TotalSalesAmountRequest) (*TotalSalesAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSalesAmount not implemented")
}
func (UnimplementedAnalyticsServer) mustEmbedUnimplementedAnalyticsServer() {}

// UnsafeAnalyticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsServer will
// result in compilation errors.
type UnsafeAnalyticsServer interface {
	mustEmbedUnimplementedAnalyticsServer()
}

func RegisterAnalyticsServer(s *grpc.Server, srv AnalyticsServer) {
	s.RegisterService(&_Analytics_serviceDesc, srv)
}

func _Analytics_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_AddStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).AddStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/AddStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).AddStore(ctx, req.(*AddStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/DeleteStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).DeleteStore(ctx, req.(*DeleteStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_PeriodicSalesAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeriodicSalesAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).PeriodicSalesAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/PeriodicSalesAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).PeriodicSalesAmount(ctx, req.(*PeriodicSalesAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_PeriodicStoreSalesAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeriodicStoreSalesAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).PeriodicStoreSalesAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/PeriodicStoreSalesAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).PeriodicStoreSalesAmount(ctx, req.(*PeriodicStoreSalesAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_TopSoldItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopSoldItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).TopSoldItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/TopSoldItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).TopSoldItems(ctx, req.(*TopSoldItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_LeastSoldItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeastSoldItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).LeastSoldItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/LeastSoldItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).LeastSoldItems(ctx, req.(*LeastSoldItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_TopRevenueItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopRevenueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).TopRevenueItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/TopRevenueItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).TopRevenueItems(ctx, req.(*TopRevenueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_LeastRevenueItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeastRevenueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).LeastRevenueItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/LeastRevenueItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).LeastRevenueItems(ctx, req.(*LeastRevenueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_StoreRevenueData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRevenueDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).StoreRevenueData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/StoreRevenueData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).StoreRevenueData(ctx, req.(*StoreRevenueDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_StoreSalesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreSalesDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).StoreSalesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/StoreSalesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).StoreSalesData(ctx, req.(*StoreSalesDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_AverageSalesAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AverageSalesAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).AverageSalesAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/AverageSalesAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).AverageSalesAmount(ctx, req.(*AverageSalesAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_AverageSoldQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AverageSoldQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).AverageSoldQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/AverageSoldQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).AverageSoldQuantity(ctx, req.(*AverageSoldQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_TotalSalesAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TotalSalesAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).TotalSalesAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/TotalSalesAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).TotalSalesAmount(ctx, req.(*TotalSalesAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analytics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.Analytics",
	HandlerType: (*AnalyticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Analytics_Ping_Handler,
		},
		{
			MethodName: "AddStore",
			Handler:    _Analytics_AddStore_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _Analytics_DeleteStore_Handler,
		},
		{
			MethodName: "PeriodicSalesAmount",
			Handler:    _Analytics_PeriodicSalesAmount_Handler,
		},
		{
			MethodName: "PeriodicStoreSalesAmount",
			Handler:    _Analytics_PeriodicStoreSalesAmount_Handler,
		},
		{
			MethodName: "TopSoldItems",
			Handler:    _Analytics_TopSoldItems_Handler,
		},
		{
			MethodName: "LeastSoldItems",
			Handler:    _Analytics_LeastSoldItems_Handler,
		},
		{
			MethodName: "TopRevenueItems",
			Handler:    _Analytics_TopRevenueItems_Handler,
		},
		{
			MethodName: "LeastRevenueItems",
			Handler:    _Analytics_LeastRevenueItems_Handler,
		},
		{
			MethodName: "StoreRevenueData",
			Handler:    _Analytics_StoreRevenueData_Handler,
		},
		{
			MethodName: "StoreSalesData",
			Handler:    _Analytics_StoreSalesData_Handler,
		},
		{
			MethodName: "AverageSalesAmount",
			Handler:    _Analytics_AverageSalesAmount_Handler,
		},
		{
			MethodName: "AverageSoldQuantity",
			Handler:    _Analytics_AverageSoldQuantity_Handler,
		},
		{
			MethodName: "TotalSalesAmount",
			Handler:    _Analytics_TotalSalesAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analytics.proto",
}

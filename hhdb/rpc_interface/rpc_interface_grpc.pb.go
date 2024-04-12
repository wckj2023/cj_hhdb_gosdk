// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: hhdb/rpc_interface/rpc_interface.proto

package rpc_interface

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	rpc "hhdb_sdk/hhdb/rpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RpcInterface_Auth_FullMethodName                       = "/hhdb.rpc_interface.RpcInterface/Auth"
	RpcInterface_QueryErrMsg_FullMethodName                = "/hhdb.rpc_interface.RpcInterface/QueryErrMsg"
	RpcInterface_InsertTable_FullMethodName                = "/hhdb.rpc_interface.RpcInterface/InsertTable"
	RpcInterface_DelTable_FullMethodName                   = "/hhdb.rpc_interface.RpcInterface/DelTable"
	RpcInterface_ClearTable_FullMethodName                 = "/hhdb.rpc_interface.RpcInterface/ClearTable"
	RpcInterface_UpdateTable_FullMethodName                = "/hhdb.rpc_interface.RpcInterface/UpdateTable"
	RpcInterface_QueryTableList_FullMethodName             = "/hhdb.rpc_interface.RpcInterface/QueryTableList"
	RpcInterface_InsertPoints_FullMethodName               = "/hhdb.rpc_interface.RpcInterface/InsertPoints"
	RpcInterface_DelPoints_FullMethodName                  = "/hhdb.rpc_interface.RpcInterface/DelPoints"
	RpcInterface_UpdatePoints_FullMethodName               = "/hhdb.rpc_interface.RpcInterface/UpdatePoints"
	RpcInterface_QueryPoints_FullMethodName                = "/hhdb.rpc_interface.RpcInterface/QueryPoints"
	RpcInterface_QueryPointIdListByNameList_FullMethodName = "/hhdb.rpc_interface.RpcInterface/QueryPointIdListByNameList"
	RpcInterface_QueryPointInfoList_FullMethodName         = "/hhdb.rpc_interface.RpcInterface/QueryPointInfoList"
	RpcInterface_UpdateRealtimeValueList_FullMethodName    = "/hhdb.rpc_interface.RpcInterface/UpdateRealtimeValueList"
	RpcInterface_QueryRealtimeValueList_FullMethodName     = "/hhdb.rpc_interface.RpcInterface/QueryRealtimeValueList"
	RpcInterface_QueryHisRangeValueList_FullMethodName     = "/hhdb.rpc_interface.RpcInterface/QueryHisRangeValueList"
	RpcInterface_QueryHisResampleValueList_FullMethodName  = "/hhdb.rpc_interface.RpcInterface/QueryHisResampleValueList"
	RpcInterface_QueryHisTimePointValueList_FullMethodName = "/hhdb.rpc_interface.RpcInterface/QueryHisTimePointValueList"
	RpcInterface_InsertHisValueList_FullMethodName         = "/hhdb.rpc_interface.RpcInterface/InsertHisValueList"
)

// RpcInterfaceClient is the client API for RpcInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcInterfaceClient interface {
	// 功能：认证接口
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，返回用户ID,失败<0
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error)
	// 功能：根据错误码查询错误信息
	// 参数说明：入参、出差参考请求、响应体注释
	QueryErrMsg(ctx context.Context, in *QueryErrMsgReq, opts ...grpc.CallOption) (*QueryErrMsgReply, error)
	// 功能：表操作--添加
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，为新增表的ID,失败<0
	InsertTable(ctx context.Context, in *rpc.TableInfo, opts ...grpc.CallOption) (*CommonReply, error)
	// 功能：表操作--删除
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，失败<0
	// 备注：通过TableInfo中的tableId删除表,tableId<0时,使用tableName进行匹配删除
	DelTable(ctx context.Context, in *rpc.TableInfo, opts ...grpc.CallOption) (*CommonReply, error)
	// 功能：表操作--清空
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，为清空点组中清空的测点数量，失败<0
	// 备注：通过TableInfo中的tableId清空表,tableId<0时,使用tableName进行匹配清空
	ClearTable(ctx context.Context, in *rpc.TableInfo, opts ...grpc.CallOption) (*CommonReply, error)
	// 功能：表操作--更新
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，失败<0
	// 备注：通过TableInfo中的tableId更新表
	UpdateTable(ctx context.Context, in *rpc.TableInfo, opts ...grpc.CallOption) (*CommonReply, error)
	// 功能：点组操作--查询点组列表
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，失败<0
	// group_id>=0时,通过id获取点组信息,当group_id<0时,通过GroupInfo.group_name进行匹配获取数据,group_name为空且group_id<0时返回全部点组
	QueryTableList(ctx context.Context, in *QueryTableReq, opts ...grpc.CallOption) (*QueryTableReply, error)
	// 功能：测点操作--添加测点
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为添加成功的个数,失败<=0
	// 备注：CommonRes中resultList返回各个测点的ID，小于0时代表添加失败的错误码，全成功时为空
	InsertPoints(ctx context.Context, in *PointInfoListReq, opts ...grpc.CallOption) (*CommonReply, error)
	// 功能：测点操作--删除测点
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为删除成功的个数,失败<=0
	// 备注：通过PointInfo中的id进行关联删除，如果首个元素id为-1，则通过使用name进行匹配删除，
	DelPoints(ctx context.Context, in *PointInfoListReq, opts ...grpc.CallOption) (*CommonReply, error)
	// 功能：测点操作--更新测点基础信息
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为更新成功的个数,失败<=0
	// 备注：通过PointInfo中的id进行关联修改，如果首个元素id为-1，则通过使用name进行匹配更新
	// CommonRes中resultList返回更新成功的各个测点ID，返回小于0时代表添加失败的错误码，全成功时为空
	UpdatePoints(ctx context.Context, in *PointInfoListReq, opts ...grpc.CallOption) (*CommonReply, error)
	// 功能：测点操作--查询测点基础信息
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，为查询成功的个数,失败<0
	QueryPoints(ctx context.Context, in *QueryPointInfoReq, opts ...grpc.CallOption) (*QueryPointInfoReply, error)
	// 功能：测点操作--使用测点名查询测点ID
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	QueryPointIdListByNameList(ctx context.Context, in *NameList, opts ...grpc.CallOption) (*QueryPointIdListByNameListReply, error)
	// 功能：测点操作--使用测点名或ID批量查询测点信息
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	QueryPointInfoList(ctx context.Context, in *QueryRealtimeValueListReq, opts ...grpc.CallOption) (*QueryPointInfoReply, error)
	// 实时值--写入实时值
	// 参数说明：入参、出差参考请求、响应体注释
	// error_code.app_code说明：成功>=0,写入成功的个数,失败<0
	// 备注：CommonRes中result_list为获取各个值的成功状态，失败<0为对应的错误码，全成功时为空
	UpdateRealtimeValueList(ctx context.Context, in *UpdateRealtimeValueListReq, opts ...grpc.CallOption) (*CommonReply, error)
	// 实时值--通获取实时值
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	// 备注：DataListRes中resultList为获取各个值的成功状态，失败<0为对应的错误码，全成功时为空
	QueryRealtimeValueList(ctx context.Context, in *QueryRealtimeValueListReq, opts ...grpc.CallOption) (*ValueListReply, error)
	// 历史值--通过测点ID查询时间段范围内的数据值
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	QueryHisRangeValueList(ctx context.Context, in *QueryHisRangeValueListReq, opts ...grpc.CallOption) (*QueryHisValueListReply, error)
	// 历史值--通过测点ID查询时间段范围内重采样的数据值
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	QueryHisResampleValueList(ctx context.Context, in *QueryHisResamplesValueListReq, opts ...grpc.CallOption) (*QueryHisValueListReply, error)
	// 历史值--通过测点ID查询时间点的数据值
	// 参数说明：入参、出差参考请求、响应体注释
	// error_code.app_code说明：为查询成功的数据个数成功>0,失败<=0
	// 备注：DataListRes中value_list与测点ID下标一一对应，result_list为每个测点查询的错误码，全成功时为空
	QueryHisTimePointValueList(ctx context.Context, in *QueryHisTimePointValueListReq, opts ...grpc.CallOption) (*ValueListReply, error)
	// 历史值--通过测点ID写入测点历史值
	// 参数说明：入参、出差参考请求、响应体注释
	// error_code.app_code说明：成功>=0,查询成功的数据个数,失败<0
	InsertHisValueList(ctx context.Context, in *InsertHisValueListReq, opts ...grpc.CallOption) (*CommonReply, error)
}

type rpcInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcInterfaceClient(cc grpc.ClientConnInterface) RpcInterfaceClient {
	return &rpcInterfaceClient{cc}
}

func (c *rpcInterfaceClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, RpcInterface_Auth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryErrMsg(ctx context.Context, in *QueryErrMsgReq, opts ...grpc.CallOption) (*QueryErrMsgReply, error) {
	out := new(QueryErrMsgReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryErrMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) InsertTable(ctx context.Context, in *rpc.TableInfo, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_InsertTable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) DelTable(ctx context.Context, in *rpc.TableInfo, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_DelTable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) ClearTable(ctx context.Context, in *rpc.TableInfo, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_ClearTable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) UpdateTable(ctx context.Context, in *rpc.TableInfo, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_UpdateTable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryTableList(ctx context.Context, in *QueryTableReq, opts ...grpc.CallOption) (*QueryTableReply, error) {
	out := new(QueryTableReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryTableList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) InsertPoints(ctx context.Context, in *PointInfoListReq, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_InsertPoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) DelPoints(ctx context.Context, in *PointInfoListReq, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_DelPoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) UpdatePoints(ctx context.Context, in *PointInfoListReq, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_UpdatePoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryPoints(ctx context.Context, in *QueryPointInfoReq, opts ...grpc.CallOption) (*QueryPointInfoReply, error) {
	out := new(QueryPointInfoReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryPoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryPointIdListByNameList(ctx context.Context, in *NameList, opts ...grpc.CallOption) (*QueryPointIdListByNameListReply, error) {
	out := new(QueryPointIdListByNameListReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryPointIdListByNameList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryPointInfoList(ctx context.Context, in *QueryRealtimeValueListReq, opts ...grpc.CallOption) (*QueryPointInfoReply, error) {
	out := new(QueryPointInfoReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryPointInfoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) UpdateRealtimeValueList(ctx context.Context, in *UpdateRealtimeValueListReq, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_UpdateRealtimeValueList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryRealtimeValueList(ctx context.Context, in *QueryRealtimeValueListReq, opts ...grpc.CallOption) (*ValueListReply, error) {
	out := new(ValueListReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryRealtimeValueList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryHisRangeValueList(ctx context.Context, in *QueryHisRangeValueListReq, opts ...grpc.CallOption) (*QueryHisValueListReply, error) {
	out := new(QueryHisValueListReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryHisRangeValueList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryHisResampleValueList(ctx context.Context, in *QueryHisResamplesValueListReq, opts ...grpc.CallOption) (*QueryHisValueListReply, error) {
	out := new(QueryHisValueListReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryHisResampleValueList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) QueryHisTimePointValueList(ctx context.Context, in *QueryHisTimePointValueListReq, opts ...grpc.CallOption) (*ValueListReply, error) {
	out := new(ValueListReply)
	err := c.cc.Invoke(ctx, RpcInterface_QueryHisTimePointValueList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcInterfaceClient) InsertHisValueList(ctx context.Context, in *InsertHisValueListReq, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, RpcInterface_InsertHisValueList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcInterfaceServer is the server API for RpcInterface service.
// All implementations must embed UnimplementedRpcInterfaceServer
// for forward compatibility
type RpcInterfaceServer interface {
	// 功能：认证接口
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，返回用户ID,失败<0
	Auth(context.Context, *AuthReq) (*AuthReply, error)
	// 功能：根据错误码查询错误信息
	// 参数说明：入参、出差参考请求、响应体注释
	QueryErrMsg(context.Context, *QueryErrMsgReq) (*QueryErrMsgReply, error)
	// 功能：表操作--添加
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，为新增表的ID,失败<0
	InsertTable(context.Context, *rpc.TableInfo) (*CommonReply, error)
	// 功能：表操作--删除
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，失败<0
	// 备注：通过TableInfo中的tableId删除表,tableId<0时,使用tableName进行匹配删除
	DelTable(context.Context, *rpc.TableInfo) (*CommonReply, error)
	// 功能：表操作--清空
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，为清空点组中清空的测点数量，失败<0
	// 备注：通过TableInfo中的tableId清空表,tableId<0时,使用tableName进行匹配清空
	ClearTable(context.Context, *rpc.TableInfo) (*CommonReply, error)
	// 功能：表操作--更新
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，失败<0
	// 备注：通过TableInfo中的tableId更新表
	UpdateTable(context.Context, *rpc.TableInfo) (*CommonReply, error)
	// 功能：点组操作--查询点组列表
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，失败<0
	// group_id>=0时,通过id获取点组信息,当group_id<0时,通过GroupInfo.group_name进行匹配获取数据,group_name为空且group_id<0时返回全部点组
	QueryTableList(context.Context, *QueryTableReq) (*QueryTableReply, error)
	// 功能：测点操作--添加测点
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为添加成功的个数,失败<=0
	// 备注：CommonRes中resultList返回各个测点的ID，小于0时代表添加失败的错误码，全成功时为空
	InsertPoints(context.Context, *PointInfoListReq) (*CommonReply, error)
	// 功能：测点操作--删除测点
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为删除成功的个数,失败<=0
	// 备注：通过PointInfo中的id进行关联删除，如果首个元素id为-1，则通过使用name进行匹配删除，
	DelPoints(context.Context, *PointInfoListReq) (*CommonReply, error)
	// 功能：测点操作--更新测点基础信息
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为更新成功的个数,失败<=0
	// 备注：通过PointInfo中的id进行关联修改，如果首个元素id为-1，则通过使用name进行匹配更新
	// CommonRes中resultList返回更新成功的各个测点ID，返回小于0时代表添加失败的错误码，全成功时为空
	UpdatePoints(context.Context, *PointInfoListReq) (*CommonReply, error)
	// 功能：测点操作--查询测点基础信息
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>=0，为查询成功的个数,失败<0
	QueryPoints(context.Context, *QueryPointInfoReq) (*QueryPointInfoReply, error)
	// 功能：测点操作--使用测点名查询测点ID
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	QueryPointIdListByNameList(context.Context, *NameList) (*QueryPointIdListByNameListReply, error)
	// 功能：测点操作--使用测点名或ID批量查询测点信息
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	QueryPointInfoList(context.Context, *QueryRealtimeValueListReq) (*QueryPointInfoReply, error)
	// 实时值--写入实时值
	// 参数说明：入参、出差参考请求、响应体注释
	// error_code.app_code说明：成功>=0,写入成功的个数,失败<0
	// 备注：CommonRes中result_list为获取各个值的成功状态，失败<0为对应的错误码，全成功时为空
	UpdateRealtimeValueList(context.Context, *UpdateRealtimeValueListReq) (*CommonReply, error)
	// 实时值--通获取实时值
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	// 备注：DataListRes中resultList为获取各个值的成功状态，失败<0为对应的错误码，全成功时为空
	QueryRealtimeValueList(context.Context, *QueryRealtimeValueListReq) (*ValueListReply, error)
	// 历史值--通过测点ID查询时间段范围内的数据值
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	QueryHisRangeValueList(context.Context, *QueryHisRangeValueListReq) (*QueryHisValueListReply, error)
	// 历史值--通过测点ID查询时间段范围内重采样的数据值
	// 参数说明：入参、出差参考请求、响应体注释
	// errMsg.code：成功>0，为查询成功的个数,失败<=0
	QueryHisResampleValueList(context.Context, *QueryHisResamplesValueListReq) (*QueryHisValueListReply, error)
	// 历史值--通过测点ID查询时间点的数据值
	// 参数说明：入参、出差参考请求、响应体注释
	// error_code.app_code说明：为查询成功的数据个数成功>0,失败<=0
	// 备注：DataListRes中value_list与测点ID下标一一对应，result_list为每个测点查询的错误码，全成功时为空
	QueryHisTimePointValueList(context.Context, *QueryHisTimePointValueListReq) (*ValueListReply, error)
	// 历史值--通过测点ID写入测点历史值
	// 参数说明：入参、出差参考请求、响应体注释
	// error_code.app_code说明：成功>=0,查询成功的数据个数,失败<0
	InsertHisValueList(context.Context, *InsertHisValueListReq) (*CommonReply, error)
	mustEmbedUnimplementedRpcInterfaceServer()
}

// UnimplementedRpcInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedRpcInterfaceServer struct {
}

func (UnimplementedRpcInterfaceServer) Auth(context.Context, *AuthReq) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryErrMsg(context.Context, *QueryErrMsgReq) (*QueryErrMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryErrMsg not implemented")
}
func (UnimplementedRpcInterfaceServer) InsertTable(context.Context, *rpc.TableInfo) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertTable not implemented")
}
func (UnimplementedRpcInterfaceServer) DelTable(context.Context, *rpc.TableInfo) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelTable not implemented")
}
func (UnimplementedRpcInterfaceServer) ClearTable(context.Context, *rpc.TableInfo) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearTable not implemented")
}
func (UnimplementedRpcInterfaceServer) UpdateTable(context.Context, *rpc.TableInfo) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTable not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryTableList(context.Context, *QueryTableReq) (*QueryTableReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTableList not implemented")
}
func (UnimplementedRpcInterfaceServer) InsertPoints(context.Context, *PointInfoListReq) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertPoints not implemented")
}
func (UnimplementedRpcInterfaceServer) DelPoints(context.Context, *PointInfoListReq) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelPoints not implemented")
}
func (UnimplementedRpcInterfaceServer) UpdatePoints(context.Context, *PointInfoListReq) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePoints not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryPoints(context.Context, *QueryPointInfoReq) (*QueryPointInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPoints not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryPointIdListByNameList(context.Context, *NameList) (*QueryPointIdListByNameListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPointIdListByNameList not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryPointInfoList(context.Context, *QueryRealtimeValueListReq) (*QueryPointInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPointInfoList not implemented")
}
func (UnimplementedRpcInterfaceServer) UpdateRealtimeValueList(context.Context, *UpdateRealtimeValueListReq) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRealtimeValueList not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryRealtimeValueList(context.Context, *QueryRealtimeValueListReq) (*ValueListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRealtimeValueList not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryHisRangeValueList(context.Context, *QueryHisRangeValueListReq) (*QueryHisValueListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHisRangeValueList not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryHisResampleValueList(context.Context, *QueryHisResamplesValueListReq) (*QueryHisValueListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHisResampleValueList not implemented")
}
func (UnimplementedRpcInterfaceServer) QueryHisTimePointValueList(context.Context, *QueryHisTimePointValueListReq) (*ValueListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHisTimePointValueList not implemented")
}
func (UnimplementedRpcInterfaceServer) InsertHisValueList(context.Context, *InsertHisValueListReq) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertHisValueList not implemented")
}
func (UnimplementedRpcInterfaceServer) mustEmbedUnimplementedRpcInterfaceServer() {}

// UnsafeRpcInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcInterfaceServer will
// result in compilation errors.
type UnsafeRpcInterfaceServer interface {
	mustEmbedUnimplementedRpcInterfaceServer()
}

func RegisterRpcInterfaceServer(s grpc.ServiceRegistrar, srv RpcInterfaceServer) {
	s.RegisterService(&RpcInterface_ServiceDesc, srv)
}

func _RpcInterface_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryErrMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryErrMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryErrMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryErrMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryErrMsg(ctx, req.(*QueryErrMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_InsertTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.TableInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).InsertTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_InsertTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).InsertTable(ctx, req.(*rpc.TableInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_DelTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.TableInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).DelTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_DelTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).DelTable(ctx, req.(*rpc.TableInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_ClearTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.TableInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).ClearTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_ClearTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).ClearTable(ctx, req.(*rpc.TableInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_UpdateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.TableInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).UpdateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_UpdateTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).UpdateTable(ctx, req.(*rpc.TableInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryTableList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTableReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryTableList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryTableList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryTableList(ctx, req.(*QueryTableReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_InsertPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointInfoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).InsertPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_InsertPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).InsertPoints(ctx, req.(*PointInfoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_DelPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointInfoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).DelPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_DelPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).DelPoints(ctx, req.(*PointInfoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_UpdatePoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointInfoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).UpdatePoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_UpdatePoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).UpdatePoints(ctx, req.(*PointInfoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPointInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryPoints(ctx, req.(*QueryPointInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryPointIdListByNameList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryPointIdListByNameList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryPointIdListByNameList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryPointIdListByNameList(ctx, req.(*NameList))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryPointInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRealtimeValueListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryPointInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryPointInfoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryPointInfoList(ctx, req.(*QueryRealtimeValueListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_UpdateRealtimeValueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRealtimeValueListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).UpdateRealtimeValueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_UpdateRealtimeValueList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).UpdateRealtimeValueList(ctx, req.(*UpdateRealtimeValueListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryRealtimeValueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRealtimeValueListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryRealtimeValueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryRealtimeValueList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryRealtimeValueList(ctx, req.(*QueryRealtimeValueListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryHisRangeValueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHisRangeValueListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryHisRangeValueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryHisRangeValueList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryHisRangeValueList(ctx, req.(*QueryHisRangeValueListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryHisResampleValueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHisResamplesValueListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryHisResampleValueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryHisResampleValueList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryHisResampleValueList(ctx, req.(*QueryHisResamplesValueListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_QueryHisTimePointValueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHisTimePointValueListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).QueryHisTimePointValueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_QueryHisTimePointValueList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).QueryHisTimePointValueList(ctx, req.(*QueryHisTimePointValueListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcInterface_InsertHisValueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertHisValueListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcInterfaceServer).InsertHisValueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcInterface_InsertHisValueList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcInterfaceServer).InsertHisValueList(ctx, req.(*InsertHisValueListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcInterface_ServiceDesc is the grpc.ServiceDesc for RpcInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hhdb.rpc_interface.RpcInterface",
	HandlerType: (*RpcInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _RpcInterface_Auth_Handler,
		},
		{
			MethodName: "QueryErrMsg",
			Handler:    _RpcInterface_QueryErrMsg_Handler,
		},
		{
			MethodName: "InsertTable",
			Handler:    _RpcInterface_InsertTable_Handler,
		},
		{
			MethodName: "DelTable",
			Handler:    _RpcInterface_DelTable_Handler,
		},
		{
			MethodName: "ClearTable",
			Handler:    _RpcInterface_ClearTable_Handler,
		},
		{
			MethodName: "UpdateTable",
			Handler:    _RpcInterface_UpdateTable_Handler,
		},
		{
			MethodName: "QueryTableList",
			Handler:    _RpcInterface_QueryTableList_Handler,
		},
		{
			MethodName: "InsertPoints",
			Handler:    _RpcInterface_InsertPoints_Handler,
		},
		{
			MethodName: "DelPoints",
			Handler:    _RpcInterface_DelPoints_Handler,
		},
		{
			MethodName: "UpdatePoints",
			Handler:    _RpcInterface_UpdatePoints_Handler,
		},
		{
			MethodName: "QueryPoints",
			Handler:    _RpcInterface_QueryPoints_Handler,
		},
		{
			MethodName: "QueryPointIdListByNameList",
			Handler:    _RpcInterface_QueryPointIdListByNameList_Handler,
		},
		{
			MethodName: "QueryPointInfoList",
			Handler:    _RpcInterface_QueryPointInfoList_Handler,
		},
		{
			MethodName: "UpdateRealtimeValueList",
			Handler:    _RpcInterface_UpdateRealtimeValueList_Handler,
		},
		{
			MethodName: "QueryRealtimeValueList",
			Handler:    _RpcInterface_QueryRealtimeValueList_Handler,
		},
		{
			MethodName: "QueryHisRangeValueList",
			Handler:    _RpcInterface_QueryHisRangeValueList_Handler,
		},
		{
			MethodName: "QueryHisResampleValueList",
			Handler:    _RpcInterface_QueryHisResampleValueList_Handler,
		},
		{
			MethodName: "QueryHisTimePointValueList",
			Handler:    _RpcInterface_QueryHisTimePointValueList_Handler,
		},
		{
			MethodName: "InsertHisValueList",
			Handler:    _RpcInterface_InsertHisValueList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hhdb/rpc_interface/rpc_interface.proto",
}

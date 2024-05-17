package cj_hhdb_gosdk

import (
	"context"
	"github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc"
	hhdbRpc "github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc_interface"
	"time"
)

type OperatorInfo struct {
	createTime   uint64 `json:"createTime"`   //创建时间
	updateTime   uint64 `json:"updateTime"`   //更新时间
	createUserId int32  `json:"createUserId"` //创建用户ID
	updateUserId int32  `json:"updateUserId"` //修改用户ID
}

type TableInfo struct {
	TableId           int32             `json:"TableId"`           //表ID
	TableName         string            `json:"TableName"`         //表名
	ExtraFiledAndDesc map[string]string `json:"ExtraFiledAndDesc"` //额外的字段与字段名
	operatorInfo      OperatorInfo      `json:"operatorInfo"`      //用户信息
}

func (table *TableInfo) go2grpcTableInfo() *rpc.TableInfo {
	grpcTable := rpc.TableInfo{TableId: table.TableId, TableName: table.TableName, ExtraFiledAndDesc: table.ExtraFiledAndDesc,
		OperatorInfo: &rpc.OperatorInfo{CreateTime: table.operatorInfo.createTime, UpdateTime: table.operatorInfo.updateTime,
			CreateUserId: uint32(table.operatorInfo.createUserId), UpdateUserId: uint32(table.operatorInfo.updateUserId)}}
	return &grpcTable
}

func (table *TableInfo) grpc2goTableInfo(grpcTableInfo *rpc.TableInfo) {
	table.TableId = grpcTableInfo.TableId
	table.TableName = grpcTableInfo.TableName
	table.ExtraFiledAndDesc = grpcTableInfo.ExtraFiledAndDesc
	table.operatorInfo.createTime = grpcTableInfo.OperatorInfo.CreateTime
	table.operatorInfo.updateTime = grpcTableInfo.OperatorInfo.UpdateTime
	table.operatorInfo.createUserId = int32(grpcTableInfo.OperatorInfo.CreateUserId)
	table.operatorInfo.updateUserId = int32(grpcTableInfo.OperatorInfo.UpdateUserId)
}

func (hhdb *HhdbConPool) InsertTable(dbName string, tableInfo TableInfo) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}

	tableInfo.operatorInfo.createTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.operatorInfo.updateTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.operatorInfo.createUserId = dbConInfo.dbId
	tableInfo.operatorInfo.updateUserId = dbConInfo.dbId

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.InsertTable(ctx, tableInfo.go2grpcTableInfo())
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) DeleteTable(dbName string, TableId int32, TableName string) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.DelTable(ctx, &rpc.TableInfo{TableId: TableId, TableName: TableName})
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) ClearTable(dbName string, TableId int32, TableName string) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.ClearTable(ctx, &rpc.TableInfo{TableId: TableId, TableName: TableName})
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) UpdateTable(dbName string, tableInfo *TableInfo) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}

	tableInfo.operatorInfo.updateTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.operatorInfo.updateUserId = dbConInfo.dbId

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.UpdateTable(ctx, tableInfo.go2grpcTableInfo())
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) QueryTableList(dbName string, TableId int32, TableName string,
	enablePage bool, page uint32, limit uint32) (*[]TableInfo, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.QueryTableList(ctx, &hhdbRpc.QueryTableReq{TableId: TableId, TableName: TableName, EnablePage: enablePage, Page: page, Limit: limit})
	if err != nil {
		return nil, hhdb.handleGrpcError(&err)
	}
	tableInfoList := make([]TableInfo, len(res.TableInfoList))
	for i, v := range res.TableInfoList {
		tableInfoList[i].grpc2goTableInfo(v)
	}
	return &tableInfoList, nil
}

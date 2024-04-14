package cj_hhdb_gosdk

import (
	"cj_hhdb_gosdk/hhdb/rpc"
	hhdbRpc "cj_hhdb_gosdk/hhdb/rpc_interface"
	"context"
	"errors"
	"time"
)

type OperatorInfo struct {
	createTime   uint64 `json:"createTime"`   //创建时间
	updateTime   uint64 `json:"updateTime"`   //更新时间
	createUserId int32  `json:"createUserId"` //创建用户ID
	updateUserId int32  `json:"updateUserId"` //修改用户ID
}

type TableInfo struct {
	tableId           int32             `json:"tableId"`           //表ID
	tableName         string            `json:"tableName"`         //表名
	extraFiledAndDesc map[string]string `json:"extraFiledAndDesc"` //额外的字段与字段名
	operatorInfo      OperatorInfo      `json:"operatorInfo"`      //用户信息
}

func (table *TableInfo) go2grpcTableInfo() *rpc.TableInfo {
	grpcTable := rpc.TableInfo{TableId: table.tableId, TableName: table.tableName, ExtraFiledAndDesc: table.extraFiledAndDesc,
		OperatorInfo: &rpc.OperatorInfo{CreateTime: table.operatorInfo.createTime, UpdateTime: table.operatorInfo.updateTime,
			CreateUserId: uint32(table.operatorInfo.createUserId), UpdateUserId: uint32(table.operatorInfo.updateUserId)}}
	return &grpcTable
}

func (table *TableInfo) grpc2goTableInfo(grpcTableInfo *rpc.TableInfo) {
	table.tableId = grpcTableInfo.TableId
	table.tableName = grpcTableInfo.TableName
	table.extraFiledAndDesc = grpcTableInfo.ExtraFiledAndDesc
	table.operatorInfo.createTime = grpcTableInfo.OperatorInfo.CreateTime
	table.operatorInfo.updateTime = grpcTableInfo.OperatorInfo.UpdateTime
	table.operatorInfo.createUserId = int32(grpcTableInfo.OperatorInfo.CreateUserId)
	table.operatorInfo.updateUserId = int32(grpcTableInfo.OperatorInfo.UpdateUserId)
}

func (hhdb *HhdbConPool) InsertTable(dbName string, tableInfo TableInfo) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, err
	}

	tableInfo.operatorInfo.createTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.operatorInfo.updateTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.operatorInfo.createUserId = dbConInfo.dbId
	tableInfo.operatorInfo.updateUserId = dbConInfo.dbId

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClinet.InsertTable(ctx, tableInfo.go2grpcTableInfo())
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) DeleteTable(dbName string, tableId int32, tableName string) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClinet.DelTable(ctx, &rpc.TableInfo{TableId: tableId, TableName: tableName})

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) ClearTable(dbName string, tableId int32, tableName string) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClinet.ClearTable(ctx, &rpc.TableInfo{TableId: tableId, TableName: tableName})

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) UpdateTable(dbName string, tableInfo *TableInfo) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, err
	}

	tableInfo.operatorInfo.updateTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.operatorInfo.updateUserId = dbConInfo.dbId

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClinet.UpdateTable(ctx, tableInfo.go2grpcTableInfo())

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) QueryTableList(dbName string, tableId int32, tableName string,
	enablePage bool, page uint32, limit uint32) (*[]TableInfo, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClinet.QueryTableList(ctx, &hhdbRpc.QueryTableReq{TableId: tableId, TableName: tableName, EnablePage: enablePage, Page: page, Limit: limit})
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}
	tableInfoList := make([]TableInfo, len(res.TableInfoList))
	for i, v := range res.TableInfoList {
		tableInfoList[i].grpc2goTableInfo(v)
	}
	return &tableInfoList, nil
}

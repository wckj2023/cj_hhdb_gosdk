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
	TableId           int32             `json:"tableId"`           //表ID
	TableName         string            `json:"tableName"`         //表名
	TableShowName     string            `json:"tableShowName"`     //表展示名
	TableRemark       string            `json:"tableRemark"`       //表备注
	TableParentId     int32             `json:"tableParentId"`     //表父节点ID
	ExtraFieldAndDesc map[string]string `json:"extraFiledAndDesc"` //额外的字段与字段名
	operatorInfo      OperatorInfo      `json:"operatorInfo"`      //用户信息
}

func (table *TableInfo) go2grpcTableInfo() *rpc.TableInfo {
	extraFieldAndDesc := make(map[string][]byte)
	for k, v := range table.ExtraFieldAndDesc {
		extraFieldAndDesc[k] = []byte(v)
	}
	grpcTable := rpc.TableInfo{TableId: table.TableId, TableName: table.TableName, TableShowName: table.TableShowName,
		TableRemark: table.TableRemark, TableParentId: table.TableParentId, ExtraFieldAndDesc: extraFieldAndDesc,
		OperatorInfo: &rpc.OperatorInfo{CreateTime: table.operatorInfo.createTime, UpdateTime: table.operatorInfo.updateTime,
			CreateUserId: uint32(table.operatorInfo.createUserId), UpdateUserId: uint32(table.operatorInfo.updateUserId)}}
	return &grpcTable
}

func (table *TableInfo) grpc2goTableInfo(grpcTableInfo *rpc.TableInfo) {
	table.TableId = grpcTableInfo.TableId
	table.TableName = grpcTableInfo.TableName
	table.TableShowName = grpcTableInfo.TableShowName
	table.TableRemark = grpcTableInfo.TableRemark
	table.TableParentId = grpcTableInfo.TableParentId
	table.operatorInfo.createTime = grpcTableInfo.OperatorInfo.CreateTime
	table.operatorInfo.updateTime = grpcTableInfo.OperatorInfo.UpdateTime
	table.operatorInfo.createUserId = int32(grpcTableInfo.OperatorInfo.CreateUserId)
	table.operatorInfo.updateUserId = int32(grpcTableInfo.OperatorInfo.UpdateUserId)
	extraFieldAndDesc := make(map[string]string)
	for k, v := range grpcTableInfo.ExtraFieldAndDesc {
		extraFieldAndDesc[k] = string(v)
	}
	table.ExtraFieldAndDesc = extraFieldAndDesc
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

func (hhdb *HhdbConPool) QueryTableList(dbName string, tableInfo *TableInfo, queryChildren bool, queryAllChildren bool,
	enablePage bool, page uint32, limit uint32) (*[]TableInfo, int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.QueryTableList(ctx, &hhdbRpc.QueryTableReq{TableId: tableInfo.TableId, TableName: tableInfo.TableName, TableShowName: tableInfo.TableShowName, TableRemark: tableInfo.TableRemark,
		QueryChildren: queryChildren, QueryAllChildren: queryAllChildren, EnablePage: enablePage, Page: page, Limit: limit})
	if err != nil {
		return nil, 0, hhdb.handleGrpcError(&err)
	}
	tableInfoList := make([]TableInfo, len(res.TableInfoList))
	for i, v := range res.TableInfoList {
		tableInfoList[i].grpc2goTableInfo(v)
	}

	return &tableInfoList, res.GetTotal(), nil
}

func (hhdb *HhdbConPool) QueryTablePointCount(dbName string, tableInfo *TableInfo, queryAllChildren bool) (total int32, switchTotal int32, analogTotal int32, packageTotal int32, err error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.QueryTablePointCount(ctx, &hhdbRpc.QueryPointCountReq{TableId: tableInfo.TableId, TableName: tableInfo.TableName, QueryAllChildren: queryAllChildren})
	if err != nil {
		return 0, 0, 0, 0, hhdb.handleGrpcError(&err)
	}

	return res.GetTotal(), res.GetSwitchTotal(), res.GetAnalogTotal(), res.GetPackageTotal(), nil
}

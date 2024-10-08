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

type TablePointCount struct {
	Total        int32 `json:"total"`        //总点数
	SwitchTotal  int32 `json:"switchTotal"`  //开关量总点数
	AnalogTotal  int32 `json:"analogTotal"`  //模拟量总点数
	PackageTotal int32 `json:"packageTotal"` //打包点总点数
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

// 功能：表操作--添加
// 参数说明：dbName：数据库名，tableInfo：新增点表信息
// 返回值：成功int32>=0，为新增表的ID,失败<0
func (hhdb *HhdbConPool) InsertTable(dbName string, tableInfo *TableInfo) (int32, error) {
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

// 功能：表操作--删除
// 参数说明：dbName：数据库名，tableInfo：删除点表信息
// 返回值：成功int32>=0，为删除表的ID,失败<0
// 备注：通过TableInfo中的tableId删除表,tableId<0时,使用tableName进行匹配删除
func (hhdb *HhdbConPool) DeleteTable(dbName string, tableInfo *TableInfo) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.DelTable(ctx, tableInfo.go2grpcTableInfo())
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：表操作--清空
// 参数说明：dbName：数据库名，tableInfo：清空点表信息
// 返回值：成功int32>=0，为清空表的ID,失败<0
// 备注：通过TableInfo中的tableId清空表,tableId<0时,使用tableName进行匹配清空
func (hhdb *HhdbConPool) ClearTable(dbName string, tableInfo *TableInfo) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.ClearTable(ctx, tableInfo.go2grpcTableInfo())
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：表操作--更新
// 参数说明：dbName：数据库名，tableInfo：删除点表信息
// 返回值：成功int32>=0，为删除表的ID,失败<0
// 备注：通过TableInfo中的tableId更新表
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

// 功能：点组操作--查询点组列表
// 参数说明：dbName：数据库名，tableInfo：查询点表信息，queryChildren：是否查询下一级子节点，queryAllChildren：是否查询所有关联层级子节点
// 返回值：点组列表,查询总数,错误信息
// group_id>=0时,通过id获取点组信息,当group_id<0时,通过GroupInfo.group_name进行匹配获取数据,group_name为空且group_id<0时返回全部点组
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

// 功能：点组操作--查询指定点组或整库下的测点数量
// 参数说明：dbName：数据库名，tableInfo：查询点表信息，queryAllChildren：是否查询所有关联层级子节点
// 返回值：查询测点数量,错误信息
// group_id>=0时,通过点组id获取测点数量信息,当group_id<0时,通过GroupInfo.group_name进行匹配获取数据,group_name为空且group_id<0时返回整库测点数据
func (hhdb *HhdbConPool) QueryTablePointCount(dbName string, tableInfo *TableInfo, queryAllChildren bool) (pointCount TablePointCount, err error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return pointCount, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.QueryTablePointCount(ctx, &hhdbRpc.QueryPointCountReq{TableId: tableInfo.TableId, TableName: tableInfo.TableName, QueryAllChildren: queryAllChildren})
	if err != nil {
		return pointCount, hhdb.handleGrpcError(&err)
	}
	pointCount.Total = res.GetTotal()
	pointCount.SwitchTotal = res.GetSwitchTotal()
	pointCount.AnalogTotal = res.GetAnalogTotal()
	pointCount.PackageTotal = res.GetPackageTotal()
	return pointCount, nil
}

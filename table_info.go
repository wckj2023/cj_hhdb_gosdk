package cj_hhdb_gosdk

import (
	"context"
	"errors"
	"github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc"
	hhdbRpc "github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc_interface"
	"time"
)

type OperatorInfo struct {
	CreateTime   uint64 `json:"CreateTime" form:"CreateTime"`     //创建时间
	UpdateTime   uint64 `json:"UpdateTime" form:"UpdateTime"`     //更新时间
	CreateUserId int32  `json:"CreateUserId" form:"CreateUserId"` //创建用户ID
	UpdateUserId int32  `json:"UpdateUserId" form:"UpdateUserId"` //修改用户ID
}

type TableInfo struct {
	TableId           int32             `json:"tableId" form:"tableId"`                     //表ID
	TableName         string            `json:"tableName" form:"tableName"`                 //表名
	TableShowName     string            `json:"tableShowName" form:"tableShowName"`         //表展示名
	TableDesc         string            `json:"tableDesc" form:"tableDesc"`                 //表描述
	TableParentId     int32             `json:"tableParentId" form:"tableParentId"`         //表父节点ID
	ExtraFieldAndDesc map[string]string `json:"extraFiledAndDesc" form:"extraFiledAndDesc"` //额外的字段与字段名
	OperatorInfo      OperatorInfo      `json:"OperatorInfo" form:"OperatorInfo"`           //用户信息
	Children          *[]TableInfo      `json:"children" form:"children" `                  //子点表
	HasChildren       bool              `json:"hasChildren" form:"hasChildren" `            //是否有子点表
}

type TablePointCount struct {
	Total        int32 `json:"total" form:"total"`               //总点数
	SwitchTotal  int32 `json:"switchTotal" form:"switchTotal"`   //开关量总点数
	AnalogTotal  int32 `json:"analogTotal" form:"analogTotal"`   //模拟量总点数
	PackageTotal int32 `json:"packageTotal" form:"packageTotal"` //打包点总点数
}

func (table *TableInfo) go2grpcTableInfo() *rpc.TableInfo {
	extraFieldAndDesc := make(map[string][]byte)
	for k, v := range table.ExtraFieldAndDesc {
		extraFieldAndDesc[k] = []byte(v)
	}
	grpcTable := rpc.TableInfo{TableId: table.TableId, TableName: table.TableName, TableShowName: table.TableShowName,
		TableDesc: table.TableDesc, TableParentId: table.TableParentId, ExtraFieldAndDesc: extraFieldAndDesc,
		OperatorInfo: &rpc.OperatorInfo{CreateTime: table.OperatorInfo.CreateTime, UpdateTime: table.OperatorInfo.UpdateTime,
			CreateUserId: uint32(table.OperatorInfo.CreateUserId), UpdateUserId: uint32(table.OperatorInfo.UpdateUserId)}}
	return &grpcTable
}

func (table *TableInfo) grpc2goTableInfo(grpcTableInfo *rpc.TableInfo) {
	table.TableId = grpcTableInfo.TableId
	table.TableName = grpcTableInfo.TableName
	table.TableShowName = grpcTableInfo.TableShowName
	table.TableDesc = grpcTableInfo.TableDesc
	table.TableParentId = grpcTableInfo.TableParentId
	table.OperatorInfo.CreateTime = grpcTableInfo.OperatorInfo.CreateTime
	table.OperatorInfo.UpdateTime = grpcTableInfo.OperatorInfo.UpdateTime
	table.OperatorInfo.CreateUserId = int32(grpcTableInfo.OperatorInfo.CreateUserId)
	table.OperatorInfo.UpdateUserId = int32(grpcTableInfo.OperatorInfo.UpdateUserId)
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
	if tableInfo == nil {
		return -1, errors.New("table info is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}

	tableInfo.OperatorInfo.CreateTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.OperatorInfo.UpdateTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.OperatorInfo.CreateUserId = dbConInfo.dbId
	tableInfo.OperatorInfo.UpdateUserId = dbConInfo.dbId

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.InsertTable(ctx, tableInfo.go2grpcTableInfo())
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：表操作--删除
// 参数说明：dbName：数据库名，tableInfo：删除点表信息
// 返回值：成功int32>=0，为删除表的ID,失败<0
// 备注：通过TableInfo中的tableId删除表,tableId<0时,使用tableName进行匹配删除
func (hhdb *HhdbConPool) DelTable(dbName string, tableInfo *TableInfo) (int32, error) {
	if tableInfo == nil {
		return -1, errors.New("table info is empty")
	}
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

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：表操作--清空
// 参数说明：dbName：数据库名，tableInfo：清空点表信息
// 返回值：成功int32>=0，为清空表的ID,失败<0
// 备注：通过TableInfo中的tableId清空表,tableId<0时,使用tableName进行匹配清空
func (hhdb *HhdbConPool) ClearTable(dbName string, tableInfo *TableInfo) (int32, error) {
	if tableInfo == nil {
		return -1, errors.New("table info is empty")
	}
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

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：表操作--更新
// 参数说明：dbName：数据库名，tableInfo：删除点表信息
// 返回值：成功int32>=0，为删除表的ID,失败<0
// 备注：通过TableInfo中的tableId更新表
func (hhdb *HhdbConPool) UpdateTable(dbName string, tableInfo *TableInfo) (int32, error) {
	if tableInfo == nil {
		return -1, errors.New("table info is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}

	tableInfo.OperatorInfo.UpdateTime = uint64(time.Now().UTC().UnixMilli())
	tableInfo.OperatorInfo.UpdateUserId = dbConInfo.dbId

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.UpdateTable(ctx, tableInfo.go2grpcTableInfo())
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：点组操作--查询点组列表
// 参数说明：dbName：数据库名，tableInfo：查询点表信息，queryChildren：是否查询下一级子节点，queryAllChildren：是否查询所有关联层级子节点
// 返回值：点组列表,查询总数,错误信息
// group_id>=0时,通过id获取点组信息,当group_id<0时,通过GroupInfo.group_name进行匹配获取数据,group_name为空且group_id<0时返回全部点组
func (hhdb *HhdbConPool) QueryTableList(dbName string, tableInfo *TableInfo, queryChildren bool, queryAllChildren bool, treeEnable bool,
	enablePage bool, page uint32, limit uint32) (*[]TableInfo, int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryTableReq{QueryChildren: queryChildren, QueryAllChildren: queryAllChildren,
		EnablePage: enablePage, Page: page, Limit: limit}
	if tableInfo == nil {
		req.TableId = -1
	} else {
		req.TableId = tableInfo.TableId
		req.TableName = tableInfo.TableName
		req.TableShowName = tableInfo.TableShowName
		req.TableDesc = tableInfo.TableDesc

	}

	res, err := dbConInfo.dbClient.QueryTableList(ctx, &req)
	if err != nil {
		return nil, 0, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return nil, 0, errors.New(res.GetErrMsg().GetMsg())
	}
	tableInfoList := make([]TableInfo, len(res.TableInfoList))
	for i, v := range res.TableInfoList {
		tableInfoList[i].grpc2goTableInfo(v)
	}

	if treeEnable {
		dataMap := make(map[int32]*TableInfo)
		var allData []*TableInfo
		var rootData []TableInfo

		for _, data := range tableInfoList {
			// 创建菜单项的映射
			dataCopy := data
			dataCopy.Children = &[]TableInfo{}
			dataCopy.HasChildren = false
			allData = append(allData, &dataCopy)
			dataMap[dataCopy.TableId] = &dataCopy
		}

		for _, data := range allData {
			if data.TableParentId < 0 {
				rootData = append(rootData, *data)
				continue
			}

			// 有父菜单项，将其添加到父菜单的 Children 中
			dataParentId := data.TableParentId
			parentTable, exists := dataMap[dataParentId]
			if exists {
				*parentTable.Children = append(*parentTable.Children, *data)
				parentTable.HasChildren = true
			} else {
				rootData = append(rootData, *data)
			}
		}

		return &rootData, res.GetTotal(), nil
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
	req := hhdbRpc.QueryPointCountReq{QueryAllChildren: queryAllChildren}
	if tableInfo == nil {
		req.TableId = -1
	} else {
		req.TableId = tableInfo.TableId
		req.TableName = tableInfo.TableName
	}
	res, err := dbConInfo.dbClient.QueryTablePointCount(ctx, &req)
	if err != nil {
		return pointCount, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return pointCount, errors.New(res.GetErrMsg().GetMsg())
	}
	pointCount.Total = res.GetTotal()
	pointCount.SwitchTotal = res.GetSwitchTotal()
	pointCount.AnalogTotal = res.GetAnalogTotal()
	pointCount.PackageTotal = res.GetPackageTotal()
	return pointCount, nil
}

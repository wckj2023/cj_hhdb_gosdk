package cj_hhdb_gosdk

import (
	"context"
	"errors"
	"github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc"
	hhdbRpc "github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc_interface"
)

// 测点类型
type PointType int32

const (
	PointType_kPtNone    PointType = -1 //无效类型
	PointType_kPtSwitch  PointType = 0  //开关量
	PointType_kPtAnalog  PointType = 1  //模拟量
	PointType_kPtPackage PointType = 2  //打包点
)

// Enum value maps for PointType.
var (
	PointType_name = map[int32]string{
		0: "开关量",
		1: "模拟量",
		2: "打包点",
	}
	PointType_value = map[string]int32{
		"开关量": 0,
		"模拟量": 1,
		"打包点": 2,
	}
)

func (m PointType) String() string {
	if str, ok := PointType_name[int32(m)]; ok {
		return str
	}
	return "模拟量"
}

func (m PointType) ParseString(s string) PointType {
	if m, ok := PointType_value[s]; ok {
		return PointType(m)
	}
	return PointType_kPtAnalog
}

// 访问模式
type AccessMode int32

const (
	AccessMode_kAmRead  AccessMode = 0 //只读
	AccessMode_kAmWrite AccessMode = 1 //只写
	AccessMode_kAmRW    AccessMode = 2 //读写
	AccessMode_kAmUnUse AccessMode = 3 //停止使用
)

var (
	AccessMode_name = map[int32]string{
		0: "只读",
		1: "只写",
		2: "读写",
		3: "停用",
	}
	AccessMode_value = map[string]int32{
		"只读": 0,
		"只写": 1,
		"读写": 2,
		"停用": 3,
	}
)

func (m AccessMode) String() string {
	if str, ok := AccessMode_name[int32(m)]; ok {
		return str
	}
	return "只读"
}

func (m AccessMode) ParseString(s string) AccessMode {
	if m, ok := AccessMode_value[s]; ok {
		return AccessMode(m)
	}
	return AccessMode_kAmRead
}

// 字节序
type ByteOrder int32

const (
	ByteOrder_kBoDEFAULT ByteOrder = 0
	ByteOrder_kBoABCD    ByteOrder = 1
	ByteOrder_kBoDCBA    ByteOrder = 2
	ByteOrder_kBoBADC    ByteOrder = 3
	ByteOrder_kBoCDAB    ByteOrder = 4
)

var (
	ByteOrder_name = map[int32]string{
		0: "默认字节序",
		1: "大端序",
		2: "小端序",
		3: "大端交换序",
		4: "小端交换序",
	}
	ByteOrder_value = map[string]int32{
		"默认字节序": 0,
		"大端序":   1,
		"小端序":   2,
		"大端交换序": 3,
		"小端交换序": 4,
	}
)

func (m ByteOrder) String() string {
	if str, ok := ByteOrder_name[int32(m)]; ok {
		return str
	}
	return "默认字节序"
}

func (m ByteOrder) ParseString(s string) ByteOrder {
	if m, ok := ByteOrder_value[s]; ok {
		return ByteOrder(m)
	}
	return ByteOrder_kBoDEFAULT
}

// 压缩模式
type CompressMode int32

const (
	CompressMode_kCmThreshold CompressMode = 0 //阈值压缩
	CompressMode_kCmLeap      CompressMode = 1 //跳变压缩
	CompressMode_kCmTime      CompressMode = 2 //定时压缩
	CompressMode_kCmLoss      CompressMode = 3 //有损压缩
	CompressMode_kCmNone      CompressMode = 4 //无损压缩
)

// Enum value maps for CompressMode.
var (
	CompressMode_name = map[int32]string{
		0: "阈值压缩",
		1: "跳变压缩",
		2: "定时压缩",
		3: "有损压缩",
		4: "无损压缩",
	}
	CompressMode_value = map[string]int32{
		"阈值压缩": 0,
		"跳变压缩": 1,
		"定时压缩": 2,
		"有损压缩": 3,
		"无损压缩": 4,
	}
)

func (m CompressMode) String() string {
	if str, ok := CompressMode_name[int32(m)]; ok {
		return str
	}
	return "阈值压缩"
}

func (m CompressMode) ParseString(s string) CompressMode {
	if m, ok := CompressMode_value[s]; ok {
		return CompressMode(m)
	}
	return CompressMode_kCmThreshold
}

// 数据值类型
type ValueType int32

const (
	ValueType_kVtBool      ValueType = 0  // true 或 false 的二进制值
	ValueType_kVtFloat     ValueType = 1  //32 位实数值浮点型 IEEE-754 标准定义
	ValueType_kVtDouble    ValueType = 2  //64 位实数值双精度 IEEE-754 标准定义
	ValueType_kVtChar      ValueType = 3  // 有符号的 8 位整数数据
	ValueType_kVtByte      ValueType = 4  //无符号的 8 位整数数据
	ValueType_kVtShort     ValueType = 5  //有符号的 16 位整数数据
	ValueType_kVtWord      ValueType = 6  //无符号的 16 位整数数据
	ValueType_kVtInt       ValueType = 7  //有符号的 32 位整数数据
	ValueType_kVtDword     ValueType = 8  // 无符号的 32 位整数数据
	ValueType_kVtLong      ValueType = 9  //有符号的 64 位整数数据
	ValueType_kVtQword     ValueType = 10 //无符号的 64 位整数数据
	ValueType_kVtString    ValueType = 11 //字符串
	ValueType_kVtBlob      ValueType = 12 //字符串
	ValueType_kVtBoolArr   ValueType = 13 //bool数组
	ValueType_kVtFloatArr  ValueType = 14 //32 位实数值浮点型数组
	ValueType_kVtDoubleArr ValueType = 15 //64 位实数值浮点型数组
	ValueType_kVtCharArr   ValueType = 16 //char数组
	ValueType_kVtByteArr   ValueType = 17 //byte数组
	ValueType_kVtShortArr  ValueType = 18 //short数组
	ValueType_kVtWordArr   ValueType = 19 //word数组
	ValueType_kVtIntArr    ValueType = 20 //有符号的 32 位整数数据数组
	ValueType_kVtDwordArr  ValueType = 21 //无符号的 32 位整数数据数组
	ValueType_kVtLongArr   ValueType = 22 //有符号的 64 位整数数据数组
	ValueType_kVtQwordArr  ValueType = 23 //无符号的 64 位整数数据数组
	ValueType_kVtStringArr ValueType = 24 //字符串数组
)

// Enum value maps for ValueType.
var (
	ValueType_name = map[int32]string{
		0:  "bool",     // true 或 false 的二进制值
		1:  "float",    //32 位实数值浮点型 IEEE-754 标准定义
		2:  "double",   //64 位实数值双精度 IEEE-754 标准定义
		3:  "char",     // 有符号的 8 位整数数据
		4:  "byte",     //无符号的 8 位整数数据
		5:  "short",    //有符号的 16 位整数数据
		6:  "word",     //无符号的 16 位整数数据
		7:  "int",      //有符号的 32 位整数数据
		8:  "dword",    // 无符号的 32 位整数数据
		9:  "long",     //有符号的 64 位整数数据
		10: "qword",    //无符号的 64 位整数数据
		11: "geo",      //地理位置信息
		12: "string",   //字符串
		13: "blob",     //二进制块
		14: "bool[]",   //bool数组
		15: "float[]",  //32 位实数值浮点型数组
		16: "double[]", //64 位实数值浮点型数组
		17: "char[]",   //char数组
		18: "byte[]",   //byte数组
		19: "short[]",  //short数组
		20: "word[]",   //word数组
		21: "int[]",    //有符号的 32 位整数数据数组
		22: "dword[]",  //无符号的 32 位整数数据数组
		23: "long[]",   //有符号的 64 位整数数据数组
		24: "qword[]",  //无符号的 64 位整数数据数组
	}
	ValueType_value = map[string]int32{
		"bool":     0,
		"float":    1,
		"double":   2,
		"char":     3,
		"byte":     4,
		"short":    5,
		"word":     6,
		"int":      7,
		"dword":    8,
		"long":     9,
		"qword":    10,
		"geo":      11,
		"string":   12,
		"blob":     13,
		"bool[]":   14,
		"float[]":  15,
		"double[]": 16,
		"char[]":   17,
		"byte[]":   18,
		"short[]":  19,
		"word[]":   20,
		"int[]":    21,
		"dword[]":  22,
		"long[]":   23,
		"qword[]":  24,
	}
)

func (m ValueType) String() string {
	if str, ok := ValueType_name[int32(m)]; ok {
		return str
	}
	return "float"
}

func (m ValueType) ParseString(s string) ValueType {
	if m, ok := ValueType_value[s]; ok {
		return ValueType(m)
	}
	return ValueType_kVtFloat
}

// 测点全量信息
type PointInfo struct {
	PointId        int32             `json:"pointId" form:"pointId"`               //测点ID，为>=0的整数
	PointName      string            `json:"pointName" form:"pointName"`           //测点名
	PointShowName  string            `json:"pointShowName" form:"pointShowName"`   //测点展示名
	PointUnit      string            `json:"pointUnit" form:"pointUnit"`           //测点单位
	PointDesc      string            `json:"pointDesc" form:"pointDesc"`           //测点描述
	PointType      PointType         `json:"pointType" form:"pointType"`           //测点类型
	Access         AccessMode        `json:"access" form:"access"`                 //访问模式
	CheckEnable    bool              `json:"checkEnable" form:"checkEnable"`       //是否进行值校验
	LowerThreshold float64           `json:"lowerThreshold" form:"lowerThreshold"` //低限阈值
	UpperThreshold float64           `json:"upperThreshold" form:"upperThreshold"` //高限阈值
	ValueOffset    float64           `json:"valueOffset" form:"valueOffset"`       //数据偏移量
	ValueRate      float64           `json:"valueRate" form:"valueRate"`           //数据倍率
	CompressMode   CompressMode      `json:"compressMode" form:"compressMode"`     //压缩模式
	CompressParam1 float64           `json:"compressParam1" form:"compressParam1"` //压缩备用参数1
	CompressParam2 float64           `json:"compressParam2" form:"compressParam2"` //压缩备用参数2
	OuttimeDay     int32             `json:"outtimeDay" form:"outtimeDay"`         //超时时间（单位：天）=0则不启用，>0为对应的超时时间，<0代表仅缓存实时数据不存储历史数据
	ValueType      ValueType         `json:"valueType" form:"valueType"`           //测点值类型
	TableId        int32             `json:"tableId" form:"tableId"`               //点组ID
	CreateTime     uint64            `json:"createTime" form:"createTime"`         //测点创建时间
	ByteOrder      ByteOrder         `json:"byteOrder" form:"byteOrder"`           //访问模式
	SecurityNum    int32             `json:"securityNum" form:"securityNum"`       //点组ID
	ExtraField     map[string]string `json:"extraField" form:"extraField"`         //自定义的拓展字段
}

func (point *PointInfo) go2grpcPointInfo() (grpc *rpc.PointInfo) {
	var pointInfo rpc.PointInfo
	pointInfo.PointId = point.PointId
	pointInfo.PointName = point.PointName
	pointInfo.PointShowName = point.PointShowName
	pointInfo.PointUnit = point.PointUnit
	pointInfo.PointDesc = point.PointDesc
	pointInfo.PointType = PointType_value[point.PointType.String()]
	pointInfo.CompressMode = CompressMode_value[point.CompressMode.String()]
	pointInfo.CompressParam1 = point.CompressParam1
	pointInfo.CompressParam2 = point.CompressParam2
	pointInfo.Access = AccessMode_value[point.Access.String()]
	pointInfo.CheckEnable = point.CheckEnable
	pointInfo.LowerThreshold = point.LowerThreshold
	pointInfo.UpperThreshold = point.UpperThreshold
	pointInfo.ValueOffset = point.ValueOffset
	pointInfo.ValueRate = point.ValueRate
	pointInfo.OuttimeDay = point.OuttimeDay
	pointInfo.ValueType = ValueType_value[point.ValueType.String()]
	pointInfo.TableId = point.TableId
	pointInfo.CreateTime = point.CreateTime
	pointInfo.ByteOrder = ByteOrder_value[point.ByteOrder.String()]
	pointInfo.SecurityNum = point.SecurityNum

	extraField := make(map[string][]byte)
	for k, v := range point.ExtraField {
		extraField[k] = []byte(v)
	}
	pointInfo.ExtraField = extraField
	return &pointInfo
}

func (point *PointInfo) go2grpcPointInfoWithTableId(tableId int32) (grpc *rpc.PointInfo) {
	var pointInfo rpc.PointInfo
	pointInfo.PointId = point.PointId
	pointInfo.PointName = point.PointName
	pointInfo.PointShowName = point.PointShowName
	pointInfo.PointUnit = point.PointUnit
	pointInfo.PointDesc = point.PointDesc
	pointInfo.PointType = PointType_value[point.PointType.String()]
	pointInfo.CompressMode = CompressMode_value[point.CompressMode.String()]
	pointInfo.CompressParam1 = point.CompressParam1
	pointInfo.CompressParam2 = point.CompressParam2
	pointInfo.Access = AccessMode_value[point.Access.String()]
	pointInfo.CheckEnable = point.CheckEnable
	pointInfo.LowerThreshold = point.LowerThreshold
	pointInfo.UpperThreshold = point.UpperThreshold
	pointInfo.ValueOffset = point.ValueOffset
	pointInfo.ValueRate = point.ValueRate
	pointInfo.OuttimeDay = point.OuttimeDay
	pointInfo.ValueType = ValueType_value[point.ValueType.String()]
	pointInfo.TableId = tableId
	pointInfo.CreateTime = point.CreateTime
	pointInfo.ByteOrder = ByteOrder_value[point.ByteOrder.String()]
	pointInfo.SecurityNum = point.SecurityNum

	extraField := make(map[string][]byte)
	for k, v := range point.ExtraField {
		extraField[k] = []byte(v)
	}
	pointInfo.ExtraField = extraField
	return &pointInfo
}

func (point *PointInfo) grpc2goPointInfo(grpc *rpc.PointInfo) {
	point.PointId = grpc.PointId
	point.PointName = grpc.PointName
	point.PointShowName = grpc.PointShowName
	point.PointDesc = grpc.PointDesc
	point.PointUnit = grpc.PointUnit
	point.PointType = PointType(grpc.PointType)
	point.CompressMode = CompressMode(grpc.CompressMode)
	point.CompressParam1 = grpc.CompressParam1
	point.CompressParam2 = grpc.CompressParam2
	point.OuttimeDay = grpc.OuttimeDay
	point.Access = AccessMode(grpc.Access)
	point.CheckEnable = grpc.CheckEnable
	point.LowerThreshold = grpc.LowerThreshold
	point.UpperThreshold = grpc.UpperThreshold
	point.ValueOffset = grpc.ValueOffset
	point.ValueRate = grpc.ValueRate
	point.ValueType = ValueType(grpc.ValueType)
	point.TableId = grpc.TableId
	point.CreateTime = grpc.CreateTime
	point.ByteOrder = ByteOrder(grpc.ByteOrder)
	point.SecurityNum = grpc.SecurityNum

	extraField := make(map[string]string)
	for k, v := range grpc.ExtraField {
		extraField[k] = string(v)
	}
	point.ExtraField = extraField
}

// 功能：测点操作--添加测点
// 参数说明：dbName：数据库名，point：创建测点
// 返回值：int32：成功>0，为添加成功的点位ID,失败<=0
func (hhdb *HhdbConPool) InsertPoint(dbName string, point *PointInfo) (int32, error) {
	if point == nil {
		return -1, errors.New("point info is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return -1, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.InsertPoint(ctx, point.go2grpcPointInfo())
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：测点操作--添加测点
// 参数说明：dbName：数据库名，tableInfo：点表信息，tableId<0时通过tableName进行匹配，pointList:插入测点列表
// 返回值：int32：成功>0，为添加成功的个数,失败<=0， []int32：返回各个测点的ID，小于0时代表添加失败的错误码，全成功时为空
func (hhdb *HhdbConPool) InsertPoints(dbName string, pointList *[]PointInfo) (int32, *[]int32, error) {
	if pointList == nil {
		return -1, nil, errors.New("point info is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, nil, err
	}
	req := hhdbRpc.PointInfoListReq{}
	for i := 0; i < len(*pointList); i++ {
		req.PointInfoList = append(req.PointInfoList, (*pointList)[i].go2grpcPointInfo())
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.InsertPoints(ctx, &req)
	if err != nil {
		return 0, nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, nil
}

// 功能：测点操作--删除测点
// 参数说明：dbName：数据库名，id:删除测点ID
// 返回值：int32：成功>0，为删除成功的个数,失败<=0
func (hhdb *HhdbConPool) DelPoint(dbName string, id int32) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}
	req := hhdbRpc.IdOrNameListReq{}
	req.IdList = append(req.IdList, id)
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.DelPoints(ctx, &req)
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：测点操作--删除测点
// 参数说明：dbName：数据库名，idList:删除测点ID列表，nameList：删除测点点名列表，idList不为空时，忽略nameList
// 返回值：int32：成功>0，为删除成功的个数,失败<=0， []int32：返回各个测点的ID，小于0时代表失败的错误码，全成功时为空
func (hhdb *HhdbConPool) DelPoints(dbName string, idList *[]int32, nameList *[]string) (int32, *[]int32, error) {
	if idList == nil {
		return -1, nil, errors.New("id[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, nil, err
	}
	req := hhdbRpc.IdOrNameListReq{}
	if idList != nil && len(*idList) > 0 {
		req.IdList = *idList
	} else if nameList != nil && len(*nameList) > 0 {
		req.NameList = *nameList
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.DelPoints(ctx, &req)
	if err != nil {
		return 0, nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, nil
}

// 功能：测点操作--通过测点ID更新测点基础信息
// 参数说明：dbName：数据库名，point:测点信息
// 返回值：int32：成功>0，为删除成功的个数,失败<=0
func (hhdb *HhdbConPool) UpdatePoint(dbName string, point *PointInfo) (int32, error) {
	if point == nil {
		return -1, errors.New("point is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.UpdatePoint(ctx, point.go2grpcPointInfo())
	if err != nil {
		return 0, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

// 功能：测点操作--更新测点基础信息
// 参数说明：dbName：数据库名，pointList:更新测点列表，通过PointId进行关联更新，如果首个元素PointId为-1，则通过使用PointName进行匹配更新
// 返回值：int32：成功>0，为更新成功的个数,失败<=0，[]int32：返回更新成功的各个测点ID，返回小于0时代表添加失败的错误码，全成功时为空
func (hhdb *HhdbConPool) UpdatePoints(dbName string, pointList *[]PointInfo) (int32, *[]int32, error) {
	if pointList == nil || len(*pointList) == 0 {
		return -1, nil, errors.New("point[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, nil, err
	}
	req := hhdbRpc.PointInfoListReq{}
	for i := 0; i < len(*pointList); i++ {
		req.PointInfoList = append(req.PointInfoList, (*pointList)[i].go2grpcPointInfo())
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClient.UpdatePoints(ctx, &req)
	if err != nil {
		return 0, nil, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, nil
}

// 功能：测点操作--查询测点基础信息
// 参数说明：dbName：数据库名，tableName：点表名,
//
//		pointSearchInfo.pointId:>=0时，查询指定测点ID的信息
//		pointSearchInfo.tableId:<0时，整库查询，>=0时在指定表内检索
//		pointSearchInfo.nameRegex:不为空时，按正则匹配点名符合的测点
//		pointSearchInfo.showNameRegex:查询的测点名，为空时使用tableId为准，不为空时，以tableName进行查找
//		pointSearchInfo.descRegex:不为空时，按正则匹配描述符合的测点，两则都不为空时取交集
//		pointSearchInfo.unitRegex:不为空时，按正则匹配
//		pointSearchInfo.pointType:测点类型，>=0时，查询指定测点类型的信息
//		pointSearchInfo.extraFields:需要检索的字段，key为字段名，value为检索的字段值
//		enablePage:是否启用分页
//		page:页数,page从0开始计数
//		limit:每页的数量
//		queryChildren:只查询一个层级的子点组
//		queryAllChildren:查询所有子层级的点组
//
//	 返回值：list：查询结果，total：符合条件的总条数，err:错误信息
func (hhdb *HhdbConPool) QueryPoints(dbName string, tableName string, pointSearchInfo *PointInfo, enablePage bool,
	page uint32, limit uint32, queryChildren bool, queryAllChildren bool) (list *[]PointInfo, total int32, err error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	var req *hhdbRpc.QueryPointInfoReq
	searchMap := make(map[string][]byte)
	if pointSearchInfo == nil {
		req = &hhdbRpc.QueryPointInfoReq{TableName: tableName, TableId: -1, PointId: -1, PointType: -1, EnablePage: enablePage,
			Page: page, Limit: limit, QueryChildren: queryChildren, QueryAllChildren: queryAllChildren}
	} else {
		for k, v := range pointSearchInfo.ExtraField {
			searchMap[k] = []byte(v)
		}
		req = &hhdbRpc.QueryPointInfoReq{TableName: tableName, TableId: pointSearchInfo.TableId, PointId: pointSearchInfo.PointId, NameRegex: pointSearchInfo.PointName,
			ShowNameRegex: pointSearchInfo.PointShowName, DescRegex: pointSearchInfo.PointDesc, UnitRegex: pointSearchInfo.PointUnit, PointType: int32(pointSearchInfo.PointType), ExtraFields: searchMap, EnablePage: enablePage,
			Page: page, Limit: limit, QueryChildren: queryChildren, QueryAllChildren: queryAllChildren}
	}
	res, err := dbConInfo.dbClient.QueryPoints(ctx, req)
	if err != nil {
		return nil, 0, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, 0, errors.New(res.GetErrMsg().GetMsg())
	}

	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	total = res.GetTotal()
	//测点太多遍历获取完
	if !enablePage && int32(len(pointList)) < total {
		pageAdd := 1
		newLimit := len(pointList)
		tempInfo := PointInfo{}
		for int32(len(pointList)) < total {
			pageAdd++
			req.Page = uint32(pageAdd)
			req.Limit = uint32(newLimit)
			req.EnablePage = true
			ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
			defer cancel()
			res, err = dbConInfo.dbClient.QueryPoints(ctx, req)
			if res.GetErrMsg().GetCode() < 0 {
				break
			}
			for i := 0; i < len(res.PointInfoList); i++ {
				tempInfo.grpc2goPointInfo(res.PointInfoList[i])
				pointList = append(pointList, tempInfo)
			}
		}
	}

	return &pointList, total, nil
}

// 功能：测点操作--查询指定ID范围或非id范围内的测点基础信息
// 参数说明：dbName：数据库名，tableName：点表名,
//
//	pointIdList:测点的id范围
//	inOrNotInFlag:是否在测点范围内，或不在测点范围内
//	pointSearchInfo.pointId:>=0时，查询指定测点ID的信息
//	pointSearchInfo.tableId:<0时，整库查询，>=0时在指定表内检索
//	pointSearchInfo.nameRegex:不为空时，按正则匹配点名符合的测点
//	pointSearchInfo.showNameRegex:查询的测点名，为空时使用tableId为准，不为空时，以tableName进行查找
//	pointSearchInfo.descRegex:不为空时，按正则匹配描述符合的测点，两则都不为空时取交集
//	pointSearchInfo.unitRegex:不为空时，按正则匹配
//	pointSearchInfo.pointType:测点类型，>=0时，查询指定测点类型的信息
//	pointSearchInfo.extraFields:需要检索的字段，key为字段名，value为检索的字段值
//	enablePage:是否启用分页
//	page:页数,page从0开始计数
//	limit:每页的数量
//
// 返回值：list：查询结果，total：符合条件的总条数，err:错误信息
func (hhdb *HhdbConPool) QueryPointsIsInList(dbName string, pointIdList *[]int32, inOrNotInFlag bool, tableName string, pointSearchInfo *PointInfo, enablePage bool,
	page uint32, limit uint32) (list *[]PointInfo, total int32, err error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	var req hhdbRpc.QueryPointsIsInIdsReq
	searchMap := make(map[string][]byte)
	if pointIdList != nil {
		req.IdList = *pointIdList
	}
	req.InOrNotInFlag = inOrNotInFlag
	if pointSearchInfo == nil {
		req.QueryInfo = &hhdbRpc.QueryPointInfoReq{TableName: tableName, TableId: -1, PointId: -1, PointType: -1, EnablePage: enablePage,
			Page: page, Limit: limit}
	} else {
		for k, v := range pointSearchInfo.ExtraField {
			searchMap[k] = []byte(v)
		}
		req.QueryInfo = &hhdbRpc.QueryPointInfoReq{TableName: tableName, TableId: pointSearchInfo.TableId, PointId: pointSearchInfo.PointId, NameRegex: pointSearchInfo.PointName,
			ShowNameRegex: pointSearchInfo.PointShowName, DescRegex: pointSearchInfo.PointDesc, UnitRegex: pointSearchInfo.PointUnit, PointType: int32(pointSearchInfo.PointType), ExtraFields: searchMap, EnablePage: enablePage,
			Page: page, Limit: limit}
	}
	res, err := dbConInfo.dbClient.QueryPointsIsInListReq(ctx, &req)
	if err != nil {
		return nil, 0, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, 0, errors.New(res.GetErrMsg().GetMsg())
	}

	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	total = res.GetTotal()
	//测点太多遍历获取完
	if !enablePage && int32(len(pointList)) < total {
		pageAdd := 1
		newLimit := len(pointList)
		tempInfo := PointInfo{}
		for int32(len(pointList)) < total {
			req.QueryInfo.Page = uint32(pageAdd)
			req.QueryInfo.Limit = uint32(newLimit)

			res, err = dbConInfo.dbClient.QueryPointsIsInListReq(ctx, &req)
			if res.GetErrMsg().GetCode() < 0 {
				break
			}
			for i := 0; i < len(res.PointInfoList); i++ {
				tempInfo.grpc2goPointInfo(res.PointInfoList[i])
				pointList = append(pointList, tempInfo)
			}
		}
	}

	return &pointList, total, nil
}

// 功能：测点操作--使用测点ID批量查询测点信息
// 参数说明：dbName：数据库名，pointIdList:测点ID列表
// 返回值：测点信息列表
func (hhdb *HhdbConPool) QueryPointInfoListByID(dbName string, pointIdList *[]int32) (*[]PointInfo, error) {
	if pointIdList == nil || len(*pointIdList) == 0 {
		return nil, errors.New("id[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.IdOrNameListReq{}
	req.IdList = *pointIdList
	res, err := dbConInfo.dbClient.QueryPointInfoList(ctx, &req)
	if err != nil {
		return nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}

	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	return &pointList, nil
}

// 功能：测点操作--使用测点名批量查询测点信息
// 参数说明：dbName：数据库名，pointNameList:测点名列表
// 返回值：测点信息列表
func (hhdb *HhdbConPool) QueryPointInfoListByName(dbName string, pointNameList *[]string) (*[]PointInfo, error) {
	if pointNameList == nil || len(*pointNameList) == 0 {
		return nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.IdOrNameListReq{}
	req.NameList = *pointNameList
	res, err := dbConInfo.dbClient.QueryPointInfoList(ctx, &req)
	if err != nil {
		return nil, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}
	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	return &pointList, nil
}

// 功能：测点操作--使用测点ID批量查询测点信息
// 参数说明：dbName：数据库名，pointIdList:测点ID列表
// 返回值：测点信息列表
func (hhdb *HhdbConPool) QueryPointTypeCountByID(dbName string, pointIdList *[]int32) (pointCount TablePointCount, err error) {
	if pointIdList == nil || len(*pointIdList) == 0 {
		return pointCount, errors.New("id[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return pointCount, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.IdOrNameListReq{}
	req.IdList = *pointIdList
	res, err := dbConInfo.dbClient.QueryPointTypeCount(ctx, &req)
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

func (hhdb *HhdbConPool) QueryPointIdListByName(dbName string, pointNameList *[]string) (*[]int32, error) {
	if pointNameList == nil || len(*pointNameList) == 0 {
		return nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.NameListReq{}
	req.NameList = *pointNameList
	res, err := dbConInfo.dbClient.QueryPointIdListByNameList(ctx, &req)
	if err != nil {
		return nil, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}

	return &res.IdList, nil
}

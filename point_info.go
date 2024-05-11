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
		0: "kPtSwitch",
		1: "kPtAnalog",
		2: "kPtPackage",
	}
	PointType_value = map[string]int32{
		"kPtSwitch":  0,
		"kPtAnalog":  1,
		"kPtPackage": 2,
	}
)

func (x PointType) Enum() *PointType {
	p := new(PointType)
	*p = x
	return p
}

func (x PointType) String() string {
	return PointType_name[int32(x)]
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
		0: "kCmThreshold",
		1: "kCmLeap",
		2: "kCmTime",
		3: "kCmLoss",
		4: "kCmNone",
	}
	CompressMode_value = map[string]int32{
		"kCmThreshold": 0,
		"kCmLeap":      1,
		"kCmTime":      2,
		"kCmLoss":      3,
		"kCmNone":      4,
	}
)

func (x CompressMode) Enum() *CompressMode {
	p := new(CompressMode)
	*p = x
	return p
}

func (x CompressMode) String() string {
	return CompressMode_name[int32(x)]
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
	ValueType_kVtBoolArr   ValueType = 12 //bool数组
	ValueType_kVtFloatArr  ValueType = 13 //32 位实数值浮点型数组
	ValueType_kVtDoubleArr ValueType = 14 //64 位实数值浮点型数组
	ValueType_kVtCharArr   ValueType = 15 //char数组
	ValueType_kVtByteArr   ValueType = 16 //byte数组
	ValueType_kVtShortArr  ValueType = 17 //short数组
	ValueType_kVtWordArr   ValueType = 18 //word数组
	ValueType_kVtIntArr    ValueType = 19 //有符号的 32 位整数数据数组
	ValueType_kVtDwordArr  ValueType = 20 //无符号的 32 位整数数据数组
	ValueType_kVtLongArr   ValueType = 21 //有符号的 64 位整数数据数组
	ValueType_kVtQwordArr  ValueType = 22 //无符号的 64 位整数数据数组
	ValueType_kVtStringArr ValueType = 23 //字符串数组
)

// Enum value maps for ValueType.
var (
	ValueType_name = map[int32]string{
		0:  "kVtBool",      // true 或 false 的二进制值
		1:  "kVtFloat",     //32 位实数值浮点型 IEEE-754 标准定义
		2:  "kVtDouble",    //64 位实数值双精度 IEEE-754 标准定义
		3:  "kVtChar",      // 有符号的 8 位整数数据
		4:  "kVtByte",      //无符号的 8 位整数数据
		5:  "kVtShort",     //有符号的 16 位整数数据
		6:  "kVtWord",      //无符号的 16 位整数数据
		7:  "kVtInt",       //有符号的 32 位整数数据
		8:  "kVtDword",     // 无符号的 32 位整数数据
		9:  "kVtLong",      //有符号的 64 位整数数据
		10: "kVtQword",     //无符号的 64 位整数数据
		11: "kVtString",    //字符串
		12: "kVtBoolArr",   //bool数组
		13: "kVtFloatArr",  //32 位实数值浮点型数组
		14: "kVtDoubleArr", //64 位实数值浮点型数组
		15: "kVtCharArr",   //char数组
		16: "kVtByteArr",   //byte数组
		17: "kVtShortArr",  //short数组
		18: "kVtWordArr",   //word数组
		19: "kVtIntArr",    //有符号的 32 位整数数据数组
		20: "kVtDwordArr",  //无符号的 32 位整数数据数组
		21: "kVtLongArr",   //有符号的 64 位整数数据数组
		22: "kVtQwordArr",  //无符号的 64 位整数数据数组
		23: "kVtStringArr", //字符串数组
	}
	ValueType_value = map[string]int32{
		"kVtBool":      0,
		"kVtFloat":     1,
		"kVtDouble":    2,
		"kVtChar":      3,
		"kVtByte":      4,
		"kVtShort":     5,
		"kVtWord":      6,
		"kVtInt":       7,
		"kVtDword":     8,
		"kVtLong":      9,
		"kVtQword":     10,
		"kVtString":    11,
		"kVtBoolArr":   12,
		"kVtFloatArr":  13,
		"kVtDoubleArr": 14,
		"kVtCharArr":   15,
		"kVtByteArr":   16,
		"kVtShortArr":  17,
		"kVtWordArr":   18,
		"kVtIntArr":    19,
		"kVtDwordArr":  20,
		"kVtLongArr":   21,
		"kVtQwordArr":  22,
		"kVtStringArr": 23,
	}
)

func (x ValueType) Enum() *ValueType {
	p := new(ValueType)
	*p = x
	return p
}

func (x ValueType) String() string {
	return ValueType_name[int32(x)]
}

// 测点全量信息
type PointInfo struct {
	PointId        int32             `json:"pointId"`        //测点ID，为>=0的整数
	PointName      string            `json:"pointName"`      //测点名
	PointUnit      string            `json:"pointUnit"`      //测点单位
	PointDesc      string            `json:"pointDesc"`      //测点描述
	PointType      PointType         `json:"pointType"`      //测点类型
	WriteEnable    bool              `json:"writeEnable"`    //是否可写
	CheckEnable    bool              `json:"checkEnable"`    //是否进行值校验
	LowerThreshold float64           `json:"lowerThreshold"` //低限阈值
	UpperThreshold float64           `json:"upperThreshold"` //高限阈值
	ValueOffset    float64           `json:"valueOffset"`    //数据偏移量
	ValueRate      float64           `json:"valueRate"`      //数据倍率
	CompressMode   CompressMode      `json:"compressMode"`   //压缩模式
	CompressParam1 float64           `json:"compressParam1"` //压缩备用参数1
	CompressParam2 float64           `json:"compressParam2"` //压缩备用参数2
	OuttimeDay     int32             `json:"outtimeDay"`     //超时时间（单位：天）=0则不启用，>0为对应的超时时间，<0代表仅缓存实时数据不存储历史数据
	ValueType      ValueType         `json:"valueType"`      //测点值类型
	TableId        int32             `json:"tableId"`        //点组ID
	CreateTime     uint64            `json:"createTime"`     //测点创建时间
	ExtraField     map[string]string `json:"extraField"`     //自定义的拓展字段
}

func (point *PointInfo) go2grpcPointInfo() (grpc *rpc.PointInfo) {
	var pointInfo rpc.PointInfo
	pointInfo.PointId = point.PointId
	pointInfo.PointName = point.PointName
	pointInfo.PointUnit = point.PointUnit
	pointInfo.PointDesc = point.PointDesc
	pointInfo.PointType = rpc.PointType(point.PointType)
	pointInfo.CompressMode = rpc.CompressMode(point.CompressMode)
	pointInfo.CompressParam1 = point.CompressParam1
	pointInfo.CompressParam2 = point.CompressParam2
	pointInfo.WriteEnable = point.WriteEnable
	pointInfo.CheckEnable = point.CheckEnable
	pointInfo.LowerThreshold = point.LowerThreshold
	pointInfo.UpperThreshold = point.UpperThreshold
	pointInfo.ValueOffset = point.ValueOffset
	pointInfo.ValueRate = point.ValueRate
	pointInfo.OuttimeDay = point.OuttimeDay
	pointInfo.ValueType = rpc.ValueType(point.ValueType)
	pointInfo.TableId = point.TableId
	pointInfo.CreateTime = point.CreateTime
	pointInfo.ExtraField = point.ExtraField
	return &pointInfo
}

func (point *PointInfo) grpc2goPointInfo(grpc *rpc.PointInfo) {
	point.PointId = grpc.PointId
	point.PointName = grpc.PointName
	point.PointDesc = grpc.PointDesc
	point.PointUnit = grpc.PointUnit
	point.PointType = PointType(grpc.PointType)
	point.CompressMode = CompressMode(grpc.CompressMode)
	point.CompressParam1 = grpc.CompressParam1
	point.CompressParam2 = grpc.CompressParam2
	point.OuttimeDay = grpc.OuttimeDay
	point.WriteEnable = grpc.WriteEnable
	point.CheckEnable = grpc.CheckEnable
	point.LowerThreshold = grpc.LowerThreshold
	point.UpperThreshold = grpc.UpperThreshold
	point.ValueOffset = grpc.ValueOffset
	point.ValueRate = grpc.ValueRate
	point.ValueType = ValueType(grpc.ValueType)
	point.TableId = grpc.TableId
	point.CreateTime = grpc.CreateTime
	point.ExtraField = grpc.ExtraField
}

func (hhdb *HhdbConPool) InsertPoints(dbName string, pointList *[]PointInfo) (int32, []int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, []int32{}, err
	}
	req := hhdbRpc.PointInfoListReq{}
	for i := 0; i < len(*pointList); i++ {
		req.PointInfoList = append(req.PointInfoList, (*pointList)[i].go2grpcPointInfo())
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.DbClinet.InsertPoints(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), res.IdOrErrCodeList, errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), res.IdOrErrCodeList, nil
}

func (hhdb *HhdbConPool) DelPoints(dbName string, pointList *[]PointInfo) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, err
	}
	req := hhdbRpc.PointInfoListReq{}
	for i := 0; i < len(*pointList); i++ {
		req.PointInfoList = append(req.PointInfoList, (*pointList)[i].go2grpcPointInfo())
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.DbClinet.DelPoints(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) UpdatePoints(dbName string, pointList *[]PointInfo) (int32, []int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, []int32{}, err
	}
	req := hhdbRpc.PointInfoListReq{}
	for i := 0; i < len(*pointList); i++ {
		req.PointInfoList = append(req.PointInfoList, (*pointList)[i].go2grpcPointInfo())
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.DbClinet.UpdatePoints(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), res.GetIdOrErrCodeList(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), res.GetIdOrErrCodeList(), nil
}

func (hhdb *HhdbConPool) QueryPoints(dbName string, tableName string, pointSearchInfo *PointInfo, enablePage bool,
	page uint32, limit uint32) (list *[]PointInfo, total int32, err error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.DbClinet.QueryPoints(ctx, &hhdbRpc.QueryPointInfoReq{TableName: tableName, TableId: pointSearchInfo.TableId, PointId: pointSearchInfo.PointId, NameRegex: pointSearchInfo.PointName,
		DescRegex: pointSearchInfo.PointDesc, UnitRegex: pointSearchInfo.PointUnit, PointType: int32(pointSearchInfo.PointType), ExtraFields: pointSearchInfo.ExtraField, EnablePage: enablePage,
		Page: page, Limit: limit})
	if res.GetErrMsg().GetCode() < 0 {
		return nil, 0, errors.New(res.GetErrMsg().GetMsg())
	}
	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	total = res.GetErrMsg().GetCode()
	if !enablePage && int32(len(pointList)) < total {
		pageAdd := 1
		newLimit := len(pointList)
		tempInfo := PointInfo{}
		for int32(len(pointList)) < total {
			res, err = dbConInfo.DbClinet.QueryPoints(ctx, &hhdbRpc.QueryPointInfoReq{TableName: tableName, TableId: pointSearchInfo.TableId, PointId: pointSearchInfo.PointId, NameRegex: pointSearchInfo.PointName,
				DescRegex: pointSearchInfo.PointDesc, UnitRegex: pointSearchInfo.PointUnit, PointType: int32(pointSearchInfo.PointType), ExtraFields: pointSearchInfo.ExtraField, EnablePage: true,
				Page: uint32(pageAdd), Limit: uint32(newLimit)})
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

func (hhdb *HhdbConPool) QueryPointInfoListByID(dbName string, pointIdList *[]int32) (*[]PointInfo, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryRealtimeValueListReq{}
	req.IdList = *pointIdList
	res, err := dbConInfo.DbClinet.QueryPointInfoList(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}
	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	return &pointList, nil
}

func (hhdb *HhdbConPool) QueryPointInfoListByName(dbName string, pointNameList *[]string) (*[]PointInfo, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryRealtimeValueListReq{}
	req.NameList = *pointNameList
	res, err := dbConInfo.DbClinet.QueryPointInfoList(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}
	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	return &pointList, nil
}

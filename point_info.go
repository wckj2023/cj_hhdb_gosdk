package hhdb_sdk

import (
	"context"
	"errors"
	"hhdb_sdk/hhdb/rpc"
	hhdbRpc "hhdb_sdk/hhdb/rpc_interface"
)

// 测点类型
type PointType int32

const (
	PointType_kPtSwitch  PointType = 0 //开关量
	PointType_kPtAnalog  PointType = 1 //模拟量
	PointType_kPtPackage PointType = 2 //打包点
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
	ValueType_kVtErr    ValueType = 0 //错误类型
	ValueType_kVtBool   ValueType = 1 //bool
	ValueType_kVtFloat  ValueType = 2 //float
	ValueType_kVtDouble ValueType = 3 //double
	ValueType_kVtInt32  ValueType = 4 //int
	ValueType_kVtUint32 ValueType = 5 //unsigned int
	ValueType_kVtInt64  ValueType = 6 //long long
	ValueType_kVtUint64 ValueType = 7 //unsigned long long
	ValueType_kVtString ValueType = 8 //string
	ValueType_kVtBinary ValueType = 9 //binary
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
	pointId        int32             `json:"pointId"`        //测点ID，为>=0的整数
	pointName      string            `json:"pointName"`      //测点名
	pointUnit      string            `json:"pointUnit"`      //测点单位
	pointDesc      string            `json:"pointDesc"`      //测点描述
	pointType      PointType         `json:"pointType"`      //测点类型
	writeEnable    bool              `json:"writeEnable"`    //是否可写
	checkEnable    bool              `json:"checkEnable"`    //是否进行值校验
	lowerThreshold float64           `json:"lowerThreshold"` //低限阈值
	upperThreshold float64           `json:"upperThreshold"` //高限阈值
	valueOffset    float64           `json:"valueOffset"`    //数据偏移量
	valueRate      float64           `json:"valueRate"`      //数据倍率
	compressMode   CompressMode      `json:"compressMode"`   //压缩模式
	compressParam1 float64           `json:"compressParam1"` //压缩备用参数1
	compressParam2 float64           `json:"compressParam2"` //压缩备用参数2
	outtimeDay     int32             `json:"outtimeDay"`     //超时时间（单位：天）=0则不启用，>0为对应的超时时间，<0代表仅缓存实时数据不存储历史数据
	valueType      ValueType         `json:"valueType"`      //测点值类型
	tableId        int32             `json:"tableId"`        //点组ID
	createTime     uint64            `json:"createTime"`     //测点创建时间
	extraField     map[string]string `json:"extraField"`     //自定义的拓展字段
}

func (point *PointInfo) go2grpcPointInfo() (grpc *rpc.PointInfo) {
	grpc.PointId = point.pointId
	grpc.ShowInfo.PointName = point.pointName
	grpc.ShowInfo.PointUnit = point.pointUnit
	grpc.ShowInfo.PointDesc = point.pointDesc
	grpc.ShowInfo.PointType = rpc.PointType(point.pointType)
	grpc.CompressMode = rpc.CompressMode(point.compressMode)
	grpc.CompressParam1 = point.compressParam1
	grpc.CompressParam2 = point.compressParam2
	grpc.OuttimeDay = point.outtimeDay
	grpc.ValueType = rpc.ValueType(point.valueType)
	grpc.TableId = point.tableId
	grpc.CreateTime = point.createTime
	grpc.ExtraField = point.extraField
	return grpc
}

func (point *PointInfo) grpc2goPointInfo(grpc *rpc.PointInfo) {
	point.pointId = grpc.PointId
	point.pointName = grpc.ShowInfo.GetPointName()
	point.pointDesc = grpc.ShowInfo.GetPointDesc()
	point.pointUnit = grpc.ShowInfo.GetPointUnit()
	point.pointType = PointType(grpc.ShowInfo.PointType)
	point.compressMode = CompressMode(grpc.CompressMode)
	point.compressParam1 = grpc.CompressParam1
	point.compressParam2 = grpc.CompressParam2
	point.outtimeDay = grpc.OuttimeDay
	point.valueType = ValueType(grpc.ValueType)
	point.tableId = grpc.TableId
	point.createTime = grpc.CreateTime
	point.extraField = grpc.ExtraField
}

func (hhdb *HhdbConPool) InsertPoints(dbName string, pointList *[]PointInfo) (int32, error) {
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
	res, err := dbConInfo.dbClinet.InsertPoints(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
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
	res, err := dbConInfo.dbClinet.DelPoints(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) UpdatePoints(dbName string, pointList *[]PointInfo) (int32, error) {
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
	res, err := dbConInfo.dbClinet.UpdatePoints(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) QueryPoints(dbName string, tableId int32, pointId int32, nameRegex string, descRegex string,
	unitRegex string, pointType int32, extraFields *map[string]string, enablePage bool,
	page uint32, limit uint32) (*[]PointInfo, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	res, err := dbConInfo.dbClinet.QueryPoints(ctx, &hhdbRpc.QueryPointInfoReq{TableId: tableId, PointId: pointId, NameRegex: nameRegex,
		DescRegex: descRegex, UnitRegex: unitRegex, PointType: pointType, ExtraFields: *extraFields, EnablePage: enablePage,
		Page: page, Limit: limit})
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}
	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	return &pointList, nil
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
	res, err := dbConInfo.dbClinet.QueryPointInfoList(ctx, &req)
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
	res, err := dbConInfo.dbClinet.QueryPointInfoList(ctx, &req)
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}
	pointList := make([]PointInfo, len(res.PointInfoList))
	for i := 0; i < len(res.PointInfoList); i++ {
		pointList[i].grpc2goPointInfo(res.PointInfoList[i])
	}
	return &pointList, nil
}

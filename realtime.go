package cj_hhdb_gosdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc"
	hhdbRpc "github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc_interface"
)

type PointState int32

const (
	PointState_kPsInit             PointState = 0   //初始化0-9为好的状态
	PointState_kPsGood             PointState = 1   //正常,表示测点的值是有效的且在正常范围内。这是最常见的状态
	PointState_kPsShowWarnning     PointState = 10  //>=10为警告状态
	PointState_kPsUncertain        PointState = 11  //不确定,表示测点的值可能有效，但存在某种不确定性或不可靠性。这通常用于表示数据质量差、未知状态或临时问题。
	PointState_kPsInavtive         PointState = 12  //不活动,表示测点当前处于非活动状态，可能是因为设备已停用或者该测点当前不可用。
	PointState_kPsQualityIssue     PointState = 13  //数据质量问题,表示测点的数据存在质量问题，可能是由于噪声、干扰或其他原因导致的不准确性。
	PointState_kPsDataTypeMisMatch PointState = 14  //数据类型不匹配
	PointState_kPsOutService       PointState = 15  //停用,表示测点当前处于停用状态，通常是因为设备维护或其他原因。
	PointState_kPsOverRange        PointState = 16  //超出范围,表示测点的值超出了其正常范围，通常用于指示测点的测量值超出了可接受的上限。
	PointState_kPsUnderRange       PointState = 17  //低于范围,表示测点的值低于其正常范围，通常用于指示测点的测量值低于可接受的下限。
	PointState_kPsShowError        PointState = 100 //>=100为错误状态
	PointState_kPsNotFound         PointState = 101 //错误,表示测点的值无效或处于错误状态。
	PointState_kPsBad              PointState = 102 //错误,表示测点的值无效或处于错误状态。
	PointState_kPsDeviceFailure    PointState = 106 //设备故障,表示测点所关联的设备发生故障，无法提供有效的数据。
	PointState_kPsSensorFailure    PointState = 107 //传感器故障,表示与测点关联的传感器发生故障，导致无法获得有效的测量数据。
	PointState_kPsCommunFailure    PointState = 108 //通信故障,表示测点与数据源之间的通信故障，导致无法获取或更新数据。
)

// Enum value maps for PointState.
var (
	PointState_name = map[int32]string{
		0:   "kPsInit",
		1:   "kPsGood",
		10:  "kPsShowWarnning",
		11:  "kPsUncertain",
		12:  "kPsInavtive",
		13:  "kPsQualityIssue",
		14:  "kPsDataTypeMisMatch",
		15:  "kPsOutService",
		16:  "kPsOverRange",
		17:  "kPsUnderRange",
		100: "kPsShowError",
		101: "kPsNotFound",
		102: "kPsBad",
		106: "kPsDeviceFailure",
		107: "kPsSensorFailure",
		108: "kPsCommunFailure",
	}
	PointState_value = map[string]int32{
		"kPsInit":             0,
		"kPsGood":             1,
		"kPsShowWarnning":     10,
		"kPsUncertain":        11,
		"kPsInavtive":         12,
		"kPsQualityIssue":     13,
		"kPsDataTypeMisMatch": 14,
		"kPsOutService":       15,
		"kPsOverRange":        16,
		"kPsUnderRange":       17,
		"kPsShowError":        100,
		"kPsNotFound":         101,
		"kPsBad":              102,
		"kPsDeviceFailure":    106,
		"kPsSensorFailure":    107,
		"kPsCommunFailure":    108,
	}
)

func (m PointState) String() string {
	if str, ok := PointState_name[int32(m)]; ok {
		return str
	}
	return "kPsInit"
}

func (m PointState) ParseString(s string) PointState {
	if m, ok := PointState_value[s]; ok {
		return PointState(m)
	}
	return PointState_kPsInit
}

type GeoValue struct {
	Longitude float32 `json:"longitude" form:"longitude"` //测点值
	Latitude  float32 `json:"latitude" form:"latitude"`   //测点值时间
	Altitude  float32 `json:"altitude" form:"altitude"`   //测点值状态
}

type BlobData []byte
type ByteArray []byte

// 👇 自定义 JSON 序列化方法
func (b BlobData) MarshalJSON() ([]byte, error) {
	hexSlice := make([]string, len(b))
	for i, v := range b {
		hexSlice[i] = fmt.Sprintf("0x%02X", v) // 使用大写16进制，可改成 "0x%02x" 小写
	}
	return json.Marshal(hexSlice)
}
func (b ByteArray) MarshalJSON() ([]byte, error) {
	// 转换为 []int，避免被当作字符串处理
	intSlice := make([]int, len(b))
	for i, v := range b {
		intSlice[i] = int(v)
	}
	return json.Marshal(intSlice)
}

type PointValue struct {
	Value  interface{} `json:"v" form:"v"` //测点值
	Mstime uint64      `json:"t" form:"t"` //测点值时间
	State  PointState  `json:"s" form:"s"` //测点值状态
}

func (pointValue *PointValue) go2grpcPointValue() (grpcValue *rpc.PointValue) {
	var pv rpc.PointValue
	pv.Mstime = pointValue.Mstime
	pv.State = int32(pointValue.State)
	switch pointValue.Value.(type) {
	case bool:
		pv.Value = &rpc.PointValue_BoolValue{BoolValue: pointValue.Value.(bool)}
	case float32:
		pv.Value = &rpc.PointValue_FloatValue{FloatValue: pointValue.Value.(float32)}
	case float64:
		pv.Value = &rpc.PointValue_DoubleValue{DoubleValue: pointValue.Value.(float64)}
	case int8:
		pv.Value = &rpc.PointValue_CharValue{CharValue: int32(pointValue.Value.(int8))}
	case uint8:
		pv.Value = &rpc.PointValue_ByteValue{ByteValue: uint32(pointValue.Value.(uint8))}
	case int16:
		pv.Value = &rpc.PointValue_ShortValue{ShortValue: int32(pointValue.Value.(int16))}
	case uint16:
		pv.Value = &rpc.PointValue_WordValue{WordValue: uint32(pointValue.Value.(uint16))}
	case int32:
		pv.Value = &rpc.PointValue_IntValue{IntValue: pointValue.Value.(int32)}
	case uint32:
		pv.Value = &rpc.PointValue_DwordValue{DwordValue: pointValue.Value.(uint32)}
	case int64:
		pv.Value = &rpc.PointValue_LongValue{LongValue: pointValue.Value.(int64)}
	case uint64:
		pv.Value = &rpc.PointValue_QwordValue{QwordValue: pointValue.Value.(uint64)}
	case GeoValue:
		pv.Value = &rpc.PointValue_GeoValue{&rpc.GeoValue{
			Longitude: pointValue.Value.(GeoValue).Longitude,
			Latitude:  pointValue.Value.(GeoValue).Latitude,
			Altitude:  pointValue.Value.(GeoValue).Altitude}}
	case string:
		pv.Value = &rpc.PointValue_StringValue{StringValue: []byte(pointValue.Value.(string))}
	case BlobData:
		pv.Value = &rpc.PointValue_BlobValue{BlobValue: pointValue.Value.([]byte)}
	case []bool:
		pv.Value = &rpc.PointValue_BoolArr{BoolArr: &rpc.BoolArr{ArrValue: pointValue.Value.([]bool)}}
	case []float32:
		pv.Value = &rpc.PointValue_FloatArr{FloatArr: &rpc.FloatArr{ArrValue: pointValue.Value.([]float32)}}
	case []float64:
		pv.Value = &rpc.PointValue_DoubleArr{DoubleArr: &rpc.DoubleArr{ArrValue: pointValue.Value.([]float64)}}
	case []int8:
		{
			tempValue := rpc.CharArr{}
			for _, v := range pointValue.Value.([]int8) {
				tempValue.ArrValue = append(tempValue.ArrValue, int32(v))
			}
			pv.Value = &rpc.PointValue_CharArr{CharArr: &tempValue}
		}
	case ByteArray:
		{
			tempValue := rpc.ByteArr{}
			for _, v := range pointValue.Value.(ByteArray) {
				tempValue.ArrValue = append(tempValue.ArrValue, uint32(v))
			}
			pv.Value = &rpc.PointValue_ByteArr{ByteArr: &tempValue}
		}
	case []int16:
		{
			tempValue := rpc.ShortArr{}
			for _, v := range pointValue.Value.([]int16) {
				tempValue.ArrValue = append(tempValue.ArrValue, int32(v))
			}
			pv.Value = &rpc.PointValue_ShortArr{ShortArr: &tempValue}
		}
	case []uint16:
		{
			tempValue := rpc.WordArr{}
			for _, v := range pointValue.Value.([]uint16) {
				tempValue.ArrValue = append(tempValue.ArrValue, uint32(v))
			}
			pv.Value = &rpc.PointValue_WordArr{WordArr: &tempValue}
		}
	case []int32:
		pv.Value = &rpc.PointValue_IntArr{IntArr: &rpc.IntArr{ArrValue: pointValue.Value.([]int32)}}
	case []uint32:
		pv.Value = &rpc.PointValue_DwordArr{DwordArr: &rpc.DwordArr{ArrValue: pointValue.Value.([]uint32)}}
	case []int64:
		pv.Value = &rpc.PointValue_LongArr{LongArr: &rpc.LongArr{ArrValue: pointValue.Value.([]int64)}}
	case []uint64:
		pv.Value = &rpc.PointValue_QwordArr{QwordArr: &rpc.QwordArr{ArrValue: pointValue.Value.([]uint64)}}
	}
	return &pv
}

func (pointValue *PointValue) grpc2goPointValue(grpcValue *rpc.PointValue) {
	pointValue.Mstime = grpcValue.Mstime
	pointValue.State = PointState(grpcValue.State)
	switch grpcValue.GetValue().(type) {
	case *rpc.PointValue_BoolValue:
		if grpcValue.GetBoolValue() {
			pointValue.Value = 1
		} else {
			pointValue.Value = 0
		}
	case *rpc.PointValue_FloatValue:
		pointValue.Value = grpcValue.GetFloatValue()
	case *rpc.PointValue_DoubleValue:
		pointValue.Value = grpcValue.GetDoubleValue()
	case *rpc.PointValue_CharValue:
		pointValue.Value = int8(grpcValue.GetCharValue())
	case *rpc.PointValue_ByteValue:
		pointValue.Value = uint8(grpcValue.GetByteValue())
	case *rpc.PointValue_ShortValue:
		pointValue.Value = int16(grpcValue.GetShortValue())
	case *rpc.PointValue_WordValue:
		pointValue.Value = uint16(grpcValue.GetWordValue())
	case *rpc.PointValue_IntValue:
		pointValue.Value = grpcValue.GetIntValue()
	case *rpc.PointValue_DwordValue:
		pointValue.Value = grpcValue.GetDwordValue()
	case *rpc.PointValue_LongValue:
		pointValue.Value = grpcValue.GetLongValue()
	case *rpc.PointValue_QwordValue:
		pointValue.Value = grpcValue.GetQwordValue()
	case *rpc.PointValue_GeoValue:
		pointValue.Value = GeoValue{grpcValue.GetGeoValue().GetLongitude(), grpcValue.GetGeoValue().GetLatitude(), grpcValue.GetGeoValue().GetAltitude()}
	case *rpc.PointValue_StringValue:
		pointValue.Value = string(grpcValue.GetStringValue())
	case *rpc.PointValue_BlobValue:
		pointValue.Value = BlobData(grpcValue.GetBlobValue())
	case *rpc.PointValue_BoolArr:
		pointValue.Value = grpcValue.GetBoolArr().GetArrValue()
	case *rpc.PointValue_FloatArr:
		pointValue.Value = grpcValue.GetFloatArr().GetArrValue()
	case *rpc.PointValue_DoubleArr:
		pointValue.Value = grpcValue.GetDoubleArr().GetArrValue()
	case *rpc.PointValue_CharArr:
		var charArr []int8
		for _, v := range grpcValue.GetCharArr().GetArrValue() {
			charArr = append(charArr, int8(v))
		}
		pointValue.Value = charArr
	case *rpc.PointValue_ByteArr:
		var byteArr ByteArray
		for _, v := range grpcValue.GetByteArr().GetArrValue() {
			byteArr = append(byteArr, uint8(v))
		}
		pointValue.Value = byteArr
	case *rpc.PointValue_ShortArr:
		var shortArr []int16
		for _, v := range grpcValue.GetShortArr().GetArrValue() {
			shortArr = append(shortArr, int16(v))
		}
		pointValue.Value = shortArr
	case *rpc.PointValue_WordArr:
		var wordArr []uint16
		for _, v := range grpcValue.GetWordArr().GetArrValue() {
			wordArr = append(wordArr, uint16(v))
		}
		pointValue.Value = wordArr
	case *rpc.PointValue_IntArr:
		pointValue.Value = grpcValue.GetIntArr().GetArrValue()
	case *rpc.PointValue_DwordArr:
		pointValue.Value = grpcValue.GetDwordArr().GetArrValue()
	case *rpc.PointValue_LongArr:
		pointValue.Value = grpcValue.GetLongArr().GetArrValue()
	case *rpc.PointValue_QwordArr:
		pointValue.Value = grpcValue.GetQwordArr().GetArrValue()
	}
}

// 实时值--写入实时值
// 参数说明：dbName：数据库名，pointIdList:测点ID列表 ，valueList：测点值列表，useSysTime：是否使用系统时间写入
// 返回值：int32:成功>=0,写入成功的个数,失败<0，*[]int32：各个值写入的状态，失败<0为对应的错误码，全成功时为空，error：错误信息
func (hhdb *HhdbConPool) UpdateRealtimeValueListByIdList(dbName string, pointIdList *[]int32, valueList *[]PointValue, useSysTime bool) (int32, *[]int32, error) {
	if pointIdList == nil || valueList == nil {
		return -1, nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.UpdateRealtimeValueListReq{}
	req.IdList = *pointIdList
	for _, v := range *valueList {
		req.ValueList = append(req.ValueList, v.go2grpcPointValue())
	}
	req.UseServerTimeFlag = useSysTime
	res, err := dbConInfo.dbClient.UpdateRealtimeValueList(ctx, &req)
	if err != nil {
		return 0, nil, hhdb.handleGrpcError(&err)
	}

	return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, nil
}

// 实时值--写入实时值
// 参数说明：dbName：数据库名，pointNameList:测点名列表 ，valueList：测点值列表，useSysTime：是否使用系统时间写入
// 返回值：int32:成功>=0,写入成功的个数,失败<0，*[]int32：各个值写入的状态，失败<0为对应的错误码，全成功时为空，error：错误信息
func (hhdb *HhdbConPool) UpdateRealtimeValueListByNameList(dbName string, pointNameList *[]string, valueList *[]PointValue, useSysTime bool) (int32, *[]int32, error) {
	if pointNameList == nil || valueList == nil {
		return -1, nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.UpdateRealtimeValueListReq{NameList: *pointNameList, UseServerTimeFlag: useSysTime}
	for _, v := range *valueList {
		req.ValueList = append(req.ValueList, v.go2grpcPointValue())
	}
	res, err := dbConInfo.dbClient.UpdateRealtimeValueList(ctx, &req)
	if err != nil {
		return 0, nil, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return 0, nil, errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, nil
}

// 实时值--查询实时值
// 参数说明：dbName：数据库名，pointIdList:测点名列表
// 返回值：valueList：测点值列表，*[]int32：各个值查询的状态，失败<0为对应的错误码，全成功时为空，error：错误信息
func (hhdb *HhdbConPool) QueryRealtimeValueListByIdList(dbName string, pointIdList *[]int32) (*[]PointValue, *[]int32, error) {
	if pointIdList == nil {
		return nil, nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.IdOrNameListReq{}
	req.IdList = *pointIdList

	res, err := dbConInfo.dbClient.QueryRealtimeValueList(ctx, &req)
	if err != nil {
		return nil, nil, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, nil, errors.New(res.GetErrMsg().GetMsg())
	}

	valueList := make([]PointValue, len(res.ValueList))
	for i, v := range res.ValueList {
		valueList[i].grpc2goPointValue(v)
	}

	return &valueList, &res.ResultList, nil
}

// 实时值--查询实时值
// 参数说明：dbName：数据库名，pointNameList:测点名列表
// 返回值：valueList：测点值列表，*[]int32：各个值查询的状态，失败<0为对应的错误码，全成功时为空，error：错误信息
func (hhdb *HhdbConPool) QueryRealtimeValueListByNameList(dbName string, pointNameList *[]string) (*[]PointValue, *[]int32, error) {
	if pointNameList == nil {
		return nil, nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.IdOrNameListReq{}
	req.NameList = *pointNameList

	res, err := dbConInfo.dbClient.QueryRealtimeValueList(ctx, &req)
	if err != nil {
		return nil, nil, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, nil, errors.New(res.GetErrMsg().GetMsg())
	}

	valueList := make([]PointValue, len(res.ValueList))
	for i, v := range res.ValueList {
		valueList[i].grpc2goPointValue(v)
	}

	return &valueList, &res.ResultList, nil
}

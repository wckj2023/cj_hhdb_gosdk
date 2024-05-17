package cj_hhdb_gosdk

import (
	"context"
	"errors"
	"github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc"
	hhdbRpc "github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc_interface"
)

type PointStatus int32

const (
	PointStatus_kPsInit             PointStatus = 0   //初始化0-9为好的状态
	PointStatus_kPsGood             PointStatus = 1   //正常,表示测点的值是有效的且在正常范围内。这是最常见的状态
	PointStatus_kPsShowWarnning     PointStatus = 10  //>=10为警告状态
	PointStatus_kPsUncertain        PointStatus = 11  //不确定,表示测点的值可能有效，但存在某种不确定性或不可靠性。这通常用于表示数据质量差、未知状态或临时问题。
	PointStatus_kPsInavtive         PointStatus = 12  //不活动,表示测点当前处于非活动状态，可能是因为设备已停用或者该测点当前不可用。
	PointStatus_kPsQualityIssue     PointStatus = 13  //数据质量问题,表示测点的数据存在质量问题，可能是由于噪声、干扰或其他原因导致的不准确性。
	PointStatus_kPsDataTypeMisMatch PointStatus = 14  //数据类型不匹配
	PointStatus_kPsOutService       PointStatus = 15  //停用,表示测点当前处于停用状态，通常是因为设备维护或其他原因。
	PointStatus_kPsOverRange        PointStatus = 16  //超出范围,表示测点的值超出了其正常范围，通常用于指示测点的测量值超出了可接受的上限。
	PointStatus_kPsUnderRange       PointStatus = 17  //低于范围,表示测点的值低于其正常范围，通常用于指示测点的测量值低于可接受的下限。
	PointStatus_kPsShowError        PointStatus = 100 //>=100为错误状态
	PointStatus_kPsNotFound         PointStatus = 101 //错误,表示测点的值无效或处于错误状态。
	PointStatus_kPsBad              PointStatus = 102 //错误,表示测点的值无效或处于错误状态。
	PointStatus_kPsDeviceFailure    PointStatus = 106 //设备故障,表示测点所关联的设备发生故障，无法提供有效的数据。
	PointStatus_kPsSensorFailure    PointStatus = 107 //传感器故障,表示与测点关联的传感器发生故障，导致无法获得有效的测量数据。
	PointStatus_kPsCommunFailure    PointStatus = 108 //通信故障,表示测点与数据源之间的通信故障，导致无法获取或更新数据。
)

// Enum value maps for PointStatus.
var (
	PointStatus_name = map[int32]string{
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
	PointStatus_value = map[string]int32{
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

func (x PointStatus) Enum() *PointStatus {
	p := new(PointStatus)
	*p = x
	return p
}

func (x PointStatus) StrEnum(str string) PointStatus {
	return PointStatus(PointStatus_value[str])
}

func (x PointStatus) String() string {
	return PointStatus_name[int32(x)]
}

type PointValue struct {
	Value  interface{} `json:"v"` //测点值
	Mstime uint64      `json:"t"` //测点值时间
	Status PointStatus `json:"s"` //测点值状态
}

func (pointValue *PointValue) go2grpcPointValue() (grpcValue *rpc.PointValue) {
	var pv rpc.PointValue
	pv.Mstime = pointValue.Mstime
	pv.Status = int32(pointValue.Status)
	switch pointValue.Value.(type) {
	case bool:
		pv.Value = &rpc.PointValue_BoolValue{BoolValue: pointValue.Value.(bool)}
	case float32:
		pv.Value = &rpc.PointValue_FloatValue{FloatValue: pointValue.Value.(float32)}
	case float64:
		pv.Value = &rpc.PointValue_DoubleValue{DoubleValue: pointValue.Value.(float64)}
	case int8:
		pv.Value = &rpc.PointValue_IntValue{IntValue: int32(pointValue.Value.(int8))}
	case int16:
		pv.Value = &rpc.PointValue_IntValue{IntValue: int32(pointValue.Value.(int16))}
	case int32:
		pv.Value = &rpc.PointValue_IntValue{IntValue: pointValue.Value.(int32)}
	case int64:
		pv.Value = &rpc.PointValue_LongValue{LongValue: pointValue.Value.(int64)}
	case uint8:
		pv.Value = &rpc.PointValue_DwordValue{DwordValue: uint32(pointValue.Value.(uint8))}
	case uint16:
		pv.Value = &rpc.PointValue_DwordValue{DwordValue: uint32(pointValue.Value.(uint16))}
	case uint32:
		pv.Value = &rpc.PointValue_DwordValue{DwordValue: pointValue.Value.(uint32)}
	case uint64:
		pv.Value = &rpc.PointValue_QwordValue{QwordValue: pointValue.Value.(uint64)}
	case string:
		pv.Value = &rpc.PointValue_StringValue{StringValue: []byte(pointValue.Value.(string))}
	case []bool:
		pv.Value = &rpc.PointValue_BoolArr{BoolArr: &rpc.BoolArr{ArrValue: pointValue.Value.([]bool)}}
	case []float32:
		pv.Value = &rpc.PointValue_FloatArr{FloatArr: &rpc.FloatArr{ArrValue: pointValue.Value.([]float32)}}
	case []float64:
		pv.Value = &rpc.PointValue_DoubleArr{DoubleArr: &rpc.DoubleArr{ArrValue: pointValue.Value.([]float64)}}
	case []int8:
		{
			tempValue := rpc.IntArr{}
			for _, v := range pointValue.Value.([]int8) {
				tempValue.ArrValue = append(tempValue.ArrValue, int32(v))
			}
			pv.Value = &rpc.PointValue_IntArr{IntArr: &tempValue}
		}
	case []int16:
		{
			tempValue := rpc.IntArr{}
			for _, v := range pointValue.Value.([]int16) {
				tempValue.ArrValue = append(tempValue.ArrValue, int32(v))
			}
			pv.Value = &rpc.PointValue_IntArr{IntArr: &tempValue}
		}
	case []int32:
		pv.Value = &rpc.PointValue_IntArr{IntArr: &rpc.IntArr{ArrValue: pointValue.Value.([]int32)}}
	case []int64:
		pv.Value = &rpc.PointValue_LongArr{LongArr: &rpc.LongArr{ArrValue: pointValue.Value.([]int64)}}
	case []uint8:
		{
			tempValue := rpc.DwordArr{}
			for _, v := range pointValue.Value.([]uint8) {
				tempValue.ArrValue = append(tempValue.ArrValue, uint32(v))
			}
			pv.Value = &rpc.PointValue_DwordArr{DwordArr: &tempValue}
		}
	case []uint16:
		{
			tempValue := rpc.DwordArr{}
			for _, v := range pointValue.Value.([]uint16) {
				tempValue.ArrValue = append(tempValue.ArrValue, uint32(v))
			}
			pv.Value = &rpc.PointValue_DwordArr{DwordArr: &tempValue}
		}
	case []uint32:
		pv.Value = &rpc.PointValue_DwordArr{DwordArr: &rpc.DwordArr{ArrValue: pointValue.Value.([]uint32)}}
	case []uint64:
		pv.Value = &rpc.PointValue_QwordArr{QwordArr: &rpc.QwordArr{ArrValue: pointValue.Value.([]uint64)}}
	case []string:
		{
			tempValue := rpc.StringArr{}
			for _, v := range pointValue.Value.([]string) {
				tempValue.ArrValue = append(tempValue.ArrValue, []byte(v))
			}
			pv.Value = &rpc.PointValue_StringArr{StringArr: &tempValue}
		}
	}
	return &pv
}

func (pointValue *PointValue) grpc2goPointValue(grpcValue *rpc.PointValue) {
	pointValue.Mstime = grpcValue.Mstime
	pointValue.Status = PointStatus(grpcValue.Status)
	switch grpcValue.GetValue().(type) {
	case *rpc.PointValue_BoolValue:
		pointValue.Value = grpcValue.GetBoolValue()
	case *rpc.PointValue_FloatValue:
		pointValue.Value = grpcValue.GetFloatValue()
	case *rpc.PointValue_DoubleValue:
		pointValue.Value = grpcValue.GetDoubleValue()
	case *rpc.PointValue_IntValue:
		pointValue.Value = grpcValue.GetIntValue()
	case *rpc.PointValue_DwordValue:
		pointValue.Value = grpcValue.GetDwordValue()
	case *rpc.PointValue_LongValue:
		pointValue.Value = grpcValue.GetLongValue()
	case *rpc.PointValue_QwordValue:
		pointValue.Value = grpcValue.GetQwordValue()
	case *rpc.PointValue_StringValue:
		pointValue.Value = grpcValue.GetStringValue()
	case *rpc.PointValue_BoolArr:
		pointValue.Value = grpcValue.GetBoolArr().GetArrValue()
	case *rpc.PointValue_FloatArr:
		pointValue.Value = grpcValue.GetFloatArr().GetArrValue()
	case *rpc.PointValue_DoubleArr:
		pointValue.Value = grpcValue.GetDoubleArr().GetArrValue()
	case *rpc.PointValue_IntArr:
		pointValue.Value = grpcValue.GetIntArr().GetArrValue()
	case *rpc.PointValue_DwordArr:
		pointValue.Value = grpcValue.GetDoubleArr().GetArrValue()
	case *rpc.PointValue_LongArr:
		pointValue.Value = grpcValue.GetLongArr().GetArrValue()
	case *rpc.PointValue_QwordArr:
		pointValue.Value = grpcValue.GetQwordArr().GetArrValue()
	case *rpc.PointValue_StringArr:
		pointValue.Value = grpcValue.GetStringArr().GetArrValue()
	}
}

func (hhdb *HhdbConPool) UpdateRealtimeValueListByIdList(dbName string, pointIdList *[]int32, valueList *[]PointValue, useSysTime bool) (int32, *[]int32, error) {
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

	res, err := dbConInfo.dbClient.UpdateRealtimeValueList(ctx, &req)
	if err != nil {
		return 0, nil, hhdb.handleGrpcError(&err)
	}

	return res.GetErrMsg().GetCode(), &res.IdOrErrCodeList, nil
}

func (hhdb *HhdbConPool) UpdateRealtimeValueListByNameList(dbName string, pointNameList *[]string, valueList *[]PointValue, useSysTime bool) (int32, *[]int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return 0, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.UpdateRealtimeValueListReq{}
	req.NameList = *pointNameList
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

func (hhdb *HhdbConPool) QueryRealtimeValueListByIdList(dbName string, pointIdList *[]int32) (*[]PointValue, *[]int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryRealtimeValueListReq{}
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

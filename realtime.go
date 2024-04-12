package hhdb_sdk

import (
	"context"
	"errors"
	"hhdb_sdk/hhdb/rpc"
	hhdbRpc "hhdb_sdk/hhdb/rpc_interface"
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

func (x PointStatus) String() string {
	return PointStatus_name[int32(x)]
}

type PointValue struct {
	value  interface{} `json:"v"` //测点值
	mstime uint64      `json:"t"` //测点值时间
	status PointStatus `json:"s"` //测点值状态
}

func (pointValue *PointValue) go2grpcPointValue() (grpcValue *rpc.PointValue) {
	switch pointValue.value.(type) {
	case bool:
		grpcValue.Value = &rpc.PointValue_BoolValue{BoolValue: pointValue.value.(bool)}
	case float32:
		grpcValue.Value = &rpc.PointValue_FloatValue{FloatValue: pointValue.value.(float32)}
	case float64:
		grpcValue.Value = &rpc.PointValue_DoubleValue{DoubleValue: pointValue.value.(float64)}
	case int32:
		grpcValue.Value = &rpc.PointValue_Int32Value{Int32Value: pointValue.value.(int32)}
	case uint32:
		grpcValue.Value = &rpc.PointValue_Uint32Value{Uint32Value: pointValue.value.(uint32)}
	case int64:
		grpcValue.Value = &rpc.PointValue_Int64Value{Int64Value: pointValue.value.(int64)}
	case uint64:
		grpcValue.Value = &rpc.PointValue_Uint64Value{Uint64Value: pointValue.value.(uint64)}
	case string:
		grpcValue.Value = &rpc.PointValue_StringValue{StringValue: pointValue.value.(string)}
	case []byte:
		grpcValue.Value = &rpc.PointValue_BinaryValue{BinaryValue: pointValue.value.([]byte)}
	}
	return grpcValue
}

func (pointValue *PointValue) grpc2goPointValue(grpcValue *rpc.PointValue) {
	switch grpcValue.GetValue().(type) {
	case *rpc.PointValue_BoolValue:
		pointValue.value = grpcValue.GetBoolValue()
	case *rpc.PointValue_FloatValue:
		pointValue.value = grpcValue.GetFloatValue()
	case *rpc.PointValue_DoubleValue:
		pointValue.value = grpcValue.GetDoubleValue()
	case *rpc.PointValue_Int32Value:
		pointValue.value = grpcValue.GetInt32Value()
	case *rpc.PointValue_Uint32Value:
		pointValue.value = grpcValue.GetUint32Value()
	case *rpc.PointValue_Int64Value:
		pointValue.value = grpcValue.GetInt64Value()
	case *rpc.PointValue_Uint64Value:
		pointValue.value = grpcValue.GetUint64Value()
	case *rpc.PointValue_StringValue:
		pointValue.value = grpcValue.GetStringValue()
	case *rpc.PointValue_BinaryValue:
		pointValue.value = grpcValue.GetBinaryValue()
	}
}

func (hhdb *HhdbConPool) UpdateRealtimeValueListByIdList(dbName string, pointIdList *[]int32, valueList *[]PointValue, useSysTime bool) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.UpdateRealtimeValueListReq{}
	req.IdList = *pointIdList
	for _, v := range *valueList {
		req.ValueList = append(req.ValueList, v.go2grpcPointValue())
		req.StatusList = append(req.StatusList, rpc.PointStatus(v.status))
		if !useSysTime {
			req.MstimeList = append(req.MstimeList, v.mstime)
		}
	}

	res, err := dbConInfo.dbClinet.UpdateRealtimeValueList(ctx, &req)
	if err != nil {
		return HHDB_RPC_REQ_ERROR, err
	}
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) UpdateRealtimeValueListByNameList(dbName string, pointNameList *[]string, valueList *[]PointValue, useSysTime bool) (int32, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return HHDB_GET_CON_ERROR, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.UpdateRealtimeValueListReq{}
	req.NameList = *pointNameList
	for _, v := range *valueList {
		req.ValueList = append(req.ValueList, v.go2grpcPointValue())
		req.StatusList = append(req.StatusList, rpc.PointStatus(v.status))
		if !useSysTime {
			req.MstimeList = append(req.MstimeList, v.mstime)
		}
	}

	res, err := dbConInfo.dbClinet.UpdateRealtimeValueList(ctx, &req)
	if err != nil {
		return HHDB_RPC_REQ_ERROR, err
	}
	if res.GetErrMsg().GetCode() < 0 {
		return res.GetErrMsg().GetCode(), errors.New(res.GetErrMsg().GetMsg())
	}
	return res.GetErrMsg().GetCode(), nil
}

func (hhdb *HhdbConPool) QueryRealtimeValueListByIdList(dbName string, pointIdList *[]int32) (*[]PointValue, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryRealtimeValueListReq{}
	req.IdList = *pointIdList

	res, err := dbConInfo.dbClinet.QueryRealtimeValueList(ctx, &req)
	if err != nil {
		return nil, err
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}

	valueList := make([]PointValue, len(res.ValueList))
	for i, v := range res.ValueList {
		valueList[i].status = PointStatus(res.StatusList[i])
		valueList[i].mstime = res.MstimeList[i]
		valueList[i].grpc2goPointValue(v)
	}
	return &valueList, nil
}

package cj_hhdb_gosdk

import (
	"context"
	"errors"
	hhdbRpc "github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc_interface"
)

// 降采样取值方式
type ResampleMode int32

const (
	ResampleMode_kRmSuggest ResampleMode = 0 //根据压缩方式自动推荐准确的取值模式
	ResampleMode_kRmBefore  ResampleMode = 1 //当前时间点有值，则取该值，无值则取时间点前的第一个值
	ResampleMode_kRmAfter   ResampleMode = 2 //当前时间点有值，则取该值，无值则取时间点后的第一个值
	ResampleMode_kRmInter   ResampleMode = 3 //当前时间点有值，则取该值，无值则取时间点前、后的第一个值后，线性运算得到时间点值
	ResampleMode_kRmNone    ResampleMode = 4 //当前时间点有值，则取该值，无值则返回默认值
)

// Enum value maps for ResampleMode.
var (
	ResampleMode_name = map[int32]string{
		0: "kRmSuggest",
		1: "kRmBefore",
		2: "kRmAfter",
		3: "kRmInter",
		4: "kRmNone",
	}
	ResampleMode_value = map[string]int32{
		"kRmSuggest": 0,
		"kRmBefore":  1,
		"kRmAfter":   2,
		"kRmInter":   3,
		"kRmNone":    4,
	}
)

func (x ResampleMode) Enum() *ResampleMode {
	p := new(ResampleMode)
	*p = x
	return p
}

func (x ResampleMode) String() string {
	return ResampleMode_name[int32(x)]
}

// 时间段查询模式
type RangeQueryMode int32

const (
	RangeQueryMode_kRqmAll       RangeQueryMode = 0  //时间段内所有数据
	RangeQueryMode_kRqmMax       RangeQueryMode = 1  //时间段内最大值
	RangeQueryMode_kRqmMin       RangeQueryMode = 2  //时间段内最小值
	RangeQueryMode_kRqmSum       RangeQueryMode = 3  //时间段内和值
	RangeQueryMode_kRqmAvg       RangeQueryMode = 4  //时间段内平均值
	RangeQueryMode_kRqmWeightSum RangeQueryMode = 5  //时间段内加权和值
	RangeQueryMode_kRqmWeightAvg RangeQueryMode = 6  //时间段内加权平均值
	RangeQueryMode_kRqmFirst     RangeQueryMode = 7  //取第一个值
	RangeQueryMode_kRqmLast      RangeQueryMode = 8  //取最后一个值
	RangeQueryMode_kRqmTimeDiff  RangeQueryMode = 9  //取时间段内开始时间点数据-结束时间点数据，若时间点数据不存在则kResSuggest模式进行查询
	RangeQueryMode_kRqmLastDiff  RangeQueryMode = 10 //取时间段内第一条数据-最后一条数据差值
	RangeQueryMode_kRqmMaxDiff   RangeQueryMode = 11 //取时间段内最大值减最小值的差值
	RangeQueryMode_kRqmMse       RangeQueryMode = 12 //均方差
	RangeQueryMode_kRqmMode      RangeQueryMode = 13 //众数
	RangeQueryMode_kRqmRate      RangeQueryMode = 14 //增长率
)

// Enum value maps for RangeQueryMode.
var (
	RangeQueryMode_name = map[int32]string{
		0:  "kRqmAll",
		1:  "kRqmMax",
		2:  "kRqmMin",
		3:  "kRqmSum",
		4:  "kRqmAvg",
		5:  "kRqmWeightSum",
		6:  "kRqmWeightAvg",
		7:  "kRqmFirst",
		8:  "kRqmLast",
		9:  "kRqmTimeDiff",
		10: "kRqmLastDiff",
		11: "kRqmMaxDiff",
		12: "kRqmMse",
		13: "kRqmMode",
		14: "kRqmRate",
	}
	RangeQueryMode_value = map[string]int32{
		"kRqmAll":       0,
		"kRqmMax":       1,
		"kRqmMin":       2,
		"kRqmSum":       3,
		"kRqmAvg":       4,
		"kRqmWeightSum": 5,
		"kRqmWeightAvg": 6,
		"kRqmFirst":     7,
		"kRqmLast":      8,
		"kRqmTimeDiff":  9,
		"kRqmLastDiff":  10,
		"kRqmMaxDiff":   11,
		"kRqmMse":       12,
		"kRqmMode":      13,
		"kRqmRate":      14,
	}
)

func (x RangeQueryMode) Enum() *RangeQueryMode {
	p := new(RangeQueryMode)
	*p = x
	return p
}

func (x RangeQueryMode) String() string {
	return RangeQueryMode_name[int32(x)]
}

func (hhdb *HhdbConPool) QueryHisRangeValueListReqByIdList(dbName string, pointIdList *[]int32, startMs uint64, endMs uint64, mode RangeQueryMode) (*[][]PointValue, *[]int32, error) {
	if pointIdList == nil {
		return nil, nil, errors.New("id[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryHisRangeValueListReq{}
	req.IdList = *pointIdList
	req.QueryMode = RangeQueryMode_value[mode.String()]
	req.StartMsTime = startMs
	req.EndMsTime = endMs
	res, err := dbConInfo.dbClient.QueryHisRangeValueList(ctx, &req)
	if err != nil {
		return nil, nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return nil, nil, errors.New(res.GetErrMsg().GetMsg())
	}
	// 创建二维数组
	valuesList := make([][]PointValue, len(res.ValueLists))
	for i, vs := range res.ValueLists {
		valuesList[i] = make([]PointValue, len(vs.ValueList))
		for j, v := range vs.ValueList {
			valuesList[i][j].grpc2goPointValue(v)
		}
	}

	result := res.ResultList
	return &valuesList, &result, nil
}

func (hhdb *HhdbConPool) QueryHisRangeValueListReqByNameList(dbName string, pointNameList *[]string, startMs uint64, endMs uint64, mode RangeQueryMode) (*[][]PointValue, *[]int32, error) {
	if pointNameList == nil {
		return nil, nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryHisRangeValueListReq{}
	req.NameList = *pointNameList
	req.QueryMode = RangeQueryMode_value[mode.String()]
	req.StartMsTime = startMs
	req.EndMsTime = endMs
	res, err := dbConInfo.dbClient.QueryHisRangeValueList(ctx, &req)
	if err != nil {
		return nil, nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return nil, nil, errors.New(res.GetErrMsg().GetMsg())
	}
	// 创建二维数组
	valuesList := make([][]PointValue, len(res.ValueLists))
	for i, vs := range res.ValueLists {
		valuesList[i] = make([]PointValue, len(vs.ValueList))
		for j, v := range vs.ValueList {
			valuesList[i][j].grpc2goPointValue(v)
		}
	}

	result := res.ResultList
	return &valuesList, &result, nil
}

func (hhdb *HhdbConPool) QueryHisResampleValueListByIdList(dbName string, pointIdList *[]int32, startMs uint64, endMs uint64, priodMs uint64, mode ResampleMode) (*[][]PointValue, *[]int32, error) {
	if pointIdList == nil {
		return nil, nil, errors.New("id [] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryHisResamplesValueListReq{}
	req.IdList = *pointIdList
	req.ResampleMode = ResampleMode_value[mode.String()]
	req.StartMsTime = startMs
	req.EndMsTime = endMs
	req.PeriodMs = priodMs
	res, err := dbConInfo.dbClient.QueryHisResampleValueList(ctx, &req)
	if err != nil {
		return nil, nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return nil, nil, errors.New(res.GetErrMsg().GetMsg())
	}
	// 创建二维数组
	valuesList := make([][]PointValue, len(res.ValueLists))
	for i, vs := range res.ValueLists {
		valuesList[i] = make([]PointValue, len(vs.ValueList))
		for j, v := range vs.ValueList {
			valuesList[i][j].grpc2goPointValue(v)
		}
	}

	result := res.ResultList
	return &valuesList, &result, nil
}

func (hhdb *HhdbConPool) QueryHisResampleValueListByNameList(dbName string, pointNameList *[]string, startMs uint64, endMs uint64, priodMs uint64, mode ResampleMode) (*[][]PointValue, *[]int32, error) {
	if pointNameList == nil {
		return nil, nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryHisResamplesValueListReq{}
	req.NameList = *pointNameList
	req.ResampleMode = ResampleMode_value[mode.String()]
	req.StartMsTime = startMs
	req.EndMsTime = endMs
	req.PeriodMs = priodMs
	res, err := dbConInfo.dbClient.QueryHisResampleValueList(ctx, &req)
	if err != nil {
		return nil, nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return nil, nil, errors.New(res.GetErrMsg().GetMsg())
	}
	// 创建二维数组
	valuesList := make([][]PointValue, len(res.ValueLists))
	for i, vs := range res.ValueLists {
		valuesList[i] = make([]PointValue, len(vs.ValueList))
		for j, v := range vs.ValueList {
			valuesList[i][j].grpc2goPointValue(v)
		}
	}

	result := res.ResultList
	return &valuesList, &result, nil
}

func (hhdb *HhdbConPool) QueryHisTimePointValueListByIdList(dbName string, pointIdList *[]int32, MsTimePointList *[]uint64, mode ResampleMode) (*[]PointValue, *[]int32, error) {
	if pointIdList == nil {
		return nil, nil, errors.New("id [] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryHisTimePointValueListReq{}
	req.IdList = *pointIdList
	req.ResampleMode = ResampleMode_value[mode.String()]
	req.MsTimePointList = *MsTimePointList
	res, err := dbConInfo.dbClient.QueryHisTimePointValueList(ctx, &req)
	if err != nil {
		return nil, nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return nil, nil, errors.New(res.GetErrMsg().GetMsg())
	}
	// 创建二维数组
	valueList := make([]PointValue, len(res.ValueList))
	for i, v := range res.ValueList {
		valueList[i].grpc2goPointValue(v)
	}

	result := res.ResultList
	return &valueList, &result, nil
}

func (hhdb *HhdbConPool) QueryHisTimePointValueListByNameList(dbName string, pointNameList *[]string, MsTimePointList *[]uint64, mode ResampleMode) (*[]PointValue, *[]int32, error) {
	if pointNameList == nil {
		return nil, nil, errors.New("name[] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryHisTimePointValueListReq{}
	req.NameList = *pointNameList
	req.ResampleMode = ResampleMode_value[mode.String()]
	req.MsTimePointList = *MsTimePointList
	res, err := dbConInfo.dbClient.QueryHisTimePointValueList(ctx, &req)
	if err != nil {
		return nil, nil, hhdb.handleGrpcError(&err)
	}

	if res.GetErrMsg().GetCode() < 0 {
		return nil, nil, errors.New(res.GetErrMsg().GetMsg())
	}
	// 创建二维数组
	valueList := make([]PointValue, len(res.ValueList))
	for i, v := range res.ValueList {
		valueList[i].grpc2goPointValue(v)
	}

	result := res.ResultList
	return &valueList, &result, nil
}

func (hhdb *HhdbConPool) InsertHisValueListByIdList(dbName string, pointIdList *[]int32, valueList *[][]PointValue) (*[]int32, error) {
	if pointIdList == nil || valueList == nil {
		return nil, errors.New("id or value [] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.InsertHisValueListReq{}
	req.IdList = *pointIdList
	for _, vs := range *valueList {
		grpcValueList := hhdbRpc.ValueList{}
		for _, v := range vs {
			grpcValueList.ValueList = append(grpcValueList.ValueList, v.go2grpcPointValue())
		}
		req.ValueLists = append(req.ValueLists, &grpcValueList)
	}

	res, err := dbConInfo.dbClient.InsertHisValueList(ctx, &req)
	if err != nil {
		return nil, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}
	result := res.IdOrErrCodeList
	return &result, nil
}

func (hhdb *HhdbConPool) InsertHisValueListByNameList(dbName string, pointNameList *[]string, valueList *[][]PointValue) (*[]int32, error) {
	if pointNameList == nil || valueList == nil {
		return nil, errors.New("name or value [] is empty")
	}
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.InsertHisValueListReq{}
	req.NameList = *pointNameList
	for _, vs := range *valueList {
		grpcValueList := hhdbRpc.ValueList{}
		for _, v := range vs {
			grpcValueList.ValueList = append(grpcValueList.ValueList, v.go2grpcPointValue())
		}
		req.ValueLists = append(req.ValueLists, &grpcValueList)
	}

	res, err := dbConInfo.dbClient.InsertHisValueList(ctx, &req)
	if err != nil {
		return nil, hhdb.handleGrpcError(&err)
	}
	if res.GetErrMsg().GetCode() < 0 {
		return nil, errors.New(res.GetErrMsg().GetMsg())
	}
	result := res.IdOrErrCodeList
	return &result, nil
}

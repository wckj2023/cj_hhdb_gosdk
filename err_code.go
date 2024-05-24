package cj_hhdb_gosdk

import (
	"context"
	"errors"
	hhdbRpc "github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc_interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var HHDB_CONNECT_ERR = errors.New("connect hhdb failed;")
var HHDB_LOSS_DB_CONNECT_PARAMS_ERR = errors.New("loss db connect params;")

func (hhdb *HhdbConPool) handleGrpcError(grpcErr *error) error {
	if *grpcErr != nil {
		st, ok := status.FromError(*grpcErr)
		if ok && st.Code() == codes.Unavailable {
			time.Sleep(hhdb.reconnectTimewait)
			return HHDB_CONNECT_ERR
		}
	}
	return *grpcErr
}

func (hhdb *HhdbConPool) QueryErrorCode(dbName string, errCodeList *[]int32) (*[]string, error) {
	dbConInfo, err := hhdb.getDbCon(dbName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
	defer cancel()
	req := hhdbRpc.QueryErrMsgReq{}
	req.ErrCodeList = *errCodeList

	res, err := dbConInfo.dbClient.QueryErrMsg(ctx, &req)
	if err != nil {
		return nil, hhdb.handleGrpcError(&err)
	}
	return &res.ErrMsgList, nil
}

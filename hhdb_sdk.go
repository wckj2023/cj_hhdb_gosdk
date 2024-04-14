package cj_hhdb_gosdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	hhdbRpc "github.com/wckj2023/cj_hhdb_gosdk/hhdb/rpc_interface"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

type DbInfo struct {
	username string `json:"username"` //用户名
	password string `json:"password"` //用户密码
	url      string `json:"url"`      //连接信息127.0.0.1:6666
	dbName   string
}

type dbConObject struct {
	dbInfo   DbInfo
	dbCon    *grpc.ClientConn
	dbClinet hhdbRpc.RpcInterfaceClient
	isAuth   bool
	token    string
	dbId     int32
}

type HhdbConPool struct {
	dbConPool sync.Map
	outtime   time.Duration
}

func (hhdb *HhdbConPool) SetDbInfo(info *DbInfo) {
	object := dbConObject{*info, nil, nil, false, "", -1}
	hhdb.dbConPool.Store(info.dbName, &object)
}

func (hhdb *HhdbConPool) SetOuttime(outtimeSec time.Duration) {
	hhdb.outtime = outtimeSec
}

func generateMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func (hhdb *HhdbConPool) getDbCon(dbName string) (*dbConObject, error) {
	value, found := hhdb.dbConPool.Load(dbName)
	if !found {
		return nil, errors.New("please call SetDbInfo to add db connect info;")
	}

	dbConInfo, typeOk := value.(*dbConObject)
	if !typeOk {
		return nil, errors.New("db pool con type is error")
	}

	//连接未建立或连接断开，则进行重连
	if dbConInfo.dbCon == nil || dbConInfo.dbCon.GetState().String() == "SHUTDOWN" {
		con, err := grpc.Dial(dbConInfo.dbInfo.url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		dbConInfo.dbCon = con
		dbConInfo.isAuth = false
		dbConInfo.dbClinet = hhdbRpc.NewRpcInterfaceClient(dbConInfo.dbCon)
		hhdb.dbConPool.Store(dbName, dbConInfo)
	}

	if !dbConInfo.isAuth {
		ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
		defer cancel()
		res, err := dbConInfo.dbClinet.Auth(ctx, &hhdbRpc.AuthReq{Username: dbConInfo.dbInfo.username, Password: generateMD5(dbConInfo.dbInfo.password)})
		if err != nil {
			return nil, err
		}
		if res.GetErrMsg().GetCode() < 0 {
			dbConInfo.isAuth = false
			hhdb.dbConPool.Store(dbName, dbConInfo)
			return nil, errors.New(res.GetErrMsg().GetMsg())
		}
		dbConInfo.dbId = res.GetErrMsg().GetCode()
		dbConInfo.isAuth = true
		dbConInfo.token = string(res.GetUserInfo().GetToken())
		hhdb.dbConPool.Store(dbName, dbConInfo)
	}

	return dbConInfo, nil
}

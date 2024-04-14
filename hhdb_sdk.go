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
	Username string `json:"username"` //用户名
	Password string `json:"password"` //用户密码
	Url      string `json:"url"`      //连接信息127.0.0.1:6666
	DbName   string
}

type dbConObject struct {
	DbInfo   DbInfo
	DbCon    *grpc.ClientConn
	DbClinet hhdbRpc.RpcInterfaceClient
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
	hhdb.dbConPool.Store(info.DbName, &object)
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
	if dbConInfo.DbCon == nil || dbConInfo.DbCon.GetState().String() == "SHUTDOWN" {
		con, err := grpc.Dial(dbConInfo.DbInfo.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		dbConInfo.DbCon = con
		dbConInfo.isAuth = false
		dbConInfo.DbClinet = hhdbRpc.NewRpcInterfaceClient(dbConInfo.DbCon)
		hhdb.dbConPool.Store(dbName, dbConInfo)
	}

	if !dbConInfo.isAuth {
		ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
		defer cancel()
		res, err := dbConInfo.DbClinet.Auth(ctx, &hhdbRpc.AuthReq{Username: dbConInfo.DbInfo.Username, Password: generateMD5(dbConInfo.DbInfo.Password)})
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

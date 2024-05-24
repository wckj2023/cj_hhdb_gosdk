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
	dbCon    *grpc.ClientConn
	dbClient hhdbRpc.RpcInterfaceClient
	isAuth   bool
	token    string
	dbId     int32
}

type HhdbConPool struct {
	dbConPool         sync.Map
	outtime           time.Duration
	reconnectTimewait time.Duration
}

func (hhdb *HhdbConPool) AddDbInfo(info *DbInfo) {
	object := dbConObject{*info, nil, nil, false, "", -1}
	hhdb.dbConPool.Store(info.DbName, &object)
}

func (hhdb *HhdbConPool) SetOuttime(outtimeSec time.Duration) {
	hhdb.outtime = outtimeSec
}

func (hhdb *HhdbConPool) SetDisconnectTimewait(reconnectTimewait time.Duration) {
	hhdb.reconnectTimewait = reconnectTimewait
}

func NewHhdbConPool() *HhdbConPool {
	return &HhdbConPool{
		outtime:           30 * time.Second, // 设置 outtime 字段的默认值为 30 秒
		reconnectTimewait: 3 * time.Second,  // 断开连接时的等待时间
	}
}

func generateMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func (hhdb *HhdbConPool) getDbCon(dbName string) (*dbConObject, error) {
	value, found := hhdb.dbConPool.Load(dbName)
	if !found {
		return nil, HHDB_LOSS_DB_CONNECT_PARAMS_ERR
	}

	dbConInfo, typeOk := value.(*dbConObject)
	if !typeOk {
		return nil, errors.New("db pool con type is error")
	}

	//连接未建立或连接断开，则进行重连
	if dbConInfo.dbCon == nil || dbConInfo.dbCon.GetState().String() == "SHUTDOWN" {
		con, err := grpc.Dial(dbConInfo.DbInfo.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, hhdb.handleGrpcError(&err)
		}
		dbConInfo.dbCon = con
		dbConInfo.isAuth = false
		dbConInfo.dbClient = hhdbRpc.NewRpcInterfaceClient(dbConInfo.dbCon)
		hhdb.dbConPool.Store(dbName, dbConInfo)
	}

	if !dbConInfo.isAuth {
		ctx, cancel := context.WithTimeout(context.Background(), hhdb.outtime)
		defer cancel()
		res, err := dbConInfo.dbClient.Auth(ctx, &hhdbRpc.AuthReq{Username: dbConInfo.DbInfo.Username, Password: generateMD5(dbConInfo.DbInfo.Password)})
		if err != nil {
			return nil, hhdb.handleGrpcError(&err)
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

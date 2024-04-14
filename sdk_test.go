package hhdb_sdk

import (
	"testing"
	"time"
)

var hhdbPool HhdbConPool
var dbName string = "hhdb"

func TestAdd(t *testing.T) {
	hhdbPool.SetOuttime(time.Second * 3)
	hhdbPool.SetDbInfo(&DbInfo{"china", "love&peace", "127.0.0.1:60000", dbName})
	tableId, err := hhdbPool.InsertTable(dbName, TableInfo{tableName: "table1", extraFiledAndDesc: map[string]string{"city": "城市"}})
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("insert table table1 success,table id:", tableId)
	}

	list, err := hhdbPool.QueryTableList(dbName, tableId, "", false, 0, 0)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("query table table1 success,table info:", list)
	}
}

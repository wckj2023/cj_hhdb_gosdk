package cj_hhdb_gosdk

import (
	"testing"
)

var dbName string = "hhdb"

func TestAdd(t *testing.T) {
	hhdbPool := NewHhdbConPool()
	hhdbPool.AddDbInfo(&DbInfo{"china", "love&peace", "127.0.0.1:60000", dbName})
	tableId, err := hhdbPool.InsertTable(dbName, TableInfo{TableName: "table1", ExtraFiledAndDesc: map[string]string{"city": "城市"}})
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("insert table table1 success,table id:", tableId)
	}

	list, _, err := hhdbPool.QueryTableList(dbName, tableId, "", false, 0, 0)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("query table table1 success,table info:", list)
	}
}

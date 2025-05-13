package cj_hhdb_gosdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

var dbName string = "hhdb"

type ByteList1 []uint8

// 👇 自定义 JSON 序列化方法
func (b ByteList1) MarshalJSON() ([]byte, error) {
	hexSlice := make([]string, len(b))
	for i, v := range b {
		hexSlice[i] = fmt.Sprintf("0x%02X", v) // 使用大写16进制，可改成 "0x%02x" 小写
	}
	return json.Marshal(hexSlice)
}

func TestAdd(t *testing.T) {
	data := ByteList1{}
	data = nil
	b, _ := json.Marshal(data)
	fmt.Println(string(b)) // 输出: [1,2,3,4]
	//hhdbPool := NewHhdbConPool()
	//hhdbPool.AddDbInfo(&DbInfo{"china", "love&peace", "127.0.0.1:60000", dbName})
	//tableId, err := hhdbPool.InsertTable(dbName, TableInfo{TableName: "table1", ExtraFiledAndDesc: map[string]string{"city": "城市"}})
	//if err != nil {
	//	t.Errorf(err.Error())
	//} else {
	//	t.Log("insert table table1 success,table id:", tableId)
	//}
	//
	//list, _, err := hhdbPool.QueryTableList(dbName, tableId, "", false, 0, 0)
	//if err != nil {
	//	t.Errorf(err.Error())
	//} else {
	//	t.Log("query table table1 success,table info:", list)
	//}
}

package main

import (
	"crypto/md5"
	"encoding/hex"
	"hhdb_sdk/hhdb/rpc"
	"log"
)

func generateMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func main() {
	grpcValue := rpc.PointValue{Value: &rpc.PointValue_BoolValue{true}}
	var _, ok = grpcValue.Value.(*rpc.PointValue_BoolValue)
	log.Printf("", ok)
}

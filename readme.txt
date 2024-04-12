go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

设置国内代理
go env -w GOPROXY=https://goproxy.cn,direct
生成的rpc_interface的文件中需要修改
	rpc "hhdb_sdk/hhdb/rpc"
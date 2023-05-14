package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	// 总配置
	MongoUri MongoUri
}

type MongoUri struct {
	Uri string
}

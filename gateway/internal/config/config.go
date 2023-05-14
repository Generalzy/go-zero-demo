package config

import (
	"github.com/generalzy/zeropan/gateway/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 映射的是配置文件
	UserRpc        zrpc.RpcClientConf
	ScoreRpc       zrpc.RpcClientConf
	AuthMiddleware middleware.AuthMiddleware
	// 校验
	Auth Auth
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

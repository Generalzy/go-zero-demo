package svc

import (
	"github.com/generalzy/zeropan/gateway/internal/config"
	"github.com/generalzy/zeropan/gateway/internal/middleware"
	"github.com/generalzy/zeropan/score/userscore"
	"github.com/generalzy/zeropan/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// 此处初始化配置
	// UserRpc写的是userclient的User
	// go-zero对client封装了一次
	UserRpc  userclient.User
	ScoreRpc userscore.UserScore
	// 中间件
	AuthMiddleware middleware.AuthMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ScoreRpc:       userscore.NewUserScore(zrpc.MustNewClient(c.ScoreRpc)),
		AuthMiddleware: *middleware.NewAuthMiddleware(),
	}
}

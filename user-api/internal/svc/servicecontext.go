package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-play-1/user-api/internal/config"
	"go-zero-play-1/user-api/internal/middleware"
	"go-zero-play-1/user-rpc/user"
	"go-zero-play-1/yys-rpc/yys"
)

type ServiceContext struct {
	Config          config.Config
	UserRpc         user.User
	YysRpc          yys.Yys
	TestMiddleware1 rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:          c,
		UserRpc:         user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		YysRpc:          yys.NewYys(zrpc.MustNewClient(c.YysRpc)),
		TestMiddleware1: middleware.NewTestMiddleware1Middleware().Handle,
	}
}

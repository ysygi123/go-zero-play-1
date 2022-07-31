package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-play-1/user-api/internal/config"
	"go-zero-play-1/user-api/internal/middleware"
	"go-zero-play-1/user-rpc/user"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     user.User
	TestMiddle1 rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:      c,
		UserRpc:     user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		TestMiddle1: middleware.NewTestMiddleware1Middleware().Handle,
	}
}

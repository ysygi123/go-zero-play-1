package main

import (
	"flag"
	"fmt"
	"go-zero-play-1/common/symysql"

	"go-zero-play-1/yys-rpc/internal/config"
	"go-zero-play-1/yys-rpc/internal/server"
	"go-zero-play-1/yys-rpc/internal/svc"
	"go-zero-play-1/yys-rpc/types/yys"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/yys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewYysServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		yys.RegisterYysServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	err := symysql.InitSyMysql(c.Mysql.MasterHost, c.Mysql.SlaveHosts)
	if err != nil {
		fmt.Println("初始化 数据库 失败 ...", err)
		return
	}
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

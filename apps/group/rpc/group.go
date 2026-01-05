package main

import (
	"flag"
	"fmt"

	"easy-chat/apps/group/rpc/group"
	"easy-chat/apps/group/rpc/internal/config"
	"easy-chat/apps/group/rpc/internal/server"
	"easy-chat/apps/group/rpc/internal/svc"
	"easy-chat/pkg/idgen"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/group.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 🔥 新增：初始化雪花算法
	if err := idgen.InitSnowflake(c.SnowflakeNodeId); err != nil {
		panic(err)
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		group.RegisterGroupServer(grpcServer, server.NewGroupServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

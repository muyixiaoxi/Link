package main

import (
	"flag"
	"fmt"

	"user/service/tag/service/internal/config"
	tagloginServer "user/service/tag/service/internal/server/taglogin"
	tagsignServer "user/service/tag/service/internal/server/tagsign"
	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/tag.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		tag.RegisterTagLoginServer(grpcServer, tagloginServer.NewTagLoginServer(ctx))
		tag.RegisterTagSignServer(grpcServer, tagsignServer.NewTagSignServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

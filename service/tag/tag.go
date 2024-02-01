package main

import (
	"flag"
	"fmt"

	"Link/service/tag/internal/config"
	tagloginServer "Link/service/tag/internal/server/taglogin"
	tagloginfrontServer "Link/service/tag/internal/server/tagloginfront"
	"Link/service/tag/internal/svc"
	"Link/service/tag/tag"

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
		tag.RegisterTagLoginFrontServer(grpcServer, tagloginfrontServer.NewTagLoginFrontServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

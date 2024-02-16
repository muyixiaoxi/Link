package main

import (
	"Link/service/tag/internal/config"
	tagloginServer "Link/service/tag/internal/server/taglogin"
	tagSignServer "Link/service/tag/internal/server/tagsign"
	"Link/service/tag/internal/svc"
	"Link/service/tag/tag"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/tag/etc/tag.yaml", "the config file")

func main() {
	flag.Parse()

	var cfg logx.LogConf
	_ = conf.FillDefault(&cfg)
	cfg.Mode = "file"
	cfg.Path = "service/tag/logs"
	logc.MustSetup(cfg)

	logc.Info(context.Background(), "hello world")

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		tag.RegisterTagLoginServer(grpcServer, tagloginServer.NewTagLoginServer(ctx))
		tag.RegisterTagSignServer(grpcServer, tagSignServer.NewTagSignServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	defer logc.Close()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

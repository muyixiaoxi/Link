package main

import (
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
	"tag/service/internal/config"
	tagloginServer "tag/service/internal/server/taglogin"
	tagSignServer "tag/service/internal/server/tagsign"
	"tag/service/internal/svc"
	"tag/service/tag"
)

var configFile = flag.String("f", "app/tag/service/etc/tag.yaml", "the config file")

func main() {
	flag.Parse()

	var cfg logx.LogConf
	_ = conf.FillDefault(&cfg)
	cfg.Mode = "file"
	cfg.Path = "app/tag/service/logs"
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

package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"tag/common/jwt"
	"tag/restful/internal/config"
	"tag/restful/internal/handler"
	"tag/restful/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/tag/restful/etc/tag-api.yaml", "the config file")

func main() {
	flag.Parse()

	var cfg logx.LogConf
	_ = conf.FillDefault(&cfg)
	cfg.Mode = "file"
	cfg.Path = "app/tag/restful/tag/logs"
	logc.MustSetup(cfg)

	//logc.Info(context.Background(), "hello world")

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(jwt.JwtUnauthorizedResult), rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"

	"user/restful/internal/config"
	"user/restful/internal/handler"
	"user/restful/internal/svc"

	_ "github.com/dtm-labs/driver-gozero" // 添加导入 `gozero` 的 `dtm` 驱动
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/user/restful/etc/user.yaml", "the config file")

func main() {

	flag.Parse()

	var cfg logx.LogConf
	_ = conf.FillDefault(&cfg)
	cfg.Mode = "file"
	cfg.Path = "app/user/restful/logs"
	logc.MustSetup(cfg)
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

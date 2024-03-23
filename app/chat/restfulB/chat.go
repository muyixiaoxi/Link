package main

import (
	"chat/restfulB/internal/config"
	"chat/restfulB/internal/handler"
	"chat/restfulB/internal/logic"
	"chat/restfulB/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/chat/restfulB/etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 启动转发服务
	var err error
	logic.Conn, err = logic.InitConnect()
	if err == nil {
		go logic.Consumer()
	} else {
		logx.Error("logic.InitConnect failed:", err)
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

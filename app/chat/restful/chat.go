package main

import (
	"chat/restful/internal/config"
	"chat/restful/internal/handler"
	"chat/restful/internal/logic"
	"chat/restful/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/chat/restful/etc/chat.yaml", "the config file")

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
	if err != nil {
		logx.Error("logic.InitConnect() failed:", err)
		return
	}
	go logic.Consumer()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

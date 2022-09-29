package main

import (
	"flag"
	"fmt"
	"micro-mall-server/service/user/api/internal/config"
	"micro-mall-server/service/user/api/internal/handler"
	"micro-mall-server/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	// 将etc/user.yaml配置解析到c
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 上下文
	ctx := svc.NewServiceContext(c)
	// server
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

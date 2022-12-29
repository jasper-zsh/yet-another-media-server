package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-queue/rabbitmq"
	"yet-another-media-server/encoder/internal/queue"

	"yet-another-media-server/encoder/internal/config"
	"yet-another-media-server/encoder/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/encoder.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	handler := queue.NewHandler(ctx, svcCtx)
	listener := rabbitmq.MustNewListener(c.Listener, handler)

	admin := rabbitmq.MustNewAdmin(c.Listener.RabbitConf)
	err := admin.DeclareQueue(rabbitmq.QueueConf{
		Name:    "yams.encoder",
		Durable: true,
	}, nil)
	if err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(listener)
	defer serviceGroup.Stop()
	serviceGroup.Start()
}

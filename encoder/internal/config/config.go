package config

import (
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	MediaLibraryRpc zrpc.RpcClientConf
	Listener        rabbitmq.RabbitListenerConf
}

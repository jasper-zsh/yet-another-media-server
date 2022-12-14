package config

import (
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MediaLibraryRpc zrpc.RpcClientConf
	EncoderSender   rabbitmq.RabbitSenderConf
}

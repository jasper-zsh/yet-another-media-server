package svc

import (
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/zrpc"
	"yet-another-media-server/media_library/medialibraryclient"
	"yet-another-media-server/scanner/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	MediaLibraryRpc medialibraryclient.MediaLibrary
	EncoderSender   rabbitmq.Sender
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		MediaLibraryRpc: medialibraryclient.NewMediaLibrary(zrpc.MustNewClient(c.MediaLibraryRpc)),
		EncoderSender:   rabbitmq.MustNewSender(c.EncoderSender),
	}
}

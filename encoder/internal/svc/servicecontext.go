package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"yet-another-media-server/encoder/internal/config"
	"yet-another-media-server/media_library/medialibraryclient"
)

type ServiceContext struct {
	Config       config.Config
	MediaLibrary medialibraryclient.MediaLibrary
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		MediaLibrary: medialibraryclient.NewMediaLibrary(zrpc.MustNewClient(c.MediaLibraryRpc)),
	}
}

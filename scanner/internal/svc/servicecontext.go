package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"yet-another-media-server/media_library/medialibraryclient"
	"yet-another-media-server/scanner/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	MediaLibraryRpc medialibraryclient.MediaLibrary
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		MediaLibraryRpc: medialibraryclient.NewMediaLibrary(zrpc.MustNewClient(c.MediaLibraryRpc)),
	}
}

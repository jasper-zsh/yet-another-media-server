package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"yet-another-media-server/gateway/internal/config"
	"yet-another-media-server/media_library/medialibraryclient"
	"yet-another-media-server/scanner/scannerclient"
)

type ServiceContext struct {
	Config       config.Config
	MediaLibrary medialibraryclient.MediaLibrary
	Scanner      scannerclient.Scanner
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		MediaLibrary: medialibraryclient.NewMediaLibrary(zrpc.MustNewClient(c.MediaLibraryRpc)),
		Scanner:      scannerclient.NewScanner(zrpc.MustNewClient(c.ScannerRpc)),
	}
}

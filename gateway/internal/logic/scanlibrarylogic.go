package logic

import (
	"context"
	"yet-another-media-server/scanner/scannerclient"

	"yet-another-media-server/gateway/internal/svc"
	"yet-another-media-server/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScanLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScanLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScanLibraryLogic {
	return &ScanLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScanLibraryLogic) ScanLibrary(req *types.ScanLibraryRequest) (resp *types.ScanLibraryResponse, err error) {
	_, err = l.svcCtx.Scanner.ScanLibrary(l.ctx, &scannerclient.ScanLibraryRequest{
		LibraryId: req.LibraryID,
		Options:   req.Options,
	})
	return
}

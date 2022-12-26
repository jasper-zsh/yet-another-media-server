package logic

import (
	"context"
	"yet-another-media-server/scanner/internal/scanner/local"

	"yet-another-media-server/scanner/internal/svc"
	"yet-another-media-server/scanner/scanner"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScanLibraryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	scanner *local.LocalScanner
}

func NewScanLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScanLibraryLogic {
	return &ScanLibraryLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		scanner: local.NewLocalScanner(context.Background(), svcCtx),
	}
}

func (l *ScanLibraryLogic) ScanLibrary(in *scanner.ScanLibraryRequest) (*scanner.ScanLibraryResponse, error) {
	go func() {
		err := l.scanner.Scan(in.LibraryId, in.Options)
		if err != nil {
			l.Logger.Errorf("Failed to scan library %d %v", in.LibraryId, err)
		}
	}()
	return &scanner.ScanLibraryResponse{}, nil
}

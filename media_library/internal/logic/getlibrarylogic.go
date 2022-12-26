package logic

import (
	"context"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLibraryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLibraryLogic {
	return &GetLibraryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLibraryLogic) GetLibrary(in *media_library.GetLibraryRequest) (*media_library.GetLibraryResponse, error) {
	libs := make([]model.Library, 0)
	result := l.svcCtx.DB.Find(&libs, in.LibraryId)
	if result.Error != nil {
		return nil, result.Error
	}
	resp := &media_library.GetLibraryResponse{}
	if len(libs) > 0 {
		resp.Library = libs[0].ToProto()
	}

	return resp, nil
}

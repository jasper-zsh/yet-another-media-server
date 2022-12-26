package logic

import (
	"context"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMediaLogic {
	return &CreateMediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMediaLogic) CreateMedia(in *media_library.CreateMediaRequest) (*media_library.Media, error) {
	media := model.Media{
		LibraryID: uint(in.LibraryId),
	}
	result := l.svcCtx.DB.Save(&media)
	if result.Error != nil {
		return nil, result.Error
	}
	return media.ToProto(), nil
}

package logic

import (
	"context"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMediaLogic {
	return &GetMediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMediaLogic) GetMedia(in *media_library.GetMediaRequest) (*media_library.Media, error) {
	medias := make([]model.Media, 0)
	result := l.svcCtx.DB.
		Preload("Metadatas", "parent_id = ?", 0).
		Preload("Metadatas.Children").
		Preload("Files", "parent_id = ?", 0).
		Preload("Files.Children").
		Where("id = ?", in.MediaId).Find(&medias)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(medias) == 0 {
		return nil, nil
	}
	return medias[0].ToProto(), nil
}

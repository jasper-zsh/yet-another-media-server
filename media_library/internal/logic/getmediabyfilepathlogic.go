package logic

import (
	"context"
	"yet-another-media-server/media_library/internal/model"
	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMediaByFilePathLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMediaByFilePathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMediaByFilePathLogic {
	return &GetMediaByFilePathLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMediaByFilePathLogic) GetMediaByFilePath(in *media_library.GetMediaByFilePathRequest) (*media_library.GetMediaByFilePathResponse, error) {
	files := make([]model.File, 0)
	result := l.svcCtx.DB.Preload("Media").
		Preload("Media.Metadatas").
		Where("library_id = ? AND file_path = ?", in.LibraryId, in.FilePath).Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}
	resp := &media_library.GetMediaByFilePathResponse{}
	if len(files) > 0 {
		resp.Media = files[0].Media.ToProto()
	}
	return resp, nil
}

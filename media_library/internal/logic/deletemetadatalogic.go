package logic

import (
	"context"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMetadataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMetadataLogic {
	return &DeleteMetadataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMetadataLogic) DeleteMetadata(in *media_library.DeleteMetadataRequest) (*media_library.DeleteMetadataResponse, error) {
	result := l.svcCtx.DB.Delete(&model.Metadata{}, in.MetadataId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &media_library.DeleteMetadataResponse{}, nil
}

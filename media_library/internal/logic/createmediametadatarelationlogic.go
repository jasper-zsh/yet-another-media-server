package logic

import (
	"context"
	"gorm.io/gorm"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMediaMetadataRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMediaMetadataRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMediaMetadataRelationLogic {
	return &CreateMediaMetadataRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMediaMetadataRelationLogic) CreateMediaMetadataRelation(in *media_library.CreateMediaMetadataRelationRequest) (*media_library.CreateMediaMetadataRelationResponse, error) {
	err := l.svcCtx.DB.Model(&model.Media{
		Model: gorm.Model{
			ID: uint(in.MediaId),
		},
	}).Association("Metadatas").Append(&model.Metadata{
		Model: gorm.Model{
			ID: uint(in.MetadataId),
		},
	})
	if err != nil {
		return nil, err
	}
	return &media_library.CreateMediaMetadataRelationResponse{}, nil
}

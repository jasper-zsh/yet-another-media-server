package logic

import (
	"context"
	"yet-another-media-server/media_library/internal/model"
	"yet-another-media-server/media_library/medialibraryclient"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginateMetadataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaginateMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginateMetadataLogic {
	return &PaginateMetadataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaginateMetadataLogic) PaginateMetadata(in *media_library.PaginateMetadataRequest) (*media_library.PaginateMetadataResponse, error) {
	metas := make([]model.Metadata, 0)
	query := l.svcCtx.DB.Model(&model.Metadata{}).
		Preload("Children").
		Where("library_id = ? AND type = ?", in.LibraryId, in.Type)
	var count int64
	result := query.Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	result = query.Limit(int(in.Pagination.Size)).Offset(int(in.Pagination.Size * (in.Pagination.Page - 1))).Find(&metas)
	if result.Error != nil {
		return nil, result.Error
	}
	resp := &media_library.PaginateMetadataResponse{
		Pagination: &media_library.PaginationResult{
			Total: count,
		},
		Metadatas: make([]*medialibraryclient.Metadata, 0, len(metas)),
	}
	for _, t := range metas {
		resp.Metadatas = append(resp.Metadatas, t.ToProto())
	}
	return resp, nil
}

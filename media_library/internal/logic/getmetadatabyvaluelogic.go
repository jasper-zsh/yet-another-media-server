package logic

import (
	"context"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMetadataByValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMetadataByValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMetadataByValueLogic {
	return &GetMetadataByValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMetadataByValueLogic) GetMetadataByValue(in *media_library.GetMetadataByValueRequest) (*media_library.GetMetadataByValueResponse, error) {
	resp := &media_library.GetMetadataByValueResponse{}
	metas := make([]*model.Metadata, 0)
	result := l.svcCtx.DB.Where("library_id = ? AND type = ? AND value = ?", in.LibraryId, in.Type, in.Value).Find(&metas)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(metas) > 0 {
		resp.Metadata = metas[0].ToProto()
	}
	return resp, nil
}

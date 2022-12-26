package logic

import (
	"context"
	"yet-another-media-server/media_library/medialibraryclient"

	"yet-another-media-server/gateway/internal/svc"
	"yet-another-media-server/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginateMediaByMetadataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaginateMediaByMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginateMediaByMetadataLogic {
	return &PaginateMediaByMetadataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaginateMediaByMetadataLogic) PaginateMediaByMetadata(req *types.PaginateMediaByMetadataRequest) (resp *types.PaginationResponse, err error) {
	pReq := &medialibraryclient.PaginateMediaByMetadataRequest{
		LibraryId: req.LibraryID,
		Filters:   make([]*medialibraryclient.Filter, 0, len(req.Filters)),
		Pagination: &medialibraryclient.PaginationOptions{
			Size: req.Size,
			Page: req.Page,
		},
	}
	for _, f := range req.Filters {
		pReq.Filters = append(pReq.Filters, &medialibraryclient.Filter{
			Type:     f.Type,
			Value:    f.Value,
			Operator: f.Operator,
		})
	}
	result, err := l.svcCtx.MediaLibrary.PaginateMediaByMetadata(l.ctx, pReq)
	if err != nil {
		return nil, err
	}
	resp = &types.PaginationResponse{
		Total: result.Pagination.Total,
		List:  make([]interface{}, 0, len(result.Medias)),
	}
	for _, p := range result.Medias {
		resp.List = append(resp.List, types.NewMediaFromProto(p))
	}
	return
}

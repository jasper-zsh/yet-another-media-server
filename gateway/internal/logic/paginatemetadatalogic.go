package logic

import (
	"context"
	"yet-another-media-server/media_library/medialibraryclient"

	"yet-another-media-server/gateway/internal/svc"
	"yet-another-media-server/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginateMetadataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaginateMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginateMetadataLogic {
	return &PaginateMetadataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaginateMetadataLogic) PaginateMetadata(req *types.PaginateMetadataRequest) (resp *types.PaginationResponse, err error) {
	ret, err := l.svcCtx.MediaLibrary.PaginateMetadata(l.ctx, &medialibraryclient.PaginateMetadataRequest{
		Pagination: &medialibraryclient.PaginationOptions{
			Size: req.Size,
			Page: req.Page,
		},
		LibraryId: req.LibraryID,
		Type:      req.Type,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.PaginationResponse{
		Total: ret.Pagination.Total,
		List:  make([]interface{}, 0, len(ret.Metadatas)),
	}
	for _, p := range ret.Metadatas {
		resp.List = append(resp.List, types.NewMetadataFromProto(p))
	}
	return
}

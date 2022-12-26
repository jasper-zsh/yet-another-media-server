package logic

import (
	"context"
	"yet-another-media-server/media_library/medialibraryclient"

	"yet-another-media-server/gateway/internal/svc"
	"yet-another-media-server/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMetadataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMetadataLogic {
	return &DeleteMetadataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMetadataLogic) DeleteMetadata(req *types.DeleteMetadataRequest) (resp *types.DeleteMetadataResponse, err error) {
	_, err = l.svcCtx.MediaLibrary.DeleteMetadata(l.ctx, &medialibraryclient.DeleteMetadataRequest{
		MetadataId: req.MetadataID,
	})
	resp = &types.DeleteMetadataResponse{}
	return
}

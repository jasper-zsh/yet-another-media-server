package logic

import (
	"context"
	"yet-another-media-server/media_library/medialibraryclient"

	"yet-another-media-server/gateway/internal/svc"
	"yet-another-media-server/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMediaLogic {
	return &GetMediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMediaLogic) GetMedia(req *types.GetMediaRequest) (resp *types.Media, err error) {
	media, err := l.svcCtx.MediaLibrary.GetMedia(l.ctx, &medialibraryclient.GetMediaRequest{
		MediaId: req.MediaID,
	})
	if err != nil {
		return nil, err
	}
	return types.NewMediaFromProto(media), nil
}

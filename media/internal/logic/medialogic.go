package logic

import (
	"context"

	"yet-another-media-server/media/internal/svc"
	"yet-another-media-server/media/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MediaLogic {
	return &MediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MediaLogic) Media(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

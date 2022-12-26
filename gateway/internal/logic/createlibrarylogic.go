package logic

import (
	"context"
	"yet-another-media-server/media_library/medialibraryclient"

	"yet-another-media-server/gateway/internal/svc"
	"yet-another-media-server/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLibraryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLibraryLogic {
	return &CreateLibraryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLibraryLogic) CreateLibrary(req *types.CreateLibraryRequest) (resp *types.Library, err error) {
	library, err := l.svcCtx.MediaLibrary.CreateLibrary(l.ctx, &medialibraryclient.CreateLibraryRequest{
		Name:     req.Name,
		BasePath: req.BasePath,
	})
	if err != nil {
		return nil, err
	}
	return types.NewLibraryFromProto(library), nil
}

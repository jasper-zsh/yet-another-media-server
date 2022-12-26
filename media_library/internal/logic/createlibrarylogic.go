package logic

import (
	"context"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLibraryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLibraryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLibraryLogic {
	return &CreateLibraryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLibraryLogic) CreateLibrary(in *media_library.CreateLibraryRequest) (*media_library.Library, error) {
	library := model.Library{
		Name:     in.Name,
		BasePath: in.BasePath,
	}
	result := l.svcCtx.DB.Save(&library)
	if result.Error != nil {
		return nil, result.Error
	}
	return library.ToProto(), nil
}

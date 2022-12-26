package logic

import (
	"context"
	"gorm.io/gorm/clause"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFileLogic {
	return &CreateFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFileLogic) CreateFile(in *media_library.CreateFileRequest) (*media_library.File, error) {
	file := model.File{
		LibraryID: uint(in.LibraryId),
		Type:      in.Type,
		FilePath:  in.FilePath,
		MediaID:   uint(in.MediaId),
		ParentID:  uint(in.ParentId),
	}
	result := l.svcCtx.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "library_id"}, {Name: "file_path"}},
		UpdateAll: true,
	}).Create(&file)
	if result.Error != nil {
		return nil, result.Error
	}
	return file.ToProto(), nil
}

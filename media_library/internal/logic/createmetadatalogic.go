package logic

import (
	"context"
	"gorm.io/gorm"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMetadataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMetadataLogic {
	return &CreateMetadataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMetadataLogic) CreateMetadata(in *media_library.CreateMetadataRequest) (resp *media_library.Metadata, err error) {
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		meta, err := model.GetMetadataByTypeAndValue(tx, uint(in.LibraryId), in.Type, in.Value)
		if err != nil {
			return err
		}
		if meta != nil {
			resp = meta.ToProto()
			return nil
		}
		metadata := model.Metadata{
			LibraryID: uint(in.LibraryId),
			ParentID:  uint(in.ParentId),
			Type:      in.Type,
			Value:     in.Value,
		}
		result := tx.Save(&metadata)
		if result.Error != nil {
			return result.Error
		}
		resp = metadata.ToProto()
		return nil
	})
	return
}

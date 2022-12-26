package logic

import (
	"context"
	"fmt"
	"github.com/juliangruber/go-intersect"
	"gorm.io/gorm"
	"yet-another-media-server/media_library/internal/model"

	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaginateMediaByMetadataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaginateMediaByMetadataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaginateMediaByMetadataLogic {
	return &PaginateMediaByMetadataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaginateMediaByMetadataLogic) PaginateMediaByMetadata(in *media_library.PaginateMediaByMetadataRequest) (*media_library.PaginateMediaByMetadataResponse, error) {
	medias := make([]model.Media, 0)
	query := l.svcCtx.DB.Model(&model.Media{}).Preload("Metadatas").
		Preload("Metadatas.Children").
		Where("media.library_id = ?", in.LibraryId)
	if len(in.Filters) > 0 {
		mediaIds := make([]interface{}, 0)
		for idx, filter := range in.Filters {
			ids := make([]interface{}, 0)
			var q *gorm.DB
			switch filter.Operator {
			case "", "eq":
				q = l.svcCtx.DB.Model(&model.Metadata{}).
					Joins("join media_metadata_relations rel on metadata.id = rel.metadata_id").
					Where("metadata.library_id = ? AND metadata.type = ? AND metadata.value = ?", in.LibraryId, filter.Type, filter.Value).
					Select("rel.media_id")
			case "absent":
				subQ := l.svcCtx.DB.Model(&model.Metadata{}).
					Joins("join media_metadata_relations rel on metadata.id = rel.metadata_id").
					Where("metadata.library_id = ? AND metadata.type = ?", in.LibraryId, filter.Type).
					Select("rel.*")
				q = l.svcCtx.DB.Model(&model.Media{}).
					Joins("left join (?) meta on media.id = meta.media_id", subQ).
					Where("meta.media_id is null").
					Select("media.id")
			case "match":
				q = l.svcCtx.DB.Model(&model.Metadata{}).
					Joins("join media_metadata_relations rel on metadata.id = rel.metadata_id").
					Where("metadata.library_id = ? AND metadata.type = ? AND metadata.value like ?", in.LibraryId, filter.Type, fmt.Sprintf("%%%s%%", filter.Value)).
					Select("rel.media_id")
			default:
				return nil, fmt.Errorf("invalid operator %s", filter.Operator)
			}
			result := q.Debug().Find(&ids)
			if result.Error != nil {
				return nil, result.Error
			}
			if idx == 0 {
				mediaIds = ids
			} else {
				mediaIds = intersect.Simple(mediaIds, ids)
			}
			if len(mediaIds) == 0 {
				break
			}
		}
		query = query.Where("id IN ?", mediaIds)
	}
	var count int64
	result := query.Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	result = query.Limit(int(in.Pagination.Size)).
		Offset(int(in.Pagination.Size * (in.Pagination.Page - 1))).Find(&medias)
	if result.Error != nil {
		return nil, result.Error
	}
	ret := make([]*media_library.Media, 0, len(medias))
	for _, m := range medias {
		ret = append(ret, m.ToProto())
	}
	return &media_library.PaginateMediaByMetadataResponse{
		Pagination: &media_library.PaginationResult{
			Total: count,
		},
		Medias: ret,
	}, nil
}

package video_manager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"yet-another-media-server/media_library/medialibraryclient"
	"yet-another-media-server/scanner/internal/svc"
)

type VideoManagerScanner struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger logx.Logger
}

func NewVideoManagerScanner(ctx context.Context, svcCtx *svc.ServiceContext) *VideoManagerScanner {
	return &VideoManagerScanner{
		ctx:    ctx,
		svcCtx: svcCtx,
		logger: logx.WithContext(ctx),
	}
}

func (s *VideoManagerScanner) Scan(libraryID int32, options map[string]string) error {
	connStr := fmt.Sprintf("file:%s?cache=shared", options["db"])
	db, err := gorm.Open(sqlite.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	videos := make([]Video, 0)
	result := db.Find(&videos)
	if result.Error != nil {
		return result.Error
	}
	for _, video := range videos {
		videoFiles := make([]map[string]interface{}, 0)
		_ = json.Unmarshal(video.VideoFiles, &videoFiles)
		var media *medialibraryclient.Media
		for _, videoFile := range videoFiles {
			if media == nil {
				m, err := s.svcCtx.MediaLibraryRpc.GetMediaByFilePath(s.ctx, &medialibraryclient.GetMediaByFilePathRequest{
					LibraryId: libraryID,
					FilePath:  videoFile["path"].(string),
				})
				if err != nil {
					s.logger.Errorf("Failed to get media by file path %v", err)
					return err
				}
				if m.Media != nil {
					media = m.Media
					break
				}
			}
		}
		if media == nil {
			m, err := s.svcCtx.MediaLibraryRpc.CreateMedia(s.ctx, &medialibraryclient.CreateMediaRequest{
				LibraryId: libraryID,
			})
			if err != nil {
				s.logger.Errorf("Failed to create media: %v", err)
				return err
			}
			media = m
		}

		err = s.handleBasicInfo(libraryID, media, video)
		if err != nil {
			s.logger.Errorf("Failed to commit video basic info %v", err)
			return err
		}

		err = s.handleCover(libraryID, media, video)
		if err != nil {
			s.logger.Errorf("Failed to commit video cover %v", err)
			return err
		}

		err = s.handleVideoFiles(libraryID, media, videoFiles)
		if err != nil {
			s.logger.Errorf("Failed to commit video files %v", err)
			return err
		}

		err = s.handleActors(db, libraryID, media, video)
		if err != nil {
			s.logger.Errorf("Failed to commit actors %v", err)
			return err
		}

		err = s.handleDirectors(db, libraryID, media, video)
		if err != nil {
			s.logger.Errorf("Failed to commit directors %v", err)
			return err
		}

		err = s.handleMakers(db, libraryID, media, video)
		if err != nil {
			s.logger.Errorf("Failed to commit makers %v", err)
			return err
		}

		err = s.handleVideoTags(db, libraryID, media, video)
		if err != nil {
			s.logger.Errorf("Failed to commit video tags %v", err)
			return err
		}
	}

	return nil
}

func (s VideoManagerScanner) handleBasicInfo(libraryID int32, media *medialibraryclient.Media, video Video) error {
	title, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
		LibraryId: libraryID,
		Type:      "title",
		Value:     video.Title,
	})
	if err != nil {
		s.logger.Errorf("Failed to create title %v", err)
		return err
	}
	_, err = s.svcCtx.MediaLibraryRpc.CreateMediaMetadataRelation(s.ctx, &medialibraryclient.CreateMediaMetadataRelationRequest{
		MediaId:    media.Id,
		MetadataId: title.Id,
	})
	if err != nil {
		s.logger.Errorf("Failed to associate title to media %v", err)
		return err
	}

	code, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
		LibraryId: libraryID,
		Type:      "code",
		Value:     video.SourceIdentifier,
	})
	if err != nil {
		s.logger.Errorf("Failed to create code %v", err)
		return err
	}
	_, err = s.svcCtx.MediaLibraryRpc.CreateMediaMetadataRelation(s.ctx, &medialibraryclient.CreateMediaMetadataRelationRequest{
		MediaId:    media.Id,
		MetadataId: code.Id,
	})
	if err != nil {
		s.logger.Errorf("Failed to associate code with media %v", err)
		return err
	}
	return nil
}

func (s VideoManagerScanner) handleCover(libraryID int32, media *medialibraryclient.Media, video Video) error {
	cover := make(map[string]interface{})
	_ = json.Unmarshal(video.Cover, &cover)
	if coverPath, ok := cover["path"]; ok {
		meta, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
			LibraryId: libraryID,
			Type:      "cover",
			Value:     coverPath.(string),
		})
		if err != nil {
			s.logger.Errorf("Failed to create cover: %v", err)
			return err
		}
		_, err = s.svcCtx.MediaLibraryRpc.CreateMediaMetadataRelation(s.ctx, &medialibraryclient.CreateMediaMetadataRelationRequest{
			MediaId:    media.Id,
			MetadataId: meta.Id,
		})
		if err != nil {
			s.logger.Errorf("Failed to create media cover relation: %v", err)
			return err
		}
	}
	return nil
}

func (s VideoManagerScanner) handleVideoFiles(libraryID int32, media *medialibraryclient.Media, videoFiles []map[string]interface{}) error {

	for _, videoFile := range videoFiles {
		if path, ok := videoFile["path"]; ok {
			file, err := s.svcCtx.MediaLibraryRpc.CreateFile(s.ctx, &medialibraryclient.CreateFileRequest{
				LibraryId: libraryID,
				MediaId:   media.Id,
				Type:      "video",
				FilePath:  path.(string),
			})
			if err != nil {
				s.logger.Errorf("Failed to create video file: %v", err)
				return err
			}
			if extra, ok := videoFile["extras"]; ok && extra != nil {
				extraMap := extra.(map[string]interface{})
				if thumbnails, ok := extraMap["thumbnails"]; ok {
					_, err := s.svcCtx.MediaLibraryRpc.CreateFile(s.ctx, &medialibraryclient.CreateFileRequest{
						LibraryId: libraryID,
						MediaId:   media.Id,
						ParentId:  file.Id,
						Type:      "video_thumbnails",
						FilePath:  thumbnails.(string),
					})
					if err != nil {
						s.logger.Errorf("Failed to create video thumbnails file: %v", err)
						return err
					}
				}
			}
		}
	}
	return nil
}

func (s VideoManagerScanner) handleActors(db *gorm.DB, libraryID int32, media *medialibraryclient.Media, video Video) error {
	actors := make([]map[string]interface{}, 0)
	result := db.Table("actors").
		Joins("join video_actor_relations rel on actors.id = rel.actor_id").
		Where("rel.video_id = ?", video.ID).Find(&actors)
	if result.Error != nil {
		return result.Error
	}
	for _, actor := range actors {
		aliases := make([]map[string]interface{}, 0)
		result = db.Table("aliases").
			Joins("join actor_alias_relations rel on aliases.id = rel.alias_id").
			Where("rel.actor_id = ?", actor["id"]).Find(&aliases)
		if result.Error != nil {
			return result.Error
		}
		var originActorId int32
		for _, alias := range aliases {
			aliasMeta, err := s.svcCtx.MediaLibraryRpc.GetMetadataByValue(s.ctx, &medialibraryclient.GetMetadataByValueRequest{
				LibraryId: libraryID,
				Type:      "actor_alias",
				Value:     alias["name"].(string),
			})
			if err != nil {
				return err
			}
			if aliasMeta.Metadata != nil && aliasMeta.Metadata.ParentId > 0 {
				originActorId = aliasMeta.Metadata.ParentId
				break
			}
		}
		if originActorId == 0 {
			actorMeta, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
				LibraryId: libraryID,
				Type:      "actor",
				Value:     aliases[0]["name"].(string),
			})
			if err != nil {
				return err
			}
			originActorId = actorMeta.Id
		}
		_, err := s.svcCtx.MediaLibraryRpc.CreateMediaMetadataRelation(s.ctx, &medialibraryclient.CreateMediaMetadataRelationRequest{
			MediaId:    media.Id,
			MetadataId: originActorId,
		})
		if err != nil {
			return err
		}
		s.logger.Infof("Created relation between actor %d media %d", originActorId, media.Id)
		for _, alias := range aliases {
			aliasMeta, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
				LibraryId: libraryID,
				Type:      "actor_alias",
				ParentId:  originActorId,
				Value:     alias["name"].(string),
			})
			if err != nil {
				return err
			}
			s.logger.Infof("Created alias %d for actor %d", aliasMeta.Id, originActorId)
		}
	}
	return nil
}

func (s VideoManagerScanner) handleDirectors(db *gorm.DB, libraryID int32, media *medialibraryclient.Media, video Video) error {
	directors := make([]map[string]interface{}, 0)
	result := db.Table("directors").
		Joins("join video_director_relations rel on directors.id = rel.director_id").
		Where("rel.video_id = ?", video.ID).Find(&directors)
	if result.Error != nil {
		return result.Error
	}
	for _, director := range directors {
		meta, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
			LibraryId: libraryID,
			Type:      "director",
			Value:     director["name"].(string),
		})
		if err != nil {
			return err
		}
		_, err = s.svcCtx.MediaLibraryRpc.CreateMediaMetadataRelation(s.ctx, &medialibraryclient.CreateMediaMetadataRelationRequest{
			MediaId:    media.Id,
			MetadataId: meta.Id,
		})
		if err != nil {
			return err
		}
		s.logger.Infof("Associated media %d with director %d", media.Id, meta.Id)
	}
	return nil
}

func (s VideoManagerScanner) handleMakers(db *gorm.DB, libraryID int32, media *medialibraryclient.Media, video Video) error {
	makers := make([]map[string]interface{}, 0)
	result := db.Table("makers").
		Joins("join video_maker_relations rel on makers.id = rel.maker_id").
		Where("rel.video_id = ?", video.ID).Find(&makers)
	if result.Error != nil {
		return result.Error
	}
	for _, maker := range makers {
		meta, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
			LibraryId: libraryID,
			Type:      "maker",
			Value:     maker["name"].(string),
		})
		if err != nil {
			return err
		}
		_, err = s.svcCtx.MediaLibraryRpc.CreateMediaMetadataRelation(s.ctx, &medialibraryclient.CreateMediaMetadataRelationRequest{
			MediaId:    media.Id,
			MetadataId: meta.Id,
		})
		if err != nil {
			return err
		}
		s.logger.Infof("Associated media %d with maker %d", media.Id, meta.Id)
	}
	return nil
}

func (s VideoManagerScanner) handleVideoTags(db *gorm.DB, libraryID int32, media *medialibraryclient.Media, video Video) error {
	tags := make([]map[string]interface{}, 0)
	result := db.Table("tags").
		Joins("join video_tag_relations rel on tags.id = rel.tag_id").
		Where("rel.video_id = ?", video.ID).Find(&tags)
	if result.Error != nil {
		return result.Error
	}
	for _, tag := range tags {
		meta, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
			LibraryId: libraryID,
			Type:      "tag",
			Value:     tag["name"].(string),
		})
		if err != nil {
			return err
		}
		_, err = s.svcCtx.MediaLibraryRpc.CreateMediaMetadataRelation(s.ctx, &medialibraryclient.CreateMediaMetadataRelationRequest{
			MediaId:    media.Id,
			MetadataId: meta.Id,
		})
		if err != nil {
			return err
		}
		s.logger.Infof("Associated media %d with tag %d", media.Id, meta.Id)
	}
	return nil
}

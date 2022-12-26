package local

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path"
	"path/filepath"
	"yet-another-media-server/media_library/medialibraryclient"
	"yet-another-media-server/scanner/internal/svc"
)

type LocalScanner struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger logx.Logger
}

func NewLocalScanner(ctx context.Context, svcCtx *svc.ServiceContext) *LocalScanner {
	return &LocalScanner{
		ctx:    ctx,
		svcCtx: svcCtx,
		logger: logx.WithContext(ctx),
	}
}

func (s *LocalScanner) Scan(libraryID int32, options map[string]string) error {
	libraryRes, err := s.svcCtx.MediaLibraryRpc.GetLibrary(s.ctx, &medialibraryclient.GetLibraryRequest{
		LibraryId: libraryID,
	})
	if err != nil {
		return err
	}
	if libraryRes.Library == nil {
		return fmt.Errorf("library %d not found", libraryID)
	}
	library := libraryRes.Library
	err = s.walkDirectory(library, library.BasePath)
	if err != nil {
		return err
	}
	return nil
}

func (s *LocalScanner) walkDirectory(library *medialibraryclient.Library, root string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if !entry.IsDir() && isVideo(entry.Name()) {
			return s.processDirectory(library, root)
		}
	}
	for _, entry := range entries {
		if entry.IsDir() {
			err := s.walkDirectory(library, path.Join(root, entry.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isVideo(name string) bool {
	switch name[len(name)-4:] {
	case ".mp4", ".wmv", ".avi", ".rmvb", ".mkv":
		return true
	}
	return false
}

func (s *LocalScanner) processDirectory(library *medialibraryclient.Library, root string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}
	videoFiles := make([]string, 0)
	var media *medialibraryclient.Media
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !isVideo(entry.Name()) {
			continue
		}
		videoFiles = append(videoFiles, path.Join(root, entry.Name()))
		basePath, _ := filepath.Rel(library.BasePath, root)
		m, err := s.svcCtx.MediaLibraryRpc.GetMediaByFilePath(s.ctx, &medialibraryclient.GetMediaByFilePathRequest{
			LibraryId: library.Id,
			FilePath:  path.Join(basePath, entry.Name()),
		})
		if err != nil {
			return err
		}
		if m.Media != nil {
			media = m.Media
		}
	}
	if media == nil {
		media, err = s.svcCtx.MediaLibraryRpc.CreateMedia(s.ctx, &medialibraryclient.CreateMediaRequest{
			LibraryId: library.Id,
		})
		if err != nil {
			return err
		}
		s.logger.Infof("Created media %d", media.Id)
	}
	hasTitle := false
	for _, meta := range media.Metadatas {
		if meta.Type == "title" {
			hasTitle = true
		}
	}
	if !hasTitle {
		title := path.Base(root)
		meta, err := s.svcCtx.MediaLibraryRpc.CreateMetadata(s.ctx, &medialibraryclient.CreateMetadataRequest{
			LibraryId: library.Id,
			Type:      "title",
			Value:     title,
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
		s.logger.Infof("Associated title %s with media %d", title, media.Id)
	}
	for _, videoFile := range videoFiles {
		relPath, _ := filepath.Rel(library.BasePath, videoFile)
		_, err := s.svcCtx.MediaLibraryRpc.CreateFile(s.ctx, &medialibraryclient.CreateFileRequest{
			LibraryId: library.Id,
			MediaId:   media.Id,
			Type:      "video",
			FilePath:  relPath,
		})
		if err != nil {
			return err
		}
		s.logger.Infof("Associated file %s with media %d", relPath, media.Id)
	}

	return nil
}

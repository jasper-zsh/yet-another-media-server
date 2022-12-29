package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"os/exec"
	"path"
	"yet-another-media-server/encoder/internal/svc"
	"yet-another-media-server/encoder/queue"
	"yet-another-media-server/media_library/medialibraryclient"
)

type Handler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger logx.Logger
}

func NewHandler(ctx context.Context, svcCtx *svc.ServiceContext) *Handler {
	return &Handler{
		ctx:    ctx,
		svcCtx: svcCtx,
		logger: logx.WithContext(ctx),
	}
}

func (h Handler) Consume(message string) error {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(message), &m)
	if err != nil {
		return err
	}
	action, ok := m["action"]
	if !ok {
		h.logger.Errorf("Illegal message: %s", message)
		return nil
	}
	switch action {
	case queue.ActionVTTSprite:
		msg := queue.VTTSpriteMessage{}
		err = mapstructure.Decode(m, &msg)
		if err != nil {
			h.logger.Errorf("Failed to parse message: %s %v", message, err)
			return err
		}
		return h.vttSprite(msg.Media, msg.File)
	default:
		h.logger.Errorf("Illegal message: %s", message)
	}
	return nil
}

func (h Handler) vttSprite(media *queue.Media, videoFile *queue.File) error {
	libraryRes, err := h.svcCtx.MediaLibrary.GetLibrary(h.ctx, &medialibraryclient.GetLibraryRequest{
		LibraryId: media.LibraryID,
	})
	if err != nil {
		return err
	}
	library := libraryRes.Library
	if library == nil {
		h.logger.Errorf("Library %d not found.", media.LibraryID)
		return nil
	}
	videoFilePath := path.Join(library.BasePath, videoFile.FilePath)
	//vttAbsPath, err := vttsprite.GenerateVttSprite(videoFilePath)
	//if err != nil {
	//	return err
	//}
	//vttRelPath, _ := filepath.Rel(library.BasePath, vttAbsPath)
	cmd := exec.Command("./vttsprite", videoFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		h.logger.Errorf("Failed to generate vtt file. %v", err)
		return err
	}
	vttRelPath := fmt.Sprintf("%s.vtt", videoFile.FilePath)
	_, err = os.Stat(path.Join(library.BasePath, vttRelPath))
	if err != nil {
		h.logger.Errorf("Failed to generate vtt file.")
		return err
	}
	_, err = h.svcCtx.MediaLibrary.CreateFile(h.ctx, &medialibraryclient.CreateFileRequest{
		LibraryId: library.Id,
		MediaId:   media.ID,
		ParentId:  videoFile.ID,
		Type:      "video_thumbnails",
		FilePath:  vttRelPath,
	})
	if err != nil {
		return err
	}
	h.logger.Infof("Generated video thumbnails %s for file %d", vttRelPath, videoFile.ID)
	return nil
}

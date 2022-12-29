package queue

import "yet-another-media-server/media_library/medialibraryclient"

const (
	ActionVTTSprite = "vtt_sprite"
)

type VTTSpriteMessage struct {
	Action string `json:"action" mapstructure:"action"`
	Media  *Media `json:"media" mapstructure:"media"`
	File   *File  `json:"file" mapstructure:"file"`
}

func NewVTTSpriteMessage(media *medialibraryclient.Media, file *medialibraryclient.File) VTTSpriteMessage {
	return VTTSpriteMessage{
		Action: ActionVTTSprite,
		Media: &Media{
			ID:        media.Id,
			LibraryID: media.LibraryId,
		},
		File: &File{
			ID:       file.Id,
			Type:     file.Type,
			FilePath: file.FilePath,
		},
	}
}

type Media struct {
	ID        int32 `json:"id" mapstructure:"id"`
	LibraryID int32 `json:"library_id" mapstructure:"library_id"`
}

type File struct {
	ID       int32  `json:"id" mapstructure:"id"`
	Type     string `json:"type" mapstructure:"type"`
	FilePath string `json:"file_path" mapstructure:"file_path"`
}

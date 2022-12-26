package model

import (
	"gorm.io/gorm"
	"yet-another-media-server/media_library/media_library"
)

type Library struct {
	gorm.Model
	Name     string
	BasePath string
}

func (l *Library) ToProto() *media_library.Library {
	return &media_library.Library{
		Id:        int32(l.ID),
		CreatedAt: l.CreatedAt.Unix(),
		UpdatedAt: l.UpdatedAt.Unix(),
		Name:      l.Name,
		BasePath:  l.BasePath,
	}
}

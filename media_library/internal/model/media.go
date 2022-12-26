package model

import (
	"gorm.io/gorm"
	"yet-another-media-server/media_library/media_library"
)

type Media struct {
	gorm.Model
	LibraryID uint
	Library   Library
	Metadatas []Metadata `gorm:"many2many:media_metadata_relations;"`
	Files     []File
}

func (m *Media) ToProto() *media_library.Media {
	ret := &media_library.Media{
		Id:        int32(m.ID),
		CreatedAt: m.CreatedAt.Unix(),
		UpdatedAt: m.UpdatedAt.Unix(),
		LibraryId: int32(m.LibraryID),
		Metadatas: make([]*media_library.Metadata, 0, len(m.Metadatas)),
		Files:     make([]*media_library.File, 0, len(m.Files)),
	}
	for _, t := range m.Metadatas {
		ret.Metadatas = append(ret.Metadatas, t.ToProto())
	}
	for _, t := range m.Files {
		ret.Files = append(ret.Files, t.ToProto())
	}
	return ret
}

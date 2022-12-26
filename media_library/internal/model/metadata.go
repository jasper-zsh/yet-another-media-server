package model

import (
	"gorm.io/gorm"
	"yet-another-media-server/media_library/media_library"
)

type Metadata struct {
	gorm.Model
	LibraryID uint
	Type      string
	ParentID  uint
	Parent    *Metadata
	Children  []Metadata `gorm:"foreignkey:ParentID"`
	Value     string
}

func (m *Metadata) ToProto() *media_library.Metadata {
	result := &media_library.Metadata{
		Id:        int32(m.ID),
		CreatedAt: m.CreatedAt.Unix(),
		UpdatedAt: m.UpdatedAt.Unix(),
		Type:      m.Type,
		Children:  make([]*media_library.Metadata, 0),
		Value:     m.Value,
		ParentId:  int32(m.ParentID),
	}
	for _, t := range m.Children {
		result.Children = append(result.Children, t.ToProto())
	}
	return result
}

func GetMetadataByTypeAndValue(db *gorm.DB, libraryID uint, metaType, value string) (*Metadata, error) {
	metas := make([]*Metadata, 0)
	result := db.Where("library_id = ? AND type = ? AND value = ?", libraryID, metaType, value).Limit(1).Find(&metas)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(metas) == 0 {
		return nil, nil
	}
	return metas[0], nil
}

package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	MediaID  uint
	Media    Media
	FilePath string
	Ext      datatypes.JSONMap
}

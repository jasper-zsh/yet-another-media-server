package model

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	LibraryID uint
	Library   Library
	Metadatas []Metadata
	Files     []File
}

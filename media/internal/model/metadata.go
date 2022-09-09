package model

import "gorm.io/gorm"

type Metadata struct {
	gorm.Model
	MetadataDefinitionID uint
	MetadataDefinition   MetadataDefinition
	MetadataID           uint
	MetadataValue        MetadataValue
	MediaID              uint
	Media                Media
}

type MetadataDefinition struct {
	gorm.Model
	LibraryID uint
	Library   Library
	Name      string
	Values    []MetadataValue
}

type MetadataValue struct {
	gorm.Model
	MetadataDefinitionID uint
	MetadataDefinition   MetadataDefinition
	AliasID              *uint
	Aliases              []MetadataValue
	Value                string
	Picture              string
}

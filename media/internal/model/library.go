package model

import "gorm.io/gorm"

type Library struct {
	gorm.Model
	Name     string
	BasePath string
}

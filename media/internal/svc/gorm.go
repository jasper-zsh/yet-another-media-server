package svc

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"yet-another-media-server/media/internal/config"
	"yet-another-media-server/media/internal/model"
)

func InitGorm(conf config.GormConf) *gorm.DB {
	var dialector gorm.Dialector
	switch conf.Driver {
	case "sqlite":
		dialector = sqlite.Open(conf.DSN)
	default:
		panic("unsupported database driver")
	}
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to init gorm")
	}
	err = db.AutoMigrate(
		&model.File{},
		&model.Library{},
		&model.Media{},
		&model.Metadata{}, &model.MetadataDefinition{}, &model.MetadataValue{},
	)
	if err != nil {
		panic("failed to migrate")
	}
	return db
}

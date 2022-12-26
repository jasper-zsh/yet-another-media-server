package svc

import (
	"gorm.io/gorm"
	"yet-another-media-server/media_library/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     InitGorm(c.Gorm),
	}
}

package global

import (
	"simplebank/pkg/logger"
	"simplebank/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
)

/*
Config
Redis
MySQL
PostgreSQL
*/
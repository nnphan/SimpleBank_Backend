package global

import (
	"simplebank/pkg/logger"
	"simplebank/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Pdb *gorm.DB
	Rdb *redis.Client
)

/*
Config
Redis
MySQL
PostgreSQL
*/
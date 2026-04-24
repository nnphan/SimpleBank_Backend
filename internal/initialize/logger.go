package initialize

import (
	"simplebank/global"
	"simplebank/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
package initialize

import (
	"simplebank/global"

	"go.uber.org/zap"
)

func Run() {
	InitConfig()
	InitLogger()
	global.Logger.Info("Logger initialized successfully", zap.String("ok", "success"))
	InitRedis()
	// InitPostgreSQL()

	//Init router
	r := InitRouter()
	r.Run(":8002")

}
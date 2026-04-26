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
	InitPostgreSQL()
	global.Logger.Info("InitPostgreSQL successfully", zap.String("ok", "success"))
	//Init router
	r := InitRouter()
	r.Run(":8002")

}
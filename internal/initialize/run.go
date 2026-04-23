package initialize

func Run() {
	InitConfig()
	InitLogger()
	InitRedis()
	InitRouter()
	// InitPostgreSQL()
}
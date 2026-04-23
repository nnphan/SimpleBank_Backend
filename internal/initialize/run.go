package initialize

func Run() {
	InitConfig()
	InitLogger()
	InitRedis()
	// InitPostgreSQL()

	
	//Init router
	r := InitRouter()
	r.Run(":8002")

}
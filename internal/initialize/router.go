package initialize

import (
	"simplebank/global"
	"simplebank/internal/router"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// Init middleware
	r.Use()// logging
	r.Use() //CORS
	r.Use() // Limiter global

	manageRouter := router.RouterGroupApp.Manage
	userRouter := router.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/checkStatus") // health check, tracking monitor
	}
	// User router
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	// Admin router
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}

	
	return r
}
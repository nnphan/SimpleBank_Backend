package initialize

import (
	"simplebank/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", controller.NewPingController().Ping)
	}

	return r
}
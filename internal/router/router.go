package router

import (
	"simplebank/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/ping", controller.NewPingController().Ping)
		v1.GET("/user/1", controller.NewUserController().GetUserById)
		v1.POST("/users", controller.NewUserController().CreateUser)
	}
	return r
}


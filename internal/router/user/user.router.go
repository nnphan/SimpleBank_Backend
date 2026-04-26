package user

import (
	"simplebank/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	// non dependency
	// ur := repository.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandler := controller.NewUserController(us)

	// use DI
	userController, _ := wire.InitUserRouterHandler()
	


	//Public router
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/login")
	}

	//Private router
	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.GET("/profile2")
		userRouterPrivate.PUT("/update-profile2")
	}

}
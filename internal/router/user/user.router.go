package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	//Public router
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/login")
	}

	//Private router
	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.GET("/profile")
		userRouterPrivate.PUT("/update-profile")
	}

}
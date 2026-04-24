package manage

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	//Private router
	userRouterPrivate := router.Group("/user")
	{
		userRouterPrivate.GET("/profile")
		userRouterPrivate.PUT("/update-profile")
		userRouterPrivate.POST("/active-user")
	}

}